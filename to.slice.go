package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toChan returns a channel of the specified reflect.Value type with a buffer of
// 1 containing the from value.
func toSlice(to reflect.Value, val any, ops Ops) (any, error) {
	//toType := reflect.New(to.Elem().Type()).Interface()
	var ret any

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
				ret = []any{}
			}
			tval, err := ToE[any](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]any), tval)
		case []bool:
			if nil == ret {
				ret = []bool{}
			}
			tval, err := ToE[bool](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]bool), tval)
		case []complex64:
			if nil == ret {
				ret = []complex64{}
			}
			tval, err := ToE[complex64](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]complex64), tval)
		case []complex128:
			if nil == ret {
				ret = []complex128{}
			}
			tval, err := ToE[complex128](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]complex128), tval)
		case []float32:
			if nil == ret {
				ret = []float32{}
			}
			tval, err := ToE[float32](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]float32), tval)

		case []float64:
			if nil == ret {
				ret = []float64{}
			}
			tval, err := ToE[float64](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]float64), tval)
		case []int:
			if nil == ret {
				ret = []int{}
			}
			tval, err := ToE[int](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int), tval)
		case []int8:
			if nil == ret {
				ret = []int8{}
			}
			tval, err := ToE[int8](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int8), tval)
		case []int16:
			if nil == ret {
				ret = []int16{}
			}
			tval, err := ToE[int16](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int16), tval)
		case []int32:
			if nil == ret {
				ret = []int32{}
			}
			tval, err := ToE[int32](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int32), tval)
		case []int64:
			if nil == ret {
				ret = []int64{}
			}
			tval, err := ToE[int64](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]int64), tval)
		case []uint:
			if nil == ret {
				ret = []uint{}
			}
			tval, err := ToE[uint](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint), tval)
		case []uint8:
			if nil == ret {
				ret = []uint8{}
			}
			tval, err := ToE[uint8](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint8), tval)
		case []uint16:
			if nil == ret {
				ret = []uint16{}
			}
			tval, err := ToE[uint16](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint16), tval)
		case []uint32:
			if nil == ret {
				ret = []uint32{}
			}
			tval, err := ToE[uint32](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint32), tval)
		case []uint64:
			if nil == ret {
				ret = []uint64{}
			}
			tval, err := ToE[uint64](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uint64), tval)
		case []uintptr:
			if nil == ret {
				ret = []uintptr{}
			}
			tval, err := ToE[uintptr](elm)
			if err != nil {
				return ret, err
			}
			ret = append(ret.([]uintptr), tval)
		case []string:
			if nil == ret {
				ret = []string{}
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
