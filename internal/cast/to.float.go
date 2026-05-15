package cast

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ToFloat casts an interface to a float type.
//
// Options:
//   - DEFAULT: float32 or float64, default 0.0. Default return value on error.
func ToFloat[TTo Float](from any, ops Ops) (TTo, error) {
	var defaultValue TTo
	var ok bool

	if ops.HasDefault {
		if defaultValue, ok = ops.DefaultVal.(TTo); !ok {
			return defaultValue, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.DefaultVal)
		}
	}

	switch typ := from.(type) {
	case nil:
		return TTo(0), nil
	case bool:
		if typ {
			return TTo(1), nil
		}
		return TTo(0), nil
	case float64:
		return TTo(typ), nil
	case float32:
		return TTo(typ), nil
	case int:
		return TTo(typ), nil
	case int64:
		return TTo(typ), nil
	case int32:
		return TTo(typ), nil
	case int16:
		return TTo(typ), nil
	case int8:
		return TTo(typ), nil
	case uint:
		return TTo(typ), nil
	case uint64:
		return TTo(typ), nil
	case uint32:
		return TTo(typ), nil
	case uint16:
		return TTo(typ), nil
	case uint8:
		return TTo(typ), nil
	case uintptr:
		return TTo(typ), nil
	case complex64:
		return TTo(real(typ)), nil
	case complex128:
		return TTo(real(typ)), nil
	case time.Time:
		secs := float64(typ.Unix()) + float64(typ.Nanosecond())/1e9
		return TTo(secs), nil
	case fmt.Stringer:
		return ToFloat[TTo](typ.String(), ops)
	case string:
		result, err := strToFloat[TTo](typ)
		if err != nil {
			if ops.HasDecode && ops.DecodeVal == "json" {
				// Fast path: unmarshal into typed targets instead of any to avoid
				// the boxing overhead of tryDecodeJSON. JSON strings ("42") decode
				// into a Go string; JSON numbers (42) decode into json.Number.
				var s string
				if json.Unmarshal([]byte(typ), &s) == nil {
					return ToFloat[TTo](s, ops.Delete(DECODE))
				}
				var n json.Number
				if json.Unmarshal([]byte(typ), &n) == nil {
					return ToFloat[TTo](string(n), ops.Delete(DECODE))
				}
			}
			return defaultValue, err
		}
		return result, nil
	}

	if s, err := ToString(from, ops.Delete(DEFAULT)); err == nil {
		result, e := ToFloat[TTo](s, ops)
		if e == nil {
			return result, nil
		}
	}
	return defaultValue, fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0))
}

// strToFloat converts a string to a float type. On initial parse failure it
// strips whitespace, drops everything after the first decimal point, and
// removes comma thousands-separators, then retries (e.g. "1,234.56" → "1234").
func strToFloat[TTo Float](from string) (TTo, error) {
	isFloat32 := reflect.TypeOf(TTo(0)).Kind() == reflect.Float32
	bitSize := 64
	if isFloat32 {
		bitSize = 32
	}

	val, err := strconv.ParseFloat(from, bitSize)
	if err != nil {
		stripped := strings.ReplaceAll(
			strings.Split(
				strings.Trim(from, "\r\n\t "),
				".",
			)[0],
			",", "",
		)
		val2, e := strconv.ParseFloat(stripped, bitSize)
		if e != nil {
			err = fmt.Errorf("%w: %v", err, e)
			return TTo(0), err
		}
		val = val2
		err = nil
	}

	if isFloat32 {
		return TTo(float32(val)), nil
	}
	return TTo(val), nil
}
