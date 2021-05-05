package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapUint8(from, to interface{}) interface{} {
	ret, _ := CastMapUint8E(from, to)
	return ret
}

func CastMapUint8E(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapUint8BoolE(from) // Bool
	case reflect.Int:
		return ToMapUint8IntE(from) // Int
	case reflect.Int8:
		return ToMapUint8Int8E(from) // Int8
	case reflect.Int16:
		return ToMapUint8Int16E(from) // Int16
	case reflect.Int32:
		return ToMapUint8Int32E(from) // Int32
	case reflect.Int64:
		return ToMapUint8Int64E(from) // Int64
	case reflect.Uint:
		return ToMapUint8UintE(from) // Uint
	case reflect.Uint8:
		return ToMapUint8Uint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapUint8Uint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapUint8Uint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapUint8Uint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapUint8UintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapUint8Float32E(from) // Float32
	case reflect.Float64:
		return ToMapUint8Float64E(from) // Float64
	case reflect.Complex64:
		return ToMapUint8Complex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapUint8Complex128E(from) // Complex128
	case reflect.Interface:
		return ToMapUint8InterfaceE(from) // Interface
	case reflect.String:
		return ToMapUint8StringE(from) // String
	}
}
