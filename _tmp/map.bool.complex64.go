package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// ToMapBoolComplex64 casts an interface to a map[bool]complex64 type, discarding
// any errors.
func ToMapBoolComplex64(i interface{}) map[bool]complex64 {
	ret, _ := ToMapBoolComplex64E(i)
	return ret
}

// ToMapBoolComplex64E casts an interface to a map[bool]complex64 type.
func ToMapBoolComplex64E(from interface{}) (map[bool]complex64, error) {
	var ret = map[bool]complex64{}
	var err error
	t := reflect.TypeOf(from)

	// Parse JSON strings.
	if reflect.String == t.Kind() {
		err = jsonStringToObject(t, &ret)
		return ret, err
	}

	// Reject non-map types.
	if reflect.Map != t.Kind() {
		return nil, errors.Errorf("cannot cast value %#v of type %T as map[bool]complex64", from, from)
	}

	// Create map
	for k, v := range reflect.ValueOf(from).MapKeys() {
		key, err := ToBoolE(k)
		if nil != err {
			return nil, err
		}
		val, err := ToComplex64E(v)
		if nil != err {
			return nil, err
		}
		ret[key] = val
	}

	return ret, nil
}
