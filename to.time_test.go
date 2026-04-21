package cast_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

func TestToETime(t *testing.T) {
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

	cases := []struct {
		name      string
		in        any
		expect    time.Time
		expectErr bool
	}{
		// String formats
		{name: "RFC3339", in: "2024-06-15T12:00:00Z", expect: time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)},
		{name: "RFC3339Nano", in: "2024-06-15T12:00:00.123456789Z", expect: time.Date(2024, 6, 15, 12, 0, 0, 123456789, time.UTC)},
		{name: "DateOnly", in: "1994-04-20", expect: time.Date(1994, 4, 20, 0, 0, 0, 0, time.UTC)},
		{name: "DateTime", in: "2024-06-15 12:00:00", expect: time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)},

		// time.Time passthrough
		{name: "time.Time", in: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), expect: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},

		// *time.Time
		{name: "*time.Time", in: func() *time.Time { t := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC); return &t }(), expect: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "*time.Time nil", in: (*time.Time)(nil), expectErr: true},

		// []byte
		{name: "[]byte RFC3339", in: []byte("2024-06-15T12:00:00Z"), expect: time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)},
		{name: "[]byte DateOnly", in: []byte("1994-04-20"), expect: time.Date(1994, 4, 20, 0, 0, 0, 0, time.UTC)},

		// Integer → Unix nanoseconds
		{name: "int zero", in: int(0), expect: epoch},
		{name: "int64 zero", in: int64(0), expect: epoch},
		{name: "int64 1s", in: int64(time.Second), expect: epoch.Add(time.Second)},
		{name: "uint64 1s", in: uint64(time.Second), expect: epoch.Add(time.Second)},

		// Float → Unix seconds
		{name: "float64 1.5s", in: float64(1.5), expect: epoch.Add(1500 * time.Millisecond)},
		{name: "float32 zero", in: float32(0), expect: epoch},

		// Error cases
		{name: "nil", in: nil, expectErr: true},
		{name: "invalid string", in: "not-a-time", expectErr: true},
		{name: "invalid []byte", in: []byte("garbage"), expectErr: true},
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

func TestToTime(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := cast.To[time.Time]("2024-01-01")
		if !result.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("unexpected: %v", result)
		}
	})
	t.Run("error returns zero", func(t *testing.T) {
		result := cast.To[time.Time]("not-a-time")
		if !result.IsZero() {
			t.Errorf("expected zero time, got %v", result)
		}
	})
}

func TestToETimeDefault(t *testing.T) {
	def := time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)
	result, err := cast.ToE[time.Time]("bad", cast.Op{Flag: cast.DEFAULT, Val: def})
	if err == nil {
		t.Error("expected error, got nil")
	}
	if !result.Equal(def) {
		t.Errorf("expected default %v, got %v", def, result)
	}
}

func TestToETimeInvalidDefault(t *testing.T) {
	_, err := cast.ToE[time.Time]("2024-01-01", cast.Op{Flag: cast.DEFAULT, Val: "wrong-type"})
	if err == nil {
		t.Error("expected error for non-time.Time DEFAULT, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %T: %v", err, err)
	}
}

func TestToStructTimeField(t *testing.T) {
	// Verify time.Time fields in structs still work after the refactor.
	type Event struct {
		Name string
		At   time.Time
	}
	result, err := cast.ToStructE[Event](map[string]any{
		"Name": "launch",
		"At":   "2024-06-15T12:00:00Z",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "launch" {
		t.Errorf("expected Name=launch, got %q", result.Name)
	}
	expected := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	if !result.At.Equal(expected) {
		t.Errorf("expected At=%v, got %v", expected, result.At)
	}
}

func TestToETimeReflectTypes(t *testing.T) {
	// Verify that the reflect.Value returned by toTime round-trips correctly
	// through the assertion logic in ToE.
	result, err := cast.ToE[time.Time]("2024-06-15")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reflect.TypeOf(result) != reflect.TypeOf(time.Time{}) {
		t.Errorf("expected time.Time, got %T", result)
	}
}
