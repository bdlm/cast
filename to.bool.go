package cast

import (
	"strconv"

	"github.com/bdlm/errors/v2"
)

// toBool casts an interface to a bool type. For numeric values, anything but
// zero is considered true.
func toBool[TTo bool](from any) (TTo, error) {
	switch from.(type) {
	case bool:
		return from.(bool) != false, nil
	case byte:
		return from.(byte) != 0 && from.(byte) != '0', nil
	case complex64:
		return from.(complex64) != 0, nil
	case complex128:
		return from.(complex128) != 0, nil
	case float32:
		return from.(float32) != 0, nil
	case float64:
		return from.(float64) != 0, nil
	case int:
		return from.(int) != 0, nil
	case int8:
		return from.(int8) != 0, nil
	case int16:
		return from.(int16) != 0, nil
	case int32: // rune
		return from.(rune) != 0 && from.(rune) != '0', nil
	case int64:
		return from.(int64) != 0, nil
	case nil:
		return false, nil
	case string:
		r, e := strconv.ParseBool(from.(string))
		if nil != e {
			i, e2 := ToE[float64](from)
			if nil != e2 {
				return false, errors.Wrap(errors.WrapE(e2, e), ErrorStrUnableToCast, from, from, false)
			}
			return i != 0, nil
		}
		return r != false, nil
	}

	return false, errors.Errorf(ErrorStrUnableToCast, from, from, false)
}
