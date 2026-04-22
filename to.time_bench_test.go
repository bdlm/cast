package cast_test

import (
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_time_fromTime(b *testing.B) {
	t := time.Now()
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Time](t)
	}
}

func BenchmarkTo_time_fromRFC3339(b *testing.B) {
	// RFC3339 is tried first in timeFormats; best-case string parse.
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Time]("2024-06-15T12:00:00Z")
	}
}

func BenchmarkTo_time_fromDateOnly(b *testing.B) {
	// time.DateOnly is tried after RFC3339 and RFC3339Nano in timeFormats.
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Time]("2024-06-15")
	}
}

func BenchmarkTo_time_fromRFC1123(b *testing.B) {
	// RFC1123 sits further down the timeFormats list; measures fallthrough cost.
	ref := time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)
	s := ref.Format(time.RFC1123)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Time](s)
	}
}

func BenchmarkTo_time_fromInt64Unix(b *testing.B) {
	// int64 → time.Unix(0, ns) path
	ns := time.Now().UnixNano()
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Time](ns)
	}
}

func BenchmarkTo_time_fromFloat64Unix(b *testing.B) {
	// float64 → time.Unix(sec, frac) path
	for i := 0; i < b.N; i++ {
		_ = cast.To[time.Time](float64(1.5))
	}
}

func BenchmarkToE_time_fromInvalid(b *testing.B) {
	// All timeFormats tried and failed — worst-case string parse.
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[time.Time]("not-a-time")
	}
}
