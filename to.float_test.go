package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestComplexToFloat(t *testing.T) {
	testSimpleCases[float32](t, []testCase{
		{complex(1, 0), 0, nil, true},
		{complex(1, 1), 0, nil, true},
		{complex(0, 0), 0, nil, true},
		{complex(-1, 0), 0, nil, true},
		{complex(-1, -1), 0, nil, true},
		{complex64(float32(1.1)), 0, nil, true},
		{complex64(float32(-1.1)), 0, nil, true},
		{complex64(float32(0.0)), 0, nil, true},
		{complex128(float32(1.1)), 0, nil, true},
		{complex128(float32(-1.1)), 0, nil, true},
		{complex128(float32(0.0)), 0, nil, true},
		{complex64(float64(1.1)), 0, nil, true},
		{complex64(float64(-1.1)), 0, nil, true},
		{complex64(float64(0.0)), 0, nil, true},
		{complex128(float64(1.1)), 0, nil, true},
		{complex128(float64(-1.1)), 0, nil, true},
		{complex128(float64(0.0)), 0, nil, true},
	})
	testSimpleCases[float64](t, []testCase{
		{complex(1, 0), 0, nil, true},
		{complex(1, 1), 0, nil, true},
		{complex(0, 0), 0, nil, true},
		{complex(-1, 0), 0, nil, true},
		{complex(-1, -1), 0, nil, true},
		{complex64(float32(1.1)), 0, nil, true},
		{complex64(float32(-1.1)), 0, nil, true},
		{complex64(float32(0.0)), 0, nil, true},
		{complex128(float32(1.1)), 0, nil, true},
		{complex128(float32(-1.1)), 0, nil, true},
		{complex128(float32(0.0)), 0, nil, true},
		{complex64(float64(1.1)), 0, nil, true},
		{complex64(float64(-1.1)), 0, nil, true},
		{complex64(float64(0.0)), 0, nil, true},
		{complex128(float64(1.1)), 0, nil, true},
		{complex128(float64(-1.1)), 0, nil, true},
		{complex128(float64(0.0)), 0, nil, true},
	})
}

func TestArrayToFloat(t *testing.T) {
	testSimpleCases[float32](t, []testCase{
		{[]int{1, 2}, 0, nil, true},
		{[]string{"1", "2"}, 0, nil, true},
	})
	testSimpleCases[float64](t, []testCase{
		{[]int{1, 2}, 0, nil, true},
		{[]string{"1", "2"}, 0, nil, true},
	})
}

func TestBoolToFloat(t *testing.T) {
	testSimpleCases[float32](t, []testCase{
		{true, float32(1), nil, false},
		{false, float32(0), nil, false},
	})
	testSimpleCases[float64](t, []testCase{
		{true, float64(1), nil, false},
		{false, float64(0), nil, false},
	})
}

func TestStrToFloat(t *testing.T) {
	testSimpleCases[float32](t, []testCase{
		{"1", float32(1), nil, false},
		{"0", float32(0), nil, false},
		{"-1", float32(-1), nil, false},
		{"1.1", float32(1.1), nil, false},
		{"1.4", float32(1.4), nil, false},
		{"1.5", float32(1.5), nil, false},
		{"1.9", float32(1.9), nil, false},
		{"0.0", float32(0.0), nil, false},
		{"0.1", float32(0.1), nil, false},
		{"-1.1", float32(-1.1), nil, false},
		{"-1.4", float32(-1.4), nil, false},
		{"-1.1", float32(-1.1), nil, false},
		{"-1.1", float32(-1.1), nil, false},
		{"1,000", float32(1000), nil, false},
		{"1,000,000", float32(1000000), nil, false},
		{"Hi", float32(0), nil, true},
	})
	testSimpleCases[float64](t, []testCase{
		{"1", float64(1), nil, false},
		{"0", float64(0), nil, false},
		{"-1", float64(-1), nil, false},
		{"1.1", float64(1.1), nil, false},
		{"1.4", float64(1.4), nil, false},
		{"1.5", float64(1.5), nil, false},
		{"1.9", float64(1.9), nil, false},
		{"0.0", float64(0.0), nil, false},
		{"0.1", float64(0.1), nil, false},
		{"-1.1", float64(-1.1), nil, false},
		{"-1.4", float64(-1.4), nil, false},
		{"-1.1", float64(-1.1), nil, false},
		{"-1.1", float64(-1.1), nil, false},
		{"1,000", float64(1000), nil, false},
		{"1,000,000", float64(1000000), nil, false},
		{"Hi", float64(0), nil, true},
	})
}

