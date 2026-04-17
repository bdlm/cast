package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toChan returns a channel of the specified reflect.Value type with a buffer of
// LENGTH containing the from value.
//
// Options:
//   - DEFAULT: channel, default return value on error.
//   - LENGTH: int, channel buffer size, default 1. Must be greater than or
//     equal to 0. A value of 0 creates an unbuffered channel.
func toChan(to reflect.Value, from any, ops Ops) (any, error) {
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
		return ret, errors.Errorf("invalid channel buffer size %d", size)
	}

	var err error
	switch to.Type().Elem().Kind() {
	//case reflect.Struct:
	//case reflect.UnsafePointer:
	//case reflect.Map:
	//case reflect.Pointer:
	default:
		return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())

	case reflect.Interface:
		ret, err = makeChan[interface{}](from, size)
	case reflect.Bool:
		ret, err = makeChan[bool](from, size)
	case reflect.Complex64:
		ret, err = makeChan[complex64](from, size)
	case reflect.Complex128:
		ret, err = makeChan[complex128](from, size)
	case reflect.Float32:
		ret, err = makeChan[float32](from, size)
	case reflect.Float64:
		ret, err = makeChan[float64](from, size)
	case reflect.Int:
		ret, err = makeChan[int](from, size)
	case reflect.Int8:
		ret, err = makeChan[int8](from, size)
	case reflect.Int16:
		ret, err = makeChan[int16](from, size)
	case reflect.Int32:
		ret, err = makeChan[int32](from, size)
	case reflect.Int64:
		ret, err = makeChan[int64](from, size)
	case reflect.Uint:
		ret, err = makeChan[uint](from, size)
	case reflect.Uint8:
		ret, err = makeChan[uint8](from, size)
	case reflect.Uint16:
		ret, err = makeChan[uint16](from, size)
	case reflect.Uint32:
		ret, err = makeChan[uint32](from, size)
	case reflect.Uint64:
		ret, err = makeChan[uint64](from, size)
	case reflect.Uintptr:
		ret, err = makeChan[uintptr](from, size)
	case reflect.String:
		ret, err = makeChan[string](from, size)

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Slices — chan []T
	////////////////////////////////////////////////////////////////////////////////////////////////
	case reflect.Slice:
		switch to.Type().Elem().Elem().Kind() {
		default:
			return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())
		case reflect.Interface:
			ret, err = makeChan[[]any](from, size)
		case reflect.Bool:
			ret, err = makeChan[[]bool](from, size)
		case reflect.Complex64:
			ret, err = makeChan[[]complex64](from, size)
		case reflect.Complex128:
			ret, err = makeChan[[]complex128](from, size)
		case reflect.Float32:
			ret, err = makeChan[[]float32](from, size)
		case reflect.Float64:
			ret, err = makeChan[[]float64](from, size)
		case reflect.Int:
			ret, err = makeChan[[]int](from, size)
		case reflect.Int8:
			ret, err = makeChan[[]int8](from, size)
		case reflect.Int16:
			ret, err = makeChan[[]int16](from, size)
		case reflect.Int32:
			ret, err = makeChan[[]int32](from, size)
		case reflect.Int64:
			ret, err = makeChan[[]int64](from, size)
		case reflect.Uint:
			ret, err = makeChan[[]uint](from, size)
		case reflect.Uint8:
			ret, err = makeChan[[]uint8](from, size)
		case reflect.Uint16:
			ret, err = makeChan[[]uint16](from, size)
		case reflect.Uint32:
			ret, err = makeChan[[]uint32](from, size)
		case reflect.Uint64:
			ret, err = makeChan[[]uint64](from, size)
		case reflect.Uintptr:
			ret, err = makeChan[[]uintptr](from, size)
		case reflect.String:
			ret, err = makeChan[[]string](from, size)
		}

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Funcs — chan Func[T]
	////////////////////////////////////////////////////////////////////////////////////////////////
	case reflect.Func:
		switch to.Type().Elem().Out(0).Kind() {
		default:
			return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())
		case reflect.Interface:
			ret, err = makeChan[Func[any]](from, size)
		case reflect.Bool:
			ret, err = makeChan[Func[bool]](from, size)
		case reflect.Complex64:
			ret, err = makeChan[Func[complex64]](from, size)
		case reflect.Complex128:
			ret, err = makeChan[Func[complex128]](from, size)
		case reflect.Float32:
			ret, err = makeChan[Func[float32]](from, size)
		case reflect.Float64:
			ret, err = makeChan[Func[float64]](from, size)
		case reflect.Int:
			ret, err = makeChan[Func[int]](from, size)
		case reflect.Int8:
			ret, err = makeChan[Func[int8]](from, size)
		case reflect.Int16:
			ret, err = makeChan[Func[int16]](from, size)
		case reflect.Int32:
			ret, err = makeChan[Func[int32]](from, size)
		case reflect.Int64:
			ret, err = makeChan[Func[int64]](from, size)
		case reflect.Uint:
			ret, err = makeChan[Func[uint]](from, size)
		case reflect.Uint8:
			ret, err = makeChan[Func[uint8]](from, size)
		case reflect.Uint16:
			ret, err = makeChan[Func[uint16]](from, size)
		case reflect.Uint32:
			ret, err = makeChan[Func[uint32]](from, size)
		case reflect.Uint64:
			ret, err = makeChan[Func[uint64]](from, size)
		case reflect.Uintptr:
			ret, err = makeChan[Func[uintptr]](from, size)
		case reflect.String:
			ret, err = makeChan[Func[string]](from, size)
		}

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Channels of channels — chan chan T
	////////////////////////////////////////////////////////////////////////////////////////////////
	case reflect.Chan:
		switch to.Type().Elem().Elem().Kind() {
		default:
			return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())
		case reflect.Interface:
			ret, err = makeChan[chan any](from, size)
		case reflect.Bool:
			ret, err = makeChan[chan bool](from, size)
		case reflect.Complex64:
			ret, err = makeChan[chan complex64](from, size)
		case reflect.Complex128:
			ret, err = makeChan[chan complex128](from, size)
		case reflect.Float32:
			ret, err = makeChan[chan float32](from, size)
		case reflect.Float64:
			ret, err = makeChan[chan float64](from, size)
		case reflect.Int:
			ret, err = makeChan[chan int](from, size)
		case reflect.Int8:
			ret, err = makeChan[chan int8](from, size)
		case reflect.Int16:
			ret, err = makeChan[chan int16](from, size)
		case reflect.Int32:
			ret, err = makeChan[chan int32](from, size)
		case reflect.Int64:
			ret, err = makeChan[chan int64](from, size)
		case reflect.Uint:
			ret, err = makeChan[chan uint](from, size)
		case reflect.Uint8:
			ret, err = makeChan[chan uint8](from, size)
		case reflect.Uint16:
			ret, err = makeChan[chan uint16](from, size)
		case reflect.Uint32:
			ret, err = makeChan[chan uint32](from, size)
		case reflect.Uint64:
			ret, err = makeChan[chan uint64](from, size)
		case reflect.Uintptr:
			ret, err = makeChan[chan uintptr](from, size)
		case reflect.String:
			ret, err = makeChan[chan string](from, size)
		}

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Arrays — chan [N]T (reflection path, N is not known at compile time)
	////////////////////////////////////////////////////////////////////////////////////////////////
	case reflect.Array:
		ret, err = makeArrayChan(to.Type(), from, size)
	}
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// makeChan casts from to T, creates a buffered channel of size, sends the
// result, and returns the channel as any.
func makeChan[T Types](from any, size int) (any, error) {
	val, err := ToE[T](from)
	if err != nil {
		return nil, err
	}
	ch := make(chan T, size)
	ch <- val
	return ch, nil
}
