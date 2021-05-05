package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapInt16(from, to interface{}) interface{} {
	ret, _ := CastMapInt16E(from, to)
	return ret
}

func CastMapInt16E(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapInt16BoolE(from) // Bool
	case reflect.Int:
		return ToMapInt16IntE(from) // Int
	case reflect.Int8:
		return ToMapInt16Int8E(from) // Int8
	case reflect.Int16:
		return ToMapInt16Int16E(from) // Int16
	case reflect.Int32:
		return ToMapInt16Int32E(from) // Int32
	case reflect.Int64:
		return ToMapInt16Int64E(from) // Int64
	case reflect.Uint:
		return ToMapInt16UintE(from) // Uint
	case reflect.Uint8:
		return ToMapInt16Uint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapInt16Uint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapInt16Uint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapInt16Uint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapInt16UintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapInt16Float32E(from) // Float32
	case reflect.Float64:
		return ToMapInt16Float64E(from) // Float64
	case reflect.Complex64:
		return ToMapInt16Complex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapInt16Complex128E(from) // Complex128
	case reflect.Interface:
		return ToMapInt16InterfaceE(from) // Interface
	case reflect.String:
		return ToMapInt16StringE(from) // String
	}
}
