package cast_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestToChanBool(t *testing.T) {
	testChanCases[bool](t, []testCase{
		{in: true, expect: true, err: nil, expectErr: false},
		{in: 1, expect: true, err: nil, expectErr: false},
		{in: 0, expect: false, err: nil, expectErr: false},
		{in: "hi", expect: false, err: nil, expectErr: true},
		{in: float64(1.1), expect: true, err: nil, expectErr: false},
		{in: float64(-1.1), expect: true, err: nil, expectErr: false},
	})
}

func TestToChanByte(t *testing.T) {
	testChanCases[byte](t, []testCase{
		{in: "a", expect: byte(0), err: nil, expectErr: true},
		{in: byte(1), expect: byte(1), err: nil, expectErr: false},
		{in: byte(0), expect: byte(0), err: nil, expectErr: false},
		{in: "hi", expect: byte(0), err: nil, expectErr: true},
		{in: float64(1.1), expect: byte(1), err: nil, expectErr: false},
		{in: float64(-1.1), expect: byte(0), err: nil, expectErr: true},
	})
}

func TestToChanInt(t *testing.T) {
	testChanCases[int](t, []testCase{
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
		{in: "Hi!", expect: int(0), err: nil, expectErr: true},
	})
	testChanCases[int8](t, []testCase{
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
		{in: "Hi!", expect: int8(0), err: nil, expectErr: true},
	})
	testChanCases[int16](t, []testCase{
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
		{in: "Hi!", expect: int16(0), err: nil, expectErr: true},
	})
	testChanCases[int32](t, []testCase{
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
		{in: "Hi!", expect: int32(0), err: nil, expectErr: true},
	})
	testChanCases[int64](t, []testCase{
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
		{in: "Hi!", expect: int64(0), err: nil, expectErr: true},
	})
}

func TestToChanUint(t *testing.T) {
	testChanCases[uint](t, []testCase{
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
		{in: "Hi!", expect: uint(0), err: nil, expectErr: true},
	})
	testChanCases[uint8](t, []testCase{
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
		{in: "Hi!", expect: uint8(0), err: nil, expectErr: true},
	})
	testChanCases[uint16](t, []testCase{
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
		{in: "Hi!", expect: uint16(0), err: nil, expectErr: true},
	})
	testChanCases[uint32](t, []testCase{
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
		{in: "Hi!", expect: uint32(0), err: nil, expectErr: true},
	})
	testChanCases[uint64](t, []testCase{
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
		{in: "Hi!", expect: uint64(0), err: nil, expectErr: true},
	})
	testChanCases[uintptr](t, []testCase{
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
	})
}

func TestToChanFloat(t *testing.T) {
	testChanCases[float32](t, []testCase{
		{in: 1, expect: float32(1), err: nil, expectErr: false},
		{in: "1", expect: float32(1), err: nil, expectErr: false},
		{in: 1.1, expect: float32(1.1), err: nil, expectErr: false},
		{in: "1.1", expect: float32(1.1), err: nil, expectErr: false},
		{in: 1.9, expect: float32(1.9), err: nil, expectErr: false},
		{in: "1.9", expect: float32(1.9), err: nil, expectErr: false},
		{in: -1, expect: float32(-1), err: nil, expectErr: false},
		{in: "-1", expect: float32(-1), err: nil, expectErr: false},
		{in: -1.9, expect: float32(-1.9), err: nil, expectErr: false},
		{in: "-1.9", expect: float32(-1.9), err: nil, expectErr: false},
		{in: "Hi!", expect: float32(0), err: nil, expectErr: true},
	})
	testChanCases[float64](t, []testCase{
		{in: 1, expect: float64(1), err: nil, expectErr: false},
		{in: "1", expect: float64(1), err: nil, expectErr: false},
		{in: 1.1, expect: float64(1.1), err: nil, expectErr: false},
		{in: "1.1", expect: float64(1.1), err: nil, expectErr: false},
		{in: 1.9, expect: float64(1.9), err: nil, expectErr: false},
		{in: "1.9", expect: float64(1.9), err: nil, expectErr: false},
		{in: -1, expect: float64(-1), err: nil, expectErr: false},
		{in: "-1", expect: float64(-1), err: nil, expectErr: false},
		{in: -1.9, expect: float64(-1.9), err: nil, expectErr: false},
		{in: "-1.9", expect: float64(-1.9), err: nil, expectErr: false},
		{in: "Hi!", expect: float64(0), err: nil, expectErr: true},
	})
}

