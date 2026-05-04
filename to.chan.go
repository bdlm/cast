package cast

import (
	"fmt"
	"reflect"
)

// toChan returns a channel of the specified reflect.Value type with a buffer of
// LENGTH containing the from value.
//
// Options:
//   - DEFAULT: channel, default return value on error.
//   - LENGTH: int, channel buffer size, default 1. Must be greater than or
//     equal to 1.
func toChan(to reflect.Value, from any, ops ops) (any, error) {
	var defaultValue any
	var returnValue any

	if ops.hasDefault {
		defaultVal := reflect.ValueOf(ops.defaultVal)
		if defaultVal.IsValid() && !defaultVal.Type().AssignableTo(to.Type()) {
			return defaultValue, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		defaultValue = ops.defaultVal
	}

	size := 1
	if ops.hasLength {
		var sizeErr error
		if size, sizeErr = ToE[int](ops.lengthVal); sizeErr != nil {
			return defaultValue, fmt.Errorf(ErrorInvalidOption, "LENGTH", ops.lengthVal)
		}
	}
	if size < 1 {
		return defaultValue, fmt.Errorf("invalid channel buffer size %d", size)
	}

	// Strip local flags before element-level casts.
	elemOps := ops.Global()

	var err error
	switch to.Type().Elem().Kind() {
	//case reflect.Struct:
	//case reflect.UnsafePointer:
	//case reflect.Map:
	//case reflect.Pointer:
	default:
		return defaultValue, fmt.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())

	case reflect.Interface:
		returnValue, err = makeChan[interface{}](from, size, elemOps)
	case reflect.Bool:
		returnValue, err = makeChan[bool](from, size, elemOps)
	case reflect.Complex64:
		returnValue, err = makeChan[complex64](from, size, elemOps)
	case reflect.Complex128:
		returnValue, err = makeChan[complex128](from, size, elemOps)
	case reflect.Float32:
		returnValue, err = makeChan[float32](from, size, elemOps)
	case reflect.Float64:
		returnValue, err = makeChan[float64](from, size, elemOps)
	case reflect.Int:
		returnValue, err = makeChan[int](from, size, elemOps)
	case reflect.Int8:
		returnValue, err = makeChan[int8](from, size, elemOps)
	case reflect.Int16:
		returnValue, err = makeChan[int16](from, size, elemOps)
	case reflect.Int32:
		returnValue, err = makeChan[int32](from, size, elemOps)
	case reflect.Int64:
		returnValue, err = makeChan[int64](from, size, elemOps)
	case reflect.Uint:
		returnValue, err = makeChan[uint](from, size, elemOps)
	case reflect.Uint8:
		returnValue, err = makeChan[uint8](from, size, elemOps)
	case reflect.Uint16:
		returnValue, err = makeChan[uint16](from, size, elemOps)
	case reflect.Uint32:
		returnValue, err = makeChan[uint32](from, size, elemOps)
	case reflect.Uint64:
		returnValue, err = makeChan[uint64](from, size, elemOps)
	case reflect.Uintptr:
		returnValue, err = makeChan[uintptr](from, size, elemOps)
	case reflect.String:
		returnValue, err = makeChan[string](from, size, elemOps)

	//
	// Slices — chan []T
	//
	case reflect.Slice:
		switch to.Type().Elem().Elem().Kind() {
		default:
			return defaultValue, fmt.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())
		case reflect.Interface:
			returnValue, err = makeChan[[]any](from, size, elemOps)
		case reflect.Bool:
			returnValue, err = makeChan[[]bool](from, size, elemOps)
		case reflect.Complex64:
			returnValue, err = makeChan[[]complex64](from, size, elemOps)
		case reflect.Complex128:
			returnValue, err = makeChan[[]complex128](from, size, elemOps)
		case reflect.Float32:
			returnValue, err = makeChan[[]float32](from, size, elemOps)
		case reflect.Float64:
			returnValue, err = makeChan[[]float64](from, size, elemOps)
		case reflect.Int:
			returnValue, err = makeChan[[]int](from, size, elemOps)
		case reflect.Int8:
			returnValue, err = makeChan[[]int8](from, size, elemOps)
		case reflect.Int16:
			returnValue, err = makeChan[[]int16](from, size, elemOps)
		case reflect.Int32:
			returnValue, err = makeChan[[]int32](from, size, elemOps)
		case reflect.Int64:
			returnValue, err = makeChan[[]int64](from, size, elemOps)
		case reflect.Uint:
			returnValue, err = makeChan[[]uint](from, size, elemOps)
		case reflect.Uint8:
			returnValue, err = makeChan[[]uint8](from, size, elemOps)
		case reflect.Uint16:
			returnValue, err = makeChan[[]uint16](from, size, elemOps)
		case reflect.Uint32:
			returnValue, err = makeChan[[]uint32](from, size, elemOps)
		case reflect.Uint64:
			returnValue, err = makeChan[[]uint64](from, size, elemOps)
		case reflect.Uintptr:
			returnValue, err = makeChan[[]uintptr](from, size, elemOps)
		case reflect.String:
			returnValue, err = makeChan[[]string](from, size, elemOps)
		}

	// Funcs — chan Func[T]
	case reflect.Func:
		if to.Type().Elem().NumOut() < 1 {
			return defaultValue, fmt.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())
		}
		switch to.Type().Elem().Out(0).Kind() {
		default:
			return defaultValue, fmt.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())
		case reflect.Interface:
			returnValue, err = makeChan[Func[any]](from, size, elemOps)
		case reflect.Bool:
			returnValue, err = makeChan[Func[bool]](from, size, elemOps)
		case reflect.Complex64:
			returnValue, err = makeChan[Func[complex64]](from, size, elemOps)
		case reflect.Complex128:
			returnValue, err = makeChan[Func[complex128]](from, size, elemOps)
		case reflect.Float32:
			returnValue, err = makeChan[Func[float32]](from, size, elemOps)
		case reflect.Float64:
			returnValue, err = makeChan[Func[float64]](from, size, elemOps)
		case reflect.Int:
			returnValue, err = makeChan[Func[int]](from, size, elemOps)
		case reflect.Int8:
			returnValue, err = makeChan[Func[int8]](from, size, elemOps)
		case reflect.Int16:
			returnValue, err = makeChan[Func[int16]](from, size, elemOps)
		case reflect.Int32:
			returnValue, err = makeChan[Func[int32]](from, size, elemOps)
		case reflect.Int64:
			returnValue, err = makeChan[Func[int64]](from, size, elemOps)
		case reflect.Uint:
			returnValue, err = makeChan[Func[uint]](from, size, elemOps)
		case reflect.Uint8:
			returnValue, err = makeChan[Func[uint8]](from, size, elemOps)
		case reflect.Uint16:
			returnValue, err = makeChan[Func[uint16]](from, size, elemOps)
		case reflect.Uint32:
			returnValue, err = makeChan[Func[uint32]](from, size, elemOps)
		case reflect.Uint64:
			returnValue, err = makeChan[Func[uint64]](from, size, elemOps)
		case reflect.Uintptr:
			returnValue, err = makeChan[Func[uintptr]](from, size, elemOps)
		case reflect.String:
			returnValue, err = makeChan[Func[string]](from, size, elemOps)
		case reflect.Slice:
			switch to.Type().Elem().Out(0).Elem().Kind() {
			default:
				return defaultValue, fmt.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())
			case reflect.Interface:
				returnValue, err = makeChan[Func[[]any]](from, size, elemOps)
			case reflect.Bool:
				returnValue, err = makeChan[Func[[]bool]](from, size, elemOps)
			case reflect.Complex64:
				returnValue, err = makeChan[Func[[]complex64]](from, size, elemOps)
			case reflect.Complex128:
				returnValue, err = makeChan[Func[[]complex128]](from, size, elemOps)
			case reflect.Float32:
				returnValue, err = makeChan[Func[[]float32]](from, size, elemOps)
			case reflect.Float64:
				returnValue, err = makeChan[Func[[]float64]](from, size, elemOps)
			case reflect.Int:
				returnValue, err = makeChan[Func[[]int]](from, size, elemOps)
			case reflect.Int8:
				returnValue, err = makeChan[Func[[]int8]](from, size, elemOps)
			case reflect.Int16:
				returnValue, err = makeChan[Func[[]int16]](from, size, elemOps)
			case reflect.Int32:
				returnValue, err = makeChan[Func[[]int32]](from, size, elemOps)
			case reflect.Int64:
				returnValue, err = makeChan[Func[[]int64]](from, size, elemOps)
			case reflect.Uint:
				returnValue, err = makeChan[Func[[]uint]](from, size, elemOps)
			case reflect.Uint8:
				returnValue, err = makeChan[Func[[]uint8]](from, size, elemOps)
			case reflect.Uint16:
				returnValue, err = makeChan[Func[[]uint16]](from, size, elemOps)
			case reflect.Uint32:
				returnValue, err = makeChan[Func[[]uint32]](from, size, elemOps)
			case reflect.Uint64:
				returnValue, err = makeChan[Func[[]uint64]](from, size, elemOps)
			case reflect.Uintptr:
				returnValue, err = makeChan[Func[[]uintptr]](from, size, elemOps)
			case reflect.String:
				returnValue, err = makeChan[Func[[]string]](from, size, elemOps)
			}
		}

	// Channels of channels — chan chan T
	case reflect.Chan:
		switch to.Type().Elem().Elem().Kind() {
		default:
			return defaultValue, fmt.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface())
		case reflect.Interface:
			returnValue, err = makeChan[chan any](from, size, elemOps)
		case reflect.Bool:
			returnValue, err = makeChan[chan bool](from, size, elemOps)
		case reflect.Complex64:
			returnValue, err = makeChan[chan complex64](from, size, elemOps)
		case reflect.Complex128:
			returnValue, err = makeChan[chan complex128](from, size, elemOps)
		case reflect.Float32:
			returnValue, err = makeChan[chan float32](from, size, elemOps)
		case reflect.Float64:
			returnValue, err = makeChan[chan float64](from, size, elemOps)
		case reflect.Int:
			returnValue, err = makeChan[chan int](from, size, elemOps)
		case reflect.Int8:
			returnValue, err = makeChan[chan int8](from, size, elemOps)
		case reflect.Int16:
			returnValue, err = makeChan[chan int16](from, size, elemOps)
		case reflect.Int32:
			returnValue, err = makeChan[chan int32](from, size, elemOps)
		case reflect.Int64:
			returnValue, err = makeChan[chan int64](from, size, elemOps)
		case reflect.Uint:
			returnValue, err = makeChan[chan uint](from, size, elemOps)
		case reflect.Uint8:
			returnValue, err = makeChan[chan uint8](from, size, elemOps)
		case reflect.Uint16:
			returnValue, err = makeChan[chan uint16](from, size, elemOps)
		case reflect.Uint32:
			returnValue, err = makeChan[chan uint32](from, size, elemOps)
		case reflect.Uint64:
			returnValue, err = makeChan[chan uint64](from, size, elemOps)
		case reflect.Uintptr:
			returnValue, err = makeChan[chan uintptr](from, size, elemOps)
		case reflect.String:
			returnValue, err = makeChan[chan string](from, size, elemOps)
		}

	}
	if err != nil {
		return defaultValue, err
	}
	// Convert to named channel type if needed (e.g. type MyChan chan int).
	if returnValue != nil {
		rv := reflect.ValueOf(returnValue)
		if rv.IsValid() && rv.Type() != to.Type() && rv.Type().ConvertibleTo(to.Type()) {
			returnValue = rv.Convert(to.Type()).Interface()
		}
	}
	return returnValue, nil
}

// makeChan casts from to T, creates a buffered channel of the given size,
// sends the cast value, and returns the channel as any. The caller is
// responsible for converting the result to a named channel type if needed;
// toChan does this via a reflect.Convert step after the switch.
func makeChan[T Types](from any, size int, ops ops) (any, error) {
	val, err := ToE[T](from, ops.List()...)
	if err != nil {
		return nil, err
	}
	ch := make(chan T, size)
	ch <- val
	return ch, nil
}
