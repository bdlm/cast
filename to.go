package cast

import (
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
	std_error "github.com/bdlm/std/v2/errors"
)

// To casts the value v to the given type, ignoring any errors. See the ToE
// documentation more information.
func To[TTo Types](v any, o ...Ops) TTo {
	ret, _ := ToE[TTo](v, o...)
	return ret
}

// ToE casts the value v to the given type, returning any errors.
//
// o (Ops) is an optional parameter providing flags that can be used to modify
// the default type conversion behavior. If ops is not provided, the default
// conversion behavior for a given type is used. Available options depend on the
// target type, see the documentation for the specific type conversion function
// for more information.
//
// Complex types have specific default behaviors, for example:
//
//   - If the given type is a channel, a channel with a buffer of 1 is created
//     and the cast value `v` is added to the the channel before it is returned.
//
//   - If the given type is an array a slice is created. To create a slice with
//     a backing array with a spcific size, set the LENGTH flag to the desired
//     size as an integer: `slice, err := cast.ToE[[]int](v, Ops{LENGTH: 10})`.
//     The value `v` is cast to the required type and appended to the returned
//     slice.
//
//   - If the given type is a map, a map is created with a zero-value key
//     containing the cast value `v` which is then returned.
//
// See the documentation for the specific type conversion function for more
// information.
func ToE[TTo Types](val any, o ...Ops) (panicTo TTo, panicErr error) {
	ops := parseOps(o)

	var err error
	var reti any
	var ret TTo
	var ok bool

	// Don't panic.
	defer func() {
		if err := recover(); err != nil {
			panicTo = ret
			panicErr = errors.Wrap(err.(error), "failure casting %T to %T (panic)", val, ret)
			fmt.Printf("% +#v", panicErr)
		}
	}()
	go func() {
		if e := recover(); e != nil {
			err = errors.WrapE(err, e.(error))
		}
	}()

	// Collapse reflection values.
	from := reflect.ValueOf(val)
	toRef := reflect.ValueOf(new(TTo))
	to := reflect.Indirect(toRef)

	switch to.Type().Kind() {
	// reflect.Array:
	// reflect.Invalid:
	// reflect.Pointer:
	// reflect.Struct:
	// reflect.UnsafePointer:
	// error
	// std_error.Error
	default:
		reti = ret
		if _, ok := reti.(error); ok {
			reti = errors.Errorf(To[string](val, ops))
		} else if _, ok := reti.(std_error.Error); ok {
			reti = errors.Errorf(To[string](val, ops))
		} else if _, ok := reti.(fmt.Stringer); ok {
			reti = errors.Errorf(To[string](val, ops))
		} else {
			return ret, errors.WrapE(Error, errors.Errorf(ErrorStrUnableToCast, from, from.Interface(), to.Interface()))
		}

	case reflect.Interface:
		reti = val

	case reflect.Bool:
		reti, err = toBool(val, ops)
	case reflect.Chan:
		reti, err = toChan(to, val, ops)
	//case reflect.Map: // TODO
	//	if reflect.Map != from.Type().Kind() {
	//		return ret, errors.WrapE(Error, errors.Errorf(ErrorStrUnableToCast, from, from.Interface(), to.Interface()))
	//	}
	//	reti, err = toMap(to, val, ops)
	case reflect.Array: // TODO
		fallthrough
	case reflect.Slice:
		if reflect.Slice != from.Type().Kind() {
			return ret, errors.WrapE(Error, errors.Errorf(ErrorStrUnableToCast, from, from.Interface(), to.Interface()))
		}
		reti, err = toSlice(to, val, ops)
	case reflect.Func:
		reti, err = toFunc[TTo](to, val, ops)
	case reflect.Complex64:
		reti, err = toComplex[complex64](from, ops)
	case reflect.Complex128:
		reti, err = toComplex[complex128](from, ops)
	case reflect.Float32:
		reti, err = toFloat[float32](from, ops)
	case reflect.Float64:
		reti, err = toFloat[float64](from, ops)
	case reflect.Int:
		reti, err = toInt[int](from, ops)
	case reflect.Int8:
		reti, err = toInt[int8](from, ops)
	case reflect.Int16:
		reti, err = toInt[int16](from, ops)
	case reflect.Int32:
		reti, err = toInt[int32](from, ops)
	case reflect.Int64:
		reti, err = toInt[int64](from, ops)
	case reflect.Uint:
		reti, err = toInt[uint](from, ops)
	case reflect.Uint8:
		reti, err = toInt[uint8](from, ops)
	case reflect.Uint16:
		reti, err = toInt[uint16](from, ops)
	case reflect.Uint32:
		reti, err = toInt[uint32](from, ops)
	case reflect.Uint64:
		reti, err = toInt[uint64](from, ops)
	case reflect.Uintptr:
		reti, err = toInt[uintptr](from, ops)
	case reflect.String:
		reti, err = toString(from, ops)
	}

	if err != nil {
		return ret, errors.WrapE(Error, err)
	}

	if ret, ok = reti.(TTo); ok {
		return ret, nil
	}

	if nil == reti {
		return *new(TTo), nil
	}

	return ret, errors.WrapE(Error, errors.Errorf("unable to cast %#.10v of type %T to %T (%#.10v %T)", from, from.Interface(), to.Interface(), ret, ret))
}
