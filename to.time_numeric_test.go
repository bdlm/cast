package cast_test

import (
	"errors"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

// TestBigTypesToTime validates that *big.Int, big.Int, *big.Float, and big.Float
// are all treated as Unix seconds when converting to time.Time.
func TestBigTypesToTime(t *testing.T) {
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

	cases := []struct {
		name      string
		in        any
		expect    time.Time
		expectErr bool
	}{
		// *big.Int sources
		{name: "*big.Int zero → epoch", in: big.NewInt(0), expect: epoch},
		{name: "*big.Int 1s → epoch+1s", in: big.NewInt(1), expect: epoch.Add(time.Second)},
		{name: "*big.Int negative → pre-epoch", in: big.NewInt(-86400), expect: epoch.Add(-24 * time.Hour)},
		{name: "*big.Int nil → error", in: (*big.Int)(nil), expectErr: true},
		{name: "*big.Int too large for int64 → error", in: new(big.Int).Add(
			new(big.Int).SetInt64(math.MaxInt64), big.NewInt(1),
		), expectErr: true},

		// big.Int value (non-pointer)
		{name: "big.Int value zero → epoch", in: *big.NewInt(0), expect: epoch},
		{name: "big.Int value 1s", in: *big.NewInt(1), expect: epoch.Add(time.Second)},
		{name: "big.Int value too large → error", in: func() big.Int {
			v := new(big.Int).Add(new(big.Int).SetInt64(math.MaxInt64), big.NewInt(1))
			return *v
		}(), expectErr: true},

		// *big.Float sources (Unix seconds, fractional ns)
		{name: "*big.Float zero → epoch", in: new(big.Float).SetFloat64(0), expect: epoch},
		{name: "*big.Float 1s", in: new(big.Float).SetFloat64(1), expect: epoch.Add(time.Second)},
		{name: "*big.Float 1.5s → fractional", in: new(big.Float).SetFloat64(1.5), expect: epoch.Add(1500 * time.Millisecond)},
		{name: "*big.Float nil → error", in: (*big.Float)(nil), expectErr: true},
		{name: "*big.Float +Inf → error", in: new(big.Float).SetInf(false), expectErr: true},
		{name: "*big.Float -Inf → error", in: new(big.Float).SetInf(true), expectErr: true},

		// big.Float value (non-pointer)
		{name: "big.Float value zero → epoch", in: *new(big.Float).SetFloat64(0), expect: epoch},
		{name: "big.Float value 2s", in: *new(big.Float).SetFloat64(2.0), expect: epoch.Add(2 * time.Second)},
		{name: "big.Float value +Inf → error", in: func() big.Float {
			var f big.Float
			f.SetInf(false)
			return f
		}(), expectErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[time.Time](tc.in)
			if err != nil && !tc.expectErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectErr {
				t.Error("expected error, got nil")
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Errorf("expected cast.Error, got %T: %v", err, err)
			} else if err == nil && !result.Equal(tc.expect) {
				t.Errorf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

// TestTimeToIntConversions validates time.Time → int*/uint* using Unix seconds,
// including overflow detection for narrow integer types.
func TestTimeToIntConversions(t *testing.T) {
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	oneSecond := epoch.Add(time.Second)
	preEpoch := epoch.Add(-time.Second)
	// Unix=1000 overflows int8 (max 127) and int16 is fine (max 32767).
	bigTime := epoch.Add(1000 * time.Second)
	// Unix=40000 overflows int16 (max 32767).
	veryBigTime := epoch.Add(40000 * time.Second)

	testSimpleCases[int64](t, []testCase{
		{epoch, int64(0), nil, false},
		{oneSecond, int64(1), nil, false},
		{preEpoch, int64(-1), nil, false},
	})
	testSimpleCases[int32](t, []testCase{
		{epoch, int32(0), nil, false},
		{oneSecond, int32(1), nil, false},
		{preEpoch, int32(-1), nil, false},
	})
	testSimpleCases[int16](t, []testCase{
		{epoch, int16(0), nil, false},
		{oneSecond, int16(1), nil, false},
		{preEpoch, int16(-1), nil, false},
		{veryBigTime, int16(0), nil, true}, // Unix=40000 overflows int16
	})
	testSimpleCases[int8](t, []testCase{
		{epoch, int8(0), nil, false},
		{oneSecond, int8(1), nil, false},
		{bigTime, int8(0), nil, true}, // Unix=1000 overflows int8
	})
	testSimpleCases[int](t, []testCase{
		{epoch, int(0), nil, false},
		{oneSecond, int(1), nil, false},
		{preEpoch, int(-1), nil, false},
	})

	// Unsigned targets: pre-epoch Unix < 0 must error.
	testSimpleCases[uint64](t, []testCase{
		{epoch, uint64(0), nil, false},
		{oneSecond, uint64(1), nil, false},
		{preEpoch, uint64(0), nil, true},
	})
	testSimpleCases[uint32](t, []testCase{
		{epoch, uint32(0), nil, false},
		{oneSecond, uint32(1), nil, false},
		{preEpoch, uint32(0), nil, true},
	})
	testSimpleCases[uint16](t, []testCase{
		{epoch, uint16(0), nil, false},
		{oneSecond, uint16(1), nil, false},
		{preEpoch, uint16(0), nil, true},
	})
	testSimpleCases[uint8](t, []testCase{
		{epoch, uint8(0), nil, false},
		{oneSecond, uint8(1), nil, false},
		{preEpoch, uint8(0), nil, true},
	})
	testSimpleCases[uint](t, []testCase{
		{epoch, uint(0), nil, false},
		{oneSecond, uint(1), nil, false},
		{preEpoch, uint(0), nil, true},
	})
}

// TestTimeToFloatConversions validates time.Time → float*, including fractional
// nanoseconds encoded in the sub-second part.
func TestTimeToFloatConversions(t *testing.T) {
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	oneSecond := epoch.Add(time.Second)
	halfSecond := epoch.Add(500 * time.Millisecond)

	testSimpleCases[float64](t, []testCase{
		{epoch, float64(0), nil, false},
		{oneSecond, float64(1.0), nil, false},
		{halfSecond, float64(0.5), nil, false},
	})
	testSimpleCases[float32](t, []testCase{
		{epoch, float32(0), nil, false},
		{oneSecond, float32(1.0), nil, false},
	})
}

// TestTimeToBigConversions validates time.Time → *big.Int and *big.Float.
func TestTimeToBigConversions(t *testing.T) {
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	oneSecond := epoch.Add(time.Second)
	halfSecond := epoch.Add(500 * time.Millisecond)
	preEpoch := epoch.Add(-time.Second)

	t.Run("*big.Int: epoch → 0", func(t *testing.T) {
		result, err := cast.ToE[*big.Int](epoch)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil || result.Int64() != 0 {
			t.Errorf("expected 0, got %v", result)
		}
	})
	t.Run("*big.Int: 1s → 1", func(t *testing.T) {
		result, err := cast.ToE[*big.Int](oneSecond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil || result.Int64() != 1 {
			t.Errorf("expected 1, got %v", result)
		}
	})
	t.Run("*big.Int: pre-epoch → -1", func(t *testing.T) {
		result, err := cast.ToE[*big.Int](preEpoch)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil || result.Int64() != -1 {
			t.Errorf("expected -1, got %v", result)
		}
	})

	t.Run("*big.Float: epoch → 0.0", func(t *testing.T) {
		result, err := cast.ToE[*big.Float](epoch)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		f64, _ := result.Float64()
		if f64 != 0.0 {
			t.Errorf("expected 0.0, got %v", f64)
		}
	})
	t.Run("*big.Float: 1s → 1.0", func(t *testing.T) {
		result, err := cast.ToE[*big.Float](oneSecond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		f64, _ := result.Float64()
		if f64 != 1.0 {
			t.Errorf("expected 1.0, got %v", f64)
		}
	})
	t.Run("*big.Float: 500ms → 0.5", func(t *testing.T) {
		result, err := cast.ToE[*big.Float](halfSecond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		f64, _ := result.Float64()
		if math.Abs(f64-0.5) > 1e-9 {
			t.Errorf("expected ~0.5, got %v", f64)
		}
	})
	t.Run("*big.Float: pre-epoch → -1.0", func(t *testing.T) {
		result, err := cast.ToE[*big.Float](preEpoch)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		f64, _ := result.Float64()
		if f64 != -1.0 {
			t.Errorf("expected -1.0, got %v", f64)
		}
	})
}
