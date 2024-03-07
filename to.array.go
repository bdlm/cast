package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toArray returns a slice containing the specified reflect.Value type
// containing the from value.
//
// Options:
//   - DEFAULT: slice, default return value on error.
//   - LENGTH: int, size of the backing array length, default 1. Must be greater
//     than 0.
func toArray(to reflect.Value, from any, ops Ops) (any, error) {
	var ret any
	var ok bool

	if _, ok = ops[DEFAULT]; ok {
		ret = ops[DEFAULT]
	}

	size := 1
	if _, ok = ops[LENGTH]; ok {
		size = To[int](ops[LENGTH])
	}
	if !ok || size < 0 {
		return ret, errors.Errorf("invalid channel buffer size %d", size)
	}

	switch to.Type().Elem().Kind() {
	//case reflect.Invalid:
	//case reflect.Array:
	//case reflect.Func:
	//case reflect.Chan:
	//	return toChan(to.Elem(), from)
	//case reflect.Struct:
	//case reflect.UnsafePointer:
	//case reflect.Map:
	//case reflect.Pointer:
	//case reflect.Slice:
	default:
		return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())

	case reflect.Interface:
		//		v := to.Interface()
		//		tval, err := ToE[interface{}](from)
		//		v = append(v, tval)
		//		if err != nil {
		//			return ret, err
		//		}
		//		ret = make([]interface{}, 1)
		//		ret = append(ret.([]interface{}), tval)
	case reflect.Bool:
		tval, err := ToE[bool](from)
		if err != nil {
			return ret, err
		}
		ret = make([]bool, size)
		ret = append(ret.([]bool), tval)
	case reflect.Complex64:
		tval, err := ToE[complex64](from)
		if err != nil {
			return ret, err
		}
		ret = make([]complex64, size)
		ret = append(ret.([]complex64), tval)
	case reflect.Complex128:
		tval, err := ToE[complex128](from)
		if err != nil {
			return ret, err
		}
		ret = make([]complex128, size)
		ret = append(ret.([]complex128), tval)
	case reflect.Float32:
		tval, err := ToE[float32](from)
		if err != nil {
			return ret, err
		}
		ret = make([]float32, size)
		ret = append(ret.([]float32), tval)
	case reflect.Float64:
		tval, err := ToE[float64](from)
		if err != nil {
			return ret, err
		}
		ret = make([]float64, size)
		ret = append(ret.([]float64), tval)
	case reflect.Int:
		tval, err := ToE[int](from)
		if err != nil {
			return ret, err
		}
		ret = make([]int, size)
		ret = append(ret.([]int), tval)
	case reflect.Int8:
		tval, err := ToE[int8](from)
		if err != nil {
			return ret, err
		}
		ret = make([]int8, size)
		ret = append(ret.([]int8), tval)
	case reflect.Int16:
		tval, err := ToE[int16](from)
		if err != nil {
			return ret, err
		}
		ret = make([]int16, size)
		ret = append(ret.([]int16), tval)
	case reflect.Int32:
		tval, err := ToE[int32](from)
		if err != nil {
			return ret, err
		}
		ret = make([]int32, size)
		ret = append(ret.([]int32), tval)
	case reflect.Int64:
		tval, err := ToE[int64](from)
		if err != nil {
			return ret, err
		}
		ret = make([]int64, size)
		ret = append(ret.([]int64), tval)
	case reflect.Uint:
		tval, err := ToE[uint](from)
		if err != nil {
			return ret, err
		}
		ret = make([]uint, size)
		ret = append(ret.([]uint), tval)
	case reflect.Uint8:
		tval, err := ToE[uint8](from)
		if err != nil {
			return ret, err
		}
		ret = make([]uint8, size)
		ret = append(ret.([]uint8), tval)
	case reflect.Uint16:
		tval, err := ToE[uint16](from)
		if err != nil {
			return ret, err
		}
		ret = make([]uint16, size)
		ret = append(ret.([]uint16), tval)
	case reflect.Uint32:
		tval, err := ToE[uint32](from)
		if err != nil {
			return ret, err
		}
		ret = make([]uint32, size)
		ret = append(ret.([]uint32), tval)
	case reflect.Uint64:
		tval, err := ToE[uint64](from)
		if err != nil {
			return ret, err
		}
		ret = make([]uint64, size)
		ret = append(ret.([]uint64), tval)
	case reflect.Uintptr:
		tval, err := ToE[uintptr](from)
		if err != nil {
			return ret, err
		}
		ret = make([]uintptr, size)
		ret = append(ret.([]uintptr), tval)
	case reflect.String:
		tval, err := ToE[string](from)
		if err != nil {
			return ret, err
		}
		ret = make([]string, size)
		ret = append(ret.([]string), tval)

	}

	return ret, nil
}
