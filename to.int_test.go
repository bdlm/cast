package cast_test

import (
	"testing"
)

func TestComplexToInt(t *testing.T) {
	var tests = []testCase{
		{complex64(1), 0, nil, true},
		{complex64(-1), 0, nil, true},
		{complex128(1), 0, nil, true},
		{complex128(-1), 0, nil, true},
	}
	testSimpleCases[int](t, tests)
	testSimpleCases[int8](t, tests)
	testSimpleCases[int16](t, tests)
	testSimpleCases[int32](t, tests)
	testSimpleCases[int64](t, tests)
}

func TestComplexToUint(t *testing.T) {
	var tests = []testCase{
		{complex64(1), 0, nil, true},
		{complex64(-1), 0, nil, true},
		{complex128(1), 0, nil, true},
		{complex128(-1), 0, nil, true},
	}
	testSimpleCases[uint](t, tests)
	testSimpleCases[uint8](t, tests)
	testSimpleCases[uint16](t, tests)
	testSimpleCases[uint32](t, tests)
	testSimpleCases[uint64](t, tests)
}

func TestSliceToInt(t *testing.T) {
	var tests = []testCase{
		{[]int{1, 2}, 0, nil, true},
		{[]string{"1", "2"}, 0, nil, true},
	}
	testSimpleCases[int](t, tests)
	testSimpleCases[int8](t, tests)
	testSimpleCases[int16](t, tests)
	testSimpleCases[int32](t, tests)
	testSimpleCases[int64](t, tests)
}

func TestSliceToUint(t *testing.T) {
	var tests = []testCase{
		{[]int{1, 2}, 0, nil, true},
		{[]string{"1", "2"}, 0, nil, true},
	}
	testSimpleCases[uint](t, tests)
	testSimpleCases[uint8](t, tests)
	testSimpleCases[uint16](t, tests)
	testSimpleCases[uint32](t, tests)
	testSimpleCases[uint64](t, tests)
}

func TestBoolToInt(t *testing.T) {
	testSimpleCases[int](t, []testCase{
		{true, int(1), nil, false},
		{false, int(0), nil, false},
	})
	testSimpleCases[int8](t, []testCase{
		{true, int8(1), nil, false},
		{false, int8(0), nil, false},
	})
	testSimpleCases[int16](t, []testCase{
		{true, int16(1), nil, false},
		{false, int16(0), nil, false},
	})
	testSimpleCases[int32](t, []testCase{
		{true, int32(1), nil, false},
		{false, int32(0), nil, false},
	})
	testSimpleCases[int64](t, []testCase{
		{true, int64(1), nil, false},
		{false, int64(0), nil, false},
	})
}

func TestBoolToUint(t *testing.T) {
	testSimpleCases[uint](t, []testCase{
		{true, uint(1), nil, false},
		{false, uint(0), nil, false},
	})
	testSimpleCases[uint8](t, []testCase{
		{true, uint8(1), nil, false},
		{false, uint8(0), nil, false},
	})
	testSimpleCases[uint16](t, []testCase{
		{true, uint16(1), nil, false},
		{false, uint16(0), nil, false},
	})
	testSimpleCases[uint32](t, []testCase{
		{true, uint32(1), nil, false},
		{false, uint32(0), nil, false},
	})
	testSimpleCases[uint64](t, []testCase{
		{true, uint64(1), nil, false},
		{false, uint64(0), nil, false},
	})
	testSimpleCases[uintptr](t, []testCase{
		{true, uintptr(1), nil, false},
		{false, uintptr(0), nil, false},
	})
}

