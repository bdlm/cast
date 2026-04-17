package cast_test

import (
	"testing"
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
