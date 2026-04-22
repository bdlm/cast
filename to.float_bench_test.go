package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_float64_fromFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[float64](float64(3.14))
	}
}

func BenchmarkTo_float64_fromInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[float64](int(42))
	}
}

func BenchmarkTo_float64_fromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[float64]("3.14159")
	}
}

func BenchmarkTo_float64_fromFormattedString(b *testing.B) {
	// Formatted strings require comma stripping before parse.
	for i := 0; i < b.N; i++ {
		_ = cast.To[float64]("1,234.56")
	}
}

func BenchmarkTo_float64_fromFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[float64](float32(1.5))
	}
}

func BenchmarkTo_float32_fromFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[float32](float64(3.14))
	}
}

func BenchmarkTo_float64_withDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[float64]("3.14", cast.Op{Flag: cast.DEFAULT, Val: float64(0)})
	}
}

func BenchmarkToE_float64_fromInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[float64]("not-a-float")
	}
}