func TestStrToInt(t *testing.T) {
	testSimpleCases[int](t, []testCase{
		{"1", int(1), nil, false},
		{"0", int(0), nil, false},
		{"-1", int(-1), nil, false},
		{"1.1", int(1), nil, false},
		{"1.4", int(1), nil, false},
		{"1.5", int(1), nil, false},
		{"1.9", int(1), nil, false},
		{"0.0", int(0), nil, false},
		{"-1.1", int(-1), nil, false},
		{"-1.4", int(-1), nil, false},
		{"-1.1", int(-1), nil, false},
		{"-1.1", int(-1), nil, false},
		{"1,000", int(1000), nil, false},
		{"1,000,000", int(1000000), nil, false},
		{"Hi", int(0), nil, true},
	})
	testSimpleCases[int8](t, []testCase{
		{"1", int8(1), nil, false},
		{"0", int8(0), nil, false},
		{"-1", int8(-1), nil, false},
		{"1.1", int8(1), nil, false},
		{"1.4", int8(1), nil, false},
		{"1.5", int8(1), nil, false},
		{"1.9", int8(1), nil, false},
		{"0.0", int8(0), nil, false},
		{"-1.1", int8(-1), nil, false},
		{"-1.4", int8(-1), nil, false},
		{"-1.1", int8(-1), nil, false},
		{"-1.1", int8(-1), nil, false},
		{"1,000", int8(-24), nil, false},
		{"1,000,000", int8(64), nil, false},
		{"Hi", int8(0), nil, true},
	})
	testSimpleCases[int16](t, []testCase{
		{"1", int16(1), nil, false},
		{"0", int16(0), nil, false},
		{"-1", int16(-1), nil, false},
		{"1.1", int16(1), nil, false},
		{"1.4", int16(1), nil, false},
		{"1.5", int16(1), nil, false},
		{"1.9", int16(1), nil, false},
		{"0.0", int16(0), nil, false},
		{"-1.1", int16(-1), nil, false},
		{"-1.4", int16(-1), nil, false},
		{"-1.1", int16(-1), nil, false},
		{"-1.1", int16(-1), nil, false},
		{"1,000", int16(1000), nil, false},
		{"1,000,000", int16(16960), nil, false},
		{"Hi", int16(0), nil, true},
	})
	testSimpleCases[int32](t, []testCase{
		{"1", int32(1), nil, false},
		{"0", int32(0), nil, false},
		{"-1", int32(-1), nil, false},
		{"1.1", int32(1), nil, false},
		{"1.4", int32(1), nil, false},
		{"1.5", int32(1), nil, false},
		{"1.9", int32(1), nil, false},
		{"0.0", int32(0), nil, false},
		{"-1.1", int32(-1), nil, false},
		{"-1.4", int32(-1), nil, false},
		{"-1.1", int32(-1), nil, false},
		{"-1.1", int32(-1), nil, false},
		{"1,000", int32(1000), nil, false},
		{"1,000,000", int32(1000000), nil, false},
		{"Hi", int32(0), nil, true},
	})
	testSimpleCases[int64](t, []testCase{
		{"1", int64(1), nil, false},
		{"0", int64(0), nil, false},
		{"-1", int64(-1), nil, false},
		{"1.1", int64(1), nil, false},
		{"1.4", int64(1), nil, false},
		{"1.5", int64(1), nil, false},
		{"1.9", int64(1), nil, false},
		{"0.0", int64(0), nil, false},
		{"-1.1", int64(-1), nil, false},
		{"-1.4", int64(-1), nil, false},
		{"-1.1", int64(-1), nil, false},
		{"-1.1", int64(-1), nil, false},
		{"1,000", int64(1000), nil, false},
		{"1,000,000", int64(1000000), nil, false},
		{"Hi", int64(0), nil, true},
	})
}

