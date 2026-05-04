package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_complex128_fromComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[complex128](complex128(1 + 2i))
	}
}

func BenchmarkTo_complex128_fromFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[complex128](float64(3.14))
	}
}

func BenchmarkTo_complex128_fromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[complex128]("(1+2i)")
	}
}

func BenchmarkTo_complex64_fromComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[complex64](complex64(1 + 2i))
	}
}

func BenchmarkTo_complex128_fromComplex64(b *testing.B) {
	// Cross-width conversion: complex64 → complex128
	for i := 0; i < b.N; i++ {
		_ = cast.To[complex128](complex64(1 + 2i))
	}
}

func BenchmarkToE_complex128_fromInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[complex128]("not-complex")
	}
}
