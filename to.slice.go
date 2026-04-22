package cast

import (
	"reflect"
	"slices"

	"github.com/bdlm/errors/v2"
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
			return defaultValue, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		defaultValue = ops.defaultVal
	}

	size := 1
	if ops.hasLength {
		var sizeErr error
		if size, sizeErr = ToE[int](ops.lengthVal); sizeErr != nil {
			return defaultValue, errors.Errorf(ErrorInvalidOption, "LENGTH", ops.lengthVal)
		}
	}
	if size < 0 {
		return defaultValue, errors.Errorf("invalid array length %d", size)
	}

	// Read local flags then strip them for element-level casts.
	uniqueVals := ops.uniqueVals
	elemOps := ops.Global()

	// Special cases mirroring Go's built-in string conversions:
	//   string → []byte / []uint8      (and named variants with Uint8 element)
	//   string → []rune / []int32      (and named variants with Int32 element)
	// These must be handled before the source-kind guard below because a string
	// is not a slice/array kind, and iterating bytes would give wrong results for
	// []rune (bytes ≠ Unicode code points for multibyte UTF-8).
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
			return defaultValue, errors.Errorf(ErrorStrUnableToCast, val, val, to.Interface())
		}
		if uniqueVals {
			rv := reflect.ValueOf(result)
			result = dedupeSliceVal(rv).Interface()
		}
		return result, nil
	}

	slice := reflect.ValueOf(val)
	if !slice.IsValid() || (slice.Kind() != reflect.Slice && slice.Kind() != reflect.Array) {
		return defaultValue, errors.Errorf(ErrorStrUnableToCast, val, val, to.Interface())
	}

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
		// For scalar element kinds, call castToKind directly — the same converter
		// the concrete switch uses above — to skip the castToType dispatch layer.
		// The scalar check is hoisted before the loop so it runs once, not per element.
		elemType := to.Type().Elem()
		elemKind := elemType.Kind()
		scalarElem := isScalarKind(elemKind)
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
			s, _ := tval.(string)
			result = append(r, s)
		}
	}

	if uniqueVals {
		rv := reflect.ValueOf(result)
		result = dedupeSliceVal(rv).Interface()
	}

	return result, nil
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

// isScalarKind reports whether k is a scalar kind handled by castToKind.
// reflect.Interface is excluded: castToType enforces assignability checks that
// castToKind skips, so interface element types must go through castToType.
func isScalarKind(k reflect.Kind) bool {
	switch k {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String:
		return true
	}
	return false
}
