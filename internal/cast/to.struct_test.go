package cast_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

// testStructCase runs a single ToStructE[TTo] test.
func testStructCase[TTo any](t *testing.T, test testCase) {
	t.Helper()
	var typ TTo
	name := fmt.Sprintf("%T", typ)
	t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
		actual, err := cast.ToE[TTo](test.in)
		testInfo := fmt.Sprintf(`
case: ToStructE[%s]
input: %v (%T)
expect error: %v; actual error: %v
expected result: %v (%T); actual result: %v (%T)
test: %#v
		`,
			name, test.in, test.in,
			test.expectErr, err,
			test.expect, test.expect,
			actual, actual,
			test,
		)

		if err != nil && !test.expectErr {
			t.Error("1. expected nil, got error", testInfo)
		} else if err == nil && test.expectErr {
			t.Error("2. expected error, got nil", testInfo)
		} else if err != nil && !errors.Is(err, cast.ErrorUnableToCast) {
			t.Error("3. expected cast.ErrorUnableToCast, got different error type", testInfo)
		} else if err == nil && !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("4. expected %v to equal %v %s", test.expect, actual, testInfo)
		}
	})
}

// ── types used across tests ───────────────────────────────────────────────────

type flatStruct struct {
	Name  string
	Age   int
	Score float64
	Valid bool
}

type nestedStruct struct {
	Title  string
	Inner  flatStruct
	Counts []int
}

type hobbyStruct struct {
	Name      string
	Equipment []string
}

type personStruct struct {
	Name     string
	Age      int
	Hobbies  []hobbyStruct
	Birthday time.Time
}

type embeddedBase struct {
	ID int
}

type embeddedStruct struct {
	embeddedBase
	Label string
}

type EmbeddedPtrBase struct {
	Name string
}

type embeddedPtrStruct struct {
	*EmbeddedPtrBase
	Age int
}

// ── ToStructE: flat struct ────────────────────────────────────────────────────

func TestStructFlat(t *testing.T) {
	testStructCase[flatStruct](t, testCase{
		in:     map[string]any{"Name": "Alice", "Age": "30", "Score": "9.5", "Valid": "true"},
		expect: flatStruct{Name: "Alice", Age: 30, Score: 9.5, Valid: true},
	})
}

func TestStructFlatStringMap(t *testing.T) {
	testStructCase[flatStruct](t, testCase{
		in:     map[string]string{"Name": "Bob", "Age": "25", "Score": "7.1", "Valid": "false"},
		expect: flatStruct{Name: "Bob", Age: 25, Score: 7.1, Valid: false},
	})
}

func TestStructPartialKeys(t *testing.T) {
	// Missing keys leave zero values; no error by default.
	testStructCase[flatStruct](t, testCase{
		in:     map[string]any{"Name": "Carol"},
		expect: flatStruct{Name: "Carol"},
	})
}

func TestStructUnknownKeys(t *testing.T) {
	// Unknown keys are silently ignored by default.
	testStructCase[flatStruct](t, testCase{
		in:     map[string]any{"Name": "Dave", "Extra": "ignored"},
		expect: flatStruct{Name: "Dave"},
	})
}

