package cast

import (
	"reflect"

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
func toSlice(to reflect.Value, val any, ops Ops) (any, error) {
	var ret any
	var ok bool

	if _, ok = ops[DEFAULT]; ok {
		defaultVal := reflect.ValueOf(ops[DEFAULT])
		if defaultVal.IsValid() && !defaultVal.Type().AssignableTo(to.Type()) {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
		}
		ret = ops[DEFAULT]
		ops = ops.Delete(DEFAULT) // Prevent DEFAULT from being passed to element casts.
	}

	size := 1
	if _, ok = ops[LENGTH]; ok {
		size = To[int](ops[LENGTH])
	}
	if size < 0 {
		return ret, errors.Errorf("invalid array length %d", size)
	}

	fromKind := reflect.TypeOf(val).Kind()
	if fromKind != reflect.Slice && fromKind != reflect.Array {
		return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", val, val, to.Interface())
	}

	// Initialize the result slice based on target element type.
	switch to.Interface().(type) {
	default:
		return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", val, val, to.Interface())
	case []interface{}:
		ret = make([]any, 0, size)
	case []bool:
		ret = make([]bool, 0, size)
	case []complex64:
		ret = make([]complex64, 0, size)
	case []complex128:
		ret = make([]complex128, 0, size)
	case []float32:
		ret = make([]float32, 0, size)
	case []float64:
		ret = make([]float64, 0, size)
	case []int:
		ret = make([]int, 0, size)
	case []int8:
		ret = make([]int8, 0, size)
	case []int16:
		ret = make([]int16, 0, size)
	case []int32:
		ret = make([]int32, 0, size)
	case []int64:
		ret = make([]int64, 0, size)
	case []uint:
		ret = make([]uint, 0, size)
	case []uint8:
		ret = make([]uint8, 0, size)
	case []uint16:
		ret = make([]uint16, 0, size)
	case []uint32:
		ret = make([]uint32, 0, size)
	case []uint64:
		ret = make([]uint64, 0, size)
	case []uintptr:
		ret = make([]uintptr, 0, size)
	case []string:
		ret = make([]string, 0, size)
	}

	slice := reflect.ValueOf(val)
	for a := 0; a < slice.Len(); a++ {
		elm := slice.Index(a).Interface()
		switch r := ret.(type) {
		case []any:
			tval, err := ToE[any](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []bool:
			tval, err := ToE[bool](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []complex64:
			tval, err := ToE[complex64](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []complex128:
			tval, err := ToE[complex128](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []float32:
			tval, err := ToE[float32](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []float64:
			tval, err := ToE[float64](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []int:
			tval, err := ToE[int](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []int8:
			tval, err := ToE[int8](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []int16:
			tval, err := ToE[int16](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []int32:
			tval, err := ToE[int32](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []int64:
			tval, err := ToE[int64](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []uint:
			tval, err := ToE[uint](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []uint8:
			tval, err := ToE[uint8](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []uint16:
			tval, err := ToE[uint16](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []uint32:
			tval, err := ToE[uint32](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []uint64:
			tval, err := ToE[uint64](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []uintptr:
			tval, err := ToE[uintptr](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		case []string:
			tval, err := ToE[string](elm, ops.List()...)
			if err != nil {
				return ret, err
			}
			ret = append(r, tval)
		}
	}

	if unique, _ := ops[UNIQUE_VALUES].(bool); unique {
		rv := reflect.ValueOf(ret)
		deduped := reflect.MakeSlice(rv.Type(), 0, rv.Len())
		seen := make(map[any]struct{})
		var seenNonComparable []any

		for i := 0; i < rv.Len(); i++ {
			elem := rv.Index(i).Interface()
			// For interface elements, comparability depends on the concrete value.
			concrete := reflect.ValueOf(elem)
			if !concrete.IsValid() || concrete.Type().Comparable() {
				if _, exists := seen[elem]; !exists {
					seen[elem] = struct{}{}
					deduped = reflect.Append(deduped, rv.Index(i))
				}
			} else {
				found := false
				for _, prev := range seenNonComparable {
					if reflect.DeepEqual(elem, prev) {
						found = true
						break
					}
				}
				if !found {
					seenNonComparable = append(seenNonComparable, elem)
					deduped = reflect.Append(deduped, rv.Index(i))
				}
			}
		}
		ret = deduped.Interface()
	}

	return ret, nil
}
