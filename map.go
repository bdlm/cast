package cast

import (
	"fmt"
)

/*
ToStringMapString casts an interface to a map[string]string type.
*/
func ToStringMapString(i interface{}) (map[string]string, error) {
	var m = map[string]string{}

	switch t := i.(type) {
	case map[string]string:
		return t, nil
	case map[string]interface{}:
		for key, val := range t {
			v, _ := ToString(val)
			k, _ := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}]string:
		for key, val := range t {
			v, _ := ToString(val)
			k, _ := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}]interface{}:
		for key, val := range t {
			v, _ := ToString(val)
			k, _ := ToString(key)
			m[k] = v
		}
		return m, nil
	case string:
		err := jsonStringToObject(t, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]string", i, i)
	}
}

/*
ToStringMapStringSlice casts an interface to a map[string][]string type.
*/
func ToStringMapStringSlice(i interface{}) (map[string][]string, error) {
	var m = map[string][]string{}

	switch t := i.(type) {
	case map[string][]string:
		return t, nil
	case map[string][]interface{}:
		for key, val := range t {
			v, _ := ToStringSlice(val)
			k, _ := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[string]string:
		for key, val := range t {
			k, _ := ToString(key)
			m[k] = []string{val}
		}
	case map[string]interface{}:
		for key, val := range t {
			k, _ := ToString(key)
			switch vt := val.(type) {
			case []interface{}:
				vs, _ := ToStringSlice(vt)
				m[k] = vs
			case []string:
				m[k] = vt
			default:
				vs, _ := ToString(val)
				m[k] = []string{vs}
			}
		}
		return m, nil
	case map[interface{}][]string:
		for key, val := range t {
			v, _ := ToStringSlice(val)
			k, _ := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}]string:
		for key, val := range t {
			v, _ := ToStringSlice(val)
			k, _ := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}][]interface{}:
		for key, val := range t {
			v, _ := ToStringSlice(val)
			k, _ := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}]interface{}:
		for key, val := range t {
			k, err := ToString(key)
			if err != nil {
				return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
			}
			v, err := ToStringSlice(val)
			if err != nil {
				return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
			}
			m[k] = v
		}
	case string:
		err := jsonStringToObject(t, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
	}
	return m, nil
}

/*
ToStringMapBool casts an interface to a map[string]bool type.
*/
func ToStringMapBool(i interface{}) (map[string]bool, error) {
	var m = map[string]bool{}

	switch t := i.(type) {
	case map[interface{}]interface{}:
		for key, val := range t {
			k, _ := ToString(key)
			v, _ := ToBool(val)
			m[k] = v
		}
		return m, nil
	case map[string]interface{}:
		for key, val := range t {
			k, _ := ToString(key)
			v, _ := ToBool(val)
			m[k] = v
		}
		return m, nil
	case map[string]bool:
		return t, nil
	case string:
		err := jsonStringToObject(t, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]bool", i, i)
	}
}

/*
ToStringMap casts an interface to a map[string]interface{} type.
*/
func ToStringMap(i interface{}) (map[string]interface{}, error) {
	var m = map[string]interface{}{}

	switch t := i.(type) {
	case map[interface{}]interface{}:
		for key, val := range t {
			k, _ := ToString(key)
			m[k] = val
		}
		return m, nil
	case map[string]interface{}:
		return t, nil
	case string:
		err := jsonStringToObject(t, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]interface{}", i, i)
	}
}
