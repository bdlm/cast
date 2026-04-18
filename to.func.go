package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toFunc returns a function that casts an interface to the specified type and
// returns the result.
//
// Options:
//   - DEFAULT: Func[T], default return value on error.
func toFunc[TTo any](to reflect.Value, from any, ops Ops) (TTo, error) {
	var err error
	var ret TTo
	var reti any
	var ok bool

	if _, ok = ops[DEFAULT]; ok {
		if ret, ok = ops[DEFAULT].(TTo); !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
		}
		ops = ops.Delete(DEFAULT) // Prevent DEFAULT from being passed to element casts.
	}

	switch to.Type().Out(0).Kind() {
	//case reflect.Struct:
	//case reflect.UnsafePointer:
	//case reflect.Map:
	//case reflect.Pointer:
	default:
		return ret, errors.Errorf(ErrorStrUnableToCast, from, from, to.Interface())

	case reflect.Interface:
		reti, err = makeFunc[any](from, ret, ops)
	case reflect.Bool:
		reti, err = makeFunc[bool](from, ret, ops)
	case reflect.Complex64:
		reti, err = makeFunc[complex64](from, ret, ops)
	case reflect.Complex128:
		reti, err = makeFunc[complex128](from, ret, ops)
	case reflect.Float32:
		reti, err = makeFunc[float32](from, ret, ops)
	case reflect.Float64:
		reti, err = makeFunc[float64](from, ret, ops)
	case reflect.Int:
		reti, err = makeFunc[int](from, ret, ops)
	case reflect.Int8:
		reti, err = makeFunc[int8](from, ret, ops)
	case reflect.Int16:
		reti, err = makeFunc[int16](from, ret, ops)
	case reflect.Int32:
		reti, err = makeFunc[int32](from, ret, ops)
	case reflect.Int64:
		reti, err = makeFunc[int64](from, ret, ops)
	case reflect.Uint:
		reti, err = makeFunc[uint](from, ret, ops)
	case reflect.Uint8:
		reti, err = makeFunc[uint8](from, ret, ops)
	case reflect.Uint16:
		reti, err = makeFunc[uint16](from, ret, ops)
	case reflect.Uint32:
		reti, err = makeFunc[uint32](from, ret, ops)
	case reflect.Uint64:
		reti, err = makeFunc[uint64](from, ret, ops)
	case reflect.Uintptr:
		reti, err = makeFunc[uintptr](from, ret, ops)
	case reflect.String:
		reti, err = makeFunc[string](from, ret, ops)

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Slices — Func[[]T]
	////////////////////////////////////////////////////////////////////////////////////////////////
	case reflect.Slice:
		switch to.Type().Out(0).Elem().Kind() {
		default:
			return ret, errors.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
		case reflect.Interface:
			reti, err = makeFunc[[]any](from, ret, ops)
		case reflect.Bool:
			reti, err = makeFunc[[]bool](from, ret, ops)
		case reflect.Complex64:
			reti, err = makeFunc[[]complex64](from, ret, ops)
		case reflect.Complex128:
			reti, err = makeFunc[[]complex128](from, ret, ops)
		case reflect.Float32:
			reti, err = makeFunc[[]float32](from, ret, ops)
		case reflect.Float64:
			reti, err = makeFunc[[]float64](from, ret, ops)
		case reflect.Int:
			reti, err = makeFunc[[]int](from, ret, ops)
		case reflect.Int8:
			reti, err = makeFunc[[]int8](from, ret, ops)
		case reflect.Int16:
			reti, err = makeFunc[[]int16](from, ret, ops)
		case reflect.Int32:
			reti, err = makeFunc[[]int32](from, ret, ops)
		case reflect.Int64:
			reti, err = makeFunc[[]int64](from, ret, ops)
		case reflect.Uint:
			reti, err = makeFunc[[]uint](from, ret, ops)
		case reflect.Uint8:
			reti, err = makeFunc[[]uint8](from, ret, ops)
		case reflect.Uint16:
			reti, err = makeFunc[[]uint16](from, ret, ops)
		case reflect.Uint32:
			reti, err = makeFunc[[]uint32](from, ret, ops)
		case reflect.Uint64:
			reti, err = makeFunc[[]uint64](from, ret, ops)
		case reflect.Uintptr:
			reti, err = makeFunc[[]uintptr](from, ret, ops)
		case reflect.String:
			reti, err = makeFunc[[]string](from, ret, ops)
		}

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Arrays — Func[[N]T] (reflection path, N is not known at compile time)
	////////////////////////////////////////////////////////////////////////////////////////////////
	case reflect.Array:
		reti, err = makeArrayFunc(to.Type(), from, ops)

	////////////////////////////////////////////////////////////////////////////////////////////////
	// Channels — Func[chan T], Func[chan []T], Func[chan chan T], Func[chan [N]T]
	////////////////////////////////////////////////////////////////////////////////////////////////
	case reflect.Chan:
		switch to.Type().Out(0).Elem().Kind() {
		//case reflect.Struct:
		//case reflect.UnsafePointer:
		//case reflect.Map:
		//case reflect.Pointer:
		default:
			return ret, errors.Errorf(ErrorStrUnableToCast, from, from, to.Interface())

		case reflect.Interface:
			reti, err = makeFunc[chan any](from, ret, ops)
		case reflect.Bool:
			reti, err = makeFunc[chan bool](from, ret, ops)
		case reflect.Complex64:
			reti, err = makeFunc[chan complex64](from, ret, ops)
		case reflect.Complex128:
			reti, err = makeFunc[chan complex128](from, ret, ops)
		case reflect.Float32:
			reti, err = makeFunc[chan float32](from, ret, ops)
		case reflect.Float64:
			reti, err = makeFunc[chan float64](from, ret, ops)
		case reflect.Int:
			reti, err = makeFunc[chan int](from, ret, ops)
		case reflect.Int8:
			reti, err = makeFunc[chan int8](from, ret, ops)
		case reflect.Int16:
			reti, err = makeFunc[chan int16](from, ret, ops)
		case reflect.Int32:
			reti, err = makeFunc[chan int32](from, ret, ops)
		case reflect.Int64:
			reti, err = makeFunc[chan int64](from, ret, ops)
		case reflect.Uint:
			reti, err = makeFunc[chan uint](from, ret, ops)
		case reflect.Uint8:
			reti, err = makeFunc[chan uint8](from, ret, ops)
		case reflect.Uint16:
			reti, err = makeFunc[chan uint16](from, ret, ops)
		case reflect.Uint32:
			reti, err = makeFunc[chan uint32](from, ret, ops)
		case reflect.Uint64:
			reti, err = makeFunc[chan uint64](from, ret, ops)
		case reflect.Uintptr:
			reti, err = makeFunc[chan uintptr](from, ret, ops)
		case reflect.String:
			reti, err = makeFunc[chan string](from, ret, ops)

		// Func[chan []T]
		case reflect.Slice:
			switch to.Type().Out(0).Elem().Elem().Kind() {
			default:
				return ret, errors.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
			case reflect.Interface:
				reti, err = makeFunc[chan []any](from, ret, ops)
			case reflect.Bool:
				reti, err = makeFunc[chan []bool](from, ret, ops)
			case reflect.Complex64:
				reti, err = makeFunc[chan []complex64](from, ret, ops)
			case reflect.Complex128:
				reti, err = makeFunc[chan []complex128](from, ret, ops)
			case reflect.Float32:
				reti, err = makeFunc[chan []float32](from, ret, ops)
			case reflect.Float64:
				reti, err = makeFunc[chan []float64](from, ret, ops)
			case reflect.Int:
				reti, err = makeFunc[chan []int](from, ret, ops)
			case reflect.Int8:
				reti, err = makeFunc[chan []int8](from, ret, ops)
			case reflect.Int16:
				reti, err = makeFunc[chan []int16](from, ret, ops)
			case reflect.Int32:
				reti, err = makeFunc[chan []int32](from, ret, ops)
			case reflect.Int64:
				reti, err = makeFunc[chan []int64](from, ret, ops)
			case reflect.Uint:
				reti, err = makeFunc[chan []uint](from, ret, ops)
			case reflect.Uint8:
				reti, err = makeFunc[chan []uint8](from, ret, ops)
			case reflect.Uint16:
				reti, err = makeFunc[chan []uint16](from, ret, ops)
			case reflect.Uint32:
				reti, err = makeFunc[chan []uint32](from, ret, ops)
			case reflect.Uint64:
				reti, err = makeFunc[chan []uint64](from, ret, ops)
			case reflect.Uintptr:
				reti, err = makeFunc[chan []uintptr](from, ret, ops)
			case reflect.String:
				reti, err = makeFunc[chan []string](from, ret, ops)
			}

		// Func[chan Func[T]]
		case reflect.Func:
			switch to.Type().Out(0).Elem().Out(0).Kind() {
			default:
				return ret, errors.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
			case reflect.Interface:
				reti, err = makeFunc[chan Func[any]](from, ret, ops)
			case reflect.Bool:
				reti, err = makeFunc[chan Func[bool]](from, ret, ops)
			case reflect.Complex64:
				reti, err = makeFunc[chan Func[complex64]](from, ret, ops)
			case reflect.Complex128:
				reti, err = makeFunc[chan Func[complex128]](from, ret, ops)
			case reflect.Float32:
				reti, err = makeFunc[chan Func[float32]](from, ret, ops)
			case reflect.Float64:
				reti, err = makeFunc[chan Func[float64]](from, ret, ops)
			case reflect.Int:
				reti, err = makeFunc[chan Func[int]](from, ret, ops)
			case reflect.Int8:
				reti, err = makeFunc[chan Func[int8]](from, ret, ops)
			case reflect.Int16:
				reti, err = makeFunc[chan Func[int16]](from, ret, ops)
			case reflect.Int32:
				reti, err = makeFunc[chan Func[int32]](from, ret, ops)
			case reflect.Int64:
				reti, err = makeFunc[chan Func[int64]](from, ret, ops)
			case reflect.Uint:
				reti, err = makeFunc[chan Func[uint]](from, ret, ops)
			case reflect.Uint8:
				reti, err = makeFunc[chan Func[uint8]](from, ret, ops)
			case reflect.Uint16:
				reti, err = makeFunc[chan Func[uint16]](from, ret, ops)
			case reflect.Uint32:
				reti, err = makeFunc[chan Func[uint32]](from, ret, ops)
			case reflect.Uint64:
				reti, err = makeFunc[chan Func[uint64]](from, ret, ops)
			case reflect.Uintptr:
				reti, err = makeFunc[chan Func[uintptr]](from, ret, ops)
			case reflect.String:
				reti, err = makeFunc[chan Func[string]](from, ret, ops)
			}

		// Func[chan chan T]
		case reflect.Chan:
			switch to.Type().Out(0).Elem().Elem().Kind() {
			default:
				return ret, errors.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
			case reflect.Interface:
				reti, err = makeFunc[chan chan any](from, ret, ops)
			case reflect.Bool:
				reti, err = makeFunc[chan chan bool](from, ret, ops)
			case reflect.Complex64:
				reti, err = makeFunc[chan chan complex64](from, ret, ops)
			case reflect.Complex128:
				reti, err = makeFunc[chan chan complex128](from, ret, ops)
			case reflect.Float32:
				reti, err = makeFunc[chan chan float32](from, ret, ops)
			case reflect.Float64:
				reti, err = makeFunc[chan chan float64](from, ret, ops)
			case reflect.Int:
				reti, err = makeFunc[chan chan int](from, ret, ops)
			case reflect.Int8:
				reti, err = makeFunc[chan chan int8](from, ret, ops)
			case reflect.Int16:
				reti, err = makeFunc[chan chan int16](from, ret, ops)
			case reflect.Int32:
				reti, err = makeFunc[chan chan int32](from, ret, ops)
			case reflect.Int64:
				reti, err = makeFunc[chan chan int64](from, ret, ops)
			case reflect.Uint:
				reti, err = makeFunc[chan chan uint](from, ret, ops)
			case reflect.Uint8:
				reti, err = makeFunc[chan chan uint8](from, ret, ops)
			case reflect.Uint16:
				reti, err = makeFunc[chan chan uint16](from, ret, ops)
			case reflect.Uint32:
				reti, err = makeFunc[chan chan uint32](from, ret, ops)
			case reflect.Uint64:
				reti, err = makeFunc[chan chan uint64](from, ret, ops)
			case reflect.Uintptr:
				reti, err = makeFunc[chan chan uintptr](from, ret, ops)
			case reflect.String:
				reti, err = makeFunc[chan chan string](from, ret, ops)
			}

		// Func[chan [N]T] (reflection path)
		case reflect.Array:
			reti, err = makeArrayChanFunc(to.Type(), from, ops)
		}
	}
	if err != nil {
		return ret, err
	}
	if ret, ok := reti.(TTo); ok {
		return ret, nil
	}
	return ret, errors.Errorf(ErrorStrErrorCastingFunc, reti, ret)
}

// makeFunc casts from to T and returns a Func[T] closure capturing the result.
func makeFunc[T Types](from any, orig any, ops Ops) (any, error) {
	val, err := ToE[T](from, ops.List()...)
	if err != nil {
		return nil, errors.Wrap(err, ErrorStrErrorCastingFunc, from, orig)
	}
	return Func[T](func() T { return val }), nil
}
