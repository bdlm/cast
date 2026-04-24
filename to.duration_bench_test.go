package cast_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_duration_fromDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Duration](5 * time.Second)
	}
}

func BenchmarkTo_duration_fromString(b *testing.B) {
	// time.ParseDuration path
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Duration]("1h30m")
	}
}

func BenchmarkTo_duration_fromInt(b *testing.B) {
	// int → nanoseconds direct conversion
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Duration](int64(time.Second))
	}
}

func BenchmarkTo_duration_fromFloat(b *testing.B) {
	// float64 → nanoseconds via truncation
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Duration](float64(time.Second))
	}
}

func BenchmarkToE_duration_fromInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[time.Duration]("not-a-duration")
	}
}

func BenchmarkTo_duration_fromBigInt(b *testing.B) {
	src := big.NewInt(int64(time.Second))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Duration](src)
	}
}

func BenchmarkTo_duration_fromBigFloat(b *testing.B) {
	src := new(big.Float).SetFloat64(float64(time.Second))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Duration](src)
	}
}

func BenchmarkToE_duration_fromComplexDefaultBranch(b *testing.B) {
	// complex128 hits the default: branch (toString + ParseDuration), which fails.
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[time.Duration](complex128(1 + 0i))
	}
}
