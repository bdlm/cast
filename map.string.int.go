package cast

import (
	"github.com/bdlm/errors/v2"
)

// ToMapStringInt casts an interface to a map[string]int type,
// discarding any errors.
func ToMapStringInt(i interface{}) map[string]int {
	ret, _ := ToMapStringIntE(i)
	return ret
}

// ToMapStringIntE casts an interface to a map[string]int type.
func ToMapStringIntE(i interface{}) (map[string]int, error) {
	var ret = map[string]int{}
	var err error

	cast := func(k, v interface{}) (string, int, error) {
		key, err := ToStringE(k)
		if nil != err {
			return "", 0, err
		}
		val, err := ToIntE(v)
		if nil != err {
			return "", 0, err
		}
		return key, val, nil
	}

	switch t := i.(type) {
	case map[string]int:
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
