package cast_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestSliceTypes(t *testing.T) {
	for name, cases := range sliceCases {
		switch name {
		case "bool":
			testSliceCases[[]bool](t, cases)
		case "byte":
			testSliceCases[[]byte](t, cases)
		case "rune":
			testSliceCases[[]rune](t, cases)
		case "int":
			testSliceCases[[]int](t, cases)
		case "int8":
			testSliceCases[[]int8](t, cases)
		case "int16":
			testSliceCases[[]int16](t, cases)
		case "int32":
			testSliceCases[[]int32](t, cases)
		case "int64":
			testSliceCases[[]int64](t, cases)
		case "uint":
			testSliceCases[[]uint](t, cases)
		case "uint8":
			testSliceCases[[]uint8](t, cases)
		case "uint16":
			testSliceCases[[]uint16](t, cases)
		case "uint32":
			testSliceCases[[]uint32](t, cases)
		case "uint64":
			testSliceCases[[]uint64](t, cases)
		case "uintptr":
			testSliceCases[[]uintptr](t, cases)
		case "float32":
			testSliceCases[[]float32](t, cases)
		case "float64":
			testSliceCases[[]float64](t, cases)
		case "complex64":
			testSliceCases[[]complex64](t, cases)
		case "complex128":
			testSliceCases[[]complex128](t, cases)
		}
	}
}

func testSliceCases[TTo any](t *testing.T, cases []testCase) {
	var typ TTo
	name := fmt.Sprintf("%T", typ)

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
			actual, err := cast.ToE[TTo](test.in)
			testInfo := fmt.Sprintf(`
case: ToE[%s]
input: %v (%T)
expect error: %v; actual error: % #+v
expected result: %v (%T); actual result: %v (%T)
test: %#v
			`,
				name,
				test.in,
				test.in,
				test.expectErr,
				err,
				test.expect,
				test.expect,
				actual,
				actual,
				test,
			)
			if nil != err && !test.expectErr {
				t.Error("1. expected nil, got error", testInfo)
			} else if nil == err && test.expectErr {
				t.Error("2. expected error, got nil", testInfo)
			} else if nil != err && !errors.Is(err, cast.Error) {
				t.Error("3. expected cast.Error, got different error type", testInfo)
			}
		})
	}
}

