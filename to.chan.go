package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toChan returns a channel of the specified reflect.Value type with a buffer of
// 1 containing the from value.
func toChan(to reflect.Value, from any, size int) (any, error) {
	var ret any
	if size < 1 {
		size = 1
	}

	switch to.Type().Elem().Kind() {
	//case reflect.Invalid:
	//case reflect.Array:
	//case reflect.Func:
	//case reflect.Chan:
	//	return toChan(to.Elem(), from) <- todo: totally explodes, turtles all the way down
	//case reflect.Struct:
	//case reflect.UnsafePointer:
	//case reflect.Map:
	//case reflect.Pointer:
	//case reflect.Slice:
	default:
		return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())

	case reflect.Interface:
		ret = make(chan interface{}, size)
		result, err := ToE[interface{}](from)
		if nil != err {
			return nil, err
		}
		ret.(chan interface{}) <- result
	case reflect.Bool:
		ret = make(chan bool, size)
		result, err := ToE[bool](from)
		if nil != err {
			return nil, err
		}
		ret.(chan bool) <- result
	case reflect.Complex64:
		ret = make(chan complex64, size)
		result, err := ToE[complex64](from)
		if nil != err {
			return nil, err
		}
		ret.(chan complex64) <- result
	case reflect.Complex128:
		ret = make(chan complex128, size)
		result, err := ToE[complex128](from)
		if nil != err {
			return nil, err
		}
		ret.(chan complex128) <- result
	case reflect.Float32:
		ret = make(chan float32, size)
		result, err := ToE[float32](from)
		if nil != err {
			return nil, err
		}
		ret.(chan float32) <- result
	case reflect.Float64:
		ret = make(chan float64, size)
		result, err := ToE[float64](from)
		if nil != err {
			return nil, err
		}
		ret.(chan float64) <- result
	case reflect.Int:
		ret = make(chan int, size)
		result, err := ToE[int](from)
		if nil != err {
			return nil, err
		}
		ret.(chan int) <- result
	case reflect.Int8:
		ret = make(chan int8, size)
		result, err := ToE[int8](from)
		if nil != err {
			return nil, err
		}
		ret.(chan int8) <- result
	case reflect.Int16:
		ret = make(chan int16, size)
		result, err := ToE[int16](from)
		if nil != err {
			return nil, err
		}
		ret.(chan int16) <- result
	case reflect.Int32:
		ret = make(chan int32, size)
		result, err := ToE[int32](from)
		if nil != err {
			return nil, err
		}
		ret.(chan int32) <- result
	case reflect.Int64:
		ret = make(chan int64, size)
		result, err := ToE[int64](from)
		if nil != err {
			return nil, err
		}
		ret.(chan int64) <- result
	case reflect.Uint:
		ret = make(chan uint, size)
		result, err := ToE[uint](from)
		if nil != err {
			return nil, err
		}
		ret.(chan uint) <- result
	case reflect.Uint8:
		ret = make(chan uint8, size)
		result, err := ToE[uint8](from)
		if nil != err {
			return nil, err
		}
		ret.(chan uint8) <- result
	case reflect.Uint16:
		ret = make(chan uint16, size)
		result, err := ToE[uint16](from)
		if nil != err {
			return nil, err
		}
		ret.(chan uint16) <- result
	case reflect.Uint32:
		ret = make(chan uint32, size)
		result, err := ToE[uint32](from)
		if nil != err {
			return nil, err
		}
		ret.(chan uint32) <- result
	case reflect.Uint64:
		ret = make(chan uint64, size)
		result, err := ToE[uint64](from)
		if nil != err {
			return nil, err
		}
		ret.(chan uint64) <- result
	case reflect.Uintptr:
		ret = make(chan uintptr, size)
		result, err := ToE[uintptr](from)
		if nil != err {
			return nil, err
		}
		ret.(chan uintptr) <- result
	case reflect.String:
		ret = make(chan string, size)
		result, err := ToE[string](from)
		if nil != err {
			return nil, err
		}
		ret.(chan string) <- result

	}
	return ret, nil
}
