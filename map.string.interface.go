package cast

import (
	"github.com/bdlm/errors/v2"
)

// ToMapStringInterface casts an interface to a map[string]interface{} type,
// discarding any errors.
func ToMapStringInterface(i interface{}) map[string]interface{} {
	ret, _ := ToMapStringInterfaceE(i)
	return ret
}

// ToMapStringInterfaceE casts an interface to a map[string]interface{} type.
func ToMapStringInterfaceE(i interface{}) (map[string]interface{}, error) {
	var ret = map[string]interface{}{}
	var err error

	cast := func(k, v interface{}) (string, interface{}, error) {
		key, err := ToStringE(k)
		if nil != err {
			return "", nil, err
		}
		return key, v, nil
	}

	switch t := i.(type) {
	case map[string]interface{}:
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
