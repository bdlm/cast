package cast_test

import (
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

var benchFloat64Time = time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)

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

func BenchmarkTo_float64_fromComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[float64](complex128(3.14 + 2i))
	}
}

func BenchmarkTo_float64_fromTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[float64](benchFloat64Time)
	}
}

func BenchmarkTo_float64_decodeJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[float64](`"3.14"`, cast.Op{Flag: cast.DECODE, Val: "json"})
	}
}
