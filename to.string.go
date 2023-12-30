package cast

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toString casts an interface to a string type.
func toString(from reflect.Value, ops Ops) (any, error) {
	var ret any
	var err error
	var ok bool

	switch from.Type().Kind() {
	case
		reflect.Bool,
		reflect.Complex128, reflect.Complex64,
		reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8,
		reflect.String,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Uintptr,
		reflect.Chan:
		ret = fmt.Sprintf("%v", from.Interface())
	default:
		if s, ok := from.Interface().(fmt.Stringer); ok || nil == from.Interface() {
			ret = fmt.Sprintf("%v", s)
		} else {
			b := []byte{}
			b, err = json.Marshal(from.Interface())
			ret = string(b)
		}
	}

	if err != nil {
		return ret, errors.WrapE(Error, err)
	}
	if ret, ok = ret.(string); ok {
		return ret, nil
	}
	if nil == ret {
		return "", nil
	}

	return ret, errors.WrapE(Error, errors.Errorf(ErrorStrUnableToCast, from, from.Interface(), ""))
}
