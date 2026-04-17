package cast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bdlm/errors/v2"
	"golang.org/x/exp/constraints"
)

// toFloat casts an interface to a float type.
//
// Options:
//   - DEFAULT: float32 or float64, default 0.0. Default return value on error.
func toFloat[TTo constraints.Float](from any, ops Ops) (TTo, error) {
	var ret TTo
	var ok bool

	if _, ok = ops[DEFAULT]; ok {
		if ret, ok = ops[DEFAULT].(TTo); !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
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
	case fmt.Stringer:
		return strToFloat[TTo](typ.String())
	case string:
		return strToFloat[TTo](typ)
	}

	ret, err := toFloat[TTo](fmt.Sprintf("%v", from), ops)
	if nil != err {
		return ret, errors.Wrap(err, ErrorStrUnableToCast, from, from, TTo(0))
	}
	return ret, nil
}

// strToFloat converts a string to a float type.
func strToFloat[TTo constraints.Float](from string) (TTo, error) {
	var e, err error
	var val any
	var bitSize = 64

	switch any(TTo(0)).(type) {
	case float32:
		bitSize = 32
	}

	if val, e = strconv.ParseFloat(from, bitSize); e != nil {
		err = e
		from = strings.ReplaceAll(
			strings.Split(
				strings.Trim(from, "\r\n\t "),
				".",
			)[0],
			",", "",
		)
		if val, e = strconv.ParseFloat(from, bitSize); e != nil {
			err = errors.WrapE(err, e)
			_, e := strconv.ParseComplex(from, 64)
			if e != nil {
				err = errors.WrapE(err, e)
			} else {
				err = errors.Wrap(err, ErrorStrUnableToCast, from, from, TTo(0))
				val = float64(0)
			}
		} else {
			err = nil
		}
	}
	switch any(TTo(0)).(type) {
	case float32:
		val = float32(val.(float64))
	}
	return val.(TTo), err
}
