package cast

import (
	"fmt"
	"reflect"
)

// ToMap returns a map of the specified reflect.Value type built from from.
//
// Options:
//   - DEFAULT: map, default return value on error.
//   - DUPLICATE_KEY_ERROR: bool, error on duplicate map key (map→map only).
//   - PRIVATE: bool, include unexported struct fields (struct→map only).
//   - STRICT: bool, return an error instead of skipping unconvertible fields.
func ToMap(to reflect.Value, from any, ops Ops) (any, error) {
	var ret any
	if ops.HasDefault {
		defaultVal := reflect.ValueOf(ops.DefaultVal)
		if defaultVal.IsValid() && !defaultVal.Type().AssignableTo(to.Type()) {
			return ret, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.DefaultVal)
		}
		ret = ops.DefaultVal
	}

	fromVal := reflect.Indirect(reflect.ValueOf(from))
	if !fromVal.IsValid() {
		return ret, fmt.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
	}

	switch fromVal.Kind() {
	case reflect.String:
		s := fromVal.String()
		if looksLikeCollection(s) {
			if decoded, ok := unmarshalCollection(s); ok {
				return ToMap(to, decoded, ops)
			}
		}
		return ret, fmt.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
	case reflect.Map:
		result, err := mapFromMap(to, fromVal, ops)
		if err != nil {
			return ret, err
		}
		return result, nil
	case reflect.Struct:
		result, err := mapFromStruct(to, fromVal, ops)
		if err != nil {
			return ret, err
		}
		return result, nil
	case reflect.Slice, reflect.Array:
		result, err := mapFromSlice(to, fromVal, ops)
		if err != nil {
			return ret, err
		}
		return result, nil
	default:
		return ret, fmt.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
	}
}

// mapFromMap converts a source map to the target map type, casting each key
// and value individually. DUPLICATE_KEY_ERROR causes it to error on collision.
func mapFromMap(to reflect.Value, src reflect.Value, ops Ops) (any, error) {
	dupKeyErr := ops.DupKeyErr
	elemOps := ops.Global()
	targetMap := reflect.MakeMap(to.Type())
	keyType := to.Type().Key()
	valType := to.Type().Elem()

	for _, srcKey := range src.MapKeys() {
		castKey, err := CastToType(srcKey.Interface(), keyType, elemOps)
		if err != nil {
			return nil, err
		}
		if dupKeyErr && targetMap.MapIndex(castKey).IsValid() {
			return nil, fmt.Errorf("duplicate key %v", castKey.Interface())
		}
		castVal, err := CastToType(src.MapIndex(srcKey).Interface(), valType, elemOps)
		if err != nil {
			return nil, err
		}
		targetMap.SetMapIndex(castKey, castVal)
	}
	return targetMap.Interface(), nil
}

// mapFromStruct converts a struct to a map using field names as keys.
// Embedded (anonymous) struct fields are recursively inlined. Unexported
// fields are skipped unless PRIVATE is set; when PRIVATE is set, unexported
// scalar fields are included, while unexported non-scalar fields are skipped
// (or return an error when STRICT is set).
func mapFromStruct(to reflect.Value, src reflect.Value, ops Ops) (any, error) {
	targetMap := reflect.MakeMap(to.Type())
	private := ops.Private
	strict := ops.Strict
	keyType := to.Type().Key()
	valType := to.Type().Elem()

	if err := collectStructFields(targetMap, src, keyType, valType, private, strict, ops.Global()); err != nil {
		return nil, err
	}
	return targetMap.Interface(), nil
}

// collectStructFields iterates struct fields and populates targetMap.
func collectStructFields(
	targetMap reflect.Value, src reflect.Value,
	keyType reflect.Type, valType reflect.Type,
	private, strict bool, ops Ops,
) error {
	for i := 0; i < src.NumField(); i++ {
		field := src.Type().Field(i)
		fieldVal := src.Field(i)

		if field.Anonymous {
			embedded := reflect.Indirect(fieldVal)
			if embedded.IsValid() && embedded.Kind() == reflect.Struct {
				if err := collectStructFields(targetMap, embedded, keyType, valType, private, strict, ops); err != nil {
					return err
				}
				continue
			}
		}

		if !field.IsExported() && !private {
			continue
		}

		rawVal, canExtract := extractFieldValue(fieldVal)
		if !canExtract {
			if strict {
				return fmt.Errorf("cannot extract unexported non-scalar field %q (%v)", field.Name, fieldVal.Kind())
			}
			continue
		}

		key, tagOk := fieldKey(field)
		if !tagOk {
			continue // tagged with "-"
		}
		castKey, err := CastToType(key, keyType, ops)
		if err != nil {
			if strict {
				return fmt.Errorf("cannot cast field name %q to map key: %v", key, err)
			}
			continue
		}

		// Nested struct: recurse when the target value type supports it.
		var castVal reflect.Value
		if fieldVal.Kind() == reflect.Struct && fieldVal.CanInterface() {
			switch valType.Kind() {
			case reflect.Interface:
				if valType.NumMethod() != 0 {
					castVal, err = CastToType(rawVal, valType, ops)
					if err != nil {
						if strict {
							return err
						}
						continue
					}
				} else {
					anyType := reflect.TypeOf((*any)(nil)).Elem()
					nestedMapType := reflect.MapOf(keyType, anyType)
					nested, nestedErr := mapFromStruct(reflect.Zero(nestedMapType), fieldVal, ops)
					if nestedErr != nil {
						if strict {
							return nestedErr
						}
						continue
					}
					castVal = reflect.ValueOf(nested)
				}
			case reflect.Map:
				nested, nestedErr := mapFromStruct(reflect.Zero(valType), fieldVal, ops)
				if nestedErr != nil {
					if strict {
						return nestedErr
					}
					continue
				}
				castVal = reflect.ValueOf(nested)
			default:
				castVal, err = CastToType(rawVal, valType, ops)
				if err != nil {
					if strict {
						return err
					}
					continue
				}
			}
		} else {
			castVal, err = CastToType(rawVal, valType, ops)
			if err != nil {
				if strict {
					return err
				}
				continue
			}
		}

		targetMap.SetMapIndex(castKey, castVal)
	}
	return nil
}

// mapFromSlice converts a slice or array to a map using element indices as keys.
func mapFromSlice(to reflect.Value, src reflect.Value, ops Ops) (any, error) {
	elemOps := ops.Global()
	targetMap := reflect.MakeMap(to.Type())
	keyType := to.Type().Key()
	valType := to.Type().Elem()

	for i := 0; i < src.Len(); i++ {
		castKey, err := CastToType(i, keyType, elemOps)
		if err != nil {
			return nil, fmt.Errorf("cannot cast index %d to map key type %v: %v", i, keyType, err)
		}
		elem := src.Index(i)
		var elemIface any
		if elem.CanInterface() {
			elemIface = elem.Interface()
		} else {
			val, ok := extractFieldValue(elem)
			if !ok {
				return nil, fmt.Errorf("cannot extract slice element at index %d", i)
			}
			elemIface = val
		}
		castVal, err := CastToType(elemIface, valType, elemOps)
		if err != nil {
			return nil, err
		}
		targetMap.SetMapIndex(castKey, castVal)
	}
	return targetMap.Interface(), nil
}
