package cast_test

import (
	"errors"
	"testing"

	"github.com/bdlm/cast/v2"
)

// DEFAULT option: returns the default value (alongside the error) when
// conversion fails. ToE still returns the error; To silently discards it.
// The DEFAULT type must match TTo exactly or an error is returned immediately.

func TestDefaultBool(t *testing.T) {
	t.Run("default returned by To on error", func(t *testing.T) {
		result := cast.To[bool]("invalid", cast.Op{cast.DEFAULT, true})
		if result != true {
			t.Errorf("expected true (default), got %v", result)
		}
	})
	t.Run("default not used on success", func(t *testing.T) {
		result := cast.To[bool]("1", cast.Op{cast.DEFAULT, false})
		if result != true {
			t.Errorf("expected true (converted), got %v", result)
		}
	})
	t.Run("wrong DEFAULT type errors immediately even for valid input", func(t *testing.T) {
		_, err := cast.ToE[bool](true, cast.Op{cast.DEFAULT, "true"})
		if err == nil {
			t.Error("expected error for wrong DEFAULT type, got nil")
		}
	})
	t.Run("nil DEFAULT errors (nil doesn't satisfy bool)", func(t *testing.T) {
		_, err := cast.ToE[bool](true, cast.Op{cast.DEFAULT, nil})
		if err == nil {
			t.Error("expected error for nil DEFAULT, got nil")
		}
	})
}

func TestDefaultInt(t *testing.T) {
	t.Run("default returned by To on error", func(t *testing.T) {
		result := cast.To[int]("bad", cast.Op{cast.DEFAULT, int(99)})
		if result != 99 {
			t.Errorf("expected 99 (default), got %v", result)
		}
	})
	t.Run("default not used on success", func(t *testing.T) {
		result := cast.To[int]("42", cast.Op{cast.DEFAULT, int(99)})
		if result != 42 {
			t.Errorf("expected 42, got %v", result)
		}
	})
	t.Run("wrong DEFAULT type errors immediately", func(t *testing.T) {
		_, err := cast.ToE[int](42, cast.Op{cast.DEFAULT, "99"})
		if err == nil {
			t.Error("expected error for wrong DEFAULT type")
		}
	})
	t.Run("int DEFAULT mismatches int8 target", func(t *testing.T) {
		_, err := cast.ToE[int8](42, cast.Op{cast.DEFAULT, int(0)})
		if err == nil {
			t.Error("expected error: int DEFAULT for int8 target")
		}
	})
}

func TestDefaultUint(t *testing.T) {
	t.Run("default returned by To on error", func(t *testing.T) {
		result := cast.To[uint]("bad", cast.Op{cast.DEFAULT, uint(7)})
		if result != 7 {
			t.Errorf("expected 7 (default), got %v", result)
		}
	})
	t.Run("int DEFAULT rejected for uint target", func(t *testing.T) {
		_, err := cast.ToE[uint](42, cast.Op{cast.DEFAULT, int(7)})
		if err == nil {
			t.Error("expected error for int DEFAULT on uint target")
		}
	})
}

func TestDefaultFloat(t *testing.T) {
	t.Run("default returned by To on error", func(t *testing.T) {
		result := cast.To[float64]("bad", cast.Op{cast.DEFAULT, float64(3.14)})
		if result != 3.14 {
			t.Errorf("expected 3.14 (default), got %v", result)
		}
	})
	t.Run("float32 DEFAULT rejected for float64 target", func(t *testing.T) {
		_, err := cast.ToE[float64](1.0, cast.Op{cast.DEFAULT, float32(3.14)})
		if err == nil {
			t.Error("expected error for float32 DEFAULT on float64 target")
		}
	})
}

func TestDefaultString(t *testing.T) {
	t.Run("wrong DEFAULT type errors immediately even for valid input", func(t *testing.T) {
		_, err := cast.ToE[string]("hello", cast.Op{cast.DEFAULT, 42})
		if err == nil {
			t.Error("expected error for wrong DEFAULT type")
		}
	})
	t.Run("default returned by To when JSON marshal fails", func(t *testing.T) {
		type unmarshalable struct{ Ch chan int }
		result := cast.To[string](unmarshalable{Ch: make(chan int)}, cast.Op{cast.DEFAULT, "fallback"})
		if result != "fallback" {
			t.Errorf("expected fallback, got %q", result)
		}
	})
}

func TestDefaultComplex(t *testing.T) {
	t.Run("default returned by To on error", func(t *testing.T) {
		result := cast.To[complex64]("bad", cast.Op{cast.DEFAULT, complex64(1 + 2i)})
		if result != complex64(1+2i) {
			t.Errorf("expected (1+2i) (default), got %v", result)
		}
	})
	t.Run("complex128 DEFAULT rejected for complex64 target", func(t *testing.T) {
		_, err := cast.ToE[complex64](complex64(0), cast.Op{cast.DEFAULT, complex128(1 + 2i)})
		if err == nil {
			t.Error("expected error for complex128 DEFAULT on complex64 target")
		}
	})
}

