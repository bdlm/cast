package cast_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestToFuncBool(t *testing.T) {
	testFuncCases[bool](t, []testCase{
		{in: true, expect: true, err: nil, expectErr: false},
		{in: 1, expect: true, err: nil, expectErr: false},
		{in: 0, expect: false, err: nil, expectErr: false},
		{in: "hi", expect: false, err: nil, expectErr: true},
		{in: float64(1.1), expect: true, err: nil, expectErr: false},
		{in: float64(-1.1), expect: true, err: nil, expectErr: false},
	})
}

func TestToFuncByte(t *testing.T) {
	testFuncCases[byte](t, []testCase{
		{in: "a", expect: byte(0), err: nil, expectErr: true},
		{in: byte(1), expect: byte(1), err: nil, expectErr: false},
		{in: byte(0), expect: byte(0), err: nil, expectErr: false},
		{in: "hi", expect: byte(0), err: nil, expectErr: true},
		{in: float64(1.1), expect: byte(1), err: nil, expectErr: false},
		{in: float64(-1.1), expect: byte(0), err: nil, expectErr: true},
	})
}

func TestToFuncInt(t *testing.T) {
	testFuncCases[int](t, []testCase{
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
	testFuncCases[int8](t, []testCase{
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
	testFuncCases[int16](t, []testCase{
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
	testFuncCases[int32](t, []testCase{
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
	testFuncCases[int64](t, []testCase{
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

func TestToFuncUint(t *testing.T) {
	testFuncCases[uint](t, []testCase{
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
	testFuncCases[uint8](t, []testCase{
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
	testFuncCases[uint16](t, []testCase{
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
	testFuncCases[uint32](t, []testCase{
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
	testFuncCases[uint64](t, []testCase{
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
	testFuncCases[uintptr](t, []testCase{
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

func TestToFuncFloat(t *testing.T) {
	testFuncCases[float32](t, []testCase{
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
	testFuncCases[float64](t, []testCase{
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

func TestToFuncComplex(t *testing.T) {
	testFuncCases[complex64](t, []testCase{
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
	testFuncCases[complex128](t, []testCase{
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

func testFuncCases[TTo any](t *testing.T, cases []testCase) {
	var typ TTo
	name := fmt.Sprintf("%T", typ)

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
			var result TTo
			actual, err := cast.ToE[cast.Func[TTo]](test.in)
			if nil == err {
				result = actual()
			}
			testInfo := fmt.Sprintf(`
case: ToE[Func[%s]]
input: %v (%T)
expect error: %v; actual error: %v
expected result: %v (%T); actual result: %v (%T)
result fn: %#v, (%T)
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
			//fmt.Println(testInfo)
			if !test.expectErr && err != nil {
				t.Error(1, testInfo)
			} else if test.expectErr && err == nil {
				t.Error(2, testInfo)
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Error(3, testInfo)
			} else if nil == err && !reflect.DeepEqual(result, test.expect.(TTo)) {
				t.Error(4, testInfo)
			}
		})
	}
}
