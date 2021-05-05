package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// ToMapBoolComplex128 casts an interface to a map[bool]complex128 type, discarding
// any errors.
func ToMapBoolComplex128(i interface{}) map[bool]complex128 {
	ret, _ := ToMapBoolComplex128E(i)
	return ret
}

// ToMapBoolComplex128E casts an interface to a map[bool]complex128 type.
func ToMapBoolComplex128E(from interface{}) (map[bool]complex128, error) {
	var ret = map[bool]complex128{}
	var err error
	t := reflect.TypeOf(from)

	// Parse JSON strings.
	if reflect.String == t.Kind() {
		err = jsonStringToObject(t, &ret)
		return ret, err
	}

	// Reject non-map types.
	if reflect.Map != t.Kind() {
		return nil, errors.Errorf("cannot cast value %#v of type %T as map[bool]complex128", from, from)
	}

	// Create map
	for k, v := range reflect.ValueOf(from).MapKeys() {
		key, err := ToBoolE(k)
		if nil != err {
			return nil, err
		}
		val, err := ToComplex128E(v)
		if nil != err {
			return nil, err
		}
		ret[key] = val
	}

	return ret, nil
}
