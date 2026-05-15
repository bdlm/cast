package cast

import (
	"fmt"
	"strconv"
)

// ToComplex casts an interface to a complex number.
//
// When the source is a complex type both the real and imaginary parts are
// preserved. When the source is a non-complex type (int, float, string, bool)
// the imaginary part is zero by definition and the result is complex(src, 0).
//
// Options:
//   - DEFAULT: complex64 or complex128, default return value on error.
func ToComplex[TTo ComplexNum](from any, ops Ops) (TTo, error) {
	var ret TTo
	var ok bool

	if ops.HasDefault {
		if ret, ok = ops.DefaultVal.(TTo); !ok {
			return ret, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.DefaultVal)
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

	// String and []byte sources: strconv.ParseComplex handles the full complex
	// number syntax including imaginary parts (e.g. "(1+2i)", "3.14+0i").
	// Falls through to ToFloat only when ParseComplex fails (e.g. comma-
	// formatted numbers or DECODE=JSON inputs).
	bitSize := 128
	if _, ok := any(TTo(0)).(complex64); ok {
		bitSize = 64
	}
	switch v := from.(type) {
	case string:
		if c, err := strconv.ParseComplex(v, bitSize); err == nil {
			return TTo(c), nil
		}
	case []byte:
		if c, err := strconv.ParseComplex(string(v), bitSize); err == nil {
			return TTo(c), nil
		}
	}

	// For non-complex sources the imaginary part is zero by definition.
	floatOps := ops.Delete(DEFAULT)
	switch any(TTo(0)).(type) {
	case complex64:
		f, err := ToFloat[float32](from, floatOps)
		if nil != err {
			return ret, err
		}
		return TTo(complex(f, 0)), nil
	case complex128:
		f, err := ToFloat[float64](from, floatOps)
		if nil != err {
			return ret, err
		}
		return TTo(complex(f, 0)), nil
	}

	// Dead code, the above switch covers all ComplexNum types but the compiler
	// doesn't know that.
	return TTo(0), fmt.Errorf(ErrorStrUnableToCast, from, from, TTo(0))
}