var sliceCases = testCases{
	"bool": {
		{in: []bool{true, false}, expect: []bool{true, false}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []bool{false, true}, err: nil, expectErr: false},
		{in: []bool{}, expect: []bool{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []bool{true, true, true, false, true}, err: nil, expectErr: false},
		{in: []bool{}, expect: []bool{}, err: nil, expectErr: false},
		{in: 1, expect: []bool{}, err: nil, expectErr: true},
		{in: "1", expect: []bool{}, err: nil, expectErr: true},
		{in: -1, expect: []bool{}, err: nil, expectErr: true},
		{in: "-1", expect: []bool{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []bool{true}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []bool{true}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []bool{true}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []bool{true}, err: nil, expectErr: false},
	},
	// int
	"int": {
		{in: []int{1, 30, 10, 0, 1}, expect: []int{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int{-1, 0, 1}, expect: []int{-1, 0, 1}, err: nil, expectErr: false},
		{in: []int{}, expect: []int{}, err: nil, expectErr: false},
		{in: 1, expect: []int{1}, err: nil, expectErr: true},
		{in: "1", expect: []int{1}, err: nil, expectErr: true},
		{in: -1, expect: []int{1}, err: nil, expectErr: true},
		{in: "-1", expect: []int{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int{1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []int{-1}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []int{-1}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []int{-1}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []int{-1}, err: nil, expectErr: false},
	},
	"int8": {
		{in: []int8{1, 30, 10, 0, 1}, expect: []int8{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int8{-1, 0, 1}, expect: []int8{-1, 0, 1}, err: nil, expectErr: false},
		{in: []int8{}, expect: []int8{}, err: nil, expectErr: false},
		{in: 1, expect: []int8{1}, err: nil, expectErr: true},
		{in: "1", expect: []int8{1}, err: nil, expectErr: true},
		{in: -1, expect: []int8{1}, err: nil, expectErr: true},
		{in: "-1", expect: []int8{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []int8{-1}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []int8{-1}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []int8{-1}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []int8{-1}, err: nil, expectErr: false},
	},
	"int16": {
		{in: []int16{1, 30, 10, 0, 1}, expect: []int16{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int16{-1, 0, 1}, expect: []int16{-1, 0, 1}, err: nil, expectErr: false},
		{in: []int16{}, expect: []int16{}, err: nil, expectErr: false},
		{in: 1, expect: []int16{1}, err: nil, expectErr: true},
		{in: "1", expect: []int16{1}, err: nil, expectErr: true},
		{in: -1, expect: []int16{1}, err: nil, expectErr: true},
		{in: "-1", expect: []int16{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []int16{-1}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []int16{-1}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []int16{-1}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []int16{-1}, err: nil, expectErr: false},
	},
	"int32": {
		{in: []int32{1, 30, 10, 0, 1}, expect: []int32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 0, 1}, expect: []int32{-1, 0, 1}, err: nil, expectErr: false},
		{in: []int32{}, expect: []int32{}, err: nil, expectErr: false},
		{in: 1, expect: []int32{1}, err: nil, expectErr: true},
		{in: "1", expect: []int32{1}, err: nil, expectErr: true},
		{in: -1, expect: []int32{1}, err: nil, expectErr: true},
		{in: "-1", expect: []int32{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []int32{-1}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []int32{-1}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []int32{-1}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []int32{-1}, err: nil, expectErr: false},
	},
	"int64": {
		{in: []int64{1, 30, 10, 0, 1}, expect: []int64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int64{-1, 0, 1}, expect: []int64{-1, 0, 1}, err: nil, expectErr: false},
		{in: []int64{}, expect: []int64{}, err: nil, expectErr: false},
		{in: 1, expect: []int64{1}, err: nil, expectErr: true},
		{in: "1", expect: []int64{1}, err: nil, expectErr: true},
		{in: -1, expect: []int64{1}, err: nil, expectErr: true},
		{in: "-1", expect: []int64{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []int64{-1}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []int64{-1}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []int64{-1}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []int64{-1}, err: nil, expectErr: false},
	},

	// uint
	"uint": {
		{in: []uint{1, 30, 10, 0, 1}, expect: []uint{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []uint{}, err: nil, expectErr: true},
		{in: []string{"-1"}, expect: []uint{}, err: nil, expectErr: true},
		{in: []float32{-1.9}, expect: []uint{}, err: nil, expectErr: true},
		{in: []string{"-1.9"}, expect: []uint{}, err: nil, expectErr: true},
		{in: []uint{}, expect: []uint{}, err: nil, expectErr: false},
		{in: 1, expect: []uint{1}, err: nil, expectErr: true},
		{in: "1", expect: []uint{1}, err: nil, expectErr: true},
		{in: -1, expect: []uint{1}, err: nil, expectErr: true},
		{in: "-1", expect: []uint{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint{1}, err: nil, expectErr: false},
	},
	"uint8": {
		{in: []uint8{1, 30, 10, 0, 1}, expect: []uint8{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []uint8{}, err: nil, expectErr: true},
		{in: []string{"-1"}, expect: []uint8{}, err: nil, expectErr: true},
		{in: []float32{-1.9}, expect: []uint8{}, err: nil, expectErr: true},
		{in: []string{"-1.9"}, expect: []uint8{}, err: nil, expectErr: true},
		{in: []uint8{}, expect: []uint8{}, err: nil, expectErr: false},
		{in: 1, expect: []uint8{1}, err: nil, expectErr: true},
		{in: "1", expect: []uint8{1}, err: nil, expectErr: true},
		{in: -1, expect: []uint8{1}, err: nil, expectErr: true},
		{in: "-1", expect: []uint8{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint8{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint8{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint8{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint8{1}, err: nil, expectErr: false},
	},
	"uint16": {
		{in: []uint16{1, 30, 10, 0, 1}, expect: []uint16{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []uint16{}, err: nil, expectErr: true},
		{in: []string{"-1"}, expect: []uint16{}, err: nil, expectErr: true},
		{in: []float32{-1.9}, expect: []uint16{}, err: nil, expectErr: true},
		{in: []string{"-1.9"}, expect: []uint16{}, err: nil, expectErr: true},
		{in: []uint16{}, expect: []uint16{}, err: nil, expectErr: false},
		{in: 1, expect: []uint16{1}, err: nil, expectErr: true},
		{in: "1", expect: []uint16{1}, err: nil, expectErr: true},
		{in: -1, expect: []uint16{1}, err: nil, expectErr: true},
		{in: "-1", expect: []uint16{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint16{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint16{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint16{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint16{1}, err: nil, expectErr: false},
	},
	"uint32": {
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []uint32{}, err: nil, expectErr: true},
		{in: []string{"-1"}, expect: []uint32{}, err: nil, expectErr: true},
		{in: []float32{-1.9}, expect: []uint32{}, err: nil, expectErr: true},
		{in: []string{"-1.9"}, expect: []uint32{}, err: nil, expectErr: true},
		{in: []uint32{}, expect: []uint32{}, err: nil, expectErr: false},
		{in: 1, expect: []uint32{1}, err: nil, expectErr: true},
		{in: "1", expect: []uint32{1}, err: nil, expectErr: true},
		{in: -1, expect: []uint32{1}, err: nil, expectErr: true},
		{in: "-1", expect: []uint32{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint32{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint32{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint32{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint32{1}, err: nil, expectErr: false},
	},
	"uint64": {
		{in: []uint64{1, 30, 10, 0, 1}, expect: []uint64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []uint64{}, err: nil, expectErr: true},
		{in: []string{"-1"}, expect: []uint64{}, err: nil, expectErr: true},
		{in: []float32{-1.9}, expect: []uint64{}, err: nil, expectErr: true},
		{in: []string{"-1.9"}, expect: []uint64{}, err: nil, expectErr: true},
		{in: []uint64{}, expect: []uint64{}, err: nil, expectErr: false},
		{in: 1, expect: []uint64{1}, err: nil, expectErr: true},
		{in: "1", expect: []uint64{1}, err: nil, expectErr: true},
		{in: -1, expect: []uint64{1}, err: nil, expectErr: true},
		{in: "-1", expect: []uint64{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint64{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint64{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint64{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint64{1}, err: nil, expectErr: false},
	},
	"uintptr": {
		{in: []uintptr{1, 30, 10, 0, 1}, expect: []uintptr{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []uintptr{}, err: nil, expectErr: true},
		{in: []string{"-1"}, expect: []uintptr{}, err: nil, expectErr: true},
		{in: []float32{-1.9}, expect: []uintptr{}, err: nil, expectErr: true},
		{in: []string{"-1.9"}, expect: []uintptr{}, err: nil, expectErr: true},
		{in: []uintptr{}, expect: []uintptr{}, err: nil, expectErr: false},
		{in: 1, expect: []uintptr{1}, err: nil, expectErr: true},
		{in: "1", expect: []uintptr{1}, err: nil, expectErr: true},
		{in: -1, expect: []uintptr{1}, err: nil, expectErr: true},
		{in: "-1", expect: []uintptr{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uintptr{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uintptr{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uintptr{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uintptr{1}, err: nil, expectErr: false},
	},
	"float32": {
		{in: []float32{1, 30, 10, 0, 1}, expect: []float32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []float32{-1}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []float32{-1}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []float32{-1.9}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []float32{-1.9}, err: nil, expectErr: false},
		{in: []float32{}, expect: []float32{}, err: nil, expectErr: false},
		{in: 1, expect: []float32{1}, err: nil, expectErr: true},
		{in: "1", expect: []float32{1}, err: nil, expectErr: true},
		{in: -1, expect: []float32{1}, err: nil, expectErr: true},
		{in: "-1", expect: []float32{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []float32{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []float32{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []float32{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []float32{1}, err: nil, expectErr: false},
	},
	"float64": {
		{in: []float64{1, 30, 10, 0, 1}, expect: []float64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []float64{}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []float64{}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []float64{}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []float64{}, err: nil, expectErr: false},
		{in: []float64{}, expect: []float64{}, err: nil, expectErr: false},
		{in: 1, expect: []float64{1}, err: nil, expectErr: true},
		{in: "1", expect: []float64{1}, err: nil, expectErr: true},
		{in: -1, expect: []float64{1}, err: nil, expectErr: true},
		{in: "-1", expect: []float64{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []float64{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []float64{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []float64{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []float64{1}, err: nil, expectErr: false},
	},
	"complex64": {
		{in: []float64{1, 30, 10, 0, 1}, expect: []complex64{complex64(1), complex64(30), complex64(10), complex64(0), complex64(1)}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []complex64{complex64(-1)}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []complex64{complex64(-1)}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []complex64{complex64(-1.9)}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []complex64{complex64(-1.9)}, err: nil, expectErr: false},
		{in: []complex64{}, expect: []float64{}, err: nil, expectErr: false},
		{in: 1, expect: []complex64{1}, err: nil, expectErr: true},
		{in: "1", expect: []complex64{1}, err: nil, expectErr: true},
		{in: -1, expect: []complex64{1}, err: nil, expectErr: true},
		{in: "-1", expect: []complex64{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []complex64{complex64(1.1)}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []complex64{complex64(1.1)}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []complex64{complex64(1.9)}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []complex64{complex64(1.9)}, err: nil, expectErr: false},
	},
	"complex128": {
		{in: []float64{1, 30, 10, 0, 1}, expect: []complex128{complex128(1), complex128(30), complex128(10), complex128(0), complex128(1)}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []complex128{complex128(-1)}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []complex128{complex128(-1)}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []complex128{complex128(-1.9)}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []complex128{complex128(-1.9)}, err: nil, expectErr: false},
		{in: []complex128{}, expect: []float64{}, err: nil, expectErr: false},
		{in: 1, expect: []complex128{1}, err: nil, expectErr: true},
		{in: "1", expect: []complex128{1}, err: nil, expectErr: true},
		{in: -1, expect: []complex128{1}, err: nil, expectErr: true},
		{in: "-1", expect: []complex128{1}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []complex128{complex128(1.1)}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []complex128{complex128(1.1)}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []complex128{complex128(1.9)}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []complex128{complex128(1.9)}, err: nil, expectErr: false},
	},
}
