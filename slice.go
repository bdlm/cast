package cast

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// ToSlice casts an interface to a []interface{} type, discarding any errors.
func ToSlice(i interface{}) []interface{} {
	ret, _ := ToSliceE(i)
	return ret
}

// ToSliceE casts an interface to a []interface{} type.
func ToSliceE(i interface{}) ([]interface{}, error) {
	var s []interface{}

	switch t := i.(type) {
	case []interface{}:
		return append(s, t...), nil
	case []map[string]interface{}:
		for _, u := range t {
			s = append(s, u)
		}
		return s, nil
	default:
		return s, fmt.Errorf("unable to cast %#v of type %T to []interface{}", i, i)
	}
}

// ToBoolSlice casts an interface to a []bool type, discarding any errors.
func ToBoolSlice(i interface{}) []bool {
	ret, _ := ToBoolSliceE(i)
	return ret
}

// ToBoolSliceE casts an interface to a []bool type.
func ToBoolSliceE(i interface{}) ([]bool, error) {
	if i == nil {
		return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}

	switch v := i.(type) {
	case []bool:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]bool, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToBoolE(s.Index(j).Interface())
			if err != nil {
				return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}
}

// ToStringSlice casts an interface to a []string type, discarding any
// errors.
func ToStringSlice(i interface{}) []string {
	ret, _ := ToStringSliceE(i)
	return ret
}

// ToStringSliceE casts an interface to a []string type.
func ToStringSliceE(i interface{}) ([]string, error) {
	var a []string

	switch t := i.(type) {
	case []interface{}:
		for _, val := range t {
			v := ToString(val)
			a = append(a, v)
		}
		return a, nil
	case []string:
		return t, nil
	case string:
		return strings.Fields(t), nil
	case interface{}:
		str, err := ToStringE(t)
		if err != nil {
			return a, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
		}
		return []string{str}, nil
	default:
		return a, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
	}
}

// ToIntSlice casts an interface to a []int type, discarding any errors.
func ToIntSlice(i interface{}) []int {
	ret, _ := ToIntSliceE(i)
	return ret
}

// ToIntSliceE casts an interface to a []int type.
func ToIntSliceE(i interface{}) ([]int, error) {
	if i == nil {
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}

	switch v := i.(type) {
	case []int:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToIntE(s.Index(j).Interface())
			if err != nil {
				return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}
}

// ToInt64Slice casts an interface to a []int type, discarding any errors.
func ToInt64Slice(i interface{}) []int64 {
	ret, _ := ToInt64SliceE(i)
	return ret
}

// ToInt64SliceE casts an interface to a []int type.
func ToInt64SliceE(i interface{}) ([]int64, error) {
	if i == nil {
		return []int64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}

	switch v := i.(type) {
	case []int64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToInt64E(s.Index(j).Interface())
			if err != nil {
				return []int64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int64{}, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}
}

// ToDurationSlice casts an interface to a []time.Duration type, discarding
// any errors.
func ToDurationSlice(i interface{}) []time.Duration {
	ret, _ := ToDurationSliceE(i)
	return ret
}

// ToDurationSliceE casts an interface to a []time.Duration type.
func ToDurationSliceE(i interface{}) ([]time.Duration, error) {
	if i == nil {
		return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}

	switch v := i.(type) {
	case []time.Duration:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]time.Duration, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToDurationE(s.Index(j).Interface())
			if err != nil {
				return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}
}
