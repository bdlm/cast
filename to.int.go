package cast

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/bdlm/errors/v2"
	"golang.org/x/exp/constraints"
)

// toInt casts an interface to an int type.
//
// Options:
//   - DEFAULT: constraints.Integer, default 0. Default return value on error.
//   - ABS: bool, default false. Return the absolute value of integers.
func toInt[TTo constraints.Integer](from reflect.Value, ops Ops) (TTo, error) {
	var default_val TTo
	var ok bool
	var abs bool

	if _, ok = ops[DEFAULT]; ok {
		if default_val, ok = ops[DEFAULT].(TTo); !ok {
			return default_val, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
		}
	}
	if _, ok = ops[ABS]; ok {
		if abs, ok = ops[ABS].(bool); !ok {
			return default_val, errors.Errorf(ErrorInvalidOption, "ABS", ops[ABS])
		}
	}

	fromVal := reflect.Indirect(from)
	if !fromVal.IsValid() || !fromVal.CanInterface() {
		return default_val, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, TTo(0))
	}

	errDetail := errors.Errorf("unable to cast %#.10v of type %T to %T", from.Interface(), from.Interface(), TTo(0))
	to := reflect.Indirect(reflect.ValueOf(new(TTo)))
	unsigned := false
	switch to.Interface().(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr:
		unsigned = true
	}

	//fmt.Printf("\n\nto: %v (%T); from: %v (%T)\n\n", to, to.Interface(), from, from.Interface())
	switch val := from.Interface().(type) {
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
			return default_val, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case int64:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return default_val, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case int32:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return default_val, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case int16:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return default_val, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case int8:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return default_val, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(val), nil
	case float64:
		if unsigned && val < 0 {
			if abs {
				return TTo(-val), nil
			}
			return default_val, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return strToInt[TTo](To[string](from.Interface()), ops)
	case float32:
		if unsigned && val < 0 {
			if abs {
				return strToInt[TTo](To[string](-val), ops)
			}
			return default_val, errors.WrapE(ErrorSignedToUnsigned, errDetail)
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

	//case complex128:
	//case complex64:
	default:
		return toInt[TTo](reflect.ValueOf(fmt.Sprintf("%#.10v", from.Interface())), ops)
	}
}

// strToInt converts a string to an integer type.
//
// Options:
//   - DEFAULT: constraints.Integer, default 0. Default return value on error.
//   - ABS: bool, default false. Return the absolute value of negative integers
//     when casting to unsigned integers.
func strToInt[TTo constraints.Integer](from string, ops Ops) (TTo, error) {
	var default_val TTo
	var ok bool
	var abs bool

	if _, ok = ops[DEFAULT]; ok {
		if default_val, ok = ops[DEFAULT].(TTo); !ok {
			return default_val, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
		}
	}
	if _, ok = ops[ABS]; ok {
		if abs, ok = ops[ABS].(bool); !ok {
			return default_val, errors.Errorf(ErrorInvalidOption, "ABS", ops[ABS])
		}
	}

	errDetail := errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, TTo(0))
	var e, err error
	var val float64
	if val, e = strconv.ParseFloat(from, 64); e != nil {
		err = e
		from = strings.ReplaceAll(
			strings.Split(
				strings.Trim(from, "\r\n\t "),
				".",
			)[0],
			",", "",
		)
		if val, e = strconv.ParseFloat(from, 64); e != nil {
			err = errors.WrapE(err, e)
			_, e := strconv.ParseComplex(from, 64)
			if e != nil {
				err = errors.WrapE(err, e)
			} else {
				err = errors.Wrap(err, "unable to cast complex to int")
				val = float64(default_val)
			}
		} else {
			err = nil
		}
	}
	if val < 0 && abs {
		val = -val
	}
	if err != nil {
		return default_val, errors.WrapE(err, errDetail)
	}
	if val >= 0 {
		return TTo(math.Floor(val)), nil
	}

	// Negative to uint error.
	ref := reflect.ValueOf(TTo(0))
	switch ref.Interface().(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr:
		if val < 0 {
			return default_val, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
	}

	return TTo(math.Ceil(val)), nil
}
