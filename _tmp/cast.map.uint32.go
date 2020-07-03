package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapUint32(from, to interface{}) interface{} {
	ret, _ := CastMapUint32E(from, to)
	return ret
}

func CastMapUint32E(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapUint32BoolE(from) // Bool
	case reflect.Int:
		return ToMapUint32IntE(from) // Int
	case reflect.Int8:
		return ToMapUint32Int8E(from) // Int8
	case reflect.Int16:
		return ToMapUint32Int16E(from) // Int16
	case reflect.Int32:
		return ToMapUint32Int32E(from) // Int32
	case reflect.Int64:
		return ToMapUint32Int64E(from) // Int64
	case reflect.Uint:
		return ToMapUint32UintE(from) // Uint
	case reflect.Uint8:
		return ToMapUint32Uint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapUint32Uint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapUint32Uint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapUint32Uint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapUint32UintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapUint32Float32E(from) // Float32
	case reflect.Float64:
		return ToMapUint32Float64E(from) // Float64
	case reflect.Complex64:
		return ToMapUint32Complex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapUint32Complex128E(from) // Complex128
	case reflect.Interface:
		return ToMapUint32InterfaceE(from) // Interface
	case reflect.String:
		return ToMapUint32StringE(from) // String
	}
}
