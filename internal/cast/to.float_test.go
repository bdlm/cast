package cast_test

import (
	"errors"
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

func TestComplexToFloat(t *testing.T) {
	// Imaginary part is discarded; only the real component is returned.
	// complex64 stores float32 components, so complex64→float64 loses precision
	// the same way float32→float64 does: float64(float32(1.1)) ≈ 1.1000000238...
	testSimpleCases[float32](t, []testCase{
		{complex(1, 0), float32(1), nil, false},
		{complex(1, 1), float32(1), nil, false},
		{complex(0, 0), float32(0), nil, false},
		{complex(-1, 0), float32(-1), nil, false},
		{complex(-1, -1), float32(-1), nil, false},
		{complex64(float32(1.1)), float32(1.1), nil, false},
		{complex64(float32(-1.1)), float32(-1.1), nil, false},
		{complex64(float32(0.0)), float32(0.0), nil, false},
		{complex128(float32(1.1)), float32(1.1), nil, false},
		{complex128(float32(-1.1)), float32(-1.1), nil, false},
		{complex128(float32(0.0)), float32(0.0), nil, false},
		{complex64(float64(1.1)), float32(1.1), nil, false},
		{complex64(float64(-1.1)), float32(-1.1), nil, false},
		{complex64(float64(0.0)), float32(0.0), nil, false},
		{complex128(float64(1.1)), float32(1.1), nil, false},
		{complex128(float64(-1.1)), float32(-1.1), nil, false},
		{complex128(float64(0.0)), float32(0.0), nil, false},
	})
	testSimpleCases[float64](t, []testCase{
		{complex(1, 0), float64(1), nil, false},
		{complex(1, 1), float64(1), nil, false},
		{complex(0, 0), float64(0), nil, false},
		{complex(-1, 0), float64(-1), nil, false},
		{complex(-1, -1), float64(-1), nil, false},
		{complex64(float32(1.1)), float64(float32(1.1)), nil, false},
		{complex64(float32(-1.1)), float64(float32(-1.1)), nil, false},
		{complex64(float32(0.0)), float64(0.0), nil, false},
		{complex128(float32(1.1)), float64(float32(1.1)), nil, false},
		{complex128(float32(-1.1)), float64(float32(-1.1)), nil, false},
		{complex128(float32(0.0)), float64(0.0), nil, false},
		{complex64(float64(1.1)), float64(float32(1.1)), nil, false},
		{complex64(float64(-1.1)), float64(float32(-1.1)), nil, false},
		{complex64(float64(0.0)), float64(0.0), nil, false},
		{complex128(float64(1.1)), float64(1.1), nil, false},
		{complex128(float64(-1.1)), float64(-1.1), nil, false},
		{complex128(float64(0.0)), float64(0.0), nil, false},
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

func TestFloatDefaultBranch(t *testing.T) {
	// A named float type without fmt.Stringer falls through to the
	// fmt.Sprintf("%v", from) default branch in toFloat. The Sprintf output
	// is a valid decimal string, so the parse succeeds.
	type myFloat float64
	result, err := cast.ToE[float64](myFloat(3.14))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 3.14 {
		t.Errorf("expected 3.14, got %v", result)
	}
}

// TestTimeToFloat validates time.Time → float* conversions using Unix seconds
// with sub-second precision encoded as fractional seconds.
func TestTimeToFloat(t *testing.T) {
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	testSimpleCases[float64](t, []testCase{
		{epoch, float64(0), nil, false},
		{epoch.Add(time.Second), float64(1.0), nil, false},
		{epoch.Add(500 * time.Millisecond), float64(0.5), nil, false},
	})
	testSimpleCases[float32](t, []testCase{
		{epoch, float32(0), nil, false},
		{epoch.Add(time.Second), float32(1.0), nil, false},
	})
}

// TestCharSeqToFloat validates that []byte and []rune reach strToFloat via the
// default: branch → toString → strToFloat.
func TestCharSeqToFloat(t *testing.T) {
	testSimpleCases[float64](t, []testCase{
		{[]byte("3.14"), float64(3.14), nil, false},
		{[]byte("-2.5"), float64(-2.5), nil, false},
		{[]byte("bad"), float64(0), nil, true},
		{[]rune("1.5"), float64(1.5), nil, false},
	})
	testSimpleCases[float32](t, []testCase{
		{[]byte("1.0"), float32(1.0), nil, false},
		{[]rune("2.5"), float32(2.5), nil, false},
	})
}

func TestFloatDefaultBranchError(t *testing.T) {
	// A type whose Sprintf representation is not a parseable float must error.
	type myStruct struct{ V int }
	_, err := cast.ToE[float64](myStruct{V: 1})
	if err == nil {
		t.Error("expected error for struct→float64, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
	}
}

// TestMapToFloatErrors verifies that map sources always error for float* targets
// (documented as ✗ in the conversion table).
func TestMapToFloatErrors(t *testing.T) {
	src := map[string]int{"a": 1}
	t.Run("map → float32", func(t *testing.T) {
		_, err := cast.ToE[float32](src)
		if err == nil {
			t.Error("expected error for map → float32, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("map → float64", func(t *testing.T) {
		_, err := cast.ToE[float64](src)
		if err == nil {
			t.Error("expected error for map → float64, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// TestSliceToFloatErrors verifies that slice sources always error for float*
// targets (documented as ✗).
func TestSliceToFloatErrors(t *testing.T) {
	src := []int{1, 2}
	t.Run("[]int → float32", func(t *testing.T) {
		_, err := cast.ToE[float32](src)
		if err == nil {
			t.Error("expected error for []int → float32, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("[]int → float64", func(t *testing.T) {
		_, err := cast.ToE[float64](src)
		if err == nil {
			t.Error("expected error for []int → float64, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}
