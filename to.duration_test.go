package cast_test

import (
	"errors"
	"math"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

func TestToEDuration(t *testing.T) {
	cases := []struct {
		name      string
		in        any
		expect    time.Duration
		expectErr bool
	}{
		// String parsing via time.ParseDuration
		{name: "nanoseconds", in: "100ns", expect: 100 * time.Nanosecond},
		{name: "microseconds us", in: "5us", expect: 5 * time.Microsecond},
		{name: "microseconds µs", in: "5µs", expect: 5 * time.Microsecond},
		{name: "milliseconds", in: "250ms", expect: 250 * time.Millisecond},
		{name: "seconds", in: "30s", expect: 30 * time.Second},
		{name: "minutes", in: "5m", expect: 5 * time.Minute},
		{name: "hours", in: "2h", expect: 2 * time.Hour},
		{name: "combined", in: "1h30m", expect: 90 * time.Minute},
		{name: "combined 2", in: "2m45s", expect: 2*time.Minute + 45*time.Second},
		{name: "zero", in: "0s", expect: 0},

		// time.Duration passthrough
		{name: "time.Duration", in: 5 * time.Second, expect: 5 * time.Second},

		// Integer → nanoseconds
		{name: "int zero", in: int(0), expect: 0},
		{name: "int 1000", in: int(1000), expect: 1000 * time.Nanosecond},
		{name: "int64 1s", in: int64(time.Second), expect: time.Second},
		{name: "int32", in: int32(500), expect: 500 * time.Nanosecond},
		{name: "uint64", in: uint64(time.Minute), expect: time.Minute},

		// Float → nanoseconds (truncated)
		{name: "float64", in: float64(time.Second), expect: time.Second},
		{name: "float32", in: float32(1000), expect: 1000 * time.Nanosecond},

		// Error cases
		{name: "nil", in: nil, expectErr: true},
		{name: "invalid string", in: "not-a-duration", expectErr: true},
		{name: "bare number string", in: "5", expectErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[time.Duration](tc.in)

			if err != nil && !tc.expectErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectErr {
				t.Error("expected error, got nil")
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Errorf("expected cast.Error, got %T: %v", err, err)
			} else if err == nil && result != tc.expect {
				t.Errorf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

func TestToDuration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := cast.To[time.Duration]("1h30m")
		if result != 90*time.Minute {
			t.Errorf("expected 90m, got %v", result)
		}
	})
	t.Run("error returns zero", func(t *testing.T) {
		result := cast.To[time.Duration]("bad")
		if result != 0 {
			t.Errorf("expected 0, got %v", result)
		}
	})
}

func TestToEDurationDefault(t *testing.T) {
	def := 5 * time.Second
	result, err := cast.ToE[time.Duration]("bad", cast.Op{Flag: cast.DEFAULT, Val: def})
	if err == nil {
		t.Error("expected error, got nil")
	}
	if result != def {
		t.Errorf("expected default %v, got %v", def, result)
	}
}

func TestToEDurationInvalidDefault(t *testing.T) {
	_, err := cast.ToE[time.Duration]("1s", cast.Op{Flag: cast.DEFAULT, Val: "wrong-type"})
	if err == nil {
		t.Error("expected error for non-time.Duration DEFAULT, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %T: %v", err, err)
	}
}

func TestToEDurationStructField(t *testing.T) {
	// Verify time.Duration fields in structs work via castToType.
	type Config struct {
		Timeout time.Duration
		Retries int
	}
	result, err := cast.ToStructE[Config](map[string]any{
		"Timeout": "30s",
		"Retries": "3",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Timeout != 30*time.Second {
		t.Errorf("expected Timeout=30s, got %v", result.Timeout)
	}
	if result.Retries != 3 {
		t.Errorf("expected Retries=3, got %d", result.Retries)
	}
}

func TestToEDurationReflectType(t *testing.T) {
	result, err := cast.ToE[time.Duration]("1h")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reflect.TypeOf(result) != reflect.TypeOf(time.Duration(0)) {
		t.Errorf("expected time.Duration, got %T", result)
	}
}

func TestToEDurationStringerDefault(t *testing.T) {
	// testStringer implements fmt.Stringer; toDuration's default: branch
	// calls toString → val.String() → "1h30m" then retries time.ParseDuration.
	result, err := cast.ToE[time.Duration](testStringer{"1h30m"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 90*time.Minute {
		t.Errorf("expected 90m, got %v", result)
	}
}

func TestToEDurationStringerDefaultFails(t *testing.T) {
	// A Stringer whose String() is not a valid duration falls through to the
	// final error return in toDuration.
	_, err := cast.ToE[time.Duration](testStringer{"not-a-duration"})
	if err == nil {
		t.Error("expected error for unparseable Stringer duration, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %T: %v", err, err)
	}
}

// TestBigIntToDuration validates *big.Int, big.Int, *big.Float, and big.Float
// sources converted to time.Duration (nanoseconds).
func TestBigIntToDuration(t *testing.T) {
	cases := []struct {
		name      string
		in        any
		expect    time.Duration
		expectErr bool
	}{
		{name: "*big.Int zero", in: big.NewInt(0), expect: 0},
		{name: "*big.Int 1µs", in: big.NewInt(1000), expect: time.Microsecond},
		{name: "*big.Int 1s in ns", in: big.NewInt(int64(time.Second)), expect: time.Second},
		{name: "*big.Int negative", in: big.NewInt(-100), expect: -100 * time.Nanosecond},
		{name: "*big.Int nil → error", in: (*big.Int)(nil), expectErr: true},
		{name: "*big.Int too large → error", in: new(big.Int).Add(
			new(big.Int).SetInt64(math.MaxInt64), big.NewInt(1),
		), expectErr: true},

		// big.Int value (non-pointer)
		{name: "big.Int value zero", in: *big.NewInt(0), expect: 0},
		{name: "big.Int value 1µs", in: *big.NewInt(1000), expect: time.Microsecond},
		{name: "big.Int value too large → error", in: func() big.Int {
			v := new(big.Int).Add(new(big.Int).SetInt64(math.MaxInt64), big.NewInt(1))
			return *v
		}(), expectErr: true},

		// *big.Float sources (nanoseconds, floor semantics)
		{name: "*big.Float zero", in: new(big.Float).SetFloat64(0), expect: 0},
		{name: "*big.Float 1s in ns", in: new(big.Float).SetFloat64(float64(time.Second)), expect: time.Second},
		{name: "*big.Float nil → error", in: (*big.Float)(nil), expectErr: true},
		{name: "*big.Float +Inf → error", in: new(big.Float).SetInf(false), expectErr: true},
		{name: "*big.Float -Inf → error", in: new(big.Float).SetInf(true), expectErr: true},

		// big.Float value (non-pointer)
		{name: "big.Float value zero", in: *new(big.Float).SetFloat64(0), expect: 0},
		{name: "big.Float value +Inf → error", in: func() big.Float {
			var f big.Float
			f.SetInf(false)
			return f
		}(), expectErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[time.Duration](tc.in)
			if err != nil && !tc.expectErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectErr {
				t.Error("expected error, got nil")
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Errorf("expected cast.Error, got %T: %v", err, err)
			} else if err == nil && result != tc.expect {
				t.Errorf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

// TestFloatDurationFloor verifies that math.Floor semantics are used for
// float→duration, so negative fractional nanoseconds round toward -∞ rather
// than toward zero (the behavior of a plain int64 cast).
func TestFloatDurationFloor(t *testing.T) {
	cases := []struct {
		name   string
		in     any
		expect time.Duration
	}{
		{name: "float64 +1.7ns floors to 1", in: float64(1.7), expect: 1},
		{name: "float64 -1.7ns floors to -2", in: float64(-1.7), expect: -2},
		{name: "float64 -0.3ns floors to -1", in: float64(-0.3), expect: -1},
		{name: "float32 -1.7ns floors to -2", in: float32(-1.7), expect: -2},
		{name: "*big.Float -1.7ns floors to -2", in: new(big.Float).SetFloat64(-1.7), expect: -2},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[time.Duration](tc.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expect {
				t.Errorf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}
