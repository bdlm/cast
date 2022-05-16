package cast

import (
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
)

// To casts the value `v` to the given type, ignoring any errors.
func To[TTo Types](v any) TTo {
	ret, _ := ToE[TTo](v)
	return ret
}

// ToE casts the value `v` to the given type, returning any errors.
//
// If the given type is a channel, a channel with a buffer of 1 is created and
// the cast value `v` is added to the the channel before it is returned.
//
// //If the given type is an array, a slice is created with a size matching the
// //length of the array and the cast value `v` is appended to the slice before it
// //is returned.
//
// If the given type is a slice, a slice with a size of 1 is created and the
// cast value `v` is appended to the the slice before it is returned.
//
// If the given type is a map, a map is created with a zero-value key containing
// the cast value `v` which is then returned.
func ToE[TTo Types](val any) (TTo, error) {
	var err error
	var reti any
	var ret TTo
	var ok bool

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
	//case reflect.Array:
	//case reflect.Invalid:
	//case reflect.Map:
	//case reflect.Pointer:
	//case reflect.Struct:
	//case reflect.UnsafePointer:
	default:
		return ret, errors.WrapE(Error, errors.Errorf("unable to cast %#v of type %T to %T", from, from, to.Interface()))

	case reflect.Interface:
		reti = val

	case reflect.Bool:
		reti, err = toBool(val)
	case reflect.Chan:
		reti, err = toChan(to, val)
	case reflect.Slice:
		reti, err = toSlice(to, val)
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

	return ret, errors.WrapE(Error, errors.Errorf("unable to cast %#v of type %T to %T", from, from, to.Interface()))
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
