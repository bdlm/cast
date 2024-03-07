package cast

import (
	"github.com/bdlm/errors/v2"
)

var (
	Error                    = errors.Errorf("unable to cast value")
	ErrorSignedToUnsigned    = errors.Wrap(Error, "cannot cast signed value to unsigned integer")
	ErrorInvalidOption       = "invalid %s value '%v'"
	ErrorStrErrorCastingFunc = "error casting %T to %T during function generation"
	ErrorStrUnableToCast     = "unable to cast %#.10v of type %T to %T"
)

type Flag int

type Ops map[Flag]any

const (
	DEFAULT Flag = iota // TTo,  default return value on error, default zero value

	ABS                 // bool, use absolute value during uint conversion, default false
	DUPLICATE_KEY_ERROR // bool, error on duplicate map key, default false
	LENGTH              // int,  number of elements in result, default 1
	UNIQUE_VALUES       // bool, dedupe slice values, default false
	JSON                // bool, encode strings as JSON, default false
)

func parseOps(o []Ops) Ops {
	ops := Ops{}
	for _, opMap := range o {
		for k, v := range opMap {
			ops[k] = v
		}
	}
	return ops
}

//////////////////////////////

type Func[TTo Types] func() TTo

type Types interface {
	Tbase | Tslice | Tchan | Tmap | Func[Tbase]
}

type Tbase interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128 |
		~string | ~bool |
		any
}

type Tslice interface {
	~[]int | ~[]int8 | ~[]int16 | ~[]int32 | ~[]int64 |
		~[]uint | ~[]uint8 | ~[]uint16 | ~[]uint32 | ~[]uint64 | ~[]uintptr |
		~[]float32 | ~[]float64 |
		~[]complex64 | ~[]complex128 |
		~[]string | ~[]bool |
		~[]any
}

type Tchan interface {
	~chan Tbase |
		~chan []int | ~chan []int8 | ~chan []int16 | ~chan []int32 | ~chan []int64 |
		~chan []uint | ~chan []uint8 | ~chan []uint16 | ~chan []uint32 | ~chan []uint64 | ~chan []uintptr |
		~chan []float32 | ~chan []float64 |
		~chan []complex64 | ~chan []complex128 |
		~chan []string | ~chan []bool |
		~chan []any | ~chan Func[Tbase]
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
