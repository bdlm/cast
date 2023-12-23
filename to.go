package cast

import (
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
	std_error "github.com/bdlm/std/v2/errors"
)

type flag int
type Ops map[flag]any

const (
	LENGTH flag = iota
	DEFAULT
)

func parseOps[TTo Types](o ...Ops) {

}

// To casts the value `v` to the given type, ignoring any errors.
func To[TTo Types](v any, o ...Ops) TTo {
	ret, _ := ToE[TTo](v, o...)
	return ret
}

// ToE casts the value `v` to the given type, returning any errors.
//
// If the given type is a channel, a channel with a buffer of 1 is created and
// the cast value `v` is added to the the channel before it is returned.
//
// If the given type is an array a slice is created. To create a slice with a
// backing array with a spcific size, set the LENGTH flag to the desired size as
// an integer: `slice, err := cast.ToE[[]int](v, Ops{LENGTH: 10})`. The value `v`
// is cast to the required type and appended to the returned slice.
//
// If the given type is a map, a map is created with a zero-value key containing
// the cast value `v` which is then returned.
//
// ops: optional, some flags may be passed to control type conversions, for
// example when creating a channel the the buffer size can be specified.
func ToE[TTo Types](val any, o ...Ops) (panicTo TTo, panicErr error) {
	var err error
	var reti any
	var ret TTo
	var ok bool

	// Don't panic.
	defer func() {
		if err := recover(); err != nil {
			panicTo = ret
			panicErr = errors.Wrap(err.(error), "failure casting %T to %T", val, ret)
		}
	}()

	ops := Ops{}
	for _, o := range o {
		for k, v := range o {
			ops[k] = v
		}
	}

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
			reti = errors.Errorf(To[string](val))
		} else if _, ok := reti.(std_error.Error); ok {
			reti = errors.Errorf(To[string](val))
		} else {
			return ret, errors.WrapE(Error, errors.Errorf("unable to cast %#v of type %T to %T", from, from, to.Interface()))
		}

	case reflect.Interface:
		reti = val

	case reflect.Bool:
		reti, err = toBool(val)
	case reflect.Chan:
		var size = 1
		if _, ok = ops[LENGTH].(int); ok {
			size = ops[LENGTH].(int)
		}
		reti, err = toChan(to, val, size)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		var size = 1
		if _, ok = ops[LENGTH].(int); ok {
			size = ops[LENGTH].(int)
		}
		reti, err = toSlice(to, val, size)
	case reflect.Func:
		reti, err = toFunc[TTo](to, val)
	case reflect.Complex64:
		reti, err = toComplex[complex64](from)
	case reflect.Complex128:
		reti, err = toComplex[complex128](from)
	case reflect.Float32:
		reti, err = toFloat[float32](from)
	case reflect.Float64:
		reti, err = toFloat[float64](from)
	case reflect.Int:
		reti, err = toInt[int](from)
	case reflect.Int8:
		reti, err = toInt[int8](from)
	case reflect.Int16:
		reti, err = toInt[int16](from)
	case reflect.Int32:
		reti, err = toInt[int32](from)
	case reflect.Int64:
		reti, err = toInt[int64](from)
	case reflect.Uint:
		reti, err = toInt[uint](from)
	case reflect.Uint8:
		reti, err = toInt[uint8](from)
	case reflect.Uint16:
		reti, err = toInt[uint16](from)
	case reflect.Uint32:
		reti, err = toInt[uint32](from)
	case reflect.Uint64:
		reti, err = toInt[uint64](from)
	case reflect.Uintptr:
		reti, err = toInt[uintptr](from)
	case reflect.String:
		reti = fmt.Sprintf("%v", val)
	}

	if err != nil {
		return ret, errors.WrapE(Error, err)
	}

	if ret, ok = reti.(TTo); ok {
		return ret, nil
	}

	return ret, errors.WrapE(Error, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, to.Interface()))
}
