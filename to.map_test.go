package cast_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
)

func testMapCase[TTo any](t *testing.T, test testCase) {
	t.Helper()
	var typ TTo
	name := fmt.Sprintf("%T", typ)
	t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
		actual, err := cast.ToE[TTo](test.in)
		testInfo := fmt.Sprintf(`
case: ToE[%s]
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
		} else if err != nil && !errors.Is(err, cast.Error) {
			t.Error("3. expected cast.Error, got different error type", testInfo)
		} else if err == nil && !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("4. expected %v to equal %v %s", test.expect, actual, testInfo)
		}
	})
}

// TestMapFromMap tests map-to-map conversions.
func TestMapFromMap(t *testing.T) {
	// identity: same key/value types
	testMapCase[map[string]int](t, testCase{
		in:        map[string]int{"a": 1, "b": 2},
		expect:    map[string]int{"a": 1, "b": 2},
		expectErr: false,
	})

	// value cast: string values → int values
	testMapCase[map[string]int](t, testCase{
		in:        map[string]string{"x": "10", "y": "20"},
		expect:    map[string]int{"x": 10, "y": 20},
		expectErr: false,
	})

	// key cast: int keys → string keys
	testMapCase[map[string]string](t, testCase{
		in:        map[int]string{1: "one", 2: "two"},
		expect:    map[string]string{"1": "one", "2": "two"},
		expectErr: false,
	})

	// target with any{} values
	testMapCase[map[string]any](t, testCase{
		in:        map[string]int{"a": 1},
		expect:    map[string]any{"a": 1},
		expectErr: false,
	})

	// empty map
	testMapCase[map[string]int](t, testCase{
		in:        map[string]int{},
		expect:    map[string]int{},
		expectErr: false,
	})

	// DUPLICATE_KEY_ERROR=true: duplicate after key cast should error
	t.Run("map[string]int: duplicate key error", func(t *testing.T) {
		// Two int keys that both cast to the same string would be unusual;
		// instead test with a map that has a duplicate at the Go level — only
		// possible through an any{} source map.
		src := map[any]int{"1": 10, 1: 20} // both keys cast to string "1"
		_, err := cast.ToE[map[string]int](src, cast.Op{Flag: cast.DUPLICATE_KEY_ERROR, Val: true})
		if err == nil {
			t.Error("expected error for duplicate key, got nil")
		}
	})

	// DUPLICATE_KEY_ERROR=false (default): last write wins, no error
	t.Run("map[string]int: duplicate key no error", func(t *testing.T) {
		src := map[any]int{"1": 10, 1: 20}
		_, err := cast.ToE[map[string]int](src)
		if err != nil {
			t.Errorf("expected no error for duplicate key, got %v", err)
		}
	})

	// nil / invalid source
	testMapCase[map[string]int](t, testCase{
		in:        nil,
		expect:    map[string]int(nil),
		expectErr: true,
	})
}

// TestMapFromStruct tests struct-to-map conversions.
func TestMapFromStruct(t *testing.T) {
	type Simple struct {
		Name   string
		Score  int
		Active bool
	}

	// exported fields only (default)
	testMapCase[map[string]any](t, testCase{
		in:     Simple{Name: "alice", Score: 42, Active: true},
		expect: map[string]any{"Name": "alice", "Score": 42, "Active": true},
	})

	// typed value map
	testMapCase[map[string]string](t, testCase{
		in:     Simple{Name: "bob", Score: 7, Active: false},
		expect: map[string]string{"Name": "bob", "Score": "7", "Active": "false"},
	})

	// pointer source
	testMapCase[map[string]any](t, testCase{
		in:     &Simple{Name: "ptr", Score: 1, Active: true},
		expect: map[string]any{"Name": "ptr", "Score": 1, "Active": true},
	})

	type WithPrivate struct {
		Public  string
		private int //nolint:unused
	}

	// PRIVATE=false: unexported fields skipped
	t.Run("map[string]any: private excluded", func(t *testing.T) {
		src := WithPrivate{Public: "hello"}
		result, err := cast.ToE[map[string]any](src)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if _, ok := result["private"]; ok {
			t.Error("private field should not be present")
		}
		if result["Public"] != "hello" {
			t.Errorf("expected Public=hello, got %v", result["Public"])
		}
	})

	// PRIVATE=true: unexported scalar fields included
	t.Run("map[string]any: private included", func(t *testing.T) {
		src := WithPrivate{Public: "hi", private: 99}
		result, err := cast.ToE[map[string]any](src, cast.Op{Flag: cast.PRIVATE, Val: true})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result["Public"] != "hi" {
			t.Errorf("expected Public=hi, got %v", result["Public"])
		}
		if result["private"] == nil {
			t.Error("expected private field to be present")
		}
	})

	// STRICT=true: unconvertible field returns error
	t.Run("map[string]int: strict mode error on bad field", func(t *testing.T) {
		type Unconvertible struct {
			Name string
			Tags []string
		}
		src := Unconvertible{Name: "x", Tags: []string{"a", "b"}}
		_, err := cast.ToE[map[string]int](src, cast.Op{Flag: cast.STRICT, Val: true})
		if err == nil {
			t.Error("expected error in strict mode for unconvertible field, got nil")
		}
	})

	// STRICT=false (default): unconvertible field skipped
	t.Run("map[string]int: skip unconvertible field", func(t *testing.T) {
		type Unconvertible struct {
			Score int
			Tags  []string
		}
		src := Unconvertible{Score: 5, Tags: []string{"a"}}
		result, err := cast.ToE[map[string]int](src)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result["Score"] != 5 {
			t.Errorf("expected Score=5, got %v", result["Score"])
		}
		if _, ok := result["Tags"]; ok {
			t.Error("Tags should have been skipped")
		}
	})

	// nested struct → map[string]any recursion
	t.Run("map[string]any: nested struct recursion", func(t *testing.T) {
		type Inner struct{ Val int }
		type Outer struct {
			Name  string
			Inner Inner
		}
		src := Outer{Name: "outer", Inner: Inner{Val: 7}}
		result, err := cast.ToE[map[string]any](src)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result["Name"] != "outer" {
			t.Errorf("expected Name=outer, got %v", result["Name"])
		}
		inner, ok := result["Inner"].(map[string]any)
		if !ok {
			t.Fatalf("expected Inner to be map[string]any, got %T", result["Inner"])
		}
		if inner["Val"] != 7 {
			t.Errorf("expected Inner.Val=7, got %v", inner["Val"])
		}
	})

	// embedded struct field promotion
	t.Run("map[string]any: embedded struct promotion", func(t *testing.T) {
		type Base struct{ ID int }
		type Derived struct {
			Base
			Name string
		}
		src := Derived{Base: Base{ID: 3}, Name: "derived"}
		result, err := cast.ToE[map[string]any](src)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result["ID"] != 3 {
			t.Errorf("expected promoted ID=3, got %v", result["ID"])
		}
		if result["Name"] != "derived" {
			t.Errorf("expected Name=derived, got %v", result["Name"])
		}
		if _, ok := result["Base"]; ok {
			t.Error("embedded Base field should be promoted, not present as 'Base'")
		}
	})
}

// TestMapFromSlice tests slice/array-to-map conversions.
func TestMapFromSlice(t *testing.T) {
	// []string → map[int]string
	testMapCase[map[int]string](t, testCase{
		in:        []string{"a", "b", "c"},
		expect:    map[int]string{0: "a", 1: "b", 2: "c"},
		expectErr: false,
	})

	// []int → map[string]int (int index keys cast to string)
	testMapCase[map[string]int](t, testCase{
		in:        []int{10, 20, 30},
		expect:    map[string]int{"0": 10, "1": 20, "2": 30},
		expectErr: false,
	})

	// empty slice
	testMapCase[map[int]string](t, testCase{
		in:        []string{},
		expect:    map[int]string{},
		expectErr: false,
	})

	// array source
	testMapCase[map[int]int](t, testCase{
		in:        [3]int{5, 6, 7},
		expect:    map[int]int{0: 5, 1: 6, 2: 7},
		expectErr: false,
	})
}
