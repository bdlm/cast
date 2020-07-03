package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapUint(from, to interface{}) interface{} {
	ret, _ := CastMapUintE(from, to)
	return ret
}

func CastMapUintE(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapUintBoolE(from) // Bool
	case reflect.Int:
		return ToMapUintIntE(from) // Int
	case reflect.Int8:
		return ToMapUintInt8E(from) // Int8
	case reflect.Int16:
		return ToMapUintInt16E(from) // Int16
	case reflect.Int32:
		return ToMapUintInt32E(from) // Int32
	case reflect.Int64:
		return ToMapUintInt64E(from) // Int64
	case reflect.Uint:
		return ToMapUintUintE(from) // Uint
	case reflect.Uint8:
		return ToMapUintUint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapUintUint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapUintUint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapUintUint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapUintUintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapUintFloat32E(from) // Float32
	case reflect.Float64:
		return ToMapUintFloat64E(from) // Float64
	case reflect.Complex64:
		return ToMapUintComplex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapUintComplex128E(from) // Complex128
	case reflect.Interface:
		return ToMapUintInterfaceE(from) // Interface
	case reflect.String:
		return ToMapUintStringE(from) // String
	}
}
