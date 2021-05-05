package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapUint64(from, to interface{}) interface{} {
	ret, _ := CastMapUint64E(from, to)
	return ret
}

func CastMapUint64E(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapUint64BoolE(from) // Bool
	case reflect.Int:
		return ToMapUint64IntE(from) // Int
	case reflect.Int8:
		return ToMapUint64Int8E(from) // Int8
	case reflect.Int16:
		return ToMapUint64Int16E(from) // Int16
	case reflect.Int32:
		return ToMapUint64Int32E(from) // Int32
	case reflect.Int64:
		return ToMapUint64Int64E(from) // Int64
	case reflect.Uint:
		return ToMapUint64UintE(from) // Uint
	case reflect.Uint8:
		return ToMapUint64Uint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapUint64Uint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapUint64Uint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapUint64Uint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapUint64UintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapUint64Float32E(from) // Float32
	case reflect.Float64:
		return ToMapUint64Float64E(from) // Float64
	case reflect.Complex64:
		return ToMapUint64Complex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapUint64Complex128E(from) // Complex128
	case reflect.Interface:
		return ToMapUint64InterfaceE(from) // Interface
	case reflect.String:
		return ToMapUint64StringE(from) // String
	}
}
