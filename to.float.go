package cast

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bdlm/errors/v2"
	"golang.org/x/exp/constraints"
)

// toFloat casts an interface to a float type.
//
// Options:
//   - DEFAULT: float32 or float64, default 0.0. Default return value on error.
func toFloat[TTo constraints.Float](from reflect.Value, ops Ops) (TTo, error) {
	var ret TTo
	var ok bool

	if _, ok = ops[DEFAULT]; ok {
		if ret, ok = ops[DEFAULT].(TTo); !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
		}
	}

	fromVal := reflect.ValueOf(from)
	if !fromVal.IsValid() || !fromVal.CanInterface() {
		return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, TTo(0))
	}

	to := reflect.Indirect(reflect.ValueOf(new(TTo)))

	switch typ := from.Interface().(type) {
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
		return strToFloat[TTo](to, typ.String())
	case string:
		return strToFloat[TTo](to, typ)
	}

	ret, err := toFloat[TTo](reflect.ValueOf(fmt.Sprintf("%v", from.Interface())), ops)
	if nil != err {
		return ret, errors.Wrap(err, ErrorStrUnableToCast, from.Interface(), from.Interface(), to.Interface())
	}
	return ret, nil
}

// strToFloat converts a string to a float type.
func strToFloat[TTo constraints.Float](to reflect.Value, from string) (TTo, error) {
	var e, err error
	var val any
	var bitSize = 64

	if to.Type().Kind() == reflect.Float32 {
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
			_, e := strconv.ParseComplex(from, bitSize)
			if e != nil {
				err = errors.WrapE(err, e)
			} else {
				err = errors.Wrap(err, ErrorStrUnableToCast, from, from, to.Interface())
				val = float64(0)
			}
		} else {
			err = nil
		}
	}
	if to.Type().Kind() == reflect.Float32 {
		val = float32(val.(float64))
	}
	return val.(TTo), err
}
