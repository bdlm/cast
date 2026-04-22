package cast_test

import (
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
