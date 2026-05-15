package cast_test

import (
	"errors"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestComplex64ToComplex64(t *testing.T) {
	testSimpleCases[complex64](t, []testCase{
		{complex64(3 + 4i), complex64(3 + 4i), nil, false},
		{complex64(-1.5 + 2.5i), complex64(-1.5 + 2.5i), nil, false},
		{complex64(0 + 0i), complex64(0 + 0i), nil, false},
		{complex64(1 + 0i), complex64(1 + 0i), nil, false},
		{complex64(0 + 1i), complex64(0 + 1i), nil, false},
	})
}

func TestComplex128ToComplex128(t *testing.T) {
	testSimpleCases[complex128](t, []testCase{
		{complex128(3 + 4i), complex128(3 + 4i), nil, false},
		{complex128(-1.5 + 2.5i), complex128(-1.5 + 2.5i), nil, false},
		{complex128(0 + 0i), complex128(0 + 0i), nil, false},
		{complex128(1 + 0i), complex128(1 + 0i), nil, false},
		{complex128(0 + 1i), complex128(0 + 1i), nil, false},
	})
}

func TestComplex64ToComplex128(t *testing.T) {
	testSimpleCases[complex128](t, []testCase{
		{complex64(3 + 4i), complex128(complex64(3 + 4i)), nil, false},
		{complex64(-1 + 2i), complex128(complex64(-1 + 2i)), nil, false},
		{complex64(0 + 0i), complex128(0 + 0i), nil, false},
		{complex64(0 + 1i), complex128(complex64(0 + 1i)), nil, false},
	})
}

func TestComplex128ToComplex64(t *testing.T) {
	testSimpleCases[complex64](t, []testCase{
		{complex128(3 + 4i), complex64(3 + 4i), nil, false},
		{complex128(-1 + 2i), complex64(-1 + 2i), nil, false},
		{complex128(0 + 0i), complex64(0 + 0i), nil, false},
		{complex128(0 + 1i), complex64(0 + 1i), nil, false},
	})
}

func TestNonComplexToComplex(t *testing.T) {
	// Non-complex sources produce complex(src, 0) — imaginary part is zero.
	testSimpleCases[complex128](t, []testCase{
		{42, complex128(42 + 0i), nil, false},
		{3.14, complex128(3.14 + 0i), nil, false},
		{"2.5", complex128(2.5 + 0i), nil, false},
		{true, complex128(1 + 0i), nil, false},
		{false, complex128(0 + 0i), nil, false},
	})
	testSimpleCases[complex64](t, []testCase{
		{42, complex64(42 + 0i), nil, false},
		{float32(1.5), complex64(1.5 + 0i), nil, false},
		{true, complex64(1 + 0i), nil, false},
	})
}

func TestStringerToComplex(t *testing.T) {
	testSimpleCases[complex128](t, []testCase{
		{testStringer{"1.5"}, complex128(1.5 + 0i), nil, false},
		{testStringer{"0"}, complex128(0 + 0i), nil, false},
		{testStringer{"bad"}, complex128(0), nil, true},
	})
	testSimpleCases[complex64](t, []testCase{
		{testStringer{"2.5"}, complex64(2.5 + 0i), nil, false},
		{testStringer{"bad"}, complex64(0), nil, true},
	})
}

func TestNilToComplex(t *testing.T) {
	testSimpleCases[complex64](t, []testCase{
		{nil, complex64(0), nil, false},
	})
	testSimpleCases[complex128](t, []testCase{
		{nil, complex128(0), nil, false},
	})
}

func TestInvalidStringToComplex(t *testing.T) {
	testSimpleCases[complex64](t, []testCase{
		{"not a number", complex64(0), nil, true},
	})
	testSimpleCases[complex128](t, []testCase{
		{"not a number", complex128(0), nil, true},
	})
}

func TestComplexDefaultOption(t *testing.T) {
	t.Run("DEFAULT used on error", func(t *testing.T) {
		result := cast.To[complex128]("bad", cast.Op{cast.DEFAULT, complex128(1 + 2i)})
		if result != complex128(1+2i) {
			t.Errorf("expected (1+2i), got %v", result)
		}
	})
	t.Run("DEFAULT wrong type errors immediately", func(t *testing.T) {
		_, err := cast.ToE[complex128]("bad", cast.Op{cast.DEFAULT, "not a complex"})
		if err == nil {
			t.Fatal("expected error for wrong DEFAULT type, got nil")
		}
	})
	t.Run("DEFAULT not used when conversion succeeds", func(t *testing.T) {
		result := cast.To[complex128]("2.5", cast.Op{cast.DEFAULT, complex128(1 + 2i)})
		if result != complex128(2.5+0i) {
			t.Errorf("expected (2.5+0i), got %v", result)
		}
	})
}

// TestCollectionToComplexErrors verifies that []T/[N]T, map[K]V, and struct
// sources always error for complex* targets (documented as ✗).
func TestCollectionToComplexErrors(t *testing.T) {
	t.Run("[]int → complex64 errors", func(t *testing.T) {
		_, err := cast.ToE[complex64]([]int{1, 2})
		if err == nil {
			t.Error("expected error for []int → complex64, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("[]int → complex128 errors", func(t *testing.T) {
		_, err := cast.ToE[complex128]([]int{1, 2})
		if err == nil {
			t.Error("expected error for []int → complex128, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("map → complex128 errors", func(t *testing.T) {
		_, err := cast.ToE[complex128](map[string]int{"a": 1})
		if err == nil {
			t.Error("expected error for map → complex128, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("struct → complex128 errors", func(t *testing.T) {
		_, err := cast.ToE[complex128](struct{ X int }{1})
		if err == nil {
			t.Error("expected error for struct → complex128, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}
