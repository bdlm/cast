package cast_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestSliceToBoolSlice(t *testing.T) {
	testSliceCases[[]bool](t, []testCase{
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
	})
}

func TestSliceToIntSlice(t *testing.T) {
	testSliceCases[[]int](t, []testCase{
		{in: []bool{true, false}, expect: []int{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []int{}, err: nil, expectErr: true},
		{in: "1", expect: []int{}, err: nil, expectErr: true},
		{in: -1, expect: []int{}, err: nil, expectErr: true},
		{in: "-1", expect: []int{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]int8](t, []testCase{
		{in: []bool{true, false}, expect: []int8{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int8{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int8{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int8{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []int8{}, err: nil, expectErr: true},
		{in: "1", expect: []int8{}, err: nil, expectErr: true},
		{in: -1, expect: []int8{}, err: nil, expectErr: true},
		{in: "-1", expect: []int8{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int8{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]int16](t, []testCase{
		{in: []bool{true, false}, expect: []int16{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int16{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int16{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int16{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []int16{}, err: nil, expectErr: true},
		{in: "1", expect: []int16{}, err: nil, expectErr: true},
		{in: -1, expect: []int16{}, err: nil, expectErr: true},
		{in: "-1", expect: []int16{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int16{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]int32](t, []testCase{
		{in: []bool{true, false}, expect: []int32{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int32{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int32{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []int32{}, err: nil, expectErr: true},
		{in: "1", expect: []int32{}, err: nil, expectErr: true},
		{in: -1, expect: []int32{}, err: nil, expectErr: true},
		{in: "-1", expect: []int32{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int32{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]int64](t, []testCase{
		{in: []bool{true, false}, expect: []int64{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int64{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int64{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []int64{}, err: nil, expectErr: true},
		{in: "1", expect: []int64{}, err: nil, expectErr: true},
		{in: -1, expect: []int64{}, err: nil, expectErr: true},
		{in: "-1", expect: []int64{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int64{1}, err: nil, expectErr: false},
	})
}

func TestSliceToUintSlice(t *testing.T) {
	testSliceCases[[]uint](t, []testCase{
		{in: []bool{true, false}, expect: []uint{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint{}, err: nil, expectErr: true},
		{in: "1", expect: []uint{}, err: nil, expectErr: true},
		{in: -1, expect: []uint{}, err: nil, expectErr: true},
		{in: "-1", expect: []uint{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]uint8](t, []testCase{
		{in: []bool{true, false}, expect: []uint8{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint8{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint8{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint8{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint8{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint8{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint8{}, err: nil, expectErr: true},
		{in: "1", expect: []uint8{}, err: nil, expectErr: true},
		{in: -1, expect: []uint8{}, err: nil, expectErr: true},
		{in: "-1", expect: []uint8{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint8{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint8{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint8{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint8{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]uint16](t, []testCase{
		{in: []bool{true, false}, expect: []uint16{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint16{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint16{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint16{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint16{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint16{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint16{}, err: nil, expectErr: true},
		{in: "1", expect: []uint16{}, err: nil, expectErr: true},
		{in: -1, expect: []uint16{}, err: nil, expectErr: true},
		{in: "-1", expect: []uint16{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint16{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint16{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint16{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint16{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]uint32](t, []testCase{
		{in: []bool{true, false}, expect: []uint32{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint32{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint32{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint32{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint32{}, err: nil, expectErr: true},
		{in: "1", expect: []uint32{}, err: nil, expectErr: true},
		{in: -1, expect: []uint32{}, err: nil, expectErr: true},
		{in: "-1", expect: []uint32{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint32{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint32{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint32{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint32{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]uint64](t, []testCase{
		{in: []bool{true, false}, expect: []uint64{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint64{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint64{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint64{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint64{}, err: nil, expectErr: true},
		{in: "1", expect: []uint64{}, err: nil, expectErr: true},
		{in: -1, expect: []uint64{}, err: nil, expectErr: true},
		{in: "-1", expect: []uint64{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint64{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint64{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint64{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint64{1}, err: nil, expectErr: false},
	})
}

func TestSliceToFloatSlice(t *testing.T) {
	testSliceCases[[]float32](t, []testCase{
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
	})
	testSliceCases[[]float64](t, []testCase{
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
	})
}

func TestSliceToComplexSlice(t *testing.T) {
	testSliceCases[[]complex64](t, []testCase{
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
	})
	testSliceCases[[]complex128](t, []testCase{
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
	})
}

func testSliceCases[TTo any](t *testing.T, cases []testCase) {
	var typ TTo
	name := fmt.Sprintf("%T", typ)

	for k, test := range cases {
		t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
			actual, err := cast.ToE[TTo](test.in)
			testInfo := fmt.Sprintf(`
case #%d: ToE[%s]
input: %v (%T)
expect error: %v; actual error: % #+v
expected result: %v (%T); actual result: %v (%T)
test: %#v
			`,
				k,
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
