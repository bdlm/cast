package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toMap returns a map of the specified reflect.Value type built from from.
//
// Options:
//   - DEFAULT: map, default return value on error.
//   - DUPLICATE_KEY_ERROR: bool, error on duplicate map key (map→map only).
//   - PRIVATE: bool, include unexported struct fields (struct→map only).
//   - STRICT: bool, return an error instead of skipping unconvertible fields.
func toMap(to reflect.Value, from any, ops ops) (any, error) {
	var ret any
	if ops.hasDefault {
		defaultVal := reflect.ValueOf(ops.defaultVal)
		if defaultVal.IsValid() && !defaultVal.Type().AssignableTo(to.Type()) {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		ret = ops.defaultVal
		ops = ops.Delete(DEFAULT) // Prevent DEFAULT from being passed to element casts.
	}

	fromVal := reflect.Indirect(reflect.ValueOf(from))
	if !fromVal.IsValid() {
		return ret, errors.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
	}

	switch fromVal.Kind() {
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
		return ret, errors.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
	}
}

// mapFromMap converts a source map to the target map type, casting each key
// and value individually. DUPLICATE_KEY_ERROR causes it to error on collision.
func mapFromMap(to reflect.Value, src reflect.Value, ops ops) (any, error) {
	dupKeyErr := ops.dupKeyErr
	targetMap := reflect.MakeMap(to.Type())
	keyType := to.Type().Key()
	valType := to.Type().Elem()

	for _, srcKey := range src.MapKeys() {
		castKey, err := castToType(srcKey.Interface(), keyType, ops)
		if err != nil {
			return nil, err
		}
		if dupKeyErr && targetMap.MapIndex(castKey).IsValid() {
			return nil, errors.Errorf("duplicate key %v", castKey.Interface())
		}
		castVal, err := castToType(src.MapIndex(srcKey).Interface(), valType, ops)
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
func mapFromStruct(to reflect.Value, src reflect.Value, ops ops) (any, error) {
	targetMap := reflect.MakeMap(to.Type())
	private := ops.private
	strict := ops.strict
	keyType := to.Type().Key()
	valType := to.Type().Elem()

	if err := collectStructFields(targetMap, src, keyType, valType, private, strict, ops); err != nil {
		return nil, err
	}
	return targetMap.Interface(), nil
}

// collectStructFields iterates struct fields and populates targetMap.
// Exported anonymous (embedded) structs are recursed into so promoted fields
// appear at the top level, matching Go's promotion semantics.
func collectStructFields(
	targetMap reflect.Value, src reflect.Value,
	keyType reflect.Type, valType reflect.Type,
	private, strict bool, ops ops,
) error {
	for i := 0; i < src.NumField(); i++ {
		field := src.Type().Field(i)
		fieldVal := src.Field(i)

		// Promoted anonymous (embedded) struct fields.
		if field.Anonymous {
			if !field.IsExported() && !private {
				continue
			}
			embedded := reflect.Indirect(fieldVal)
			if embedded.IsValid() && embedded.Kind() == reflect.Struct {
				if err := collectStructFields(targetMap, embedded, keyType, valType, private, strict, ops); err != nil {
					return err
				}
				continue
			}
			// Non-struct anonymous field: fall through to normal field processing.
		}

		if !field.IsExported() && !private {
			continue
		}

		rawVal, canExtract := extractFieldValue(fieldVal)
		if !canExtract {
			if strict {
				return errors.Errorf("cannot extract unexported non-scalar field %q (%v)", field.Name, fieldVal.Kind())
			}
			continue
		}

		castKey, err := castToType(field.Name, keyType, ops)
		if err != nil {
			if strict {
				return errors.Errorf("cannot cast field name %q to map key: %v", field.Name, err)
			}
			continue
		}

		// Nested struct: recurse when the target value type supports it.
		var castVal reflect.Value
		if fieldVal.Kind() == reflect.Struct && fieldVal.CanInterface() {
			switch valType.Kind() {
			case reflect.Interface:
				if valType.NumMethod() != 0 {
					// Non-empty interface (e.g. io.Reader): cast the raw value
					// directly rather than wrapping the struct in a nested map.
					castVal, err = castToType(rawVal, valType, ops)
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
				castVal, err = castToType(rawVal, valType, ops)
				if err != nil {
					if strict {
						return err
					}
					continue
				}
			}
		} else {
			castVal, err = castToType(rawVal, valType, ops)
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
func mapFromSlice(to reflect.Value, src reflect.Value, ops ops) (any, error) {
	targetMap := reflect.MakeMap(to.Type())
	keyType := to.Type().Key()
	valType := to.Type().Elem()

	for i := 0; i < src.Len(); i++ {
		castKey, err := castToType(i, keyType, ops)
		if err != nil {
			return nil, errors.Errorf("cannot cast index %d to map key type %v: %v", i, keyType, err)
		}
		elem := src.Index(i)
		var elemIface any
		if elem.CanInterface() {
			elemIface = elem.Interface()
		} else {
			val, ok := extractFieldValue(elem)
			if !ok {
				return nil, errors.Errorf("cannot extract slice element at index %d", i)
			}
			elemIface = val
		}
		castVal, err := castToType(elemIface, valType, ops)
		if err != nil {
			return nil, err
		}
		targetMap.SetMapIndex(castKey, castVal)
	}
	return targetMap.Interface(), nil
}

// extractFieldValue returns a struct field's value as any, handling unexported
// fields via kind-specific reflect methods. Returns (nil, false) for unexported
// non-scalar fields that cannot be read without unsafe.
func extractFieldValue(v reflect.Value) (any, bool) {
	if v.CanInterface() {
		return v.Interface(), true
	}
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	case reflect.String:
		return v.String(), true
	default:
		return nil, false
	}
}
