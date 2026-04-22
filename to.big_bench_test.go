package cast_test

import (
	"math/big"
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_bigInt_fromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[*big.Int]("999999999999999999999")
	}
}

func BenchmarkTo_bigInt_fromInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[*big.Int](int64(9999999999))
	}
}

func BenchmarkTo_bigInt_fromBigInt(b *testing.B) {
	// passthrough: *big.Int → *big.Int (copy)
	src := big.NewInt(42)
	for i := 0; i < b.N; i++ {
		_ = cast.To[*big.Int](src)
	}
}

func BenchmarkTo_bigFloat_fromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[*big.Float]("3.14159265358979323846")
	}
}

func BenchmarkTo_bigFloat_fromFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[*big.Float](float64(3.14))
	}
}

func BenchmarkTo_bigFloat_fromBigFloat(b *testing.B) {
	// passthrough: *big.Float → *big.Float (copy)
	src := new(big.Float).SetFloat64(3.14)
	for i := 0; i < b.N; i++ {
		_ = cast.To[*big.Float](src)
	}
}

func BenchmarkToE_bigInt_fromInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[*big.Int]("not-a-number")
	}
}
