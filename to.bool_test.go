package cast_test

import (
	"testing"
)

func TestBoolToBool(t *testing.T) {
	testSimpleCases[bool](t, []testCase{
		{true, true, nil, false},
		{false, false, nil, false},
	})
}

// ToBool casts an interface to a bool type.
func TestVariousToBool(t *testing.T) {
	var nilInterface interface{} = nil
	testSimpleCases[bool](t, []testCase{
		{true, true, nil, false},
		{false, false, nil, false},
		{"0", false, nil, false},
		{"-1", true, nil, false},
		{"1.1", true, nil, false},
		{"1.4", true, nil, false},
		{"1.5", true, nil, false},
		{"1.9", true, nil, false},
		{"0.0", false, nil, false},
		{"0.1", true, nil, false},
		{"-1.1", true, nil, false},
		{"-1.4", true, nil, false},
		{"-1.1", true, nil, false},
		{"-1.1", true, nil, false},
		{"1,000", true, nil, false},
		{"1,000,000", true, nil, false},
		{'a', true, nil, false},
		{'1', true, nil, false},
		{'0', false, nil, false},
		{nilInterface, false, nil, false},
		{"Hi", false, nil, true},
	})
}

func TestByteToBool(t *testing.T) {
	testSimpleCases[bool](t, []testCase{
		{byte(1), true, nil, false},
		{byte(1), true, nil, false},
		{byte(2), true, nil, false},
		{byte(0), false, nil, false},
		{'a', true, nil, false},
		{'1', true, nil, false},
		{'0', false, nil, false},
	})
}

func TestComplexToBool(t *testing.T) {
	testSimpleCases[bool](t, []testCase{
		{complex(1, 0), true, nil, false},
		{complex(1, 1), true, nil, false},
		{complex(0, 0), false, nil, false},
		{complex(-1, 0), true, nil, false},
		{complex(-1, -1), true, nil, false},
		{complex64(float32(1.1)), true, nil, false},
		{complex64(float32(-1.1)), true, nil, false},
		{complex64(float32(0.0)), false, nil, false},
		{complex128(float32(1.1)), true, nil, false},
		{complex128(float32(-1.1)), true, nil, false},
		{complex128(float32(0.0)), false, nil, false},
		{complex64(float64(1.1)), true, nil, false},
		{complex64(float64(-1.1)), true, nil, false},
		{complex64(float64(0.0)), false, nil, false},
		{complex128(float64(1.1)), true, nil, false},
		{complex128(float64(-1.1)), true, nil, false},
		{complex128(float64(0.0)), false, nil, false},
	})
}

func TestFloatToBool(t *testing.T) {
	testSimpleCases[bool](t, []testCase{
		{1.0, true, nil, false},
		{1.1, true, nil, false},
		{0.0, false, nil, false},
		{-1.0, true, nil, false},
		{-1.1, true, nil, false},
		{float32(1.0), true, nil, false},
		{float32(1.1), true, nil, false},
		{float32(0.0), false, nil, false},
		{float32(-1.0), true, nil, false},
		{float32(-1.1), true, nil, false},
		{float64(1.0), true, nil, false},
		{float64(1.1), true, nil, false},
		{float64(0.0), false, nil, false},
		{float64(-1.0), true, nil, false},
		{float64(-1.1), true, nil, false},
	})
}

func TestIntToBool(t *testing.T) {
	testSimpleCases[bool](t, []testCase{
		{int(1), true, nil, false},
		{int(0), false, nil, false},
		{int(-1), true, nil, false},
		{int8(1), true, nil, false},
		{int8(0), false, nil, false},
		{int8(-1), true, nil, false},
		{int16(1), true, nil, false},
		{int16(0), false, nil, false},
		{int16(-1), true, nil, false},
		{int32(1), true, nil, false},
		{int32(0), false, nil, false},
		{int32(-1), true, nil, false},
		{int64(1), true, nil, false},
		{int64(0), false, nil, false},
		{int64(-1), true, nil, false},
	})
}

func TestStringToBool(t *testing.T) {
	testSimpleCases[bool](t, []testCase{
		{"1", true, nil, false},
		{"0", false, nil, false},
		{"-1", true, nil, false},
		{"Hi!", false, nil, true},
		{'a', true, nil, false},
		{'1', true, nil, false},
		{'0', false, nil, false},
	})
}
