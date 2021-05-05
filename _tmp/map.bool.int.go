package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// ToMapBoolInt casts an interface to a map[bool]int type, discarding
// any errors.
func ToMapBoolInt(i interface{}) map[bool]int {
	ret, _ := ToMapBoolIntE(i)
	return ret
}

// ToMapBoolIntE casts an interface to a map[bool]int type.
func ToMapBoolIntE(from interface{}) (map[bool]int, error) {
	var ret = map[bool]int{}
	var err error
	t := reflect.TypeOf(from)

	// Parse JSON strings.
	if reflect.String == t.Kind() {
		err = jsonStringToObject(t, &ret)
		return ret, err
	}

	// Reject non-map types.
	if reflect.Map != t.Kind() {
		return nil, errors.Errorf("cannot cast value %#v of type %T as map[bool]int", from, from)
	}

	// Create map
	for k, v := range reflect.ValueOf(from).MapKeys() {
		key, err := ToBoolE(k)
		if nil != err {
			return nil, err
		}
		val, err := ToIntE(v)
		if nil != err {
			return nil, err
		}
		ret[key] = val
	}

	return ret, nil
}
