package cast

import (
	"strconv"

	"github.com/bdlm/errors/v2"
)

// toBool casts an interface to a bool type. For numeric values, anything but
// zero is considered true.
//
// Options:
//   - DEFAULT: bool, default false. Default return value on error.
func toBool[TTo bool](from any, ops Ops) (TTo, error) {
	var ret TTo
	var ok bool

	if _, ok = ops[DEFAULT]; ok {
		if ret, ok = ops[DEFAULT].(TTo); !ok {
			return false, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
		}
	}

	switch from := from.(type) {
	case bool:
		return TTo(from), nil
	case byte:
		return from != 0 && from != '0', nil
	case complex64:
		return from != 0, nil
	case complex128:
		return from != 0, nil
	case float32:
		return from != 0, nil
	case float64:
		return from != 0, nil
	case int:
		return from != 0, nil
	case int8:
		return from != 0, nil
	case int16:
		return from != 0, nil
	case int32: // rune
		return from != 0 && from != '0', nil
	case int64:
		return from != 0, nil
	case nil:
		return false, nil
	case string:
		r, e := strconv.ParseBool(from)
		if nil != e {
			i, e2 := ToE[int](from)
			if nil != e2 {
				return ret, errors.Wrap(errors.WrapE(e2, e), ErrorStrUnableToCast, from, from, false)
			}
			return i != 0, nil
		}
		return TTo(r), nil
	}

	return false, errors.Errorf(ErrorStrUnableToCast, from, from, false)
}
