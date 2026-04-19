package cast

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bdlm/errors/v2"
)

// toFloat casts an interface to a float type.
//
// Options:
//   - DEFAULT: float32 or float64, default 0.0. Default return value on error.
func toFloat[TTo float](from any, ops Ops) (TTo, error) {
	var defaultValue TTo
	var ok bool

	if _, ok = ops[DEFAULT]; ok {
		if defaultValue, ok = ops[DEFAULT].(TTo); !ok {
			return defaultValue, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
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
	case fmt.Stringer:
		return strToFloat[TTo](typ.String())
	case string:
		return strToFloat[TTo](typ)
	}

	// Fall back to string conversion for any other type (e.g. named numerics,
	// structs with a String method that were not matched above).
	result, err := toFloat[TTo](fmt.Sprintf("%v", from), ops)
	if nil != err {
		return defaultValue, errors.Wrap(err, ErrorStrUnableToCast, from, from, TTo(0))
	}
	return result, nil
}

// strToFloat converts a string to a float type. On initial parse failure it
// strips whitespace, drops everything after the first decimal point, and
// removes comma thousands-separators, then retries (e.g. "1,234.56" → "1234").
func strToFloat[TTo float](from string) (TTo, error) {
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
			err = errors.WrapE(err, e)
			return TTo(0), err
		}
		val = val2
		err = nil
	}

	if isFloat32 {
		return TTo(float32(val)), nil
	}
	return TTo(val), err
}
