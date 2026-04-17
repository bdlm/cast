package cast_test

import (
	"errors"
	"fmt"
	"testing"

	bdlmerrors "github.com/bdlm/errors/v2"
	std_error "github.com/bdlm/std/v2/errors"

	"github.com/bdlm/cast/v2"
)

// testStringer is a local type implementing fmt.Stringer.
type testStringer struct{ val string }

func (s testStringer) String() string { return s.val }

func TestToErrorTarget(t *testing.T) {
	t.Run("stdlib error returned as-is", func(t *testing.T) {
		input := fmt.Errorf("test error")
		result, err := cast.ToE[error](input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != input {
			t.Errorf("expected %v, got %v", input, result)
		}
	})
	t.Run("bdlm error returned as-is", func(t *testing.T) {
		input := bdlmerrors.Errorf("bdlm error")
		result, err := cast.ToE[error](input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Error() != input.Error() {
			t.Errorf("expected %q, got %q", input.Error(), result.Error())
		}
	})
	t.Run("nil input returns nil error", func(t *testing.T) {
		result, err := cast.ToE[error](nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})
	t.Run("string input fails", func(t *testing.T) {
		_, err := cast.ToE[error]("not an error")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("int input fails", func(t *testing.T) {
		_, err := cast.ToE[error](42)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

func TestToStdErrorTarget(t *testing.T) {
	t.Run("bdlm error satisfies std_error.Error", func(t *testing.T) {
		input := bdlmerrors.Errorf("std error test")
		result, err := cast.ToE[std_error.Error](input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.Error() != input.Error() {
			t.Errorf("expected %q, got %q", input.Error(), result.Error())
		}
	})
	t.Run("nil input returns nil", func(t *testing.T) {
		result, err := cast.ToE[std_error.Error](nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})
	t.Run("stdlib error lacks Is/Unwrap, fails", func(t *testing.T) {
		// fmt.Errorf returns *errors.errorString which has Error() but not Is() or
		// Unwrap(), so it does not satisfy std_error.Error.
		_, err := cast.ToE[std_error.Error](fmt.Errorf("plain error"))
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("non-error string input fails", func(t *testing.T) {
		_, err := cast.ToE[std_error.Error]("not an error")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

func TestToStringerTarget(t *testing.T) {
	t.Run("Stringer input returned as-is", func(t *testing.T) {
		input := testStringer{"hello"}
		result, err := cast.ToE[fmt.Stringer](input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.String() != input.String() {
			t.Errorf("expected %q, got %q", input.String(), result.String())
		}
	})
	t.Run("nil input returns nil", func(t *testing.T) {
		result, err := cast.ToE[fmt.Stringer](nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})
	t.Run("plain string does not satisfy fmt.Stringer, fails", func(t *testing.T) {
		// string has no String() method so does not implement fmt.Stringer.
		_, err := cast.ToE[fmt.Stringer]("plain string")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("int does not satisfy fmt.Stringer, fails", func(t *testing.T) {
		_, err := cast.ToE[fmt.Stringer](42)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("bdlm error has Format but not String, fails", func(t *testing.T) {
		// *bdlmerrors.E implements fmt.Formatter (Format method) but not
		// fmt.Stringer (String method), so it does not satisfy fmt.Stringer.
		_, err := cast.ToE[fmt.Stringer](bdlmerrors.Errorf("test"))
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}
