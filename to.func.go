package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

//func ToFunc[TTo Func[Typ], Typ Types](to reflect.Kind, from reflect.Value) (TTo, error) {
//	return toFunc[TTo, Typ](to, from)
//}

// toFunc returns a function that casts an interface to the specified
// type and returns the result.
func toFunc[TTo any](to reflect.Value, from interface{}) (TTo, error) {
	var err error
	var ret TTo
	var reti interface{}
	var f interface{}

	switch to.Type().Out(0).Kind() {
	//case reflect.Array:
	//case reflect.Invalid:
	//case reflect.Map:
	//case reflect.Pointer:
	//case reflect.Struct:
	//case reflect.UnsafePointer:
	default:
		return ret, errors.Errorf("unable to cast %#v of type %T to %T", from, from, to.Interface())

	case reflect.Interface:
		f, err = ToE[interface{}](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[interface{}]
		fn = func() interface{} {
			return f.(interface{})
		}
		reti = fn
	case reflect.Bool:
		f, err = ToE[bool](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[bool]
		fn = func() bool {
			return f.(bool)
		}
		reti = fn
	case reflect.Complex64:
		f, err = ToE[complex64](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[complex64]
		fn = func() complex64 {
			return f.(complex64)
		}
		reti = fn
	case reflect.Complex128:
		f, err = ToE[complex128](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[complex128]
		fn = func() complex128 {
			return f.(complex128)
		}
		reti = fn
	case reflect.Float32:
		f, err = ToE[float32](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[float32]
		fn = func() float32 {
			return f.(float32)
		}
		reti = fn
	case reflect.Float64:
		f, err = ToE[float64](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[float64]
		fn = func() float64 {
			return f.(float64)
		}
		reti = fn
	case reflect.Int:
		f, err = ToE[int](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[int]
		fn = func() int {
			return f.(int)
		}
		reti = fn
	case reflect.Int8:
		f, err = ToE[int8](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[int8]
		fn = func() int8 {
			return f.(int8)
		}
		reti = fn
	case reflect.Int16:
		f, err = ToE[int16](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[int16]
		fn = func() int16 {
			return f.(int16)
		}
		reti = fn
	case reflect.Int32:
		f, err = ToE[int32](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[int32]
		fn = func() int32 {
			return f.(int32)
		}
		reti = fn
	case reflect.Int64:
		f, err = ToE[int64](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[int64]
		fn = func() int64 {
			return f.(int64)
		}
		reti = fn
	case reflect.Uint:
		f, err = ToE[uint](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[uint]
		fn = func() uint {
			return f.(uint)
		}
		reti = fn
	case reflect.Uint8:
		f, err = ToE[uint8](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[uint8]
		fn = func() uint8 {
			return f.(uint8)
		}
		reti = fn
	case reflect.Uint16:
		f, err = ToE[uint16](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[uint16]
		fn = func() uint16 {
			return f.(uint16)
		}
		reti = fn
	case reflect.Uint32:
		f, err = ToE[uint32](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[uint32]
		fn = func() uint32 {
			return f.(uint32)
		}
		reti = fn
	case reflect.Uint64:
		f, err = ToE[uint64](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[uint64]
		fn = func() uint64 {
			return f.(uint64)
		}
		reti = fn
	case reflect.Uintptr:
		f, err = ToE[uintptr](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[uintptr]
		fn = func() uintptr {
			return f.(uintptr)
		}
		reti = fn
	case reflect.String:
		f, err = ToE[string](from)
		if nil != err {
			return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
		}
		var fn Func[string]
		fn = func() string {
			return f.(string)
		}
		reti = fn

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Channels
	////////////////////////////////////////////////////////////////////////////////////////////////
	case reflect.Chan:
		switch to.Type().Out(0).Elem().Kind() {
		//case reflect.Array:
		//case reflect.Invalid:
		//case reflect.Map:
		//case reflect.Pointer:
		//case reflect.Struct:
		//case reflect.UnsafePointer:
		//case reflect.Slice:
		//case reflect.Func:
		default:
			return ret, errors.Errorf("unable to cast %#v of type %T to %T", from, from, to.Interface())
		case reflect.Interface:
			f, err = ToE[chan interface{}](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan interface{}]
			fn = func() chan interface{} {
				return f.(chan interface{})
			}
			reti = fn
		case reflect.Bool:
			f, err = ToE[chan bool](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan bool]
			fn = func() chan bool {
				return f.(chan bool)
			}
			reti = fn
		case reflect.Complex64:
			f, err = ToE[chan complex64](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan complex64]
			fn = func() chan complex64 {
				return f.(chan complex64)
			}
			reti = fn
		case reflect.Complex128:
			f, err = ToE[chan complex128](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan complex128]
			fn = func() chan complex128 {
				return f.(chan complex128)
			}
			reti = fn
		case reflect.Float32:
			f, err = ToE[chan float32](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan float32]
			fn = func() chan float32 {
				return f.(chan float32)
			}
			reti = fn
		case reflect.Float64:
			f, err = ToE[chan float64](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan float64]
			fn = func() chan float64 {
				return f.(chan float64)
			}
			reti = fn
		case reflect.Int:
			f, err = ToE[chan int](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan int]
			fn = func() chan int {
				return f.(chan int)
			}
			reti = fn
		case reflect.Int8:
			f, err = ToE[chan int8](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan int8]
			fn = func() chan int8 {
				return f.(chan int8)
			}
			reti = fn
		case reflect.Int16:
			f, err = ToE[chan int16](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan int16]
			fn = func() chan int16 {
				return f.(chan int16)
			}
			reti = fn
		case reflect.Int32:
			f, err = ToE[chan int32](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan int32]
			fn = func() chan int32 {
				return f.(chan int32)
			}
			reti = fn
		case reflect.Int64:
			f, err = ToE[chan int64](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan int64]
			fn = func() chan int64 {
				return f.(chan int64)
			}
			reti = fn
		case reflect.Uint:
			f, err = ToE[chan uint](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan uint]
			fn = func() chan uint {
				return f.(chan uint)
			}
			reti = fn
		case reflect.Uint8:
			f, err = ToE[chan uint8](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan uint8]
			fn = func() chan uint8 {
				return f.(chan uint8)
			}
			reti = fn
		case reflect.Uint16:
			f, err = ToE[chan uint16](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan uint16]
			fn = func() chan uint16 {
				return f.(chan uint16)
			}
			reti = fn
		case reflect.Uint32:
			f, err = ToE[chan uint32](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan uint32]
			fn = func() chan uint32 {
				return f.(chan uint32)
			}
			reti = fn
		case reflect.Uint64:
			f, err = ToE[chan uint64](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan uint64]
			fn = func() chan uint64 {
				return f.(chan uint64)
			}
			reti = fn
		case reflect.Uintptr:
			f, err = ToE[chan uintptr](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan uintptr]
			fn = func() chan uintptr {
				return f.(chan uintptr)
			}
			reti = fn
		case reflect.String:
			f, err = ToE[chan string](from)
			if nil != err {
				return ret, errors.Wrap(err, "error casting %T to %T during function generation", from, ret)
			}
			var fn Func[chan string]
			fn = func() chan string {
				return f.(chan string)
			}
			reti = fn
		}
		//	case reflect.Slice:
		//		reti = func() chan string {
		//			return f.(chan string)
		//		}
		//	case reflect.Func:
		//		reti = func() chan string {
		//			return f.(chan string)
		//		}
	}
	if ret, ok := reti.(TTo); ok {
		return ret, nil
	}
	return ret, errors.Errorf("error casting %T to %T during function generation", reti, ret)
}
