package cast

import (
	"fmt"
	"reflect"
	"time"
)

var (
	// base types
	Bool       bool
	Byte       byte
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint       uint
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Uintptr    uintptr
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
	Interface  interface{}
	String     string

	// common types
	Duration time.Duration
	Time     time.Time
)

func Cast(from, to interface{}) interface{} {
	ret, _ := CastE(from, to)
	return ret
}

func CastE(from, to interface{}) (interface{}, error) {
	// Collapse reflection values.
	if _, ok := from.(reflect.Value); ok {
		from = from.(reflect.Value).Interface()
	}
	if _, ok := to.(reflect.Value); ok {
		to = to.(reflect.Value).Interface()
	}
	from, to = indirect(from), indirect(to)

	fmt.Printf("	CastE - from: %v (%T); to: %v (%T); kind: %v (%T)\n", from, from, to, to, reflect.TypeOf(to), reflect.TypeOf(to))
	switch to.(type) {
	case time.Duration:
		return ToDurationE(from)
	case time.Time:
		return ToTimeE(from)
	default:
		switch reflect.TypeOf(to).Kind() {
		case reflect.Map:
			return CastMapE(from, to)
		case reflect.Slice:
			//return CastSliceE(from, to)
		case reflect.Bool:
			return ToBoolE(from)
		//case reflect.Duration:
		//	return ToDurationE(from)
		case reflect.Int:
			return ToIntE(from)
		case reflect.Int8:
			return ToInt8E(from)
		case reflect.Int16:
			return ToInt16E(from)
		case reflect.Int32:
			return ToInt32E(from)
		case reflect.Int64:
			return ToInt64E(from)
		case reflect.Uint:
			return ToUintE(from)
		case reflect.Uint8:
			return ToUint8E(from)
		case reflect.Uint16:
			return ToUint16E(from)
		case reflect.Uint32:
			return ToUint32E(from)
		case reflect.Uint64:
			return ToUint64E(from)
		case reflect.Uintptr:
			//return ToUintptrE(from)
		case reflect.Float32:
			return ToFloat32E(from)
		case reflect.Float64:
			return ToFloat64E(from)
		case reflect.Array:
			//return ToArrayE(from)
		case reflect.Chan:
			//return ToChanE(from)
		case reflect.Func:
			//return ToFuncE(from)
		case reflect.Interface:
			//return ToInterfaceE(from)
		case reflect.String:
			return ToStringE(from)
		case reflect.Struct:
			fmt.Println("SHOULDN'T GET HERE")
			//return ToStructE(from)
			//case reflect.Ptr:
			//	return ToPtrE(from)
			//case reflect.UnsafePointer:
			//	return ToUnsafePointerE(from)

		}
	}

	//fmt.Printf("FROM: %T %#v, TO: %T %#v, reflect.Int: %T, %#v, t.Kind: %T, %#v\n\n\n", from, from, to, to, reflect.Int, reflect.Int, t.String(), t.Kind().String())
	return nil, nil
}
