package cast

import (
	"fmt"
	"strconv"

	"github.com/bdlm/errors/v2"
)

// toBool casts an interface to a bool type. For numeric values, anything but
// zero is considered true. The string path delegates to [strconv.ParseBool]
// first, then falls back to integer parsing so "1"→true and "0"→false.
//
// Options:
//   - DEFAULT: bool, default false. Default return value on error.
func toBool[TTo bool](from any, ops ops) (TTo, error) {
	var ret TTo
	var ok bool

	if ops.hasDefault {
		if ret, ok = ops.defaultVal.(TTo); !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
	}

	switch from := from.(type) {
	case bool:
		return TTo(from), nil
	case byte:
		// The character '0' (ASCII 48) represents the digit zero, so it maps to false.
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
	case int32: // rune: '0' (48) is the digit zero, not a true value
		return from != 0 && from != '0', nil
	case int64:
		return from != 0, nil
	case uint:
		return from != 0, nil
	case uint16:
		return from != 0, nil
	case uint32:
		return from != 0, nil
	case uint64:
		return from != 0, nil
	case uintptr:
		return from != 0, nil
	case nil:
		return false, nil
	case fmt.Stringer:
		return toBool[TTo](from.String(), ops)
	case error:
		return toBool[TTo](from.Error(), ops)
	case string:
		r, e := strconv.ParseBool(from)
		if nil != e {
			i, e2 := ToE[int](from)
			if nil != e2 {
				if decoded, ok, _ := tryDecodeJSON(from, ops); ok {
					return toBool[TTo](decoded, ops.Delete(DECODE))
				}
				return ret, errors.Wrap(errors.WrapE(e2, e), ErrorStrUnableToCast, from, from, false)
			}
			return i != 0, nil
		}
		return TTo(r), nil
	default:
		if s, err := toString(from, ops.Delete(DEFAULT)); err == nil {
			return toBool[TTo](s.(string), ops)
		}
		return ret, errors.Errorf(ErrorStrUnableToCast, from, from, false)
	}
}
