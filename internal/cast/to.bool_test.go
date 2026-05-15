package cast_test

import (
	"errors"
	"testing"

	"github.com/bdlm/cast/v2"
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
		{"0.1", false, nil, false},
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

func TestUintToBool(t *testing.T) {
	testSimpleCases[bool](t, []testCase{
		{uint(1), true, nil, false},
		{uint(0), false, nil, false},
		{uint16(5), true, nil, false},
		{uint16(0), false, nil, false},
		{uint32(100), true, nil, false},
		{uint32(0), false, nil, false},
		{uint64(1), true, nil, false},
		{uint64(0), false, nil, false},
		{uintptr(1), true, nil, false},
		{uintptr(0), false, nil, false},
	})
}

// strconv.ParseBool accepts "1","t","T","TRUE","true","True",
// "0","f","F","FALSE","false","False" — these are distinct from the
// numeric-string path that handles "-1", "1.5", etc.
func TestParseBoolStrings(t *testing.T) {
	testSimpleCases[bool](t, []testCase{
		{"true", true, nil, false},
		{"True", true, nil, false},
		{"TRUE", true, nil, false},
		{"t", true, nil, false},
		{"T", true, nil, false},
		{"false", false, nil, false},
		{"False", false, nil, false},
		{"FALSE", false, nil, false},
		{"f", false, nil, false},
		{"F", false, nil, false},
	})
}

func TestStringerToBool(t *testing.T) {
	t.Run("Stringer returning \"1\" → true", func(t *testing.T) {
		result, err := cast.ToE[bool](testStringer{"1"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != true {
			t.Errorf("expected true, got %v", result)
		}
	})
	t.Run("Stringer returning \"false\" → false", func(t *testing.T) {
		result, err := cast.ToE[bool](testStringer{"false"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != false {
			t.Errorf("expected false, got %v", result)
		}
	})
	t.Run("Stringer returning invalid string → error", func(t *testing.T) {
		_, err := cast.ToE[bool](testStringer{"maybe"})
		if err == nil {
			t.Error("expected error for unparseable Stringer, got nil")
		}
	})
}

// TestCharSeqToBool validates that []byte and []rune reach toBool via the
// default: branch → toString → toBool.
func TestCharSeqToBool(t *testing.T) {
	testSimpleCases[bool](t, []testCase{
		{[]byte("1"), true, nil, false},
		{[]byte("0"), false, nil, false},
		{[]byte("true"), true, nil, false},
		{[]byte("false"), false, nil, false},
		{[]rune("1"), true, nil, false},
		{[]rune("0"), false, nil, false},
		{[]byte("bad"), false, nil, true},
	})
}

// TestInt32RuneZeroConsistency documents that int32(48) (== rune '0') maps to
// false due to the rune-semantics special case, while int64(48) maps to true.
// byte(48) shares the same special case as rune because byte == uint8 and
// rune == int32 both have explicit '0'-aware branches.
func TestInt32RuneZeroConsistency(t *testing.T) {
	// int32 == rune: '0' (48) is treated as the digit zero → false.
	testSimpleCases[bool](t, []testCase{
		{int32(48), false, nil, false}, // rune '0' → false
		{int32(49), true, nil, false},  // rune '1' → true
		{int32(0), false, nil, false},  // zero → false
		{int32(1), true, nil, false},   // non-zero → true
	})

	// int64 does NOT have the rune special case: 48 is treated as non-zero → true.
	testSimpleCases[bool](t, []testCase{
		{int64(48), true, nil, false},  // 48 → true (no char special case)
		{int64(0), false, nil, false},  // zero → false
	})

	// byte == uint8: '0' (48) → false, same as rune.
	testSimpleCases[bool](t, []testCase{
		{byte(48), false, nil, false}, // byte '0' → false
		{byte(49), true, nil, false},  // byte '1' → true
	})
}

// TestCollectionToBoolErrors verifies that []T/[N]T, map[K]V, and struct sources
// always error for bool targets (documented as ✗ in the conversion table).
func TestCollectionToBoolErrors(t *testing.T) {
	t.Run("[]int → bool errors", func(t *testing.T) {
		_, err := cast.ToE[bool]([]int{1, 2})
		if err == nil {
			t.Error("expected error for []int → bool, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("map[string]int → bool errors", func(t *testing.T) {
		_, err := cast.ToE[bool](map[string]int{"a": 1})
		if err == nil {
			t.Error("expected error for map[string]int → bool, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("struct → bool errors", func(t *testing.T) {
		type point struct{ X, Y int }
		_, err := cast.ToE[bool](point{X: 1, Y: 2})
		if err == nil {
			t.Error("expected error for struct → bool, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}
