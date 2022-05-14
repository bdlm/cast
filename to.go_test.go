package cast_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
)

type testCase struct {
	in        any
	expect    any
	err       error
	expectErr bool
}

type testCases map[string][]testCase

func TestSimpleTypes(t *testing.T) {
	for name, cases := range simpleCases {
		switch name {
		case "bool":
			testSimpleCases[bool](t, cases)
		case "byte":
			testSimpleCases[byte](t, cases)
		case "rune":
			testSimpleCases[rune](t, cases)
		case "int":
			testSimpleCases[int](t, cases)
		case "int8":
			testSimpleCases[int8](t, cases)
		case "int16":
			testSimpleCases[int16](t, cases)
		case "int32":
			testSimpleCases[int32](t, cases)
		case "int64":
			testSimpleCases[int64](t, cases)
		case "uint":
			testSimpleCases[uint](t, cases)
		case "uint8":
			testSimpleCases[uint8](t, cases)
		case "uint16":
			testSimpleCases[uint16](t, cases)
		case "uint32":
			testSimpleCases[uint32](t, cases)
		case "uint64":
			testSimpleCases[uint64](t, cases)
		case "uintptr":
			testSimpleCases[uintptr](t, cases)
		case "float32":
			testSimpleCases[float32](t, cases)
		case "float64":
			testSimpleCases[float64](t, cases)
		case "complex64":
			testSimpleCases[complex64](t, cases)
		case "complex128":
			testSimpleCases[complex128](t, cases)
		}
	}
}

func testSimpleCases[TTo any](t *testing.T, cases []testCase) {
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

			if err != nil && !test.expectErr {
				t.Error(1, testInfo)
			} else if err == nil && test.expectErr {
				t.Error(2, testInfo)
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Error(3, testInfo)
			} else if !reflect.DeepEqual(actual, test.expect) {
				t.Error(4, testInfo)
			}
		})
	}
}

