package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// ToMapBoolBool casts an interface to a map[bool]bool type, discarding
// any errors.
func ToMapBoolBool(i interface{}) map[bool]bool {
	ret, _ := ToMapBoolBoolE(i)
	return ret
}

// ToMapBoolBoolE casts an interface to a map[bool]bool type.
func ToMapBoolBoolE(from interface{}) (map[bool]bool, error) {
	var ret = map[bool]bool{}
	var err error
	t := reflect.TypeOf(from)

	// Parse JSON strings.
	if reflect.String == t.Kind() {
		err = jsonStringToObject(from.(string), &ret)
		return ret, err
	}

	// Reject non-map types.
	if reflect.Map != t.Kind() {
		return nil, errors.Errorf("cannot cast value %#v of type %T as map[bool]bool", from, from)
	}

	// Create map
	for k, v := range reflect.ValueOf(from).MapKeys() {
		key, err := ToBoolE(k)
		if nil != err {
			return nil, err
		}
		val, err := ToBoolE(v)
		if nil != err {
			return nil, err
		}
		ret[key] = val
	}

	return ret, nil
}

// ToMapBoolIntE casts an interface to a map[bool]int type.
func ToMapBoolIntE(from interface{}) (map[bool]int, error) {
	var ret = map[bool]int{}
	var err error
	t := reflect.TypeOf(from)

	// Parse JSON strings.
	if reflect.String == t.Kind() {
		err = jsonStringToObject(from.(string), &ret)
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
