package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
	"golang.org/x/exp/constraints"
)

// toComplex casts an interface to a complex number.
//
// Options:
//   - DEFAULT: complex64 or complex128, default return value on error.
func toComplex[TTo constraints.Complex](from reflect.Value, ops Ops) (TTo, error) {
	var ret TTo
	var ok bool

	if _, ok = ops[DEFAULT]; ok {
		if ret, ok = ops[DEFAULT].(TTo); !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops[DEFAULT])
		}
	}

	to := reflect.Indirect(reflect.ValueOf(new(TTo)))
	switch to.Type().Kind() {
	case reflect.Complex64:
		f, err := toFloat[float32](from, ops)
		if nil != err {
			return ret, err
		}
		return TTo(complex(f, 0)), nil
	case reflect.Complex128:
		f, err := toFloat[float64](from, ops)
		if nil != err {
			return ret, err
		}
		return TTo(complex(f, 0)), nil
	}
	return TTo(0), errors.Errorf(ErrorStrUnableToCast, from, from, TTo(0))
}
