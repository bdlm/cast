package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapInterface(from, to interface{}) interface{} {
	ret, _ := CastMapInterfaceE(from, to)
	return ret
}

func CastMapInterfaceE(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapInterfaceBoolE(from) // Bool
	case reflect.Int:
		return ToMapInterfaceIntE(from) // Int
	case reflect.Int8:
		return ToMapInterfaceInt8E(from) // Int8
	case reflect.Int16:
		return ToMapInterfaceInt16E(from) // Int16
	case reflect.Int32:
		return ToMapInterfaceInt32E(from) // Int32
	case reflect.Int64:
		return ToMapInterfaceInt64E(from) // Int64
	case reflect.Uint:
		return ToMapInterfaceUintE(from) // Uint
	case reflect.Uint8:
		return ToMapInterfaceUint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapInterfaceUint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapInterfaceUint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapInterfaceUint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapInterfaceUintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapInterfaceFloat32E(from) // Float32
	case reflect.Float64:
		return ToMapInterfaceFloat64E(from) // Float64
	case reflect.Complex64:
		return ToMapInterfaceComplex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapInterfaceComplex128E(from) // Complex128
	case reflect.Interface:
		return ToMapInterfaceInterfaceE(from) // Interface
	case reflect.String:
		return ToMapInterfaceStringE(from) // String
	}
}
