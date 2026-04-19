package cast

import (
	"fmt"
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

	result, err := toFloat[TTo](fmt.Sprintf("%v", from), ops)
	if nil != err {
		return defaultValue, errors.Wrap(err, ErrorStrUnableToCast, from, from, TTo(0))
	}
	return result, nil
}

// strToFloat converts a string to a float type.
func strToFloat[TTo float](from string) (TTo, error) {
	_, isFloat32 := any(TTo(0)).(float32)
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
