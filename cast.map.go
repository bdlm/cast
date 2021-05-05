package cast

import (
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
)

func CastMap(from, to interface{}) interface{} {
	ret, _ := CastMapE(from, to)
	return ret
}

func CastMapE(from, to interface{}) (interface{}, error) {
	// Collapse reflection values.
	if _, ok := from.(reflect.Value); ok {
		from = from.(reflect.Value).Interface()
	}
	if _, ok := to.(reflect.Value); ok {
		to = to.(reflect.Value).Interface()
	}
	from, to = indirect(from), indirect(to)

	if reflect.Map != reflect.TypeOf(from).Kind() || reflect.Map != reflect.TypeOf(to).Kind() {
		return nil, errors.Errorf("cannot cast value of type %T as %T: invalid map type", from, to)
	}

	t := reflect.TypeOf(to)
	kto := reflect.Type.Key(t)
	vto := reflect.Type.Elem(t)

	toType := reflect.MapOf(kto, vto)
	retMap := reflect.MakeMapWithSize(toType, 0)

	fmt.Printf("CastMapE - from: %v (%T); to: %v (%T);\n", from, from, to, to)
	for _, k := range reflect.ValueOf(from).MapKeys() {
		key, err := CastE(k, reflect.Zero(kto))
		if nil != err {
			return nil, err
		}
		val, err := CastE(reflect.ValueOf(from).MapIndex(k), reflect.Zero(vto))
		if nil != err {
			return nil, err
		}
		retMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	}
	return retMap.Interface(), nil
}
