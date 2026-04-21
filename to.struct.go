package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// ToStruct casts from into a struct of type T, ignoring errors. See ToStructE
// for full documentation.
func ToStruct[T any](from any, ops ...Op) T {
	ret, _ := ToStructE[T](from, ops...)
	return ret
}

// ToStructE casts from into a struct of type T, returning any errors.
//
// T must be a struct type. The source may be a map with keys convertible to
// string, or another struct whose exported field names overlap with T.
//
// Field matching is case-sensitive. For each exported field in T:
//   - If the source has no matching key, the field retains its zero value
//     (or returns an error when STRICT is set).
//   - If the source value cannot be cast to the field type, the field is
//     skipped (or returns an error when STRICT is set).
//
// Options:
//   - DEFAULT: T, value to return on error.
//   - STRICT: bool, error on unknown source keys or unconvertible field values.
func ToStructE[T any](from any, ops ...Op) (T, error) {
	var zero T
	to := reflect.ValueOf(&zero).Elem()
	if to.Kind() != reflect.Struct {
		return zero, errors.WrapE(Error, errors.Errorf("ToStructE requires a struct target type, got %T", zero))
	}
	options := parseOps(ops)
	raw, err := toStruct(to, from, options)
	if err != nil {
		return zero, errors.WrapE(Error, err)
	}
	result, ok := raw.(T)
	if !ok {
		return zero, errors.WrapE(Error, errors.Errorf("internal: toStruct returned %T, want %T", raw, zero))
	}
	return result, nil
}

// toStruct populates a struct value of to's type from from.
//
// Options:
//   - DEFAULT: target type, value to return on error.
//   - STRICT: bool, error on unmatched source keys or unconvertible field values.
func toStruct(to reflect.Value, from any, ops ops) (any, error) {
	var defaultVal any
	if ops.hasDefault {
		defaultVal = ops.defaultVal
	} else {
		defaultVal = reflect.New(to.Type()).Elem().Interface()
	}

	result := reflect.New(to.Type()).Elem()

	// Struct source fast paths — avoid materializing an intermediate map[string]any.
	if fromVal := reflect.Indirect(reflect.ValueOf(from)); fromVal.IsValid() && fromVal.Kind() == reflect.Struct {
		// Same-type fast path: identical type means identical field set; no
		// field-by-field work is needed, and STRICT never fails.
		if fromVal.Type() == to.Type() {
			return fromVal.Interface(), nil
		}

		// Different-type struct: build a reflect.Value index over the source
		// (O(M) t.Field calls), then hydrate the target (O(N) t.Field calls).
		// Boxing via Interface() is deferred until each field is actually used.
		srcFields := make(map[string]reflect.Value, fromVal.Type().NumField())
		collectSourceFieldValues(srcFields, fromVal)

		var usedKeys map[string]bool
		if ops.strict {
			usedKeys = make(map[string]bool, len(srcFields))
		}
		if err := hydrateFromValues(result, srcFields, usedKeys, ops.Global()); err != nil {
			return defaultVal, err
		}
		if ops.strict {
			for k := range srcFields {
				if !usedKeys[k] {
					return defaultVal, errors.Errorf("source key %q has no matching field in %v", k, to.Type())
				}
			}
		}
		return result.Interface(), nil
	}

	// General path: normalize source to map[string]any then hydrate.
	srcMap, err := normalizeToStringMap(from)
	if err != nil {
		return defaultVal, errors.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
	}

	var usedKeys map[string]bool
	if ops.strict {
		usedKeys = make(map[string]bool, len(srcMap))
	}

	if err := hydrateStruct(result, srcMap, usedKeys, ops.Global()); err != nil {
		return defaultVal, err
	}

	if ops.strict {
		// Embedded type names (e.g. "Base" for an anonymous embedded Base
		// struct) are never present as srcMap keys, so they will not appear
		// here as unused keys and will not trigger a STRICT error.
		for k := range srcMap {
			if !usedKeys[k] {
				return defaultVal, errors.Errorf("source key %q has no matching field in %v", k, to.Type())
			}
		}
	}

	return result.Interface(), nil
}

// hydrateStruct sets exported fields of result from srcMap, recording matched
// keys in usedKeys (nil when STRICT is not set). Anonymous (embedded) struct
// fields are recursed into so their promoted field names are resolved at the
// top level.
func hydrateStruct(result reflect.Value, srcMap map[string]any, usedKeys map[string]bool, ops ops) error {
	t := result.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := result.Field(i)

		// Recurse into anonymous (embedded) struct fields regardless of whether
		// the embedded type is exported. Exported fields within an unexported
		// anonymous struct are still promoted and settable.
		if field.Anonymous {
			fv := fieldVal
			// For a nil pointer-to-struct anonymous field, allocate so that
			// promoted fields can be hydrated. Skip if the field is not
			// settable (unexported embedded pointer from another package).
			if fv.Kind() == reflect.Pointer && fv.Type().Elem().Kind() == reflect.Struct && fv.IsNil() && fv.CanSet() {
				newPtr := reflect.New(fv.Type().Elem())
				fv.Set(newPtr)
			}
			embedded := reflect.Indirect(fv)
			if embedded.IsValid() && embedded.Kind() == reflect.Struct {
				if err := hydrateStruct(embedded, srcMap, usedKeys, ops); err != nil {
					return err
				}
				continue
			}
			// Non-struct anonymous field (e.g. embedded named scalar type):
			// fall through to normal field matching so its promoted name is
			// found in srcMap.
		}

		if !field.IsExported() {
			continue
		}

		raw, ok := srcMap[field.Name]
		if !ok {
			if ops.strict {
				return errors.Errorf("no source key for required field %q", field.Name)
			}
			continue
		}
		if usedKeys != nil {
			usedKeys[field.Name] = true
		}

		castVal, err := castToType(raw, field.Type, ops)
		if err != nil {
			if ops.strict {
				return errors.Errorf("field %q: %v", field.Name, err)
			}
			continue
		}
		fieldVal.Set(castVal)
	}
	return nil
}

