package cast_test

import (
	"errors"
	"testing"

	"github.com/bdlm/cast/v2"
)

// panickyStringer implements fmt.Stringer with a String() method that panics
// with a plain string value.
type panickyStringer struct{}

func (p panickyStringer) String() string {
	panic("intentional test panic from String()")
}

// panickyErrorStringer panics with an error value from String(), covering the
// "case error:" branch in ToE's defer/recover.
type panickyErrorStringer struct{}

func (p panickyErrorStringer) String() string {
	panic(errors.New("intentional error panic from String()"))
}

func TestToERecoversPanicWithErrorValue(t *testing.T) {
	// panickyErrorStringer panics with an error value; ToE's recover hits
	// the "case error:" branch which wraps via errors.Wrap instead of Errorf.
	result, err := cast.ToE[string](panickyErrorStringer{})
	if err == nil {
		t.Error("expected error from recovered error-panic, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
	if result != "" {
		t.Errorf("expected zero value on panic, got %q", result)
	}
}

func TestToERecoversPanic(t *testing.T) {
	t.Run("panic in Stringer.String() → error, not panic (string target)", func(t *testing.T) {
		result, err := cast.ToE[string](panickyStringer{})
		// Must not panic; must return an error.
		if err == nil {
			t.Error("expected error from recovered panic, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
		if result != "" {
			t.Errorf("expected zero value on panic, got %q", result)
		}
	})

	t.Run("panic in Stringer.String() → error, not panic (bool target)", func(t *testing.T) {
		result, err := cast.ToE[bool](panickyStringer{})
		if err == nil {
			t.Error("expected error from recovered panic, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
		if result != false {
			t.Errorf("expected zero value on panic, got %v", result)
		}
	})

	t.Run("To silently returns zero value on panic", func(t *testing.T) {
		// To discards the error; the zero value should be returned without panicking.
		result := cast.To[string](panickyStringer{})
		if result != "" {
			t.Errorf("expected empty string (zero value), got %q", result)
		}
	})
}
