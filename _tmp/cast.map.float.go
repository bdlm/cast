package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapFloat32(from, to interface{}) interface{} {
	ret, _ := CastMapFloat32E(from, to)
	return ret
}

func CastMapFloat32E(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapFloat32BoolE(from) // Bool
	case reflect.Int:
		return ToMapFloat32IntE(from) // Int
	case reflect.Int8:
		return ToMapFloat32Int8E(from) // Int8
	case reflect.Int16:
		return ToMapFloat32Int16E(from) // Int16
	case reflect.Int32:
		return ToMapFloat32Int32E(from) // Int32
	case reflect.Int64:
		return ToMapFloat32Int64E(from) // Int64
	case reflect.Uint:
		return ToMapFloat32UintE(from) // Uint
	case reflect.Uint8:
		return ToMapFloat32Uint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapFloat32Uint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapFloat32Uint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapFloat32Uint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapFloat32UintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapFloat32Float32E(from) // Float32
	case reflect.Float64:
		return ToMapFloat32Float64E(from) // Float64
	case reflect.Complex64:
		return ToMapFloat32Complex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapFloat32Complex128E(from) // Complex128
	case reflect.Interface:
		return ToMapFloat32InterfaceE(from) // Interface
	case reflect.String:
		return ToMapFloat32StringE(from) // String
	}
}

func CastMapFloat64(from, to interface{}) interface{} {
	ret, _ := CastMapFloat64E(from, to)
	return ret
}

func CastMapFloat64E(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapFloat64BoolE(from) // Bool
	case reflect.Int:
		return ToMapFloat64IntE(from) // Int
	case reflect.Int8:
		return ToMapFloat64Int8E(from) // Int8
	case reflect.Int16:
		return ToMapFloat64Int16E(from) // Int16
	case reflect.Int32:
		return ToMapFloat64Int32E(from) // Int32
	case reflect.Int64:
		return ToMapFloat64Int64E(from) // Int64
	case reflect.Uint:
		return ToMapFloat64UintE(from) // Uint
	case reflect.Uint8:
		return ToMapFloat64Uint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapFloat64Uint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapFloat64Uint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapFloat64Uint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapFloat64UintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapFloat64Float32E(from) // Float32
	case reflect.Float64:
		return ToMapFloat64Float64E(from) // Float64
	case reflect.Complex64:
		return ToMapFloat64Complex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapFloat64Complex128E(from) // Complex128
	case reflect.Interface:
		return ToMapFloat64InterfaceE(from) // Interface
	case reflect.String:
		return ToMapFloat64StringE(from) // String
	}
}