// collectSourceFieldValues populates dst with the exported fields of v as
// reflect.Values, deferring Interface() boxing until each field is actually
// used during hydration. Anonymous struct fields are inlined recursively to
// match Go promotion semantics.
func collectSourceFieldValues(dst map[string]reflect.Value, v reflect.Value) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)
		if field.Anonymous {
			embedded := reflect.Indirect(fieldVal)
			if embedded.IsValid() && embedded.Kind() == reflect.Struct {
				collectSourceFieldValues(dst, embedded)
				continue
			}
		}
		if !field.IsExported() || !fieldVal.CanInterface() {
			continue
		}
		dst[field.Name] = fieldVal
	}
}

// hydrateFromValues sets exported fields of result from srcFields
// (map[string]reflect.Value), recording matched keys in usedKeys (nil when
// STRICT is not set). Anonymous embedded target fields are recursed into so
// their promoted field names are resolved at the top level.
func hydrateFromValues(result reflect.Value, srcFields map[string]reflect.Value, usedKeys map[string]bool, ops ops) error {
	t := result.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := result.Field(i)

		if field.Anonymous {
			fv := fieldVal
			if fv.Kind() == reflect.Pointer && fv.Type().Elem().Kind() == reflect.Struct && fv.IsNil() && fv.CanSet() {
				newPtr := reflect.New(fv.Type().Elem())
				fv.Set(newPtr)
			}
			embedded := reflect.Indirect(fv)
			if embedded.IsValid() && embedded.Kind() == reflect.Struct {
				if err := hydrateFromValues(embedded, srcFields, usedKeys, ops); err != nil {
					return err
				}
				continue
			}
		}

		if !field.IsExported() {
			continue
		}

		srcVal, ok := srcFields[field.Name]
		if !ok {
			if ops.strict {
				return errors.Errorf("no source key for required field %q", field.Name)
			}
			continue
		}
		if usedKeys != nil {
			usedKeys[field.Name] = true
		}

		castVal, err := castToType(srcVal.Interface(), field.Type, ops)
		if err != nil {
			if ops.strict {
				return errors.Errorf("field %q: %v", field.Name, err)
			}
			continue
		}
		fieldVal.Set(castVal)
	}
	return nil
}

// normalizeToStringMap converts from to a map[string]any. Accepts:
//   - map[string]any (returned as-is)
//   - any map with string-convertible keys
//   - struct or *struct (exported fields become keys; embedded fields inlined)
func normalizeToStringMap(from any) (map[string]any, error) {
	if from == nil {
		return nil, errors.Errorf("cannot convert nil to struct fields")
	}
	if m, ok := from.(map[string]any); ok {
		return m, nil
	}
	if m, ok := from.(map[string]string); ok {
		result := make(map[string]any, len(m))
		for k, v := range m {
			result[k] = v
		}
		return result, nil
	}

	v := reflect.Indirect(reflect.ValueOf(from))
	if !v.IsValid() {
		return nil, errors.Errorf("cannot convert %T to string map", from)
	}

	switch v.Kind() {
	case reflect.Map:
		result := make(map[string]any, v.Len())
		for _, key := range v.MapKeys() {
			k, err := ToE[string](key.Interface())
			if err != nil {
				return nil, errors.Errorf("map key %v cannot be cast to string: %v", key.Interface(), err)
			}
			if elem := v.MapIndex(key); elem.CanInterface() {
				result[k] = elem.Interface()
			}
		}
		return result, nil
	case reflect.Struct:
		result := make(map[string]any)
		collectExportedFields(result, v)
		return result, nil
	default:
		return nil, errors.Errorf("cannot convert %T (kind %v) to struct fields", from, v.Kind())
	}
}

// collectExportedFields populates dst with the exported fields of v. Anonymous
// struct fields (exported or unexported) are inlined recursively to match Go
// promotion semantics; exported fields within unexported anonymous structs are
// still included.
func collectExportedFields(dst map[string]any, v reflect.Value) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)
		if field.Anonymous {
			embedded := reflect.Indirect(fieldVal)
			if embedded.IsValid() && embedded.Kind() == reflect.Struct {
				collectExportedFields(dst, embedded)
				continue
			}
		}
		if !field.IsExported() || !fieldVal.CanInterface() {
			continue
		}
		dst[field.Name] = fieldVal.Interface()
	}
}
