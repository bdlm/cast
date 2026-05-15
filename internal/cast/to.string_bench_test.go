package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_string_fromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[string]("hello world")
	}
}

func BenchmarkTo_string_fromInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[string](12345)
	}
}

func BenchmarkTo_string_fromFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[string](3.14)
	}
}

func BenchmarkTo_string_fromBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[string](true)
	}
}

func BenchmarkTo_string_JSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[string](`hello "world"`, cast.Op{Flag: cast.JSON, Val: true})
	}
}

// Measures option-parsing overhead: same cast with and without options.
func BenchmarkTo_string_noOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[string](42)
	}
}

func BenchmarkTo_string_withDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[string](42, cast.Op{Flag: cast.DEFAULT, Val: "fallback"})
	}
}