func TestIntToFloat(t *testing.T) {
	testSimpleCases[float32](t, []testCase{
		{int(1), float32(1), nil, false},
		{int(0), float32(0), nil, false},
		{int(-1), float32(-1), nil, false},
		{int8(1), float32(1), nil, false},
		{int8(0), float32(0), nil, false},
		{int8(-1), float32(-1), nil, false},
		{int16(1), float32(1), nil, false},
		{int16(0), float32(0), nil, false},
		{int16(-1), float32(-1), nil, false},
		{int32(1), float32(1), nil, false},
		{int32(0), float32(0), nil, false},
		{int32(-1), float32(-1), nil, false},
		{int64(1), float32(1), nil, false},
		{int64(0), float32(0), nil, false},
		{int64(-1), float32(-1), nil, false},
	})
	testSimpleCases[float64](t, []testCase{
		{int(1), float64(1), nil, false},
		{int(0), float64(0), nil, false},
		{int(-1), float64(-1), nil, false},
		{int8(1), float64(1), nil, false},
		{int8(0), float64(0), nil, false},
		{int8(-1), float64(-1), nil, false},
		{int16(1), float64(1), nil, false},
		{int16(0), float64(0), nil, false},
		{int16(-1), float64(-1), nil, false},
		{int32(1), float64(1), nil, false},
		{int32(0), float64(0), nil, false},
		{int32(-1), float64(-1), nil, false},
		{int64(1), float64(1), nil, false},
		{int64(0), float64(0), nil, false},
		{int64(-1), float64(-1), nil, false},
	})
}

func TestUintToFloat(t *testing.T) {
	testSimpleCases[float32](t, []testCase{
		{uint(1), float32(1), nil, false},
		{uint(0), float32(0), nil, false},
		{uint8(1), float32(1), nil, false},
		{uint8(0), float32(0), nil, false},
		{uint16(1), float32(1), nil, false},
		{uint16(0), float32(0), nil, false},
		{uint32(1), float32(1), nil, false},
		{uint32(0), float32(0), nil, false},
		{uint64(1), float32(1), nil, false},
		{uint64(0), float32(0), nil, false},
		{uintptr(1), float32(1), nil, false},
		{uintptr(0), float32(0), nil, false},
	})
	testSimpleCases[float64](t, []testCase{
		{uint(1), float64(1), nil, false},
		{uint(0), float64(0), nil, false},
		{uint8(1), float64(1), nil, false},
		{uint8(0), float64(0), nil, false},
		{uint16(1), float64(1), nil, false},
		{uint16(0), float64(0), nil, false},
		{uint32(1), float64(1), nil, false},
		{uint32(0), float64(0), nil, false},
		{uint64(1), float64(1), nil, false},
		{uint64(0), float64(0), nil, false},
		{uintptr(1), float64(1), nil, false},
		{uintptr(0), float64(0), nil, false},
	})
}

func TestFloatToFloat(t *testing.T) {
	testSimpleCases[float32](t, []testCase{
		{float32(1.5), float32(1.5), nil, false},
		{float32(0.0), float32(0.0), nil, false},
		{float32(-1.5), float32(-1.5), nil, false},
		{float64(3.14), float32(float64(3.14)), nil, false},
	})
	testSimpleCases[float64](t, []testCase{
		{float64(3.14), float64(3.14), nil, false},
		{float64(0.0), float64(0.0), nil, false},
		{float64(-3.14), float64(-3.14), nil, false},
		{float32(1.5), float64(float32(1.5)), nil, false},
	})
}

func TestNilToFloat(t *testing.T) {
	testSimpleCases[float32](t, []testCase{
		{nil, float32(0), nil, false},
	})
	testSimpleCases[float64](t, []testCase{
		{nil, float64(0), nil, false},
	})
}

func TestStringerToFloat(t *testing.T) {
	testSimpleCases[float32](t, []testCase{
		{testStringer{"1.5"}, float32(1.5), nil, false},
		{testStringer{"bad"}, float32(0), nil, true},
	})
	testSimpleCases[float64](t, []testCase{
		{testStringer{"3.14"}, float64(3.14), nil, false},
		{testStringer{"bad"}, float64(0), nil, true},
	})
}

func TestFloatDefaultOption(t *testing.T) {
	t.Run("DEFAULT used on error", func(t *testing.T) {
		result := cast.To[float64]("bad", cast.Op{cast.DEFAULT, float64(9.9)})
		if result != 9.9 {
			t.Errorf("expected 9.9, got %v", result)
		}
	})
	t.Run("DEFAULT wrong type errors immediately", func(t *testing.T) {
		_, err := cast.ToE[float64]("bad", cast.Op{cast.DEFAULT, "not a float"})
		if err == nil {
			t.Fatal("expected error for wrong DEFAULT type, got nil")
		}
	})
	t.Run("DEFAULT not used when conversion succeeds", func(t *testing.T) {
		result := cast.To[float64]("1.5", cast.Op{cast.DEFAULT, float64(9.9)})
		if result != 1.5 {
			t.Errorf("expected 1.5, got %v", result)
		}
	})
}
