package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapInt64(from, to interface{}) interface{} {
	ret, _ := CastMapInt64E(from, to)
	return ret
}

func CastMapInt64E(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapInt64BoolE(from) // Bool
	case reflect.Int:
		return ToMapInt64IntE(from) // Int
	case reflect.Int8:
		return ToMapInt64Int8E(from) // Int8
	case reflect.Int16:
		return ToMapInt64Int16E(from) // Int16
	case reflect.Int32:
		return ToMapInt64Int32E(from) // Int32
	case reflect.Int64:
		return ToMapInt64Int64E(from) // Int64
	case reflect.Uint:
		return ToMapInt64UintE(from) // Uint
	case reflect.Uint8:
		return ToMapInt64Uint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapInt64Uint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapInt64Uint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapInt64Uint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapInt64UintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapInt64Float32E(from) // Float32
	case reflect.Float64:
		return ToMapInt64Float64E(from) // Float64
	case reflect.Complex64:
		return ToMapInt64Complex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapInt64Complex128E(from) // Complex128
	case reflect.Interface:
		return ToMapInt64InterfaceE(from) // Interface
	case reflect.String:
		return ToMapInt64StringE(from) // String
	}
}