func TestToChanComplex(t *testing.T) {
	testChanCases[complex64](t, []testCase{
		{in: 1, expect: complex64(1), err: nil, expectErr: false},
		{in: "1", expect: complex64(1), err: nil, expectErr: false},
		{in: 1.1, expect: complex64(1.1), err: nil, expectErr: false},
		{in: "1.1", expect: complex64(1.1), err: nil, expectErr: false},
		{in: 1.9, expect: complex64(1.9), err: nil, expectErr: false},
		{in: "1.9", expect: complex64(1.9), err: nil, expectErr: false},
		{in: -1, expect: complex64(-1), err: nil, expectErr: false},
		{in: "-1", expect: complex64(-1), err: nil, expectErr: false},
		{in: -1.9, expect: complex64(-1.9), err: nil, expectErr: false},
		{in: "-1.9", expect: complex64(-1.9), err: nil, expectErr: false},
		{in: "Hi!", expect: complex64(0), err: nil, expectErr: true},
	})
	testChanCases[complex128](t, []testCase{
		{in: 1, expect: complex128(1), err: nil, expectErr: false},
		{in: "1", expect: complex128(1), err: nil, expectErr: false},
		{in: 1.1, expect: complex128(1.1), err: nil, expectErr: false},
		{in: "1.1", expect: complex128(1.1), err: nil, expectErr: false},
		{in: 1.9, expect: complex128(1.9), err: nil, expectErr: false},
		{in: "1.9", expect: complex128(1.9), err: nil, expectErr: false},
		{in: -1, expect: complex128(-1), err: nil, expectErr: false},
		{in: "-1", expect: complex128(-1), err: nil, expectErr: false},
		{in: -1.9, expect: complex128(-1.9), err: nil, expectErr: false},
		{in: "-1.9", expect: complex128(-1.9), err: nil, expectErr: false},
		{in: "Hi!", expect: complex128(0), err: nil, expectErr: true},
	})
}

func testChanCases[TTo any](t *testing.T, cases []testCase) {
	var typ TTo
	name := fmt.Sprintf("%T", typ)

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
			var result TTo
			actual, err := cast.ToE[chan TTo](test.in)
			if nil == err {
				result = <-actual
			}
			testInfo := fmt.Sprintf(`
case: ToE[chan %s]
input: %v (%T)
expect error: %v; actual error: %v
expected result: %v (%T); actual result: %v (%T)
result chan: %#v, (%T)
test: %#v
			`,
				name,
				test.in,
				test.in,
				test.expectErr,
				err,
				test.expect,
				test.expect,
				result,
				result,
				actual,
				actual,
				test,
			)
			// fmt.Println(testInfo)
			// fmt.Printf("\n-----------------\n%v (%T): %v (%T)\n-----------------\n", result, result, test.expect, test.expect)
			if err != nil && !test.expectErr {
				t.Error("1. expected nil, got error", testInfo)
			} else if err == nil && test.expectErr {
				t.Error("2. expected error, got nil", testInfo)
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Error("3. expected cast.Error, got different error type", testInfo)
			} else if nil == err && !reflect.DeepEqual(result, test.expect.(TTo)) {
				t.Errorf("4. expected %v to equal %v %s", test.expect, actual, testInfo)
			}
		})
	}
}
