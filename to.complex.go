package cast

import (
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
	"golang.org/x/exp/constraints"
)

// toComplex64 casts an interface to a complex number.
func toComplex[TTo constraints.Complex](from reflect.Value) (TTo, error) {
	var nilval TTo
	fmt.Printf("complex from: '%#v' %T\n\n", from, from)

	to := reflect.Indirect(reflect.ValueOf(new(TTo)))
	switch to.Type().Kind() {
	case reflect.Complex64:
		f, err := toFloat[float32](from)
		if nil != err {
			return nilval, err
		}
		return TTo(complex(f, 0)), nil
	case reflect.Complex128:
		f, err := toFloat[float64](from)
		if nil != err {
			return nilval, err
		}
		return TTo(complex(f, 0)), nil
	}
	return TTo(0), errors.Errorf("unable to cast %#v of type %T to complex", from, from)
}
