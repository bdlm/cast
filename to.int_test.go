package cast_test

import (
	"errors"
	"math"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestComplexToInt(t *testing.T) {
	testSimpleCases[int](t, []testCase{
		{complex64(1), int(1), nil, false},
		{complex64(-1), int(-1), nil, false},
		{complex128(1), int(1), nil, false},
		{complex128(-1), int(-1), nil, false},
	})
	testSimpleCases[int8](t, []testCase{
		{complex64(1), int8(1), nil, false},
		{complex64(-1), int8(-1), nil, false},
		{complex128(1), int8(1), nil, false},
		{complex128(-1), int8(-1), nil, false},
	})
	testSimpleCases[int16](t, []testCase{
		{complex64(1), int16(1), nil, false},
		{complex64(-1), int16(-1), nil, false},
		{complex128(1), int16(1), nil, false},
		{complex128(-1), int16(-1), nil, false},
	})
	testSimpleCases[int32](t, []testCase{
		{complex64(1), int32(1), nil, false},
		{complex64(-1), int32(-1), nil, false},
		{complex128(1), int32(1), nil, false},
		{complex128(-1), int32(-1), nil, false},
	})
	testSimpleCases[int64](t, []testCase{
		{complex64(1), int64(1), nil, false},
		{complex64(-1), int64(-1), nil, false},
		{complex128(1), int64(1), nil, false},
		{complex128(-1), int64(-1), nil, false},
	})
}

func TestComplexToUint(t *testing.T) {
	testSimpleCases[uint](t, []testCase{
		{complex64(1), uint(1), nil, false},
		{complex64(-1), uint(0), nil, true},
		{complex128(1), uint(1), nil, false},
		{complex128(-1), uint(0), nil, true},
	})
	testSimpleCases[uint8](t, []testCase{
		{complex64(1), uint8(1), nil, false},
		{complex64(-1), uint8(0), nil, true},
		{complex128(1), uint8(1), nil, false},
		{complex128(-1), uint8(0), nil, true},
	})
	testSimpleCases[uint16](t, []testCase{
		{complex64(1), uint16(1), nil, false},
		{complex64(-1), uint16(0), nil, true},
		{complex128(1), uint16(1), nil, false},
		{complex128(-1), uint16(0), nil, true},
	})
	testSimpleCases[uint32](t, []testCase{
		{complex64(1), uint32(1), nil, false},
		{complex64(-1), uint32(0), nil, true},
		{complex128(1), uint32(1), nil, false},
		{complex128(-1), uint32(0), nil, true},
	})
	testSimpleCases[uint64](t, []testCase{
		{complex64(1), uint64(1), nil, false},
		{complex64(-1), uint64(0), nil, true},
		{complex128(1), uint64(1), nil, false},
		{complex128(-1), uint64(0), nil, true},
	})
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

// TestABSWithSignedIntTypes covers the TTo(-uint(val)) branch in toInt for each
// signed integer type when casting to an unsigned target, including the minimum
// value of each type (which would overflow with signed negation).
func TestABSWithSignedIntTypes(t *testing.T) {
	cases := []struct {
		name   string
		input  any
		expect uint
	}{
		{"int8 negative", int8(-3), uint(3)},
		{"int16 negative", int16(-10), uint(10)},
		{"int32 negative", int32(-100), uint(100)},
		{"int64 negative", int64(-999), uint(999)},
	}
	for _, tc := range cases {
		t.Run(tc.name+" → uint with ABS", func(t *testing.T) {
			result, err := cast.ToE[uint](tc.input, cast.Op{cast.ABS, true})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expect {
				t.Errorf("expected %v, got %v", tc.expect, result)
			}
		})
		t.Run(tc.name+" → uint without ABS errors", func(t *testing.T) {
			_, err := cast.ToE[uint](tc.input)
			if err == nil {
				t.Errorf("expected error for negative %T → uint without ABS", tc.input)
			}
		})
	}
}

// TestABSMinValues verifies that ABS with the minimum value of each signed type
// does not overflow. Signed negation of MinInt wraps; unsigned negation (-uint(v))
// produces the correct absolute value for all minimum values.
func TestABSMinValues(t *testing.T) {
	t.Run("int8 MinInt8 → uint8", func(t *testing.T) {
		result, err := cast.ToE[uint8](int8(math.MinInt8), cast.Op{cast.ABS, true})
		if err != nil || result != 128 {
			t.Errorf("expected 128, got %v (err %v)", result, err)
		}
	})
	t.Run("int16 MinInt16 → uint16", func(t *testing.T) {
		result, err := cast.ToE[uint16](int16(math.MinInt16), cast.Op{cast.ABS, true})
		if err != nil || result != 32768 {
			t.Errorf("expected 32768, got %v (err %v)", result, err)
		}
	})
	t.Run("int32 MinInt32 → uint32", func(t *testing.T) {
		result, err := cast.ToE[uint32](int32(math.MinInt32), cast.Op{cast.ABS, true})
		if err != nil || result != 2147483648 {
			t.Errorf("expected 2147483648, got %v (err %v)", result, err)
		}
	})
	t.Run("int64 MinInt64 → uint64", func(t *testing.T) {
		result, err := cast.ToE[uint64](int64(math.MinInt64), cast.Op{cast.ABS, true})
		if err != nil || result != 9223372036854775808 {
			t.Errorf("expected 9223372036854775808, got %v (err %v)", result, err)
		}
	})
}

// TestABSWithFloat32 covers the TTo(math.Floor(float64(-val))) branch in toInt
// for float32 inputs.
func TestABSWithFloat32(t *testing.T) {
	t.Run("negative float32 → uint with ABS=true", func(t *testing.T) {
		result, err := cast.ToE[uint](float32(-7.9), cast.Op{cast.ABS, true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != uint(7) {
			t.Errorf("expected 7 (floor of 7.9), got %v", result)
		}
	})
	t.Run("negative float32 → uint without ABS errors", func(t *testing.T) {
		_, err := cast.ToE[uint](float32(-7.9))
		if err == nil {
			t.Error("expected error for negative float32 → uint without ABS")
		}
	})
}

// TestToIntStructSource covers toInt's default: case — source types that match
// none of the explicit cases fall through to fmt.Sprintf, producing an
// unparseable string that strToInt rejects.
func TestToIntStructSource(t *testing.T) {
	_, err := cast.ToE[int](struct{}{})
	if err == nil {
		t.Error("expected error for struct{} source → int, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

// TestStrToIntDefaultWrongType covers the error-immediately path in strToInt
// when the DEFAULT option value has the wrong type for the target.
func TestStrToIntDefaultWrongType(t *testing.T) {
	_, err := cast.ToE[int]("bad", cast.Op{cast.DEFAULT, "fallback"})
	if err == nil {
		t.Fatal("expected error for string DEFAULT on int target, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

func TestStringerToInt(t *testing.T) {
	testSimpleCases[int](t, []testCase{
		{testStringer{"42"}, int(42), nil, false},
		{testStringer{"-7"}, int(-7), nil, false},
		{testStringer{"3.9"}, int(3), nil, false},
		{testStringer{"bad"}, int(0), nil, true},
	})
	testSimpleCases[uint](t, []testCase{
		{testStringer{"42"}, uint(42), nil, false},
		{testStringer{"-1"}, uint(0), nil, true},
		{testStringer{"bad"}, uint(0), nil, true},
	})
}

func TestNilToInt(t *testing.T) {
	testSimpleCases[int](t, []testCase{{nil, int(0), nil, false}})
	testSimpleCases[int8](t, []testCase{{nil, int8(0), nil, false}})
	testSimpleCases[int64](t, []testCase{{nil, int64(0), nil, false}})
	testSimpleCases[uint](t, []testCase{{nil, uint(0), nil, false}})
	testSimpleCases[uint64](t, []testCase{{nil, uint64(0), nil, false}})
}

func TestIntToInt(t *testing.T) {
	// Same-type identity and cross-width signed casts.
	testSimpleCases[int](t, []testCase{
		{int(7), int(7), nil, false},
		{int8(7), int(7), nil, false},
		{int16(7), int(7), nil, false},
		{int32(7), int(7), nil, false},
		{int64(7), int(7), nil, false},
	})
}

func TestUintToInt(t *testing.T) {
	testSimpleCases[int](t, []testCase{
		{uint(5), int(5), nil, false},
		{uint8(255), int(255), nil, false},
		{uint64(1000), int(1000), nil, false},
	})
	testSimpleCases[uint](t, []testCase{
		{uint(5), uint(5), nil, false},
		{uint8(255), uint(255), nil, false},
	})
}

func TestFloatToInt(t *testing.T) {
	// float64 → int routes through the string path for precision; float32 too.
	testSimpleCases[int](t, []testCase{
		{float64(3.7), int(3), nil, false},
		{float64(-3.7), int(-3), nil, false},
		{float32(2.9), int(2), nil, false},
		{float32(-2.9), int(-2), nil, false},
	})
	testSimpleCases[int64](t, []testCase{
		{float64(1000000.9), int64(1000000), nil, false},
	})
}

func TestFloatToUint(t *testing.T) {
	testSimpleCases[uint](t, []testCase{
		{float64(3.9), uint(3), nil, false},
		{float64(-1.0), uint(0), nil, true},
	})
}

func TestIntDefaultOption(t *testing.T) {
	t.Run("DEFAULT returned when input is invalid", func(t *testing.T) {
		result := cast.To[int]("not-a-number", cast.Op{cast.DEFAULT, int(42)})
		if result != 42 {
			t.Errorf("expected 42 (default), got %v", result)
		}
	})
	t.Run("DEFAULT not used when input is valid", func(t *testing.T) {
		result := cast.To[int]("7", cast.Op{cast.DEFAULT, int(42)})
		if result != 7 {
			t.Errorf("expected 7, got %v", result)
		}
	})
}

// Float-to-int truncation uses math.Floor for positive values and math.Ceil
// for negative values, which together produce truncation-toward-zero (same as
// C-style integer division, not banker's rounding). These tests lock that
// contract in place.

func TestFloatToIntTruncation(t *testing.T) {
	cases := []struct {
		name   string
		input  any
		expect int
	}{
		// positive: floor (truncate toward zero == truncate toward -∞)
		{"1.1 → 1 (floor)", "1.1", 1},
		{"1.5 → 1 (floor, not round)", "1.5", 1},
		{"1.9 → 1 (floor, not 2)", "1.9", 1},
		// negative: ceil (truncate toward zero == truncate toward +∞)
		{"-1.1 → -1 (ceil)", "-1.1", -1},
		{"-1.5 → -1 (ceil, not -2)", "-1.5", -1},
		{"-1.9 → -1 (ceil, not -2)", "-1.9", -1},
		// zero fractional part
		{"1.0 → 1", "1.0", 1},
		{"-1.0 → -1", "-1.0", -1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[int](tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expect {
				t.Errorf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

// Direct float64 → int (not via string) also truncates toward zero.
func TestFloat64DirectToIntTruncation(t *testing.T) {
	cases := []struct {
		name   string
		input  float64
		expect int
	}{
		{"1.5 → 1", 1.5, 1},
		{"1.9 → 1", 1.9, 1},
		{"-1.5 → -1", -1.5, -1},
		{"-1.9 → -1", -1.9, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[int](tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expect {
				t.Errorf("ToE[int](%v): expected %d, got %d", tc.input, tc.expect, result)
			}
		})
	}
}

// Negative float → uint always errors (no ABS flag), regardless of fraction.
func TestNegativeFloatToUintErrors(t *testing.T) {
	for _, v := range []float64{-0.1, -1.0, -1.9} {
		_, err := cast.ToE[uint](v)
		if err == nil {
			t.Errorf("ToE[uint](%v): expected error, got nil", v)
		}
	}
}

// TestToIntFromUintTypes covers the uintptr and uint16 arms in toInt's type
// switch, which are only reached when the source value is already one of those
// exact types.
func TestToIntFromUintTypes(t *testing.T) {
	if v, err := cast.ToE[int](uintptr(42)); err != nil || v != 42 {
		t.Errorf("ToE[int](uintptr(42)): expected 42/nil, got %v/%v", v, err)
	}
	if v, err := cast.ToE[int](uint16(7)); err != nil || v != 7 {
		t.Errorf("ToE[int](uint16(7)): expected 7/nil, got %v/%v", v, err)
	}
}

// TestStrToIntDefaultWrongType2 covers the DEFAULT type-assertion failure in
// strToInt via a decimal-string source (not an integer literal), exercising the
// strToInt code path specifically.
func TestStrToIntDefaultWrongType2(t *testing.T) {
	_, err := cast.ToE[int]("3.14", cast.Op{cast.DEFAULT, "not an int"})
	if err == nil {
		t.Fatal("expected error for wrong DEFAULT type, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}
