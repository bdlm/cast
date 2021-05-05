package cast

import (
	"github.com/bdlm/errors/v2"
)

// ToMapStringBool casts an interface to a map[string]bool type, discarding
// any errors.
func ToMapStringBool(i interface{}) map[string]bool {
	ret, _ := ToMapStringBoolE(i)
	return ret
}

// ToMapStringBoolE casts an interface to a map[string]bool type.
func ToMapStringBoolE(i interface{}) (map[string]bool, error) {
	var ret = map[string]bool{}
	var err error

	cast := func(k, v interface{}) (string, bool, error) {
		key, err := ToStringE(k)
		if nil != err {
			return "", false, err
		}
		val, err := ToBoolE(v)
		if nil != err {
			return "", false, err
		}
		return key, val, nil
	}

	switch t := i.(type) {
	case map[string]bool:
		ret, err = t, nil
	case map[interface{}]bool:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[interface{}]string:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[interface{}]interface{}:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[string]string:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[string]interface{}:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[int]bool:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[int]string:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[int]interface{}:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[uint]bool:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[uint]string:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case map[uint]interface{}:
		for key, val := range t {
			k, v, err := cast(key, val)
			if nil != err {
				break
			}
			ret[k] = v
		}
	case string:
		err = jsonStringToObject(t, &ret)
	default:
		ret, err = nil, errors.Errorf("unable to cast %#v of type %T to map[string]bool", i, i)
	}

	if nil != err {
		return nil, errors.Wrap(err, "unable to cast %#v of type %T to map[string]bool", i, i)
	}
	return ret, nil
}
