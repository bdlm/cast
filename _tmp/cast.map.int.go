package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapInt(from, to interface{}) interface{} {
	ret, _ := CastMapIntE(from, to)
	return ret
}

func CastMapIntE(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapIntBoolE(from) // Bool
	case reflect.Int:
		return ToMapIntIntE(from) // Int
	case reflect.Int8:
		return ToMapIntInt8E(from) // Int8
	case reflect.Int16:
		return ToMapIntInt16E(from) // Int16
	case reflect.Int32:
		return ToMapIntInt32E(from) // Int32
	case reflect.Int64:
		return ToMapIntInt64E(from) // Int64
	case reflect.Uint:
		return ToMapIntUintE(from) // Uint
	case reflect.Uint8:
		return ToMapIntUint8E(from) // Uint8
	case reflect.Uint16:
		return ToMapIntUint16E(from) // Uint16
	case reflect.Uint32:
		return ToMapIntUint32E(from) // Uint32
	case reflect.Uint64:
		return ToMapIntUint64E(from) // Uint64
	case reflect.Uintptr:
		return ToMapIntUintptrE(from) // Uintptr
	case reflect.Float32:
		return ToMapIntFloat32E(from) // Float32
	case reflect.Float64:
		return ToMapIntFloat64E(from) // Float64
	case reflect.Complex64:
		return ToMapIntComplex64E(from) // Complex64
	case reflect.Complex128:
		return ToMapIntComplex128E(from) // Complex128
	case reflect.Interface:
		return ToMapIntInterfaceE(from) // Interface
	case reflect.String:
		return ToMapIntStringE(from) // String
	}
}
