package cast

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
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
			return defaultValue, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
	}
	abs := ops.abs

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
				return TTo(-uint(val)), nil
			}
			return defaultValue, fmt.Errorf("%w: %w", ErrorSignedToUnsigned, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
		}
		return TTo(val), nil
	case int64:
		if unsigned && val < 0 {
			if abs {
				return TTo(-uint64(val)), nil
			}
			return defaultValue, fmt.Errorf("%w: %w", ErrorSignedToUnsigned, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
		}
		return TTo(val), nil
	case int32:
		if unsigned && val < 0 {
			if abs {
				return TTo(-uint32(val)), nil
			}
			return defaultValue, fmt.Errorf("%w: %w", ErrorSignedToUnsigned, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
		}
		return TTo(val), nil
	case int16:
		if unsigned && val < 0 {
			if abs {
				return TTo(-uint16(val)), nil
			}
			return defaultValue, fmt.Errorf("%w: %w", ErrorSignedToUnsigned, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
		}
		return TTo(val), nil
	case int8:
		if unsigned && val < 0 {
			if abs {
				return TTo(-uint8(val)), nil
			}
			return defaultValue, fmt.Errorf("%w: %w", ErrorSignedToUnsigned, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
		}
		return TTo(val), nil
	case float64:
		if unsigned && val < 0 {
			if abs {
				return TTo(math.Floor(-val)), nil
			}
			return defaultValue, fmt.Errorf("%w: %w", ErrorSignedToUnsigned, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
		}
		if val >= 0 {
			return TTo(math.Floor(val)), nil
		}
		return TTo(math.Ceil(val)), nil
	case float32:
		if unsigned && val < 0 {
			if abs {
				return TTo(math.Floor(float64(-val))), nil
			}
			return defaultValue, fmt.Errorf("%w: %w", ErrorSignedToUnsigned, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
		}
		v64 := float64(val)
		if v64 >= 0 {
			return TTo(math.Floor(v64)), nil
		}
		return TTo(math.Ceil(v64)), nil
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
	case time.Time:
		secs := val.Unix()
		if unsigned && secs < 0 {
			if abs {
				secs = -secs
			} else {
				return defaultValue, fmt.Errorf("%w: %w", ErrorSignedToUnsigned, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
			}
		}
		result := TTo(secs)
		if int64(result) != secs {
			return defaultValue, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0))
		}
		return result, nil
	case fmt.Stringer:
		return toInt[TTo](val.String(), ops)
	case string:
		result, err := strToInt[TTo](val, unsigned, ops)
		if err != nil {
			if ops.hasDecode && ops.decodeVal == "json" {
				// Fast path: unmarshal into typed targets instead of any to avoid
				// the boxing overhead of tryDecodeJSON. JSON strings ("42") decode
				// into a Go string; JSON numbers (42) decode into json.Number.
				var s string
				if json.Unmarshal([]byte(val), &s) == nil {
					return toInt[TTo](s, ops.Delete(DECODE))
				}
				var n json.Number
				if json.Unmarshal([]byte(val), &n) == nil {
					return toInt[TTo](string(n), ops.Delete(DECODE))
				}
			}
		}
		return result, err
	case complex64:
		return toInt[TTo](float32(real(val)), ops)
	case complex128:
		return toInt[TTo](float64(real(val)), ops)
	default:
		if s, err := toString(from, ops.Delete(DEFAULT)); err == nil {
			return strToInt[TTo](s, unsigned, ops)
		}
		return defaultValue, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0))
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
func strToInt[TTo integer](from string, unsigned bool, ops ops) (TTo, error) {
	var defaultValue TTo
	var ok bool

	if ops.hasDefault {
		if defaultValue, ok = ops.defaultVal.(TTo); !ok {
			return defaultValue, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
	}
	abs := ops.abs

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
			err = fmt.Errorf("%w: %w", err, e)
			return defaultValue, fmt.Errorf("%w: %w", err, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
		}
		err = nil
	}
	if val < 0 && abs {
		val = -val
	}
	if val >= 0 {
		return TTo(math.Floor(val)), nil
	}

	if unsigned {
		return defaultValue, fmt.Errorf("%w: %w", ErrorSignedToUnsigned, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0)))
	}

	return TTo(math.Ceil(val)), nil
}
