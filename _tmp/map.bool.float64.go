package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// ToMapBoolFloat64 casts an interface to a map[bool]float64 type, discarding
// any errors.
func ToMapBoolFloat64(i interface{}) map[bool]float64 {
	ret, _ := ToMapBoolFloat64E(i)
	return ret
}

// ToMapBoolFloat64E casts an interface to a map[bool]float64 type.
func ToMapBoolFloat64E(from interface{}) (map[bool]float64, error) {
	var ret = map[bool]float64{}
	var err error
	t := reflect.TypeOf(from)

	// Parse JSON strings.
	if reflect.String == t.Kind() {
		err = jsonStringToObject(t, &ret)
		return ret, err
	}

	// Reject non-map types.
	if reflect.Map != t.Kind() {
		return nil, errors.Errorf("cannot cast value %#v of type %T as map[bool]float64", from, from)
	}

	// Create map
	for k, v := range reflect.ValueOf(from).MapKeys() {
		key, err := ToBoolE(k)
		if nil != err {
			return nil, err
		}
		val, err := ToFloat64E(v)
		if nil != err {
			return nil, err
		}
		ret[key] = val
	}

	return ret, nil
}
