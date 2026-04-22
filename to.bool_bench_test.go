package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_bool_fromBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[bool](true)
	}
}

func BenchmarkTo_bool_fromInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[bool](1)
	}
}

func BenchmarkTo_bool_fromString_parseBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[bool]("true")
	}
}

func BenchmarkTo_bool_fromString_numericFallback(b *testing.B) {
	// "1" hits the numeric fallback branch in toBool (strconv.ParseFloat → != 0)
	for i := 0; i < b.N; i++ {
		_ = cast.To[bool]("1")
	}
}

func BenchmarkTo_bool_fromByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[bool](byte(1))
	}
}

func BenchmarkTo_bool_withDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[bool]("true", cast.Op{Flag: cast.DEFAULT, Val: false})
	}
}

func BenchmarkToE_bool_fromInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[bool]("not-a-bool")
	}
}
