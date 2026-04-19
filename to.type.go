package cast

import (
	"maps"

	"github.com/bdlm/errors/v2"
)

// Sentinel errors and reusable format strings used throughout the package.
var (
	Error                    = errors.Errorf("unable to cast value")
	ErrorSignedToUnsigned    = errors.Wrap(Error, "cannot cast signed value to unsigned integer")
	ErrorInvalidOption       = "invalid %s value '%v'"
	ErrorStrErrorCastingFunc = "error casting %T to %T during function generation"
	ErrorStrUnableToCast     = "unable to cast %#.10v of type %T to %T"
)

// integer, float, and complexNum are internal constraints used by the
// per-kind conversion functions. They accept named types with the matching
// underlying type (e.g. type Celsius float32).
type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type float interface{ ~float32 | ~float64 }

type complexNum interface{ ~complex64 | ~complex128 }

// Flag is the key type for conversion options passed to [To] and [ToE].
type Flag int

// Ops is the parsed representation of options after [parseOps] collapses the
// variadic []Op slice. Internal conversion functions receive Ops directly.
type Ops map[Flag]any

// Op is a single key/value option passed to [To] or [ToE]. Build one with a
// [Flag] constant and the appropriate value type for that flag.
type Op struct {
	Flag Flag
	Val  any
}

// Available option flags. Not all flags apply to every target type; see the
// doc comment on each internal conversion function for which flags it honours.
const (
	DEFAULT Flag = iota // TTo,  default: TTo zero value, value to return on error

	ABS                 // bool, default: false, use absolute value during uint conversion
	DUPLICATE_KEY_ERROR // bool, default: false, error on duplicate map key
	LENGTH              // int,  default: 1,     initial capacity for slices / buffer size for channels (slices allow 0; channels require >= 1)
	UNIQUE_VALUES       // bool, default: false, dedupe slice values
	JSON                // bool, default: false, encode strings as JSON
	PRIVATE             // bool, default: false, include unexported struct fields in map output
	STRICT              // bool, default: false, return error instead of skipping unconvertible fields
)

// List converts Ops back to a []Op slice, used when forwarding options to
// recursive ToE calls for element-level casts inside containers.
func (o Ops) List() []Op {
	opList := []Op{}
	for flag, val := range o {
		opList = append(opList, Op{flag, val})
	}
	return opList
}

// Delete returns a copy of Ops with the given flag removed. Container
// conversion functions call this to strip DEFAULT before forwarding ops to
// element casts, so a container default is not mistakenly applied to elements.
func (o Ops) Delete(key Flag) Ops {
	n := Ops(maps.Clone(map[Flag]any(o)))
	delete(n, key)
	return n
}

// parseOps collapses the public variadic []Op into the internal Ops map used
// by all conversion functions.
func parseOps(o []Op) Ops {
	ops := Ops{}
	for _, op := range o {
		ops[op.Flag] = op.Val
	}
	return ops
}

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

// Tbase covers all scalar types plus any (the empty interface). The any term
// enables interface{} as a valid target, not all types.
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
