package cast

import (
	"fmt"
)

// ToStringMapString casts an interface to a map[string]string type,
// discarding any errors.
func ToStringMapString(i interface{}) map[string]string {
	ret, _ := ToStringMapStringE(i)
	return ret
}

// ToStringMapStringE casts an interface to a map[string]string type.
func ToStringMapStringE(i interface{}) (map[string]string, error) {
	var m = map[string]string{}

	switch t := i.(type) {
	case map[string]string:
		return t, nil
	case map[string]interface{}:
		for key, val := range t {
			v := ToString(val)
			k := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}]string:
		for key, val := range t {
			v := ToString(val)
			k := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}]interface{}:
		for key, val := range t {
			v := ToString(val)
			k := ToString(key)
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

// ToStringMapStringSlice casts an interface to a map[string][]string type,
// discarding any errors.
func ToStringMapStringSlice(i interface{}) map[string][]string {
	ret, _ := ToStringMapStringSliceE(i)
	return ret
}

// ToStringMapStringSliceE casts an interface to a map[string][]string type.
func ToStringMapStringSliceE(i interface{}) (map[string][]string, error) {
	var m = map[string][]string{}

	switch t := i.(type) {
	case map[string][]string:
		return t, nil
	case map[string][]interface{}:
		for key, val := range t {
			v := ToStringSlice(val)
			k := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[string]string:
		for key, val := range t {
			k := ToString(key)
			m[k] = []string{val}
		}
	case map[string]interface{}:
		for key, val := range t {
			k := ToString(key)
			switch vt := val.(type) {
			case []interface{}:
				vs := ToStringSlice(vt)
				m[k] = vs
			case []string:
				m[k] = vt
			default:
				vs := ToString(val)
				m[k] = []string{vs}
			}
		}
		return m, nil
	case map[interface{}][]string:
		for key, val := range t {
			v := ToStringSlice(val)
			k := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}]string:
		for key, val := range t {
			v := ToStringSlice(val)
			k := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}][]interface{}:
		for key, val := range t {
			v := ToStringSlice(val)
			k := ToString(key)
			m[k] = v
		}
		return m, nil
	case map[interface{}]interface{}:
		for key, val := range t {
			k, err := ToStringE(key)
			if err != nil {
				return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
			}
			v, err := ToStringSliceE(val)
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

// ToStringMapBool casts an interface to a map[string]bool type, discarding
// any errors.
func ToStringMapBool(i interface{}) map[string]bool {
	ret, _ := ToStringMapBoolE(i)
	return ret
}

// ToStringMapBoolE casts an interface to a map[string]bool type.
func ToStringMapBoolE(i interface{}) (map[string]bool, error) {
	var m = map[string]bool{}

	switch t := i.(type) {
	case map[interface{}]interface{}:
		for key, val := range t {
			k := ToString(key)
			v := ToBool(val)
			m[k] = v
		}
		return m, nil
	case map[string]interface{}:
		for key, val := range t {
			k := ToString(key)
			v := ToBool(val)
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

// ToStringMap casts an interface to a map[string]interface{} type,
// discarding any errors.
func ToStringMap(i interface{}) map[string]interface{} {
	ret, _ := ToStringMapE(i)
	return ret
}

// ToStringMapE casts an interface to a map[string]interface{} type.
func ToStringMapE(i interface{}) (map[string]interface{}, error) {
	var m = map[string]interface{}{}

	switch t := i.(type) {
	case map[interface{}]interface{}:
		for key, val := range t {
			k := ToString(key)
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
