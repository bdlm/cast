package cast

import (
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
	std_error "github.com/bdlm/std/v2/errors"
)

type Flag int
type Op [2]any
type Ops map[Flag]any

const (
	LENGTH Flag = iota
	DEFAULT
)

// Parse ops
func parseOps(o ...Op) Ops {
	ops := Ops{}
	for _, op := range o {
		flag := Flag(To[int](op[0]))
		val := op[1]
		ops[flag] = val
	}
	return ops
}

// To casts the value `v` to the given type, ignoring any errors.
func To[TTo Types](v any, o ...Op) TTo {
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
func ToE[TTo Types](val any, o ...Op) (panicTo TTo, panicErr error) {
	var err error
	var reti any
	var ret TTo
	var ok bool

	// // Don't panic.
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		panicTo = ret
	// 		panicErr = errors.Wrap(err.(error), "failure casting %T to %T", val, ret)
	// 		fmt.Printf("% +#v", panicErr)
	// 	}
	// }()

	// go func() {
	// 	if e := recover(); e != nil {
	// 		err = errors.WrapE(err, e.(error))
	// 	}
	// }()

	// Parse option flags.
	ops := parseOps(o...)

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
			return ret, errors.WrapE(Error, errors.Errorf("unable to cast %#v of type %T to %T", from, from.Interface(), to.Interface()))
		}

	case reflect.Interface:
		reti = val

	case reflect.Bool:
		reti, err = toBool(val)
	case reflect.Chan:
		reti, err = toChan(to, val, ops)
		// TODO
	case reflect.Map:
		if reflect.Map != from.Type().Kind() {
			return ret, errors.WrapE(Error, errors.Errorf("unable to cast %#v of type %T to %T", from, from.Interface(), to.Interface()))
		}
		reti, err = toMap(to, val, ops)

	case reflect.Array:
		fallthrough
	case reflect.Slice:
		if reflect.Slice != from.Type().Kind() {
			return ret, errors.WrapE(Error, errors.Errorf("unable to cast %#v of type %T to %T", from, from.Interface(), to.Interface()))
		}
		reti, err = toSlice(to, val, ops)
	case reflect.Func:
		reti, err = toFunc[TTo](to, val)
		//reti, err = toFunc[TTo](from)
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

	if nil == reti {
		return *new(TTo), nil
	}

	return ret, errors.WrapE(Error, errors.Errorf("unable to cast %#.10v of type %T to %T (%#.10v %T)", from, from.Interface(), to.Interface(), ret, ret))
}

//func toType(typ reflect.Type, from any) (reflect.Value, error) {
//	var err error
//	var ret = reflect.New(typ)
//	var reterr = func(r any, e error) (reflect.Value, error) {
//		if nil != e {
//			return ret, e
//		}
//		if v, ok := r.(reflect.Value); ok {
//			return reflect.Append(ret, v), nil
//		}
//		return reflect.Append(ret, reflect.ValueOf(r)), nil
//	}
//	fmt.Printf("toType: %s\n", typ)
//
//	switch typ.Kind() {
//	case reflect.Invalid:
//		err = errors.Errorf("toType: unknown type %s (%T)", ret.Type(), ret.Interface())
//
//	case reflect.Chan:
//		fmt.Println("	not implemented")
//	case reflect.Func:
//		fmt.Println("	not implemented")
//	case reflect.Map:
//		fmt.Println("	not implemented")
//	case reflect.Pointer:
//		fmt.Println("	not implemented")
//	case reflect.Struct:
//		fmt.Println("	not implemented")
//	case reflect.UnsafePointer:
//		fmt.Println("	not implemented")
//
//	case reflect.Bool:
//		ret, err = reterr(toBool(from))
//	case reflect.Array:
//		fallthrough
//	case reflect.Slice:
//		r, err := toType(ret.Elem().Elem().Type(), from)
//		if nil != err {
//			ret = reflect.Append(ret, r)
//		}
//
//	// Numbers
//	case reflect.Complex64:
//		ret, err = reterr(toComplex[complex64](reflect.ValueOf(from)))
//	case reflect.Complex128:
//		ret, err = reterr(toComplex[complex128](reflect.ValueOf(from)))
//	case reflect.Float32:
//		ret, err = reterr(toFloat[float32](reflect.ValueOf(from)))
//	case reflect.Float64:
//		ret, err = reterr(toFloat[float64](reflect.ValueOf(from)))
//	case reflect.Int:
//		ret, err = reterr(toInt[int](reflect.ValueOf(from)))
//	case reflect.Int8:
//		ret, err = reterr(toInt[int8](reflect.ValueOf(from)))
//	case reflect.Int16:
//		ret, err = reterr(toInt[int16](reflect.ValueOf(from)))
//	case reflect.Int32:
//		ret, err = reterr(toInt[int32](reflect.ValueOf(from)))
//	case reflect.Int64:
//		ret, err = reterr(toInt[int64](reflect.ValueOf(from)))
//	case reflect.Uint:
//		ret, err = reterr(toInt[uint](reflect.ValueOf(from)))
//	case reflect.Uint8:
//		ret, err = reterr(toInt[uint8](reflect.ValueOf(from)))
//	case reflect.Uint16:
//		ret, err = reterr(toInt[uint16](reflect.ValueOf(from)))
//	case reflect.Uint32:
//		ret, err = reterr(toInt[uint32](reflect.ValueOf(from)))
//	case reflect.Uint64:
//		ret, err = reterr(toInt[uint64](reflect.ValueOf(from)))
//	case reflect.Uintptr:
//		ret, err = reterr(toInt[uintptr](reflect.ValueOf(from)))
//		//	case reflect.Complex64:
//		//		fallthrough
//		//	case reflect.Complex128:
//		//		fallthrough
//		//	case reflect.Float32:
//		//		fallthrough
//		//	case reflect.Float64:
//		//		fallthrough
//		//	case reflect.Int:
//		//		fallthrough
//		//	case reflect.Int8:
//		//		fallthrough
//		//	case reflect.Int16:
//		//		fallthrough
//		//	case reflect.Int32:
//		//		fallthrough
//		//	case reflect.Int64:
//		//		fallthrough
//		//	case reflect.Uint:
//		//		fallthrough
//		//	case reflect.Uint8:
//		//		fallthrough
//		//	case reflect.Uint16:
//		//		fallthrough
//		//	case reflect.Uint32:
//		//		fallthrough
//		//	case reflect.Uint64:
//		//		fallthrough
//		//	case reflect.Uintptr:
//		//		ret, err = toInt[TTo](from)
//
//	case reflect.Interface:
//		ret = reflect.ValueOf(interface{}(from))
//
//	case reflect.String:
//		ret = reflect.ValueOf(fmt.Sprintf("%#v", from))
//	}
//
//	fmt.Printf("toType: %#v (%T); %#v (%T)\n", ret, ret, err, err)
//	return ret, err
//
//}