var simpleCases = testCases{
	"bool": {
		{in: true, expect: true, err: nil, expectErr: false},
		{in: 1, expect: true, err: nil, expectErr: false},
		{in: 0, expect: false, err: nil, expectErr: false},
		{in: "hi", expect: false, err: nil, expectErr: true},
		{in: float64(1.1), expect: true, err: nil, expectErr: false},
		{in: float64(-1.1), expect: true, err: nil, expectErr: false},
	},
	// int
	"int": {
		{in: 1, expect: int(1), err: nil, expectErr: false},
		{in: "1", expect: int(1), err: nil, expectErr: false},
		{in: 1.1, expect: int(1), err: nil, expectErr: false},
		{in: "1.1", expect: int(1), err: nil, expectErr: false},
		{in: 1.9, expect: int(1), err: nil, expectErr: false},
		{in: "1.9", expect: int(1), err: nil, expectErr: false},
		{in: -1, expect: int(-1), err: nil, expectErr: false},
		{in: "-1", expect: int(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int(-1), err: nil, expectErr: false},
	},
	"int8": {
		{in: 1, expect: int8(1), err: nil, expectErr: false},
		{in: "1", expect: int8(1), err: nil, expectErr: false},
		{in: 1.1, expect: int8(1), err: nil, expectErr: false},
		{in: "1.1", expect: int8(1), err: nil, expectErr: false},
		{in: 1.9, expect: int8(1), err: nil, expectErr: false},
		{in: "1.9", expect: int8(1), err: nil, expectErr: false},
		{in: -1, expect: int8(-1), err: nil, expectErr: false},
		{in: "-1", expect: int8(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int8(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int8(-1), err: nil, expectErr: false},
	},
	"int16": {
		{in: 1, expect: int16(1), err: nil, expectErr: false},
		{in: "1", expect: int16(1), err: nil, expectErr: false},
		{in: 1.1, expect: int16(1), err: nil, expectErr: false},
		{in: "1.1", expect: int16(1), err: nil, expectErr: false},
		{in: 1.9, expect: int16(1), err: nil, expectErr: false},
		{in: "1.9", expect: int16(1), err: nil, expectErr: false},
		{in: -1, expect: int16(-1), err: nil, expectErr: false},
		{in: "-1", expect: int16(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int16(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int16(-1), err: nil, expectErr: false},
	},
	"int32": {
		{in: 1, expect: int32(1), err: nil, expectErr: false},
		{in: "1", expect: int32(1), err: nil, expectErr: false},
		{in: 1.1, expect: int32(1), err: nil, expectErr: false},
		{in: "1.1", expect: int32(1), err: nil, expectErr: false},
		{in: 1.9, expect: int32(1), err: nil, expectErr: false},
		{in: "1.9", expect: int32(1), err: nil, expectErr: false},
		{in: -1, expect: int32(-1), err: nil, expectErr: false},
		{in: "-1", expect: int32(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int32(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int32(-1), err: nil, expectErr: false},
	},
	"int64": {
		{in: 1, expect: int64(1), err: nil, expectErr: false},
		{in: "1", expect: int64(1), err: nil, expectErr: false},
		{in: 1.1, expect: int64(1), err: nil, expectErr: false},
		{in: "1.1", expect: int64(1), err: nil, expectErr: false},
		{in: 1.9, expect: int64(1), err: nil, expectErr: false},
		{in: "1.9", expect: int64(1), err: nil, expectErr: false},
		{in: -1, expect: int64(-1), err: nil, expectErr: false},
		{in: "-1", expect: int64(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int64(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int64(-1), err: nil, expectErr: false},
	},

	// uint
	"uint": {
		{in: 1, expect: uint(1), err: nil, expectErr: false},
		{in: "1", expect: uint(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint(1), err: nil, expectErr: false},
		{in: -1, expect: uint(0), err: nil, expectErr: true},
		{in: "-1", expect: uint(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint(0), err: nil, expectErr: true},
	},
	"uint8": {
		{in: 1, expect: uint8(1), err: nil, expectErr: false},
		{in: "1", expect: uint8(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint8(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint8(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint8(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint8(1), err: nil, expectErr: false},
		{in: -1, expect: uint8(0), err: nil, expectErr: true},
		{in: "-1", expect: uint8(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint8(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint8(0), err: nil, expectErr: true},
	},
	"uint16": {
		{in: 1, expect: uint16(1), err: nil, expectErr: false},
		{in: "1", expect: uint16(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint16(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint16(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint16(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint16(1), err: nil, expectErr: false},
		{in: -1, expect: uint16(0), err: nil, expectErr: true},
		{in: "-1", expect: uint16(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint16(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint16(0), err: nil, expectErr: true},
	},
	"uint32": {
		{in: 1, expect: uint32(1), err: nil, expectErr: false},
		{in: "1", expect: uint32(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint32(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint32(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint32(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint32(1), err: nil, expectErr: false},
		{in: -1, expect: uint32(0), err: nil, expectErr: true},
		{in: "-1", expect: uint32(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint32(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint32(0), err: nil, expectErr: true},
	},
	"uint64": {
		{in: 1, expect: uint64(1), err: nil, expectErr: false},
		{in: "1", expect: uint64(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint64(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint64(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint64(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint64(1), err: nil, expectErr: false},
		{in: -1, expect: uint64(0), err: nil, expectErr: true},
		{in: "-1", expect: uint64(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint64(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint64(0), err: nil, expectErr: true},
	},
	"uintptr": {
		{in: 1, expect: uintptr(1), err: nil, expectErr: false},
		{in: "1", expect: uintptr(1), err: nil, expectErr: false},
		{in: 1.1, expect: uintptr(1), err: nil, expectErr: false},
		{in: "1.1", expect: uintptr(1), err: nil, expectErr: false},
		{in: 1.9, expect: uintptr(1), err: nil, expectErr: false},
		{in: "1.9", expect: uintptr(1), err: nil, expectErr: false},
		{in: -1, expect: uintptr(0), err: nil, expectErr: true},
		{in: "-1", expect: uintptr(0), err: nil, expectErr: true},
		{in: -1.9, expect: uintptr(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uintptr(0), err: nil, expectErr: true},
	},
}
