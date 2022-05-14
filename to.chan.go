package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toChan returns a channel of the specified reflect.Value type with a buffer of
// 1 containing the from value.
func toChan(to reflect.Value, from any) (any, error) {
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
		ret = make(chan interface{}, 1)
		ret.(chan interface{}) <- To[interface{}](from)

	case reflect.Bool:
		ret = make(chan bool, 1)
		ret.(chan bool) <- To[bool](from)
	case reflect.Complex64:
		ret = make(chan complex64, 1)
		ret.(chan complex64) <- To[complex64](from)
	case reflect.Complex128:
		ret = make(chan complex128, 1)
		ret.(chan complex128) <- To[complex128](from)
	case reflect.Float32:
		ret = make(chan float32, 1)
		ret.(chan float32) <- To[float32](from)
	case reflect.Float64:
		ret = make(chan float64, 1)
		ret.(chan float64) <- To[float64](from)
	case reflect.Int:
		ret = make(chan int, 1)
		ret.(chan int) <- To[int](from)
	case reflect.Int8:
		ret = make(chan int8, 1)
		ret.(chan int8) <- To[int8](from)
	case reflect.Int16:
		ret = make(chan int16, 1)
		ret.(chan int16) <- To[int16](from)
	case reflect.Int32:
		ret = make(chan int32, 1)
		ret.(chan int32) <- To[int32](from)
	case reflect.Int64:
		ret = make(chan int64, 1)
		ret.(chan int64) <- To[int64](from)
	case reflect.Uint:
		ret = make(chan uint, 1)
		ret.(chan uint) <- To[uint](from)
	case reflect.Uint8:
		ret = make(chan uint8, 1)
		ret.(chan uint8) <- To[uint8](from)
	case reflect.Uint16:
		ret = make(chan uint16, 1)
		ret.(chan uint16) <- To[uint16](from)
	case reflect.Uint32:
		ret = make(chan uint32, 1)
		ret.(chan uint32) <- To[uint32](from)
	case reflect.Uint64:
		ret = make(chan uint64, 1)
		ret.(chan uint64) <- To[uint64](from)
	case reflect.Uintptr:
		ret = make(chan uintptr, 1)
		ret.(chan uintptr) <- To[uintptr](from)
	case reflect.String:
		ret = make(chan string, 1)
		ret.(chan string) <- To[string](from)

	}
	return ret, nil
}

// ToChan casts an interface to a bool type.
func _toChan[TTo Tchan](from any) (TTo, error) {
	var reti any
	var ret TTo
	var ok bool

	to := reflect.Indirect(reflect.ValueOf(new(TTo)))
	switch to.Type().Elem().Kind() {
	//case reflect.Invalid:
	//case reflect.Array:
	//case reflect.Func:
	//case reflect.Chan:
	//case reflect.Struct:
	//case reflect.UnsafePointer:
	//case reflect.Map:
	//case reflect.Pointer:
	//case reflect.Slice:
	default:
		return ret, errors.Errorf("unable to cast %#v of type %T to %T", from, from, to.Interface())

	case reflect.Interface:
		reti = make(chan interface{}, 1)
		reti.(chan interface{}) <- To[interface{}](from)

	case reflect.Bool:
		reti = make(chan bool, 1)
		reti.(chan bool) <- To[bool](from)
	case reflect.Complex64:
		reti = make(chan complex64, 1)
		reti.(chan complex64) <- To[complex64](from)
	case reflect.Complex128:
		reti = make(chan complex128, 1)
		reti.(chan complex128) <- To[complex128](from)
	case reflect.Float32:
		reti = make(chan float32, 1)
		reti.(chan float32) <- To[float32](from)
	case reflect.Float64:
		reti = make(chan float64, 1)
		reti.(chan float64) <- To[float64](from)
	case reflect.Int:
		reti = make(chan int, 1)
		reti.(chan int) <- To[int](from)
	case reflect.Int8:
		reti = make(chan int8, 1)
		reti.(chan int8) <- To[int8](from)
	case reflect.Int16:
		reti = make(chan int16, 1)
		reti.(chan int16) <- To[int16](from)
	case reflect.Int32:
		reti = make(chan int32, 1)
		reti.(chan int32) <- To[int32](from)
	case reflect.Int64:
		reti = make(chan int64, 1)
		reti.(chan int64) <- To[int64](from)
	case reflect.Uint:
		reti = make(chan uint, 1)
		reti.(chan uint) <- To[uint](from)
	case reflect.Uint8:
		reti = make(chan uint8, 1)
		reti.(chan uint8) <- To[uint8](from)
	case reflect.Uint16:
		reti = make(chan uint16, 1)
		reti.(chan uint16) <- To[uint16](from)
	case reflect.Uint32:
		reti = make(chan uint32, 1)
		reti.(chan uint32) <- To[uint32](from)
	case reflect.Uint64:
		reti = make(chan uint64, 1)
		reti.(chan uint64) <- To[uint64](from)
	case reflect.Uintptr:
		reti = make(chan uintptr, 1)
		reti.(chan uintptr) <- To[uintptr](from)
	case reflect.String:
		reti = make(chan string, 1)
		reti.(chan string) <- To[string](from)

	}

	if ret, ok = reti.(TTo); ok {
		return ret, nil
	}
	return ret, errors.Errorf("unable to cast %#v of type %T to %T", from, from, to.Interface())
}
