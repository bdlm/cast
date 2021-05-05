package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapUintptr(from, to interface{}) interface{} {
	ret, _ := CastMapUintptrE(from, to)
	return ret
}

func CastMapUintptrE(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapUintptrBoolE(from) // Bool
	case reflect.Int:
		return ToMapUintptrIntE(from) // Int
	case reflect.Int8:
		return ToMapUintptrInt8E(from) // Int8
	case reflect.Int16:
		return ToMapUintptrInt16E(from) // Int16
	case reflect.Int32:
		return ToMapUintptrInt32E(from) // Int32
	case reflect.Int64:
		return ToMapUintptrInt64E(from) // Int64
	case reflect.Uint:
		return ToMapUintptrUintE(from) // Uint
	case reflect.Uint8:
		return ToMapUintptrUint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapUintptrUint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapUintptrUint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapUintptrUint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapUintptrUintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapUintptrFloat32E(from) // Float32
	case reflect.Float64:
		return ToMapUintptrFloat64E(from) // Float64
	case reflect.Complex64:
		return ToMapUintptrComplex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapUintptrComplex128E(from) // Complex128
	case reflect.Interface:
		return ToMapUintptrInterfaceE(from) // Interface
	case reflect.String:
		return ToMapUintptrStringE(from) // String
	}
}
