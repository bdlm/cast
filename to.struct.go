package cast

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

// fieldKey returns the source-map key to use when matching this struct field.
// Priority: cast tag > json tag (name portion only) > field name.
// Returns ("", false) when the tag value is "-", meaning skip this field.
func fieldKey(field reflect.StructField) (string, bool) {
	if tag, ok := field.Tag.Lookup("cast"); ok {
		if tag == "-" {
			return "", false
		}
		return tag, true
	}
	if tag, ok := field.Tag.Lookup("json"); ok {
		name := strings.SplitN(tag, ",", 2)[0]
		if name == "-" {
			return "", false
		}
		if name != "" {
			return name, true
		}
	}
	return field.Name, true
}

// ToStruct casts from into a struct of type T, ignoring errors. See ToStructE
// for full documentation.
func ToStruct[T any](from any, ops ...Op) T {
	ret, _ := ToStructE[T](from, ops...)
	return ret
}

// ToStructE casts from into a struct of type T, returning any errors.
//
// T must be a struct type. The source may be a map with keys convertible to
// string, or another struct whose field names overlap with T.
//
// Field matching is case-sensitive and tag-aware. The lookup key for each
// target field is resolved by checking a cast tag first, then a json tag,
// then the field name. For each target field in T:
//   - If the source has no matching key, the field retains its zero value
//     (or returns an error when STRICT is set).
//   - If the source value cannot be cast to the field type, the field is
//     skipped (or returns an error when STRICT is set).
//
// Options:
//   - DEFAULT: T, value to return on error.
//   - PRIVATE: bool, include unexported fields in source collection and target hydration.
//   - STRICT: bool, error on unknown source keys or unconvertible field values.
func ToStructE[T any](from any, ops ...Op) (T, error) {
	var zero T
	to := reflect.ValueOf(&zero).Elem()
	if to.Kind() != reflect.Struct {
		return zero, fmt.Errorf("%w, %w", Error, fmt.Errorf("ToStructE requires a struct target type, got %T", zero))
	}
	options := parseOps(ops)
	raw, err := toStruct(to, from, options)
	if err != nil {
		return zero, fmt.Errorf("%w, %w", Error, err)
	}
	result, ok := raw.(T)
	if !ok {
		return zero, fmt.Errorf("%w, %w", Error, fmt.Errorf("cast failed: returned %T, want %T", raw, zero))
	}
	return result, nil
}

// toStruct populates a struct value of to's type from from.
//
// Options:
//   - DEFAULT: target type, value to return on error.
//   - PRIVATE: bool, include unexported fields.
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

		// Different-type struct: build a value index over the source fields
		// (O(M) field calls), then hydrate the target (O(N) field calls).
		srcFields := make(map[string]any, fromVal.Type().NumField())
		collectSourceFieldValues(srcFields, fromVal, ops.private)

		var usedKeys map[string]bool
		if ops.strict {
			usedKeys = make(map[string]bool, len(srcFields))
		}
		if err := hydrateStruct(result, srcFields, usedKeys, ops.Global()); err != nil {
			return defaultVal, err
		}
		if ops.strict {
			for k := range srcFields {
				if !usedKeys[k] {
					return defaultVal, fmt.Errorf("source key %q has no matching field in %v", k, to.Type())
				}
			}
		}
		return result.Interface(), nil
	}

	// General path: normalize source to map[string]any then hydrate.
	srcMap, err := normalizeToStringMap(from)
	if err != nil {
		return defaultVal, fmt.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
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
				return defaultVal, fmt.Errorf("source key %q has no matching field in %v", k, to.Type())
			}
		}
	}

	return result.Interface(), nil
}

// hydrateStruct sets fields of result from srcMap, recording matched keys in
// usedKeys (nil when STRICT is not set). Anonymous (embedded) struct fields are
// recursed into so their promoted field names are resolved at the top level.
// Field lookup keys honour fieldKey priority (cast tag > json tag > field name).
// When ops.private is true, unexported fields are set via unsafe.Pointer.
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

		if !field.IsExported() && !ops.private {
			continue
		}

		key, tagOk := fieldKey(field)
		if !tagOk {
			continue // tagged with "-"
		}

		raw, ok := srcMap[key]
		if !ok {
			if ops.strict && field.IsExported() {
				return fmt.Errorf("no source key for required field %q", field.Name)
			}
			continue
		}
		if usedKeys != nil {
			usedKeys[key] = true
		}

		castVal, err := castToType(raw, field.Type, ops)
		if err != nil {
			if ops.strict {
				return fmt.Errorf("field %q: %v", field.Name, err)
			}
			continue
		}
		if fieldVal.CanSet() {
			fieldVal.Set(castVal)
		} else if ops.private && fieldVal.CanAddr() {
			// Bypass the export check for unexported fields using unsafe.
			reflect.NewAt(fieldVal.Type(), unsafe.Pointer(fieldVal.UnsafeAddr())).Elem().Set(castVal)
		}
	}
	return nil
}

// collectSourceFieldValues populates dst with each field's value from v.
// When private is true, unexported scalar fields are included via extractFieldValue.
// Anonymous struct fields are inlined recursively to match Go promotion semantics.
// Keys are resolved via fieldKey (cast tag > json tag > field name).
func collectSourceFieldValues(dst map[string]any, v reflect.Value, private bool) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		if field.Anonymous {
			embedded := reflect.Indirect(fieldVal)
			if embedded.IsValid() && embedded.Kind() == reflect.Struct {
				collectSourceFieldValues(dst, embedded, private)
				continue
			}
		}

		if !field.IsExported() && !private {
			continue
		}

		key, ok := fieldKey(field)
		if !ok {
			continue // tagged with "-"
		}

		if fieldVal.CanInterface() {
			dst[key] = fieldVal.Interface()
		} else if private {
			if val, canExtract := extractFieldValue(fieldVal); canExtract {
				dst[key] = val
			}
		}
	}
}

// normalizeToStringMap converts from to a map[string]any. Accepts:
//   - map[string]any (returned as-is)
//   - any map with string-convertible keys
//   - struct or *struct (exported fields become keys; embedded fields inlined)
func normalizeToStringMap(from any) (map[string]any, error) {
	if from == nil {
		return nil, fmt.Errorf("cannot convert nil to struct fields")
	}
	if mapVal, ok := from.(map[string]any); ok {
		return mapVal, nil
	}
	if mapVal, ok := from.(map[string]string); ok {
		result := make(map[string]any, len(mapVal))
		for k, v := range mapVal {
			result[k] = v
		}
		return result, nil
	}

	v := reflect.Indirect(reflect.ValueOf(from))
	if !v.IsValid() {
		return nil, fmt.Errorf("cannot convert %T to string map", from)
	}

	switch v.Kind() {
	case reflect.Map:
		result := make(map[string]any, v.Len())
		for _, key := range v.MapKeys() {
			k, err := ToE[string](key.Interface())
			if err != nil {
				return nil, fmt.Errorf("map key %v cannot be cast to string: %v", key.Interface(), err)
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
		return nil, fmt.Errorf("cannot convert %T (kind %v) to struct fields", from, v.Kind())
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
		key, ok := fieldKey(field)
		if !ok {
			continue
		}
		dst[key] = fieldVal.Interface()
	}
}
