package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_funcInt_fromInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn := cast.To[cast.Func[int]](42)
		_ = fn
	}
}

func BenchmarkTo_funcInt_fromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn := cast.To[cast.Func[int]]("42")
		_ = fn
	}
}

func BenchmarkTo_funcFloat64_fromFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn := cast.To[cast.Func[float64]](3.14)
		_ = fn
	}
}

func BenchmarkTo_funcString_fromInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn := cast.To[cast.Func[string]](42)
		_ = fn
	}
}

// BenchmarkTo_funcInt_call measures the cost of creating the closure AND
// calling it once, to capture the full allocation + dispatch overhead.
func BenchmarkTo_funcInt_call(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn := cast.To[cast.Func[int]](42)
		_ = fn()
	}
}
