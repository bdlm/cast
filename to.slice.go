package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toChan returns a channel of the specified reflect.Value type with a buffer of
// 1 containing the from value.
func toSlice(to reflect.Value, from any) (any, error) {
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
		return ret, errors.Errorf("unable to cast %#v of type %T to %T", from, from, to.Interface())

	case reflect.Interface:
		tval, err := ToE[interface{}](from)
		if err != nil {
			return ret, err
		}
		ret = []interface{}{}
		ret = append(ret.([]interface{}), tval)
	case reflect.Bool:
		tval, err := ToE[bool](from)
		if err != nil {
			return ret, err
		}
		ret = []bool{}
		ret = append(ret.([]bool), tval)
	case reflect.Complex64:
		tval, err := ToE[complex64](from)
		if err != nil {
			return ret, err
		}
		ret = []complex64{}
		ret = append(ret.([]complex64), tval)
	case reflect.Complex128:
		tval, err := ToE[complex128](from)
		if err != nil {
			return ret, err
		}
		ret = []complex128{}
		ret = append(ret.([]complex128), tval)
	case reflect.Float32:
		tval, err := ToE[float32](from)
		if err != nil {
			return ret, err
		}
		ret = []float32{}
		ret = append(ret.([]float32), tval)
	case reflect.Float64:
		tval, err := ToE[float64](from)
		if err != nil {
			return ret, err
		}
		ret = []float64{}
		ret = append(ret.([]float64), tval)
	case reflect.Int:
		tval, err := ToE[int](from)
		if err != nil {
			return ret, err
		}
		ret = []int{}
		ret = append(ret.([]int), tval)
	case reflect.Int8:
		tval, err := ToE[int8](from)
		if err != nil {
			return ret, err
		}
		ret = []int8{}
		ret = append(ret.([]int8), tval)
	case reflect.Int16:
		tval, err := ToE[int16](from)
		if err != nil {
			return ret, err
		}
		ret = []int16{}
		ret = append(ret.([]int16), tval)
	case reflect.Int32:
		tval, err := ToE[int32](from)
		if err != nil {
			return ret, err
		}
		ret = []int32{}
		ret = append(ret.([]int32), tval)
	case reflect.Int64:
		tval, err := ToE[int64](from)
		if err != nil {
			return ret, err
		}
		ret = []int64{}
		ret = append(ret.([]int64), tval)
	case reflect.Uint:
		tval, err := ToE[uint](from)
		if err != nil {
			return ret, err
		}
		ret = []uint{}
		ret = append(ret.([]uint), tval)
	case reflect.Uint8:
		tval, err := ToE[uint8](from)
		if err != nil {
			return ret, err
		}
		ret = []uint8{}
		ret = append(ret.([]uint8), tval)
	case reflect.Uint16:
		tval, err := ToE[uint16](from)
		if err != nil {
			return ret, err
		}
		ret = []uint16{}
		ret = append(ret.([]uint16), tval)
	case reflect.Uint32:
		tval, err := ToE[uint32](from)
		if err != nil {
			return ret, err
		}
		ret = []uint32{}
		ret = append(ret.([]uint32), tval)
	case reflect.Uint64:
		tval, err := ToE[uint64](from)
		if err != nil {
			return ret, err
		}
		ret = []uint64{}
		ret = append(ret.([]uint64), tval)
	case reflect.Uintptr:
		tval, err := ToE[uintptr](from)
		if err != nil {
			return ret, err
		}
		ret = []uintptr{}
		ret = append(ret.([]uintptr), tval)
	case reflect.String:
		tval, err := ToE[string](from)
		if err != nil {
			return ret, err
		}
		ret = []string{}
		ret = append(ret.([]string), tval)

	}

	return ret, nil
}
