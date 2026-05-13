package cast

import (
	"fmt"
	"reflect"
	"slices"
)

// toSlice returns a slice containing the specified reflect.Value type
// containing the from value.
//
// Options:
//   - DEFAULT: slice, default return value on error.
//   - LENGTH: int, initial backing-array capacity, default 1. Must be greater
//     than or equal to 0.
//   - UNIQUE_VALUES: bool, deduplicate slice elements after conversion.
func toSlice(to reflect.Value, val any, ops ops) (any, error) {
	var defaultValue any

	if ops.hasDefault {
		defaultVal := reflect.ValueOf(ops.defaultVal)
		if defaultVal.IsValid() && !defaultVal.Type().AssignableTo(to.Type()) {
			return defaultValue, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		defaultValue = ops.defaultVal
	}

	size := 1
	if ops.hasLength {
		var sizeErr error
		if size, sizeErr = ToE[int](ops.lengthVal); sizeErr != nil {
			return defaultValue, fmt.Errorf(ErrorInvalidOption, "LENGTH", ops.lengthVal)
		}
	}
	if size < 0 {
		return defaultValue, fmt.Errorf("invalid array length %d", size)
	}

	// Read local flags then strip them for element-level casts.
	uniqueVals := ops.uniqueVals
	elemOps := ops.Global()

	// DECODE=json: JSON-decode the source before conversion. This applies to
	// string, error, and non-named-scalar Stringer sources; named scalar struct
	// types (time.Time, etc.) are excluded. A decoded array becomes a normal
	// slice source; a decoded scalar becomes a scalar-wrap source; this bypasses
	// the []byte/[]rune special cases below.
	if decoded, applied, decErr := tryDecodeJSON(val, ops); decErr != nil {
		return defaultValue, decErr
	} else if applied {
		val = decoded
		ops = ops.Delete(DECODE)
		elemOps = ops.Global()
	}

	// Special cases mirroring Go's built-in string conversions:
	//   string → []byte / []uint8      (and named variants with Uint8 element)
	//   string → []rune / []int32      (and named variants with Int32 element)
	// These must be handled before the source-kind dispatch below because a
	// string is not a slice/array kind, and iterating bytes would give wrong
	// results for []rune (bytes ≠ Unicode code points for multibyte UTF-8).
	// For other element types, try scalar wrap first, then JSON as a last resort.
	if s, ok := val.(string); ok {
		var result any
		switch to.Type().Elem().Kind() {
		case reflect.Uint8: // []byte / []uint8 and named variants
			bs := []byte(s)
			if to.Type() == reflect.TypeOf(bs) {
				result = bs
			} else {
				rv := reflect.MakeSlice(to.Type(), len(bs), len(bs))
				for i, b := range bs {
					rv.Index(i).SetUint(uint64(b))
				}
				result = rv.Interface()
			}
		case reflect.Int32: // []rune / []int32 and named variants (rune = int32)
			rs := []rune(s)
			if to.Type() == reflect.TypeOf(rs) {
				result = rs
			} else {
				rv := reflect.MakeSlice(to.Type(), len(rs), len(rs))
				for i, r := range rs {
					rv.Index(i).SetInt(int64(r))
				}
				result = rv.Interface()
			}
		default:
			// Try scalar wrap (e.g. "true" → []bool{true}, "42" → []int{42}).
			if r, err := sliceFromSingle(s, to, elemOps, nil); err == nil {
				result = r
			} else if looksLikeCollection(s) {
				// Last resort: decode JSON array/object and recurse.
				if decoded, ok := unmarshalCollection(s); ok {
					return toSlice(to, decoded, ops)
				}
				return defaultValue, fmt.Errorf(ErrorStrUnableToCast, val, val, to.Interface())
			} else {
				return defaultValue, fmt.Errorf(ErrorStrUnableToCast, val, val, to.Interface())
			}
		}
		if uniqueVals {
			rv := reflect.ValueOf(result)
			result = dedupeSliceVal(rv).Interface()
		}
		return result, nil
	}

	// Nil source: wrap as single-element slice containing the zero value.
	if val == nil {
		return sliceFromSingle(nil, to, elemOps, defaultValue)
	}

	fromVal := reflect.Indirect(reflect.ValueOf(val))
	if !fromVal.IsValid() {
		// Nil pointer or nil interface after dereferencing.
		return sliceFromSingle(nil, to, elemOps, defaultValue)
	}

	switch fromVal.Kind() {
	case reflect.Map:
		result, err := sliceFromMap(to, fromVal, uniqueVals, elemOps)
		if err != nil {
			return defaultValue, err
		}
		return result, nil
	case reflect.Struct:
		if isNamedScalarStructType(fromVal.Type()) {
			// Well-known named types (time.Time, big.Int, url.URL, etc.) have
			// meaningful scalar representations registered in namedConverters.
			// Scalar-wrap them directly rather than iterating their (mostly
			// unexported) fields.
			r, err := sliceFromSingle(fromVal.Interface(), to, elemOps, nil)
			if err != nil {
				return defaultValue, err
			}
			if uniqueVals {
				rv := reflect.ValueOf(r)
				r = dedupeSliceVal(rv).Interface()
			}
			return r, nil
		}
		// General struct: iterate exported fields first. If no fields convert
		// to the target type, fall back to scalar wrap (Stringer, then JSON).
		result, err := sliceFromStruct(to, fromVal, uniqueVals, elemOps)
		if err != nil {
			return defaultValue, err
		}
		if reflect.ValueOf(result).Len() == 0 {
			if r, sErr := sliceFromSingle(fromVal.Interface(), to, elemOps, nil); sErr == nil {
				return r, nil
			}
		}
		return result, nil
	case reflect.Slice, reflect.Array:
		// fall through to element-wise conversion below
	default:
		// Scalar or other non-iterable kind: wrap as single-element slice.
		return sliceFromSingle(fromVal.Interface(), to, elemOps, defaultValue)
	}

	slice := fromVal

	// Initialize the result slice based on target element type.
	var result any
	switch to.Interface().(type) {
	case []interface{}:
		result = make([]any, 0, size)
	case []bool:
		result = make([]bool, 0, size)
	case []complex64:
		result = make([]complex64, 0, size)
	case []complex128:
		result = make([]complex128, 0, size)
	case []float32:
		result = make([]float32, 0, size)
	case []float64:
		result = make([]float64, 0, size)
	case []int:
		result = make([]int, 0, size)
	case []int8:
		result = make([]int8, 0, size)
	case []int16:
		result = make([]int16, 0, size)
	case []int32:
		result = make([]int32, 0, size)
	case []int64:
		result = make([]int64, 0, size)
	case []uint:
		result = make([]uint, 0, size)
	case []uint8:
		result = make([]uint8, 0, size)
	case []uint16:
		result = make([]uint16, 0, size)
	case []uint32:
		result = make([]uint32, 0, size)
	case []uint64:
		result = make([]uint64, 0, size)
	case []uintptr:
		result = make([]uintptr, 0, size)
	case []string:
		result = make([]string, 0, size)
	default:
		// Named slice type (e.g. type MyInts []int): use reflection for slice
		// construction since the concrete type is unknown at compile time.
		elemType := to.Type().Elem()
		elemKind := elemType.Kind()
		scalarElem := isScalarKind(elemKind)

		// For named slice types whose element is an unnamed base scalar
		// (PkgPath == ""), redirect through the concrete type-switch path
		// above and convert the result. The Convert is O(1): named and
		// unnamed slices share the backing array.
		if scalarElem && elemType.PkgPath() == "" {
			baseTarget := reflect.Zero(reflect.SliceOf(elemType))
			raw, err := toSlice(baseTarget, val, ops.Delete(DEFAULT))
			if err != nil {
				return defaultValue, err
			}
			return reflect.ValueOf(raw).Convert(to.Type()).Interface(), nil
		}

		// Non-base-scalar element types (named scalars like type MyInt int,
		// structs, pointers, etc.): use the general reflection path.
		sliceVal := reflect.MakeSlice(to.Type(), 0, size)
		for a := 0; a < slice.Len(); a++ {
			elm := slice.Index(a).Interface()
			var elem reflect.Value
			var err error
			if scalarElem {
				elem, err = castToKind(elm, elemKind, elemOps)
				if err == nil && elem.Type() != elemType {
					elem = elem.Convert(elemType)
				}
			} else {
				elem, err = castToType(elm, elemType, elemOps)
			}
			if err != nil {
				return defaultValue, err
			}
			sliceVal = reflect.Append(sliceVal, elem)
		}
		if uniqueVals {
			sliceVal = dedupeSliceVal(sliceVal)
		}
		return sliceVal.Interface(), nil
	}

	for a := 0; a < slice.Len(); a++ {
		elm := slice.Index(a).Interface()
		switch r := result.(type) {
		case []any:
			result = append(r, elm)
		case []bool:
			tval, err := toBool[bool](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []complex64:
			tval, err := toComplex[complex64](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []complex128:
			tval, err := toComplex[complex128](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []float32:
			tval, err := toFloat[float32](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []float64:
			tval, err := toFloat[float64](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int:
			tval, err := toInt[int](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int8:
			tval, err := toInt[int8](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int16:
			tval, err := toInt[int16](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int32:
			tval, err := toInt[int32](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int64:
			tval, err := toInt[int64](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint:
			tval, err := toInt[uint](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint8:
			tval, err := toInt[uint8](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint16:
			tval, err := toInt[uint16](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint32:
			tval, err := toInt[uint32](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint64:
			tval, err := toInt[uint64](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uintptr:
			tval, err := toInt[uintptr](elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []string:
			tval, err := toString(elm, elemOps)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		}
	}

	if uniqueVals {
		rv := reflect.ValueOf(result)
		result = dedupeSliceVal(rv).Interface()
	}

	return result, nil
}

// sliceFromSingle casts v to the element type of to and returns a one-element
// slice. Returns (defaultValue, err) on cast failure.
func sliceFromSingle(v any, to reflect.Value, ops ops, defaultValue any) (any, error) {
	elemType := to.Type().Elem()
	elemKind := elemType.Kind()

	var elem reflect.Value
	var err error
	if isScalarKind(elemKind) {
		elem, err = castToKind(v, elemKind, ops)
		if err == nil && elem.Type() != elemType {
			elem = elem.Convert(elemType)
		}
	} else {
		elem, err = castToType(v, elemType, ops)
	}
	if err != nil {
		return defaultValue, err
	}

	result := reflect.MakeSlice(to.Type(), 0, 1)
	result = reflect.Append(result, elem)
	return result.Interface(), nil
}

// sliceFromMap builds a slice from a map's values, casting each to the
// target element type.
func sliceFromMap(to reflect.Value, src reflect.Value, uniqueVals bool, ops ops) (any, error) {
	elemType := to.Type().Elem()
	elemKind := elemType.Kind()
	scalarElem := isScalarKind(elemKind)
	sliceVal := reflect.MakeSlice(to.Type(), 0, src.Len())

	for _, key := range src.MapKeys() {
		mv := src.MapIndex(key)
		var rawVal any
		if mv.CanInterface() {
			rawVal = mv.Interface()
		} else {
			v, ok := extractFieldValue(mv)
			if !ok {
				continue
			}
			rawVal = v
		}

		var elem reflect.Value
		var err error
		if scalarElem {
			elem, err = castToKind(rawVal, elemKind, ops)
			if err == nil && elem.Type() != elemType {
				elem = elem.Convert(elemType)
			}
		} else {
			elem, err = castToType(rawVal, elemType, ops)
		}
		if err != nil {
			return nil, err
		}
		sliceVal = reflect.Append(sliceVal, elem)
	}

	if uniqueVals {
		sliceVal = dedupeSliceVal(sliceVal)
	}
	return sliceVal.Interface(), nil
}

// sliceFromStruct builds a slice from a struct's exported field values,
// casting each to the target element type. Unconvertible fields are skipped.
func sliceFromStruct(to reflect.Value, src reflect.Value, uniqueVals bool, ops ops) (any, error) {
	elemType := to.Type().Elem()
	elemKind := elemType.Kind()
	scalarElem := isScalarKind(elemKind)
	sliceVal := reflect.MakeSlice(to.Type(), 0, src.NumField())

	for i := 0; i < src.NumField(); i++ {
		field := src.Type().Field(i)
		if !field.IsExported() {
			continue
		}
		fieldVal := src.Field(i)
		rawVal, ok := extractFieldValue(fieldVal)
		if !ok {
			continue
		}

		var elem reflect.Value
		var err error
		if scalarElem {
			elem, err = castToKind(rawVal, elemKind, ops)
			if err == nil && elem.Type() != elemType {
				elem = elem.Convert(elemType)
			}
		} else {
			elem, err = castToType(rawVal, elemType, ops)
		}
		if err != nil {
			continue
		}
		sliceVal = reflect.Append(sliceVal, elem)
	}

	if uniqueVals {
		sliceVal = dedupeSliceVal(sliceVal)
	}
	return sliceVal.Interface(), nil
}

// dedupeSliceVal removes duplicate elements from rv, preserving first-seen
// order. Comparable elements are tracked in a map; non-comparable elements
// (e.g. slices) fall back to reflect.DeepEqual.
func dedupeSliceVal(rv reflect.Value) reflect.Value {
	deduped := reflect.MakeSlice(rv.Type(), 0, rv.Len())
	seen := make(map[any]struct{})
	var seenNonComparable []any

	for i := 0; i < rv.Len(); i++ {
		elem := rv.Index(i).Interface()
		concrete := reflect.ValueOf(elem)
		if !concrete.IsValid() || concrete.Type().Comparable() {
			if _, exists := seen[elem]; !exists {
				seen[elem] = struct{}{}
				deduped = reflect.Append(deduped, rv.Index(i))
			}
		} else {
			if !slices.ContainsFunc(seenNonComparable, func(prev any) bool {
				return reflect.DeepEqual(elem, prev)
			}) {
				seenNonComparable = append(seenNonComparable, elem)
				deduped = reflect.Append(deduped, rv.Index(i))
			}
		}
	}
	return deduped
}

