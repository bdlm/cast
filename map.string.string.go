package cast

import (
	"github.com/bdlm/errors/v2"
)

// ToMapStringString casts an interface to a map[string]string type,
// discarding any errors.
func ToMapStringString(i interface{}) map[string]string {
	ret, _ := ToMapStringStringE(i)
	return ret
}

// ToMapStringStringE casts an interface to a map[string]string type.
func ToMapStringStringE(i interface{}) (map[string]string, error) {
	var ret = map[string]string{}
	var err error

	cast := func(k, v interface{}) (string, string, error) {
		key, err := ToStringE(k)
		if nil != err {
			return "", "", err
		}
		val, err := ToStringE(v)
		if nil != err {
			return "", "", err
		}
		return key, val, nil
	}

	switch t := i.(type) {
	case map[string]string:
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
	case map[string]bool:
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
		ret, err = nil, errors.Errorf("unable to cast %#v of type %T to map[string]string", i, i)
	}

	if nil != err {
		return nil, errors.Wrap(err, "unable to cast %#v of type %T to map[string]string", i, i)
	}
	return ret, nil
}
