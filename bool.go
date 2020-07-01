package cast

import (
	"strconv"

	"github.com/bdlm/errors/v2"
)

// ToBool casts an interface to a bool type, discarding any errors.
func ToBool(i interface{}) bool {
	ret, _ := ToBoolE(i)
	return ret
}

// ToBoolE casts an interface to a bool type.
func ToBoolE(i interface{}) (bool, error) {
	var ret bool
	var err error

	i = indirect(i)
	switch t := i.(type) {
	case bool:
		ret, err = t, nil
	case nil:
		ret, err = false, nil
	case int:
		if i.(int) != 0 {
			ret, err = true, nil
		}
	case string:
		ret, err = strconv.ParseBool(i.(string))
	default:
		err = errors.Errorf("type %T cannot be cast to bool", i)
	}

	if nil != err {
		return false, errors.Wrap(err, "unable to cast %#v of type %T to bool", i, i)
	}
	return ret, nil
}
