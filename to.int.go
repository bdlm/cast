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
func toInt[TTo constraints.Integer](from reflect.Value) (TTo, error) {
	fromVal := reflect.Indirect(from)
	if !fromVal.IsValid() || !fromVal.CanInterface() {
		return TTo(0), errors.Errorf("unable to cast %#.10v of type %T to %T", from, from, TTo(0))
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
			return 0, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(from.Interface().(int)), nil
	case int64:
		if unsigned && typ < 0 {
			return 0, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(from.Interface().(int64)), nil
	case int32:
		if unsigned && typ < 0 {
			return 0, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(from.Interface().(int32)), nil
	case int16:
		if unsigned && typ < 0 {
			return 0, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(from.Interface().(int16)), nil
	case int8:
		if unsigned && typ < 0 {
			return 0, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		return TTo(from.Interface().(int8)), nil
	case float64:
		if unsigned && typ < 0 {
			return 0, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		v, e := strToInt[TTo](to, To[string](from.Interface()))
		return TTo(v), e
	case float32:
		if unsigned && typ < 0 {
			return 0, errors.WrapE(ErrorSignedToUnsigned, errDetail)
		}
		v, e := strToInt[TTo](to, To[string](from.Interface()))
		return TTo(v), e
	case uint:
		return TTo(from.Interface().(uint)), nil
	case uintptr:
		return TTo(from.Interface().(uintptr)), nil
	case uint64:
		return TTo(from.Interface().(uint64)), nil
	case uint32:
		return TTo(from.Interface().(uint32)), nil
	case uint16:
		return TTo(from.Interface().(uint16)), nil
	case uint8:
		return TTo(from.Interface().(uint8)), nil
	case fmt.Stringer:
		return strToInt[TTo](to, typ.String())
	case string:
		return strToInt[TTo](to, typ)

	//case complex128:
	//case complex64:
	default:
		return toInt[TTo](reflect.ValueOf(fmt.Sprintf("%#.10v", from.Interface())))
	}
}

// strToInt converts a string to an integer type.
func strToInt[TTo constraints.Integer](to reflect.Value, from string) (TTo, error) {
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