// ABS option: when casting a negative signed value to an unsigned integer,
// ABS=true takes the absolute value instead of returning an error.

func TestABSOption(t *testing.T) {
	t.Run("negative int → positive uint with ABS=true", func(t *testing.T) {
		result, err := cast.ToE[uint](int(-5), cast.Op{cast.ABS, true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != uint(5) {
			t.Errorf("expected 5, got %v", result)
		}
	})
	t.Run("negative int → uint errors without ABS", func(t *testing.T) {
		_, err := cast.ToE[uint](int(-5))
		if err == nil {
			t.Error("expected error for negative int → uint without ABS")
		}
	})
	t.Run("ABS has no effect on signed target", func(t *testing.T) {
		result, err := cast.ToE[int](int(-5), cast.Op{cast.ABS, true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != int(-5) {
			t.Errorf("expected -5 (ABS irrelevant on signed target), got %v", result)
		}
	})
	t.Run("ABS=true with negative string source", func(t *testing.T) {
		result, err := cast.ToE[uint]("-7", cast.Op{cast.ABS, true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != uint(7) {
			t.Errorf("expected 7, got %v", result)
		}
	})
	t.Run("ABS=true with negative float source → uint", func(t *testing.T) {
		result, err := cast.ToE[uint](float64(-5.5), cast.Op{cast.ABS, true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != uint(5) {
			t.Errorf("expected 5 (floor of 5.5), got %v", result)
		}
	})
}

// LENGTH option on slices: sets the initial backing-array capacity.
// 0 is valid for slices; negative is not.

func TestLengthOptionSlice(t *testing.T) {
	t.Run("LENGTH pre-allocates capacity", func(t *testing.T) {
		result, err := cast.ToE[[]int]([]string{"1", "2", "3"}, cast.Op{cast.LENGTH, 10})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 3 {
			t.Errorf("expected len=3, got %d", len(result))
		}
		if cap(result) < 10 {
			t.Errorf("expected cap >= 10, got %d", cap(result))
		}
	})
	t.Run("LENGTH=0 is valid for slices", func(t *testing.T) {
		_, err := cast.ToE[[]int]([]string{"1", "2"}, cast.Op{cast.LENGTH, 0})
		if err != nil {
			t.Errorf("expected no error for LENGTH=0 on slice, got %v", err)
		}
	})
	t.Run("LENGTH=-1 errors for slices", func(t *testing.T) {
		_, err := cast.ToE[[]int]([]string{"1"}, cast.Op{cast.LENGTH, -1})
		if err == nil {
			t.Error("expected error for LENGTH=-1 on slice")
		}
	})
	t.Run("non-int LENGTH errors", func(t *testing.T) {
		_, err := cast.ToE[[]int]([]string{"1"}, cast.Op{cast.LENGTH, "five"})
		if err == nil {
			t.Error("expected error for non-int LENGTH")
		}
	})
}

// LENGTH option on channels: sets the buffer capacity.
// 0 is invalid for channels (would deadlock); minimum is 1.

func TestLengthOptionChan(t *testing.T) {
	t.Run("LENGTH sets channel buffer capacity", func(t *testing.T) {
		ch, err := cast.ToE[chan int](42, cast.Op{cast.LENGTH, 5})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cap(ch) != 5 {
			t.Errorf("expected cap=5, got %d", cap(ch))
		}
		val := <-ch
		if val != 42 {
			t.Errorf("expected value 42, got %v", val)
		}
	})
	t.Run("LENGTH=0 errors for channels", func(t *testing.T) {
		_, err := cast.ToE[chan int](42, cast.Op{cast.LENGTH, 0})
		if err == nil {
			t.Error("expected error for LENGTH=0 on chan")
		}
	})
	t.Run("default channel capacity is 1", func(t *testing.T) {
		ch, err := cast.ToE[chan int](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if cap(ch) != 1 {
			t.Errorf("expected default cap=1, got %d", cap(ch))
		}
		<-ch
	})
}

// DEFAULT with ToE: on conversion failure, ToE returns (defaultValue, error) —
// both the default AND the error, so callers can inspect the error while still
// getting a usable fallback value.

func TestDefaultToEReturnsBothValueAndError(t *testing.T) {
	t.Run("ToE returns default value AND error on int failure", func(t *testing.T) {
		val, err := cast.ToE[int]("bad", cast.Op{cast.DEFAULT, int(99)})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if val != 99 {
			t.Errorf("expected default value 99, got %v", val)
		}
	})
	t.Run("ToE returns default value AND error on float failure", func(t *testing.T) {
		val, err := cast.ToE[float64]("bad", cast.Op{cast.DEFAULT, float64(3.14)})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if val != 3.14 {
			t.Errorf("expected default value 3.14, got %v", val)
		}
	})
	t.Run("ToE returns default value AND error on bool failure", func(t *testing.T) {
		val, err := cast.ToE[bool]("bad", cast.Op{cast.DEFAULT, true})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if val != true {
			t.Errorf("expected default value true, got %v", val)
		}
	})
}

// UNIQUE_VALUES=false (explicit) must not deduplicate.

func TestUniqueValuesFalseNoDedup(t *testing.T) {
	result, err := cast.ToE[[]int]([]int{1, 2, 2, 3}, cast.Op{cast.UNIQUE_VALUES, false})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 4 {
		t.Errorf("expected 4 elements (no dedup), got %d: %v", len(result), result)
	}
}

// The following tests exercise ops.List() branches by passing options through a
// chan target — makeChan calls ToE[T](from, ops.List()...) forwarding the flags.

func TestListWithABSOption(t *testing.T) {
	// ABS on a chan int target: List() emits Op{ABS, true}, forwarded to toInt.
	// Negative int → signed chan int target, so ABS is irrelevant to the result
	// but it exercises the List() abs branch.
	ch, err := cast.ToE[chan int](int(-5), cast.Op{cast.ABS, true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	val := <-ch
	if val != int(-5) {
		t.Errorf("expected -5 (ABS irrelevant for signed target), got %v", val)
	}
}

func TestListWithUniqueValuesOption(t *testing.T) {
	// UNIQUE_VALUES on a chan []int target: List() emits Op{UNIQUE_VALUES, true},
	// forwarded to toSlice which deduplicates the slice.
	ch, err := cast.ToE[chan []int]([]int{1, 2, 2, 3}, cast.Op{cast.UNIQUE_VALUES, true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result := <-ch
	if len(result) != 3 {
		t.Errorf("expected 3 unique elements, got %d: %v", len(result), result)
	}
}

func TestListWithJSONOption(t *testing.T) {
	// JSON on a chan string target: List() emits Op{JSON, true}, forwarded to toString
	// which JSON-encodes the result string.
	ch, err := cast.ToE[chan string]("hello", cast.Op{cast.JSON, true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	val := <-ch
	if val != `"hello"` {
		t.Errorf(`expected "\"hello\"", got %q`, val)
	}
}

// TestChanLengthInvalidType covers the sizeErr path in toChan when the LENGTH
// option value cannot be parsed as an int.
func TestChanLengthInvalidType(t *testing.T) {
	_, err := cast.ToE[chan int](42, cast.Op{cast.LENGTH, "not-a-number"})
	if err == nil {
		t.Error("expected error for non-int LENGTH value on chan, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

// DECODE=JSON for scalar targets: when normal parsing fails, the string is
// treated as a JSON value, decoded, and then re-cast to the target type.
// This is a fallback — it only fires after the fast parse path fails.

func TestDecodeOptionInt(t *testing.T) {
	t.Run(`JSON-encoded string "1" → int`, func(t *testing.T) {
		// `"1"` is the JSON encoding of the string "1", which parses to int 1.
		result, err := cast.ToE[int](`"1"`, cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != 1 {
			t.Errorf("expected 1, got %v", result)
		}
	})
	t.Run("plain numeric string skips DECODE (fast path)", func(t *testing.T) {
		// "42" parses directly without ever firing json.Unmarshal.
		result, err := cast.ToE[int]("42", cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("expected 42, got %v", result)
		}
	})
	t.Run("invalid JSON with DECODE → error", func(t *testing.T) {
		_, err := cast.ToE[int]("not-json", cast.Op{cast.DECODE, "JSON"})
		if err == nil {
			t.Error("expected error for non-JSON string with DECODE=json, got nil")
		}
	})
	t.Run("JSON number decoded directly → int", func(t *testing.T) {
		// The JSON value 42 (no quotes) — fast path already handles this.
		result, err := cast.ToE[int]("42", cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("expected 42, got %v", result)
		}
	})
}

func TestDecodeOptionFloat(t *testing.T) {
	t.Run(`JSON-encoded string "1.5" → float64`, func(t *testing.T) {
		result, err := cast.ToE[float64](`"1.5"`, cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != 1.5 {
			t.Errorf("expected 1.5, got %v", result)
		}
	})
	t.Run("invalid JSON with DECODE → error", func(t *testing.T) {
		_, err := cast.ToE[float64]("not-json", cast.Op{cast.DECODE, "JSON"})
		if err == nil {
			t.Error("expected error for non-JSON string with DECODE=json, got nil")
		}
	})
}

func TestDecodeOptionBool(t *testing.T) {
	t.Run(`JSON-encoded string "true" → bool`, func(t *testing.T) {
		result, err := cast.ToE[bool](`"true"`, cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != true {
			t.Errorf("expected true, got %v", result)
		}
	})
	t.Run(`JSON boolean true → bool (fast path)`, func(t *testing.T) {
		result, err := cast.ToE[bool]("true", cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != true {
			t.Errorf("expected true, got %v", result)
		}
	})
	t.Run("invalid JSON with DECODE → error", func(t *testing.T) {
		_, err := cast.ToE[bool]("notbool", cast.Op{cast.DECODE, "JSON"})
		if err == nil {
			t.Error("expected error for unparseable string with DECODE=json, got nil")
		}
	})
}
