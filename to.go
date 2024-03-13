package cast

import (
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
	std_error "github.com/bdlm/std/v2/errors"
)

// To casts the value v to the given type, ignoring any errors. See the ToE
// documentation more information.
func To[TTo Types](v any, o ...Op) TTo {
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
//   - If the target type is a channel, a channel with a buffer of 1 is created
//     and the cast value `v` is added to the the channel before it is returned.
//
//   - If the target type is an array a slice is created. To create a slice with
//     a backing array with a spcific size, set the LENGTH flag to the desired
//     size as an integer: `slice, err := cast.ToE[[]int](v, Ops{LENGTH: 10})`.
//     The value `v` is cast to the required type and appended to the returned
//     slice.
//
//   - If the target type is a map, a map is created with a zero-value key
//     containing the cast value `v` which is then returned.
//
// See the documentation for the specific type conversion function for more
// information.
func ToE[TTo Types](val any, ops ...Op) (panicTo TTo, panicErr error) {

	var err error
	var ok bool
	var retIface any
	var ret0Val TTo
	var retVal TTo

	// Don't panic.
	defer func() {
		if err := recover(); err != nil {
			panicTo = ret0Val
			panicErr = errors.Wrap(err.(error), "failure casting %T to %T (panic)", val, ret0Val)
			fmt.Printf("% +#v", panicErr)
		}
	}()
	go func() {
		if e := recover(); e != nil {
			err = errors.WrapE(err, e.(error))
		}
	}()

	options := parseOps(ops)

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
		retIface = ret0Val
		if _, ok := retIface.(error); ok {
			retIface = errors.Errorf(To[string](val, ops...))
		} else if _, ok := retIface.(std_error.Error); ok {
			retIface = errors.Errorf(To[string](val, ops...))
		} else if _, ok := retIface.(fmt.Stringer); ok {
			retIface = errors.Errorf(To[string](val, ops...))
		} else {
			return ret0Val, errors.WrapE(Error, errors.Errorf(ErrorStrUnableToCast, from, from.Interface(), to.Interface()))
		}

	case reflect.Interface:
		retIface = val

	case reflect.Bool:
		retIface, err = toBool(val, options)
	case reflect.Chan:
		retIface, err = toChan(to, val, options)
	//case reflect.Map: // TODO
	//	if reflect.Map != from.Type().Kind() {
	//		return ret0Val, errors.WrapE(Error, errors.Errorf(ErrorStrUnableToCast, from, from.Interface(), to.Interface()))
	//	}
	//	retIface, err = toMap(to, val, options)
	case reflect.Array: // TODO
		fallthrough
	case reflect.Slice:
		if reflect.Slice != from.Type().Kind() {
			return ret0Val, errors.WrapE(Error, errors.Errorf(ErrorStrUnableToCast, from, from.Interface(), to.Interface()))
		}
		retIface, err = toSlice(to, val, options)
	case reflect.Func:
		retIface, err = toFunc[TTo](to, val, options)
	case reflect.Complex64:
		retIface, err = toComplex[complex64](from, options)
	case reflect.Complex128:
		retIface, err = toComplex[complex128](from, options)
	case reflect.Float32:
		retIface, err = toFloat[float32](from, options)
	case reflect.Float64:
		retIface, err = toFloat[float64](from, options)
	case reflect.Int:
		retIface, err = toInt[int](from, options)
	case reflect.Int8:
		retIface, err = toInt[int8](from, options)
	case reflect.Int16:
		retIface, err = toInt[int16](from, options)
	case reflect.Int32:
		retIface, err = toInt[int32](from, options)
	case reflect.Int64:
		retIface, err = toInt[int64](from, options)
	case reflect.Uint:
		retIface, err = toInt[uint](from, options)
	case reflect.Uint8:
		retIface, err = toInt[uint8](from, options)
	case reflect.Uint16:
		retIface, err = toInt[uint16](from, options)
	case reflect.Uint32:
		retIface, err = toInt[uint32](from, options)
	case reflect.Uint64:
		retIface, err = toInt[uint64](from, options)
	case reflect.Uintptr:
		retIface, err = toInt[uintptr](from, options)
	case reflect.String:
		retIface, err = toString(from, options)
	}

	if retVal, ok = retIface.(TTo); !ok && retIface != nil {
		return ret0Val, errors.WrapE(Error, errors.Errorf("unable to cast %#.10v of type %T to %T (%#.10v %T)", from, from.Interface(), *new(TTo), retVal, retVal))
	}

	if err != nil {
		return retVal, errors.WrapE(Error, err)
	}

	if retVal, ok = retIface.(TTo); ok {
		return retVal, nil
	}

	if nil == retIface {
		return ret0Val, nil
	}

	return retVal, errors.WrapE(Error, errors.Errorf("unable to cast %#.10v of type %T to %T (%#.10v %T)", from, from.Interface(), to.Interface(), retVal, retVal))
}
