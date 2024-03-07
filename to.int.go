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
//   - ABS: bool, default false. Return the absolute value of negative integers
//     when casting to unsigned integers.
func toInt[TTo constraints.Integer](from reflect.Value, ops Ops) (TTo, error) {
	var ret TTo
	var ok bool
	var abs bool

	if _, ok = ops[DEFAULT]; ok {
		if ret, ok = ops[DEFAULT].(TTo); !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
		}
	}
	if _, ok = ops[ABS]; ok {
		if abs, ok = ops[ABS].(bool); !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "ABS", ops[ABS])
		}
	}

	fromVal := reflect.Indirect(from)
	if !fromVal.IsValid() || !fromVal.CanInterface() {
		return ret, errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, TTo(0))
	}

	errDetail := errors.Errorf("unable to cast %#.10v of type %T to %T", from.Interface(), from.Interface(), TTo(0))
	to := reflect.Indirect(reflect.ValueOf(new(TTo)))
	unsigned := false
	switch to.Interface().(type) {
	case uint, uint8, uint16, uint32, uint64, uintptr:
		unsigned = true
	}

	//fmt.Printf("\n\nto: %v (%T); from: %v (%T)\n\n", to, to.Interface(), from, from.Interface())
	switch typ := from.Interface().(type) {
	case nil:
		return TTo(0), nil
	case bool:
		if typ {
			return TTo(1), nil
		}
		return TTo(0), nil
	case int:
		if unsigned && typ < 0 {
			if abs {
				return TTo(-typ), nil
			}
			return ret, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(typ), nil
	case int64:
		if unsigned && typ < 0 {
			if abs {
				return TTo(-typ), nil
			}
			return ret, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(typ), nil
	case int32:
		if unsigned && typ < 0 {
			if abs {
				return TTo(-typ), nil
			}
			return ret, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(typ), nil
	case int16:
		if unsigned && typ < 0 {
			if abs {
				return TTo(-typ), nil
			}
			return ret, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(typ), nil
	case int8:
		if unsigned && typ < 0 {
			if abs {
				return TTo(-typ), nil
			}
			return ret, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(typ), nil
	case float64:
		if unsigned && typ < 0 {
			if abs {
				return TTo(-typ), nil
			}
			return ret, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		v, e := strToInt[TTo](To[string](from.Interface()))
		return TTo(v), e
	case float32:
		if unsigned && typ < 0 {
			if abs {
				return strToInt[TTo](To[string](-typ))
			}
			return ret, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		v, e := strToInt[TTo](To[string](typ))
		return TTo(v), e
	case uint:
		return TTo(typ), nil
	case uintptr:
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
		return strToInt[TTo](typ.String())
	case string:
		return strToInt[TTo](typ)

	//case complex128:
	//case complex64:
	default:
		return toInt[TTo](reflect.ValueOf(fmt.Sprintf("%#.10v", from.Interface())), ops)
	}
}

// strToInt converts a string to an integer type.
func strToInt[TTo constraints.Integer](from string) (TTo, error) {
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
				val = 0
			}
		} else {
			err = nil
		}
	}

	//fmt.Printf("to: %T, from: %v (%v), val: %d (%T)\n", TTo(0), from, from.Type().Kind(), int(math.Floor(val)), int(math.Floor(val)))
	if err != nil {
		return 0, errors.WrapE(err, errDetail)
	}
	if val >= 0 {
		return TTo(math.Floor(val)), nil
	}

	// Negative to uint error.
	ref := reflect.ValueOf(TTo(0))
	if ref.Kind() == reflect.Uint || ref.Kind() == reflect.Uint8 || ref.Kind() == reflect.Uint16 || ref.Kind() == reflect.Uint32 || ref.Kind() == reflect.Uint64 || ref.Kind() == reflect.Uintptr {
		return 0, errors.WrapE(ErrorSignedToUnsigned, errDetail)
	}

	return TTo(math.Ceil(val)), nil
}
