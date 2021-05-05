package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapString(from, to interface{}) interface{} {
	ret, _ := CastMapStringE(from, to)
	return ret
}

func CastMapStringE(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapStringBoolE(from) // Bool
	case reflect.Int:
		return ToMapStringIntE(from) // Int
	case reflect.Int8:
		return ToMapStringInt8E(from) // Int8
	case reflect.Int16:
		return ToMapStringInt16E(from) // Int16
	case reflect.Int32:
		return ToMapStringInt32E(from) // Int32
	case reflect.Int64:
		return ToMapStringInt64E(from) // Int64
	case reflect.Uint:
		return ToMapStringUintE(from) // Uint
	case reflect.Uint8:
		return ToMapStringUint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapStringUint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapStringUint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapStringUint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapStringUintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapStringFloat32E(from) // Float32
	case reflect.Float64:
		return ToMapStringFloat64E(from) // Float64
	case reflect.Complex64:
		return ToMapStringComplex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapStringComplex128E(from) // Complex128
	case reflect.Interface:
		return ToMapStringInterfaceE(from) // Interface
	case reflect.String:
		return ToMapStringStringE(from) // String
	}
}
