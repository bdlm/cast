package cast

import (
	internal "github.com/bdlm/cast/v2/internal/cast"
)

// Sentinel errors and reusable format strings used throughout the package.
// These are re-exports of the values defined in the internal package so that
// callers see them through the cast namespace while the conversion code lives
// behind the internal/ boundary.
var (
	// ErrorUnableToCast is a sentinel error returned when a value cannot be
	// cast to the requested type.
	ErrorUnableToCast = internal.ErrorUnableToCast

	// ErrorSignedToUnsigned is returned when a negative value is cast to an
	// unsigned integer type and the ABS flag is not set.
	ErrorSignedToUnsigned = internal.ErrorSignedToUnsigned

	// ErrorInvalidOption is a format string (not an error value) used to build
	// error messages when an Op flag carries an unexpected value type.
	ErrorInvalidOption = internal.ErrorInvalidOption

	// ErrorStrErrorCastingFunc is a format string (not an error value) used
	// when an element cast fails inside a Func[T] closure generator.
	ErrorStrErrorCastingFunc = internal.ErrorStrErrorCastingFunc

	// ErrorStrUnableToCast is a format string (not an error value) used when
	// a value cannot be converted to the requested target type.
	ErrorStrUnableToCast = internal.ErrorStrUnableToCast

	// Error is a deprecated alias for [ErrorUnableToCast].
	//
	// Deprecated: use [ErrorUnableToCast].
	Error = ErrorUnableToCast
)

// Flag is the key type for conversion options passed to [To] and [ToE].
type Flag = internal.Flag

// Op is a single key/value option passed to [To] or [ToE]. Build one with a
// [Flag] constant and the appropriate value type for that flag.
//
// The Val type must exactly match the target type T for [DEFAULT], or the
// converter returns an error immediately — before the input is even inspected.
type Op = internal.Op

// Available option flags. See the internal package for full documentation.
const (
	DEFAULT             = internal.DEFAULT
	ABS                 = internal.ABS
	DECODE              = internal.DECODE
	DUPLICATE_KEY_ERROR = internal.DUPLICATE_KEY_ERROR
	FORMAT              = internal.FORMAT
	JSON                = internal.JSON
	LENGTH              = internal.LENGTH
	PRIVATE             = internal.PRIVATE
	STRICT              = internal.STRICT
	UNIQUE_VALUES       = internal.UNIQUE_VALUES
)

// Func is a named zero-argument function type that returns a T. A named type
// is required because Go generics cannot use plain function literals as type
// parameters directly.
type Func[TTo Types] func() TTo

// Types is the top-level constraint that [To] and [ToE] accept as TTo.
// It unions all supported target categories. Func variants for slices and
// channels are enumerated explicitly because Go does not expand
// Func[Tslice] into all individual Func[[]T] terms automatically.
type Types interface {
	Tbase | Tslice | Tchan | Tmap | Func[Tbase] |
		Func[[]int] | Func[[]int8] | Func[[]int16] | Func[[]int32] | Func[[]int64] |
		Func[[]uint] | Func[[]uint8] | Func[[]uint16] | Func[[]uint32] | Func[[]uint64] | Func[[]uintptr] |
		Func[[]float32] | Func[[]float64] |
		Func[[]complex64] | Func[[]complex128] |
		Func[[]string] | Func[[]bool] | Func[[]any] |
		Func[chan int] | Func[chan int8] | Func[chan int16] | Func[chan int32] | Func[chan int64] |
		Func[chan uint] | Func[chan uint8] | Func[chan uint16] | Func[chan uint32] | Func[chan uint64] | Func[chan uintptr] |
		Func[chan float32] | Func[chan float64] |
		Func[chan complex64] | Func[chan complex128] |
		Func[chan string] | Func[chan bool] | Func[chan any]
}

// Tbase covers all scalar types. The any term makes this constraint
// effectively unconstrained — all types satisfy it. This is intentional:
// interface targets like error and fmt.Stringer need to be expressible as TTo,
// and there is no way to enumerate all interface types. Unsupported kinds
// are rejected at runtime by ToE's dispatch switch.
type Tbase interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128 |
		~string | ~bool |
		any
}

// Tslice covers slice types of every scalar element kind, plus named types
// with the same underlying slice type (e.g. type Tags []string).
type Tslice interface {
	~[]int | ~[]int8 | ~[]int16 | ~[]int32 | ~[]int64 |
		~[]uint | ~[]uint8 | ~[]uint16 | ~[]uint32 | ~[]uint64 | ~[]uintptr |
		~[]float32 | ~[]float64 |
		~[]complex64 | ~[]complex128 |
		~[]string | ~[]bool |
		~[]any
}

// Tchan covers channels of scalars (~chan Tbase covers all basic-type channels
// in one term), channels of slices, channels of Func values, and nested
// channels (chan chan T).
type Tchan interface {
	~chan Tbase |
		~chan []int | ~chan []int8 | ~chan []int16 | ~chan []int32 | ~chan []int64 |
		~chan []uint | ~chan []uint8 | ~chan []uint16 | ~chan []uint32 | ~chan []uint64 | ~chan []uintptr |
		~chan []float32 | ~chan []float64 |
		~chan []complex64 | ~chan []complex128 |
		~chan []string | ~chan []bool |
		~chan []any | ~chan Func[Tbase] |
		~chan chan Tbase
}

// Tmap covers map types whose keys are any [Tbase] type and whose values are
// either a scalar [Tbase] or a slice of scalars. Named types with a matching
// underlying map type (e.g. type Attrs map[string]any) also satisfy Tmap.
type Tmap interface {
	~map[Tbase]Tbase |
		~map[Tbase][]int |
		~map[Tbase][]int8 |
		~map[Tbase][]int16 |
		~map[Tbase][]int32 |
		~map[Tbase][]int64 |
		~map[Tbase][]uint |
		~map[Tbase][]uint8 |
		~map[Tbase][]uint16 |
		~map[Tbase][]uint32 |
		~map[Tbase][]uint64 |
		~map[Tbase][]uintptr |
		~map[Tbase][]float32 |
		~map[Tbase][]float64 |
		~map[Tbase][]complex64 |
		~map[Tbase][]complex128 |
		~map[Tbase][]string |
		~map[Tbase][]bool |
		~map[Tbase][]any |
		~map[Tbase][]Func[Tbase]
}
