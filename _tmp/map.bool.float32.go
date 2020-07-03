package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// ToMapBoolFloat32 casts an interface to a map[bool]float32 type, discarding
// any errors.
func ToMapBoolFloat32(i interface{}) map[bool]float32 {
	ret, _ := ToMapBoolFloat32E(i)
	return ret
}

// ToMapBoolFloat32E casts an interface to a map[bool]float32 type.
func ToMapBoolFloat32E(from interface{}) (map[bool]float32, error) {
	var ret = map[bool]float32{}
	var err error
	t := reflect.TypeOf(from)

	// Parse JSON strings.
	if reflect.String == t.Kind() {
		err = jsonStringToObject(t, &ret)
		return ret, err
	}

	// Reject non-map types.
	if reflect.Map != t.Kind() {
		return nil, errors.Errorf("cannot cast value %#v of type %T as map[bool]float32", from, from)
	}

	// Create map
	for k, v := range reflect.ValueOf(from).MapKeys() {
		key, err := ToBoolE(k)
		if nil != err {
			return nil, err
		}
		val, err := ToFloat32E(v)
		if nil != err {
			return nil, err
		}
		ret[key] = val
	}

	return ret, nil
}
