package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_int_fromInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[int](12345)
	}
}

func BenchmarkTo_int_fromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[int]("12345")
	}
}

func BenchmarkTo_int_fromFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[int](float64(3.14))
	}
}

func BenchmarkTo_int_fromStringDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[int]("3.14")
	}
}

func BenchmarkTo_int_withDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[int]("12345", cast.Op{Flag: cast.DEFAULT, Val: int(0)})
	}
}

func BenchmarkTo_uint_fromNegative_ABS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[uint](-42, cast.Op{Flag: cast.ABS, Val: true})
	}
}

func BenchmarkToE_int_fromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[int]("12345")
	}
}

func BenchmarkToE_int_fromInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[int]("not-a-number")
	}
}
