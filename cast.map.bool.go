package cast

/*
import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMapBool(from, to interface{}) interface{} {
	ret, _ := CastMapBoolE(from, to)
	return ret
}

func CastMapBoolE(from, to interface{}) (interface{}, error) {
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	switch reflect.Type.Elem(t).Kind() {
	case reflect.Bool:
		return ToMapBoolBoolE(from) // Bool
	case reflect.Int:
		return ToMapBoolIntE(from) // Int
		//	case reflect.Int8:
		//		return ToMapBoolInt8E(from) // Int8
		//	case reflect.Int16:
		//		return ToMapBoolInt16E(from) // Int16
		//	case reflect.Int32:
		//		return ToMapBoolInt32E(from) // Int32
		//	case reflect.Int64:
		//		return ToMapBoolInt64E(from) // Int64
		//	case reflect.Uint:
		//		return ToMapBoolUintE(from) // Uint
		//	case reflect.Uint8:
		//		return ToMapBoolUint8E(from) // Uint8
		//	case reflect.Uint16:
		//		return ToMapBoolUint16E(from) // Uint16
		//	case reflect.Uint32:
		//		return ToMapBoolUint32E(from) // Uint32
		//	case reflect.Uint64:
		//		return ToMapBoolUint64E(from) // Uint64
		//	case reflect.Uintptr:
		//		return ToMapBoolUintptrE(from) // Uintptr
		//	case reflect.Float32:
		//		return ToMapBoolFloat32E(from) // Float32
		//	case reflect.Float64:
		//		return ToMapBoolFloat64E(from) // Float64
		//	case reflect.Complex64:
		//		return ToMapBoolComplex64E(from) // Complex64
		//	case reflect.Complex128:
		//		return ToMapBoolComplex128E(from) // Complex128
		//	case reflect.Interface:
		//		return ToMapBoolInterfaceE(from) // Interface
		//	case reflect.String:
		//		return ToMapBoolStringE(from) // String
	}
	return nil, nil
}
*/
