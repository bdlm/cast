package cast

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/bdlm/errors/v2"
)

// toInt casts an interface to an int type.
//
// Options:
//   - DEFAULT: integer, default 0. Default return value on error.
//   - ABS: bool, default false. Return the absolute value of integers.
func toInt[TTo integer](from any, ops ops) (TTo, error) {
	var defaultValue TTo
	var ok bool

	if ops.hasDefault {
		if defaultValue, ok = ops.defaultVal.(TTo); !ok {
			return defaultValue, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
	}
	abs := ops.abs

	errDetail := errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, TTo(0))
	unsigned := false
	switch reflect.TypeOf(TTo(0)).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		unsigned = true
	}

	switch val := from.(type) {
	case nil:
		return TTo(0), nil
	case bool:
		if val {
			return TTo(1), nil
		}
		return TTo(0), nil
	case int:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return defaultValue, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case int64:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return defaultValue, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case int32:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return defaultValue, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case int16:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return defaultValue, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case int8:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return defaultValue, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case float64:
		if unsigned && val < 0 {
			if abs {
				return TTo(math.Floor(-val)), nil
			}
			return defaultValue, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		// Route through string to avoid float→int precision surprises; strToInt
		// handles truncation consistently via math.Floor.
		return strToInt[TTo](To[string](val), ops)
	case float32:
		if unsigned && val < 0 {
			if abs {
				return TTo(math.Floor(float64(-val))), nil
			}
			return defaultValue, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return strToInt[TTo](To[string](val), ops)
	case uint:
		return TTo(val), nil
	case uintptr:
		return TTo(val), nil
	case uint64:
		return TTo(val), nil
	case uint32:
		return TTo(val), nil
	case uint16:
		return TTo(val), nil
	case uint8:
		return TTo(val), nil
	case fmt.Stringer:
		return strToInt[TTo](val.String(), ops)
	case string:
		return strToInt[TTo](val, ops)
	case complex64:
		return toInt[TTo](float32(real(val)), ops)
	case complex128:
		return toInt[TTo](float64(real(val)), ops)
	default:
		return toInt[TTo](fmt.Sprintf("%v", from), ops)
	}
}

// strToInt converts a string to an integer type. It parses via float64 to
// handle decimal strings like "3.7" (truncated to 3). On failure it strips
// whitespace, drops the fractional part, and removes comma thousands-separators
// before retrying (e.g. "1,234.56 kg" fails; "1,234" → 1234 succeeds).
//
// Options:
//   - DEFAULT: integer, default 0. Default return value on error.
//   - ABS: bool, default false. Return the absolute value of negative integers
//     when casting to unsigned integers.
func strToInt[TTo integer](from string, ops ops) (TTo, error) {
	var defaultValue TTo
	var ok bool

	if ops.hasDefault {
		if defaultValue, ok = ops.defaultVal.(TTo); !ok {
			return defaultValue, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
	}
	abs := ops.abs

	errDetail := errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, TTo(0))
	var e, err error
	var val float64
	if val, e = strconv.ParseFloat(from, 64); e != nil {
		err = e
		stripped := strings.ReplaceAll(
			strings.Split(
				strings.Trim(from, "\r\n\t "),
				".",
			)[0],
			",", "",
		)
		if val, e = strconv.ParseFloat(stripped, 64); e != nil {
			err = errors.WrapE(err, e)
			return defaultValue, errors.WrapE(err, errDetail)
		}
		err = nil
	}
	if val < 0 && abs {
		val = -val
	}
	if err != nil {
		return defaultValue, errors.WrapE(err, errDetail)
	}
	if val >= 0 {
		return TTo(math.Floor(val)), nil
	}

	switch reflect.TypeOf(TTo(0)).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return defaultValue, errors.WrapE(ErrorSignedToUnsigned, errDetail)
	}

	return TTo(math.Ceil(val)), nil
}
