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
		ops = ops.Delete(DEFAULT) // Prevent DEFAULT from being passed to element casts.
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

	slice := reflect.ValueOf(val)

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
		// Named slice type: use reflection.
		sliceVal := reflect.MakeSlice(to.Type(), 0, size)
		for a := 0; a < slice.Len(); a++ {
			elm := slice.Index(a).Interface()
			elem, err := castToType(elm, to.Type().Elem(), ops)
			if err != nil {
				return defaultValue, err
			}
			sliceVal = reflect.Append(sliceVal, elem)
		}
		if ops.uniqueVals {
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
			tval, err := toBool[bool](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []complex64:
			tval, err := toComplex[complex64](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []complex128:
			tval, err := toComplex[complex128](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []float32:
			tval, err := toFloat[float32](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []float64:
			tval, err := toFloat[float64](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int:
			tval, err := toInt[int](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int8:
			tval, err := toInt[int8](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int16:
			tval, err := toInt[int16](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int32:
			tval, err := toInt[int32](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []int64:
			tval, err := toInt[int64](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint:
			tval, err := toInt[uint](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint8:
			tval, err := toInt[uint8](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint16:
			tval, err := toInt[uint16](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint32:
			tval, err := toInt[uint32](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uint64:
			tval, err := toInt[uint64](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []uintptr:
			tval, err := toInt[uintptr](elm, ops)
			if err != nil {
				return defaultValue, err
			}
			result = append(r, tval)
		case []string:
			tval, err := toString(elm, ops)
			if err != nil {
				return defaultValue, err
			}
			s, _ := tval.(string)
			result = append(r, s)
		}
	}

	if ops.uniqueVals {
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