func TestStructUnknownKeysStrict(t *testing.T) {
	// STRICT causes an error when source has keys with no matching field.
	t.Run("strict unknown key", func(t *testing.T) {
		_, err := cast.ToE[flatStruct](
			map[string]any{"Name": "Eve", "Extra": "bad"},
			cast.Op{Flag: cast.STRICT, Val: true},
		)
		if err == nil {
			t.Error("expected error for unknown key in STRICT mode, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

func TestStructMissingKeyStrict(t *testing.T) {
	// STRICT causes an error when a required field has no matching key.
	t.Run("strict missing key", func(t *testing.T) {
		_, err := cast.ToE[flatStruct](
			map[string]any{"Name": "Frank"},
			cast.Op{Flag: cast.STRICT, Val: true},
		)
		if err == nil {
			t.Error("expected error for missing key in STRICT mode, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

func TestStructUnconvertibleField(t *testing.T) {
	// A value that can't cast to the field type is skipped (zero retained).
	testStructCase[flatStruct](t, testCase{
		in:     map[string]any{"Name": "Grace", "Age": "notanumber"},
		expect: flatStruct{Name: "Grace", Age: 0},
	})
}

func TestStructUnconvertibleFieldStrict(t *testing.T) {
	t.Run("strict unconvertible field", func(t *testing.T) {
		_, err := cast.ToE[flatStruct](
			map[string]any{"Name": "Hank", "Age": "notanumber", "Score": "0", "Valid": "false"},
			cast.Op{Flag: cast.STRICT, Val: true},
		)
		if err == nil {
			t.Error("expected error for unconvertible field in STRICT mode, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// ── ToStructE: nested structs and slices ──────────────────────────────────────

func TestStructNested(t *testing.T) {
	testStructCase[nestedStruct](t, testCase{
		in: map[string]any{
			"Title":  "test",
			"Inner":  map[string]any{"Name": "x", "Age": "1", "Score": "2.0", "Valid": "true"},
			"Counts": []any{1, 2, 3},
		},
		expect: nestedStruct{
			Title:  "test",
			Inner:  flatStruct{Name: "x", Age: 1, Score: 2.0, Valid: true},
			Counts: []int{1, 2, 3},
		},
	})
}

func TestStructSliceOfStructs(t *testing.T) {
	testStructCase[personStruct](t, testCase{
		in: map[string]any{
			"Name": "John Doe",
			"Age":  "32",
			"Hobbies": []map[string]any{
				{"Name": "golf", "Equipment": []string{"clubs"}},
				{"Name": "biking", "Equipment": []string{"bicycle"}},
				{"Name": "camping", "Equipment": []string{"tent", "sleeping bag"}},
			},
		},
		expect: personStruct{
			Name: "John Doe",
			Age:  32,
			Hobbies: []hobbyStruct{
				{Name: "golf", Equipment: []string{"clubs"}},
				{Name: "biking", Equipment: []string{"bicycle"}},
				{Name: "camping", Equipment: []string{"tent", "sleeping bag"}},
			},
		},
	})
}

// ── ToStructE: time.Time field ────────────────────────────────────────────────

func TestStructTimeFieldDateOnly(t *testing.T) {
	result, err := cast.ToE[personStruct](map[string]any{
		"Name":     "Jane",
		"Age":      "28",
		"Birthday": "1996-06-15",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := time.Date(1996, 6, 15, 0, 0, 0, 0, time.UTC)
	if !result.Birthday.Equal(expected) {
		t.Errorf("expected Birthday=%v, got %v", expected, result.Birthday)
	}
}

func TestStructTimeFieldRFC3339(t *testing.T) {
	result, err := cast.ToE[personStruct](map[string]any{
		"Name":     "Jane",
		"Birthday": "1996-06-15T12:00:00Z",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := time.Date(1996, 6, 15, 12, 0, 0, 0, time.UTC)
	if !result.Birthday.Equal(expected) {
		t.Errorf("expected Birthday=%v, got %v", expected, result.Birthday)
	}
}

func TestStructTimeFieldPassThrough(t *testing.T) {
	// time.Time value passed directly (no string parsing needed).
	ts := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	result, err := cast.ToE[personStruct](map[string]any{
		"Birthday": ts,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.Birthday.Equal(ts) {
		t.Errorf("expected Birthday=%v, got %v", ts, result.Birthday)
	}
}

// ── ToStructE: source is a struct ────────────────────────────────────────────

func TestStructFromStruct(t *testing.T) {
	src := flatStruct{Name: "Ivan", Age: 40, Score: 3.14, Valid: true}
	testStructCase[flatStruct](t, testCase{
		in:     src,
		expect: src,
	})
}

func TestStructFromStructPointer(t *testing.T) {
	src := &flatStruct{Name: "Julia", Age: 22, Score: 1.0, Valid: false}
	testStructCase[flatStruct](t, testCase{
		in:     src,
		expect: *src,
	})
}

// ── ToStructE: embedded fields ────────────────────────────────────────────────

func TestStructEmbedded(t *testing.T) {
	testStructCase[embeddedStruct](t, testCase{
		in:     map[string]any{"ID": "7", "Label": "hello"},
		expect: embeddedStruct{embeddedBase: embeddedBase{ID: 7}, Label: "hello"},
	})
}

func TestStructEmbeddedPointer(t *testing.T) {
	// Nil pointer-to-struct anonymous fields should be allocated so promoted
	// fields can be hydrated.
	result, err := cast.ToE[embeddedPtrStruct](map[string]any{
		"Name": "Alice",
		"Age":  "30",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.EmbeddedPtrBase == nil {
		t.Fatal("expected embedded pointer to be allocated")
	}
	if result.Name != "Alice" {
		t.Errorf("expected Name=Alice, got %q", result.Name)
	}
	if result.Age != 30 {
		t.Errorf("expected Age=30, got %d", result.Age)
	}
}

// ── ToStructE: invalid target type ───────────────────────────────────────────

func TestStructInvalidTargetType(t *testing.T) {
	t.Run("non-struct target", func(t *testing.T) {
		_, err := cast.ToE[int](map[string]any{"x": 1})
		if err == nil {
			t.Error("expected error for non-struct target type, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

func TestStructNilSource(t *testing.T) {
	t.Run("nil source", func(t *testing.T) {
		_, err := cast.ToE[flatStruct](nil)
		if err == nil {
			t.Error("expected error for nil source, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// ── ToStruct: thin wrapper ────────────────────────────────────────────────────

func TestToStructWrapper(t *testing.T) {
	result := cast.To[flatStruct](map[string]any{"Name": "Kurt", "Age": "55"})
	if result.Name != "Kurt" || result.Age != 55 {
		t.Errorf("unexpected result: %+v", result)
	}
}

func TestToStructWrapperError(t *testing.T) {
	// On error, ToStruct returns the zero value.
	result := cast.To[flatStruct](nil)
	if !reflect.DeepEqual(result, flatStruct{}) {
		t.Errorf("expected zero value on error, got %+v", result)
	}
}

// ── PRIVATE: unexported field hydration via unsafe ───────────────────────────

func TestStructPrivateFieldHydration(t *testing.T) {
	// With PRIVATE=true, ToStructE must set unexported struct fields via the
	// unsafe.Pointer path in hydrateStruct. Round-trip through ToE[map] with
	// PRIVATE=true to verify the field was actually written.
	type withPrivate struct {
		Public  string
		private string //nolint:unused
	}
	src := map[string]any{
		"Public":  "hello",
		"private": "secret",
	}
	result, err := cast.ToE[withPrivate](src, cast.Op{Flag: cast.PRIVATE, Val: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Public != "hello" {
		t.Errorf("expected Public=hello, got %q", result.Public)
	}
	// Round-trip: convert the struct back to a map with PRIVATE=true so that
	// extractFieldValue reads the unexported field. If the unsafe set worked,
	// the value is "secret"; if not, it would be "" (zero value).
	back, mapErr := cast.ToE[map[string]any](result, cast.Op{Flag: cast.PRIVATE, Val: true})
	if mapErr != nil {
		t.Fatalf("round-trip to map error: %v", mapErr)
	}
	if back["private"] != "secret" {
		t.Errorf("expected private=secret (unsafe set succeeded), got %v", back["private"])
	}
}

// ── castToType pointer branch ─────────────────────────────────────────────────

func TestStructPointerField(t *testing.T) {
	// A struct field of type *string causes castToType to hit the reflect.Pointer
	// branch (lines 173-180 of util.reflect.go), which allocates a new pointer
	// and sets its element.
	type withPtr struct {
		Name *string
		Age  int
	}
	result, err := cast.ToE[withPtr](map[string]any{
		"Name": "Alice",
		"Age":  30,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name == nil {
		t.Fatal("expected Name to be non-nil pointer")
	}
	if *result.Name != "Alice" {
		t.Errorf("expected *Name=Alice, got %q", *result.Name)
	}
	if result.Age != 30 {
		t.Errorf("expected Age=30, got %d", result.Age)
	}
}

// ── To[T] path: struct via ToE dispatch ──────────────────────────────────────

func TestToEStructDispatch(t *testing.T) {
	result, err := cast.ToE[flatStruct](map[string]any{"Name": "Leo", "Age": "33"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "Leo" || result.Age != 33 {
		t.Errorf("unexpected result: %+v", result)
	}
}

func TestToStructDispatch(t *testing.T) {
	result := cast.To[flatStruct](map[string]any{"Name": "Mia", "Age": "28"})
	if result.Name != "Mia" || result.Age != 28 {
		t.Errorf("unexpected result: %+v", result)
	}
}

// TestScalarToStructErrors verifies that scalar sources (bool, int, uint, float,
// complex, string) always error for struct targets (documented as ✗).
func TestScalarToStructErrors(t *testing.T) {
	type point struct{ X, Y int }
	for _, tc := range []struct {
		name string
		src  any
	}{
		{"bool → struct", true},
		{"int → struct", 42},
		{"uint → struct", uint(5)},
		{"float64 → struct", float64(1.5)},
		{"complex128 → struct", complex128(1 + 2i)},
		{"string → struct", "hello"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := cast.ToE[point](tc.src)
			if err == nil {
				t.Errorf("expected error for %s, got nil", tc.name)
			}
			if !errors.Is(err, cast.ErrorUnableToCast) {
				t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
			}
		})
	}
}