func TestStrToUint(t *testing.T) {
	testSimpleCases[uint](t, []testCase{
		{"1", uint(1), nil, false},
		{"0", uint(0), nil, false},
		{"-1", uint(0), nil, true},
		{"1.1", uint(1), nil, false},
		{"1.4", uint(1), nil, false},
		{"1.5", uint(1), nil, false},
		{"1.9", uint(1), nil, false},
		{"0.0", uint(0), nil, false},
		{"-1.1", uint(0), nil, true},
		{"-1.4", uint(0), nil, true},
		{"-1.5", uint(0), nil, true},
		{"-1.9", uint(0), nil, true},
		{"100", uint(100), nil, false},
		{"1,000", uint(1000), nil, false},
		{"1,000,000", uint(1000000), nil, false},
		{"Hi", uint(0), nil, true},
	})
	testSimpleCases[uint8](t, []testCase{
		{"1", uint8(1), nil, false},
		{"0", uint8(0), nil, false},
		{"-1", uint8(0), nil, true},
		{"1.1", uint8(1), nil, false},
		{"1.4", uint8(1), nil, false},
		{"1.5", uint8(1), nil, false},
		{"1.9", uint8(1), nil, false},
		{"0.0", uint8(0), nil, false},
		{"-1.1", uint8(0), nil, true},
		{"-1.4", uint8(0), nil, true},
		{"-1.5", uint8(0), nil, true},
		{"-1.9", uint8(0), nil, true},
		{"100", uint8(100), nil, false},
		{"1,000", uint8(232), nil, false},
		{"1,000,000", uint8(64), nil, false},
		{"Hi", uint8(0), nil, true},
	})
	testSimpleCases[uint16](t, []testCase{
		{"1", uint16(1), nil, false},
		{"0", uint16(0), nil, false},
		{"-1", uint16(0), nil, true},
		{"1.1", uint16(1), nil, false},
		{"1.4", uint16(1), nil, false},
		{"1.5", uint16(1), nil, false},
		{"1.9", uint16(1), nil, false},
		{"0.0", uint16(0), nil, false},
		{"-1.1", uint16(0), nil, true},
		{"-1.4", uint16(0), nil, true},
		{"-1.5", uint16(0), nil, true},
		{"-1.9", uint16(0), nil, true},
		{"100", uint16(100), nil, false},
		{"1,000", uint16(1000), nil, false},
		{"1,000,000", uint16(16960), nil, false},
		{"Hi", uint16(0), nil, true},
	})
	testSimpleCases[uint32](t, []testCase{
		{"1", uint32(1), nil, false},
		{"0", uint32(0), nil, false},
		{"-1", uint32(0), nil, true},
		{"1.1", uint32(1), nil, false},
		{"1.4", uint32(1), nil, false},
		{"1.5", uint32(1), nil, false},
		{"1.9", uint32(1), nil, false},
		{"0.0", uint32(0), nil, false},
		{"-1.1", uint32(0), nil, true},
		{"-1.4", uint32(0), nil, true},
		{"-1.5", uint32(0), nil, true},
		{"-1.9", uint32(0), nil, true},
		{"100", uint32(100), nil, false},
		{"1,000", uint32(1000), nil, false},
		{"1,000,000", uint32(1000000), nil, false},
		{"Hi", uint32(0), nil, true},
	})
	testSimpleCases[uint64](t, []testCase{
		{"1", uint64(1), nil, false},
		{"0", uint64(0), nil, false},
		{"-1", uint64(0), nil, true},
		{"1.1", uint64(1), nil, false},
		{"1.4", uint64(1), nil, false},
		{"1.5", uint64(1), nil, false},
		{"1.9", uint64(1), nil, false},
		{"0.0", uint64(0), nil, false},
		{"-1.1", uint64(0), nil, true},
		{"-1.4", uint64(0), nil, true},
		{"-1.5", uint64(0), nil, true},
		{"-1.9", uint64(0), nil, true},
		{"100", uint64(100), nil, false},
		{"1,000", uint64(1000), nil, false},
		{"1,000,000", uint64(1000000), nil, false},
		{"Hi", uint64(0), nil, true},
	})
	testSimpleCases[uintptr](t, []testCase{
		{"1", uintptr(1), nil, false},
		{"0", uintptr(0), nil, false},
		{"-1", uintptr(0), nil, true},
		{"1.1", uintptr(1), nil, false},
		{"1.4", uintptr(1), nil, false},
		{"1.5", uintptr(1), nil, false},
		{"1.9", uintptr(1), nil, false},
		{"0.0", uintptr(0), nil, false},
		{"-1.1", uintptr(0), nil, true},
		{"-1.4", uintptr(0), nil, true},
		{"-1.5", uintptr(0), nil, true},
		{"-1.9", uintptr(0), nil, true},
		{"100", uintptr(100), nil, false},
		{"1,000", uintptr(1000), nil, false},
		{"1,000,000", uintptr(1000000), nil, false},
		{"Hi", uintptr(0), nil, true},
	})
}
