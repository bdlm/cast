package cast

import (
	"github.com/bdlm/errors/v2"
)

// toComplex casts an interface to a complex number.
//
// When the source is a complex type both the real and imaginary parts are
// preserved. When the source is a non-complex type (int, float, string, bool)
// the imaginary part is zero by definition and the result is complex(src, 0).
//
// Options:
//   - DEFAULT: complex64 or complex128, default return value on error.
func toComplex[TTo complexNum](from any, ops ops) (TTo, error) {
	var ret TTo
	var ok bool

	if ops.hasDefault {
		if ret, ok = ops.defaultVal.(TTo); !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
	}

	// Preserve both real and imaginary parts for complex sources.
	switch v := from.(type) {
	case complex64:
		switch any(TTo(0)).(type) {
		case complex64:
			return TTo(v), nil
		case complex128:
			return TTo(complex128(v)), nil
		}
	case complex128:
		switch any(TTo(0)).(type) {
		case complex64:
			return TTo(complex64(v)), nil
		case complex128:
			return TTo(v), nil
		}
	}

	// For non-complex sources the imaginary part is zero by definition.
	floatOps := ops.Delete(DEFAULT)
	switch any(TTo(0)).(type) {
	case complex64:
		f, err := toFloat[float32](from, floatOps)
		if nil != err {
			return ret, err
		}
		return TTo(complex(f, 0)), nil
	case complex128:
		f, err := toFloat[float64](from, floatOps)
		if nil != err {
			return ret, err
		}
		return TTo(complex(f, 0)), nil
	}

	// Dead code, the above switch covers all complexNum types but the compiler
	// doesn't know that.
	return TTo(0), errors.Errorf(ErrorStrUnableToCast, from, from, TTo(0))
}
