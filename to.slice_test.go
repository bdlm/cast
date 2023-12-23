package cast_test

import (
	"errors"
	"fmt"
	"reflect"
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
expect error: %v; actual error: %v
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
			} else if nil == err && !reflect.DeepEqual(actual, test.expect.(TTo)) {
				t.Errorf("4. expected %v to equal %v %s", test.expect, actual, testInfo)
			}
		})
	}
}

var sliceCases = testCases{
	"bool": {
		{in: true, expect: []bool{true}, err: nil, expectErr: false},
		{in: false, expect: []bool{false}, err: nil, expectErr: false},
		{in: 1, expect: []bool{true}, err: nil, expectErr: false},
		{in: 0, expect: []bool{false}, err: nil, expectErr: false},
		{in: "hi", expect: []bool{}, err: nil, expectErr: true},
		{in: float64(1.1), expect: []bool{true}, err: nil, expectErr: false},
		{in: float64(-1.1), expect: []bool{true}, err: nil, expectErr: false},
	},
	// int
	"int": {
		{in: 1, expect: []int{1}, err: nil, expectErr: false},
		{in: "1", expect: []int{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []int{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []int{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []int{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []int{1}, err: nil, expectErr: false},
		{in: -1, expect: []int{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []int{-1}, err: nil, expectErr: false},
		{in: -1.9, expect: []int{-1}, err: nil, expectErr: false},
		{in: "-1.9", expect: []int{-1}, err: nil, expectErr: false},
	},
	"int8": {
		{in: 1, expect: []int8{1}, err: nil, expectErr: false},
		{in: "1", expect: []int8{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []int8{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []int8{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []int8{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []int8{1}, err: nil, expectErr: false},
		{in: -1, expect: []int8{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []int8{-1}, err: nil, expectErr: false},
		{in: -1.9, expect: []int8{-1}, err: nil, expectErr: false},
		{in: "-1.9", expect: []int8{-1}, err: nil, expectErr: false},
	},
	"int16": {
		{in: 1, expect: []int16{1}, err: nil, expectErr: false},
		{in: "1", expect: []int16{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []int16{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []int16{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []int16{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []int16{1}, err: nil, expectErr: false},
		{in: -1, expect: []int16{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []int16{-1}, err: nil, expectErr: false},
		{in: -1.9, expect: []int16{-1}, err: nil, expectErr: false},
		{in: "-1.9", expect: []int16{-1}, err: nil, expectErr: false},
	},
	"int32": {
		{in: 1, expect: []int32{1}, err: nil, expectErr: false},
		{in: "1", expect: []int32{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []int32{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []int32{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []int32{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []int32{1}, err: nil, expectErr: false},
		{in: -1, expect: []int32{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []int32{-1}, err: nil, expectErr: false},
		{in: -1.9, expect: []int32{-1}, err: nil, expectErr: false},
		{in: "-1.9", expect: []int32{-1}, err: nil, expectErr: false},
	},
	"int64": {
		{in: 1, expect: []int64{1}, err: nil, expectErr: false},
		{in: "1", expect: []int64{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []int64{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []int64{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []int64{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []int64{1}, err: nil, expectErr: false},
		{in: -1, expect: []int64{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []int64{-1}, err: nil, expectErr: false},
		{in: -1.9, expect: []int64{-1}, err: nil, expectErr: false},
		{in: "-1.9", expect: []int64{-1}, err: nil, expectErr: false},
	},

	// uint
	"uint": {
		{in: 1, expect: []uint{1}, err: nil, expectErr: false},
		{in: "1", expect: []uint{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []uint{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []uint{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []uint{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []uint{1}, err: nil, expectErr: false},
		{in: -1, expect: []uint{0}, err: nil, expectErr: true},
		{in: "-1", expect: []uint{0}, err: nil, expectErr: true},
		{in: -1.9, expect: []uint{0}, err: nil, expectErr: true},
		{in: "-1.9", expect: []uint{0}, err: nil, expectErr: true},
	},
	"uint8": {
		{in: 1, expect: []uint8{1}, err: nil, expectErr: false},
		{in: "1", expect: []uint8{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []uint8{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []uint8{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []uint8{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []uint8{1}, err: nil, expectErr: false},
		{in: -1, expect: []uint8{0}, err: nil, expectErr: true},
		{in: "-1", expect: []uint8{0}, err: nil, expectErr: true},
		{in: -1.9, expect: []uint8{0}, err: nil, expectErr: true},
		{in: "-1.9", expect: []uint8{0}, err: nil, expectErr: true},
	},
	"uint16": {
		{in: 1, expect: []uint16{1}, err: nil, expectErr: false},
		{in: "1", expect: []uint16{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []uint16{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []uint16{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []uint16{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []uint16{1}, err: nil, expectErr: false},
		{in: -1, expect: []uint16{0}, err: nil, expectErr: true},
		{in: "-1", expect: []uint16{0}, err: nil, expectErr: true},
		{in: -1.9, expect: []uint16{0}, err: nil, expectErr: true},
		{in: "-1.9", expect: []uint16{0}, err: nil, expectErr: true},
	},
	"uint32": {
		{in: 1, expect: []uint32{1}, err: nil, expectErr: false},
		{in: "1", expect: []uint32{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []uint32{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []uint32{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []uint32{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []uint32{1}, err: nil, expectErr: false},
		{in: -1, expect: []uint32{0}, err: nil, expectErr: true},
		{in: "-1", expect: []uint32{0}, err: nil, expectErr: true},
		{in: -1.9, expect: []uint32{0}, err: nil, expectErr: true},
		{in: "-1.9", expect: []uint32{0}, err: nil, expectErr: true},
	},
	"uint64": {
		{in: 1, expect: []uint64{1}, err: nil, expectErr: false},
		{in: "1", expect: []uint64{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []uint64{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []uint64{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []uint64{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []uint64{1}, err: nil, expectErr: false},
		{in: -1, expect: []uint64{0}, err: nil, expectErr: true},
		{in: "-1", expect: []uint64{0}, err: nil, expectErr: true},
		{in: -1.9, expect: []uint64{0}, err: nil, expectErr: true},
		{in: "-1.9", expect: []uint64{0}, err: nil, expectErr: true},
	},
	"uintptr": {
		{in: 1, expect: []uintptr{1}, err: nil, expectErr: false},
		{in: "1", expect: []uintptr{1}, err: nil, expectErr: false},
		{in: 1.1, expect: []uintptr{1}, err: nil, expectErr: false},
		{in: "1.1", expect: []uintptr{1}, err: nil, expectErr: false},
		{in: 1.9, expect: []uintptr{1}, err: nil, expectErr: false},
		{in: "1.9", expect: []uintptr{1}, err: nil, expectErr: false},
		{in: -1, expect: []uintptr{0}, err: nil, expectErr: true},
		{in: "-1", expect: []uintptr{0}, err: nil, expectErr: true},
		{in: -1.9, expect: []uintptr{0}, err: nil, expectErr: true},
		{in: "-1.9", expect: []uintptr{0}, err: nil, expectErr: true},
	},
}
