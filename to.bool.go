package cast

import (
	"strconv"

	"github.com/bdlm/errors/v2"
)

// ToBool casts an interface to a bool type.
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
	case float32, float64:
		return from.(float64) != 0, nil
	case int, int8, int16, int64:
		return from.(int) != 0, nil
	case nil:
		return false, nil
	case int32: // rune
		return from.(rune) != 0 && from.(rune) != '0', nil
	case string:
		r, e := strconv.ParseBool(from.(string))
		if nil != e {
			i, e2 := ToE[int](from)
			if nil != e2 {
				return false, errors.Wrap(errors.WrapE(e2, e), "unable to cast %#v of type %T to bool", from, from)
			}
			return i != 0, nil
		}
		return r != false, nil
	}

	return false, errors.Errorf("unable to cast %#v of type %T to %T", from, from, false)
}
