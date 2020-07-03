package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapInt8(from, to interface{}) interface{} {
	ret, _ := CastMapInt8E(from, to)
	return ret
}

func CastMapInt8E(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapInt8BoolE(from) // Bool
	case reflect.Int:
		return ToMapInt8IntE(from) // Int
	case reflect.Int8:
		return ToMapInt8Int8E(from) // Int8
	case reflect.Int16:
		return ToMapInt8Int16E(from) // Int16
	case reflect.Int32:
		return ToMapInt8Int32E(from) // Int32
	case reflect.Int64:
		return ToMapInt8Int64E(from) // Int64
	case reflect.Uint:
		return ToMapInt8UintE(from) // Uint
	case reflect.Uint8:
		return ToMapInt8Uint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapInt8Uint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapInt8Uint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapInt8Uint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapInt8UintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapInt8Float32E(from) // Float32
	case reflect.Float64:
		return ToMapInt8Float64E(from) // Float64
	case reflect.Complex64:
		return ToMapInt8Complex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapInt8Complex128E(from) // Complex128
	case reflect.Interface:
		return ToMapInt8InterfaceE(from) // Interface
	case reflect.String:
		return ToMapInt8StringE(from) // String
	}
}
