package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toChan returns a channel of the specified reflect.Value type with a buffer of
// 1 containing the from value.
func toSlice(to reflect.Value, from any, size int) (any, error) {
	var ret any

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
		tval, err := ToE[interface{}](from)
		if err != nil {
			return ret, err
		}
		slice := make([]interface{}, size)
		slice[0] = tval
		ret = slice
	case reflect.Bool:
		tval, err := ToE[bool](from)
		if err != nil {
			return ret, err
		}
		slice := make([]bool, size)
		slice[0] = tval
		ret = slice
	case reflect.Complex64:
		tval, err := ToE[complex64](from)
		if err != nil {
			return ret, err
		}
		slice := make([]complex64, size)
		slice[0] = tval
		ret = slice
	case reflect.Complex128:
		tval, err := ToE[complex128](from)
		if err != nil {
			return ret, err
		}
		slice := make([]complex128, size)
		slice[0] = tval
		ret = slice
	case reflect.Float32:
		tval, err := ToE[float32](from)
		if err != nil {
			return ret, err
		}
		slice := make([]float32, size)
		slice[0] = tval
		ret = slice
	case reflect.Float64:
		tval, err := ToE[float64](from)
		if err != nil {
			return ret, err
		}
		slice := make([]float64, size)
		slice[0] = tval
		ret = slice
	case reflect.Int:
		tval, err := ToE[int](from)
		if err != nil {
			return ret, err
		}
		slice := make([]int, size)
		slice[0] = tval
		ret = slice
	case reflect.Int8:
		tval, err := ToE[int8](from)
		if err != nil {
			return ret, err
		}
		slice := make([]int8, size)
		slice[0] = tval
		ret = slice
	case reflect.Int16:
		tval, err := ToE[int16](from)
		if err != nil {
			return ret, err
		}
		slice := make([]int16, size)
		slice[0] = tval
		ret = slice
	case reflect.Int32:
		tval, err := ToE[int32](from)
		if err != nil {
			return ret, err
		}
		slice := make([]int32, size)
		slice[0] = tval
		ret = slice
	case reflect.Int64:
		tval, err := ToE[int64](from)
		if err != nil {
			return ret, err
		}
		slice := make([]int64, size)
		slice[0] = tval
		ret = slice
	case reflect.Uint:
		tval, err := ToE[uint](from)
		if err != nil {
			return ret, err
		}
		slice := make([]uint, size)
		slice[0] = tval
		ret = slice
	case reflect.Uint8:
		tval, err := ToE[uint8](from)
		if err != nil {
			return ret, err
		}
		slice := make([]uint8, size)
		slice[0] = tval
		ret = slice
	case reflect.Uint16:
		tval, err := ToE[uint16](from)
		if err != nil {
			return ret, err
		}
		slice := make([]uint16, size)
		slice[0] = tval
		ret = slice
	case reflect.Uint32:
		tval, err := ToE[uint32](from)
		if err != nil {
			return ret, err
		}
		slice := make([]uint32, size)
		slice[0] = tval
		ret = slice
	case reflect.Uint64:
		tval, err := ToE[uint64](from)
		if err != nil {
			return ret, err
		}
		slice := make([]uint64, size)
		slice[0] = tval
		ret = slice
	case reflect.Uintptr:
		tval, err := ToE[uintptr](from)
		if err != nil {
			return ret, err
		}
		slice := make([]uintptr, size)
		slice[0] = tval
		ret = slice
	case reflect.String:
		tval, err := ToE[string](from)
		if err != nil {
			return ret, err
		}
		slice := make([]string, size)
		slice[0] = tval
		ret = slice

	}

	return ret, nil
}
