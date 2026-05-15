package cast_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/bdlm/cast/v2"
	bdlmerrors "github.com/bdlm/errors/v2"
	std_error "github.com/bdlm/std/v2/errors"
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
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("int input fails", func(t *testing.T) {
		_, err := cast.ToE[error](42)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
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
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("non-error string input fails", func(t *testing.T) {
		_, err := cast.ToE[std_error.Error]("not an error")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
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
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("int does not satisfy fmt.Stringer, fails", func(t *testing.T) {
		_, err := cast.ToE[fmt.Stringer](42)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("bdlm error has Format but not String, fails", func(t *testing.T) {
		// *bdlmerrors.E implements fmt.Formatter (Format method) but not
		// fmt.Stringer (String method), so it does not satisfy fmt.Stringer.
		_, err := cast.ToE[fmt.Stringer](bdlmerrors.Errorf("test"))
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

func TestStringFromNil(t *testing.T) {
	result, err := cast.ToE[string](nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestStringFromBytes(t *testing.T) {
	// []byte has a dedicated case: return string(val), nil
	result, err := cast.ToE[string]([]byte("hello"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "hello" {
		t.Errorf("expected \"hello\", got %q", result)
	}
}

func TestStringFromStringer(t *testing.T) {
	t.Run("positive: Stringer returns String() value", func(t *testing.T) {
		result, err := cast.ToE[string](testStringer{"world"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != "world" {
			t.Errorf("expected \"world\", got %q", result)
		}
	})
}

func TestStringFromScalars(t *testing.T) {
	cases := []struct {
		input  any
		expect string
	}{
		{int(42), "42"},
		{int8(-5), "-5"},
		{int64(1000), "1000"},
		{uint(7), "7"},
		{float32(1.5), "1.5"},
		{float64(3.14), "3.14"},
		{bool(true), "true"},
		{bool(false), "false"},
	}
	for _, tc := range cases {
		t.Run(tc.expect, func(t *testing.T) {
			result, err := cast.ToE[string](tc.input)
			if err != nil {
				t.Fatalf("unexpected error for %v: %v", tc.input, err)
			}
			if result != tc.expect {
				t.Errorf("expected %q, got %q", tc.expect, result)
			}
		})
	}
}

func TestStringFromComplex(t *testing.T) {
	// Complex numbers use fmt.Sprintf("%v", val) → "(real+imagi)"
	result, err := cast.ToE[string](complex(1.0, 2.0))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "(1+2i)" {
		t.Errorf("expected \"(1+2i)\", got %q", result)
	}

	result, err = cast.ToE[string](complex64(0 + 0i))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "(0+0i)" {
		t.Errorf("expected \"(0+0i)\", got %q", result)
	}
}

func TestStringFromChan(t *testing.T) {
	// Channels use fmt.Sprintf("%v", val), producing a hex address string.
	// We can't predict the exact value, but it should be a non-empty string
	// starting with "0x".
	ch := make(chan int, 1)
	result, err := cast.ToE[string](ch)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(result, "0x") {
		t.Errorf("expected hex address string, got %q", result)
	}
}

func TestStringFromMap(t *testing.T) {
	t.Run("map marshals to JSON", func(t *testing.T) {
		result, err := cast.ToE[string](map[string]int{"a": 1})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != `{"a":1}` {
			t.Errorf("expected {\"a\":1}, got %q", result)
		}
	})
	t.Run("map with chan value fails JSON marshal → error", func(t *testing.T) {
		_, err := cast.ToE[string](map[string]chan int{"a": make(chan int)})
		if err == nil {
			t.Fatal("expected error for unmarshalable map, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

func TestStringFromSlice(t *testing.T) {
	t.Run("int slice marshals to JSON array", func(t *testing.T) {
		result, err := cast.ToE[string]([]int{1, 2, 3})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != `[1,2,3]` {
			t.Errorf("expected [1,2,3], got %q", result)
		}
	})
	t.Run("string slice marshals to JSON array", func(t *testing.T) {
		result, err := cast.ToE[string]([]string{"a", "b"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != `["a","b"]` {
			t.Errorf(`expected ["a","b"], got %q`, result)
		}
	})
}

func TestStringFromStructJSONMarshal(t *testing.T) {
	t.Run("positive: plain struct → JSON string", func(t *testing.T) {
		type point struct {
			X int
			Y int
		}
		result, err := cast.ToE[string](point{X: 1, Y: 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != `{"X":1,"Y":2}` {
			t.Errorf("expected JSON, got %q", result)
		}
	})
	t.Run("negative: struct with chan field fails JSON marshal", func(t *testing.T) {
		type bad struct{ Ch chan int }
		_, err := cast.ToE[string](bad{Ch: make(chan int)})
		if err == nil {
			t.Fatal("expected error for unmarshalable struct, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// TestStringJSONOptionUnmarshalableSource covers the sErr path in toString's
// JSON branch: if the inner (non-JSON) toString fails, the error is propagated.
func TestStringJSONOptionUnmarshalableSource(t *testing.T) {
	// A map with a chan value causes json.Marshal to fail in the non-JSON toString
	// call, which then bubbles up as sErr in the JSON branch.
	src := map[string]chan int{"a": make(chan int)}
	_, err := cast.ToE[string](src, cast.Op{cast.JSON, true})
	if err == nil {
		t.Fatal("expected error for JSON option with unmarshalable source, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

func TestStringJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       any
		expectJSON  string
		expectPlain string
	}{
		{"string", "hello", `"hello"`, "hello"},
		{"string with quotes", `say "hi"`, `"say \"hi\""`, `say "hi"`},
		{"int", 42, `"42"`, "42"},
		{"bool true", true, `"true"`, "true"},
		{"bool false", false, `"false"`, "false"},
		{"float", 3.14, `"3.14"`, "3.14"},
		{"nil", nil, `""`, ""},
		{"empty string", "", `""`, ""},
	}

	for _, test := range tests {
		t.Run(test.name+"/json=true", func(t *testing.T) {
			actual, err := cast.ToE[string](test.input, cast.Op{Flag: cast.JSON, Val: true})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if actual != test.expectJSON {
				t.Errorf("expected %q, got %q", test.expectJSON, actual)
			}
		})
		t.Run(test.name+"/json=false", func(t *testing.T) {
			actual, err := cast.ToE[string](test.input, cast.Op{Flag: cast.JSON, Val: false})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if actual != test.expectPlain {
				t.Errorf("expected %q, got %q", test.expectPlain, actual)
			}
		})
	}
}
