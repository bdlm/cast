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
//   - LENGTH: int, size of the backing array length, default 1. Must be greater
//     than 0.
func toSlice(to reflect.Value, val any, ops Ops) (any, error) {
	var ret any
	var ok bool

	if _, ok = ops[DEFAULT]; ok {
		ret = ops[DEFAULT]
	}

	size := 1
	if _, ok = ops[LENGTH]; ok {
		size = To[int](ops[LENGTH])
	}
	if size < 0 {
		return ret, errors.Errorf("invalid array length %d", size)
	}

	if reflect.TypeOf(val).Kind() != reflect.Slice {
		return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", val, val, to.Interface())
	}

	slice := reflect.ValueOf(val)
	for a := 0; a < slice.Len(); a++ {
		elm := slice.Index(a).Interface()
		switch to.Interface().(type) {
		//case reflect.Invalid:
		//case reflect.Array:
		//case reflect.Func:
		//case reflect.Chan:
		//	return toChan(to.Elem(), from)
		//case reflect.Slice:
		//case reflect.Struct:
		//case reflect.UnsafePointer:
		//case reflect.Map:
		//case reflect.Pointer:
		default:
			return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", val, val, to.Interface())

		case []interface{}:
			if nil == ret {
				ret = make([]any, size)
			}
			tval, err := ToE[any](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]any), tval)
		case []bool:
			if nil == ret {
				ret = make([]bool, size)
			}
			tval, err := ToE[bool](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]bool), tval)
		case []complex64:
			if nil == ret {
				ret = make([]complex64, size)
			}
			tval, err := ToE[complex64](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]complex64), tval)
		case []complex128:
			if nil == ret {
				ret = make([]complex128, size)
			}
			tval, err := ToE[complex128](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]complex128), tval)
		case []float32:
			if nil == ret {
				ret = make([]float32, size)
			}
			tval, err := ToE[float32](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]float32), tval)

		case []float64:
			if nil == ret {
				ret = make([]float64, size)
			}
			tval, err := ToE[float64](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]float64), tval)
		case []int:
			if nil == ret {
				ret = make([]int, size)
			}
			tval, err := ToE[int](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int), tval)
		case []int8:
			if nil == ret {
				ret = make([]int8, size)
			}
			tval, err := ToE[int8](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int8), tval)
		case []int16:
			if nil == ret {
				ret = make([]int16, size)
			}
			tval, err := ToE[int16](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int16), tval)
		case []int32:
			if nil == ret {
				ret = make([]int32, size)
			}
			tval, err := ToE[int32](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int32), tval)
		case []int64:
			if nil == ret {
				ret = make([]int64, size)
			}
			tval, err := ToE[int64](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int64), tval)
		case []uint:
			if nil == ret {
				ret = make([]uint, size)
			}
			tval, err := ToE[uint](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint), tval)
		case []uint8:
			if nil == ret {
				ret = make([]uint8, size)
			}
			tval, err := ToE[uint8](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint8), tval)
		case []uint16:
			if nil == ret {
				ret = make([]uint16, size)
			}
			tval, err := ToE[uint16](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint16), tval)
		case []uint32:
			if nil == ret {
				ret = make([]uint32, size)
			}
			tval, err := ToE[uint32](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint32), tval)
		case []uint64:
			if nil == ret {
				ret = make([]uint64, size)
			}
			tval, err := ToE[uint64](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint64), tval)
		case []uintptr:
			if nil == ret {
				ret = make([]uintptr, size)
			}
			tval, err := ToE[uintptr](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uintptr), tval)
		case []string:
			if nil == ret {
				ret = make([]string, size)
			}
			tval, err := ToE[string](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]string), tval)
		}
	}

	return ret, nil
}
