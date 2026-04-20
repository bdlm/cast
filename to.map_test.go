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

func TestMapInvalidDefault(t *testing.T) {
	t.Run("string DEFAULT for map[string]int", func(t *testing.T) {
		_, err := cast.ToE[map[string]int](map[string]string{"a": "1"}, cast.Op{cast.DEFAULT, "not a map"})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("int DEFAULT for map[string]string", func(t *testing.T) {
		_, err := cast.ToE[map[string]string](map[string]string{"a": "1"}, cast.Op{cast.DEFAULT, 42})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("map[string]string DEFAULT for map[string]int", func(t *testing.T) {
		_, err := cast.ToE[map[string]int](map[string]string{"a": "1"}, cast.Op{cast.DEFAULT, map[string]string{"x": "fallback"}})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("compatible DEFAULT for map[string]int does not error", func(t *testing.T) {
		_, err := cast.ToE[map[string]int](map[string]string{"a": "1"}, cast.Op{cast.DEFAULT, map[string]int{"x": -1}})
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})
}

// TestMapFromMapErrors covers the key-cast and value-cast failure paths in mapFromMap.
func TestMapFromMapErrors(t *testing.T) {
	t.Run("key cast failure errors", func(t *testing.T) {
		// "not_an_int" cannot be parsed as an int key.
		src := map[string]int{"not_an_int": 5}
		_, err := cast.ToE[map[int]int](src)
		if err == nil {
			t.Error("expected error for uncastable map key, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("value cast failure errors", func(t *testing.T) {
		// "bad" cannot be parsed as an int value.
		src := map[string]string{"a": "bad"}
		_, err := cast.ToE[map[string]int](src)
		if err == nil {
			t.Error("expected error for uncastable map value, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

// TestMapFromSliceValueError covers the value-cast failure path in mapFromSlice.
func TestMapFromSliceValueError(t *testing.T) {
	_, err := cast.ToE[map[int]int]([]string{"bad"})
	if err == nil {
		t.Error("expected error for uncastable slice element, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

// TestMapFromScalarSourceError covers toMap's default: case — scalar inputs are not
// map, struct, or slice/array and must return an error.
func TestMapFromScalarSourceError(t *testing.T) {
	_, err := cast.ToE[map[string]int](42)
	if err == nil {
		t.Error("expected error for scalar source → map, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

// TestMapFromStructUnexportedScalars covers extractFieldValue's per-kind branches
// (Bool, Uint, Float, Complex, String) for unexported fields when PRIVATE=true.
func TestMapFromStructUnexportedScalars(t *testing.T) {
	type allPrivate struct {
		PubName    string
		privBool    bool       //nolint:unused
		privUint    uint       //nolint:unused
		privFloat   float64    //nolint:unused
		privComplex complex128 //nolint:unused
		privString  string     //nolint:unused
	}
	src := allPrivate{
		PubName:     "test",
		privBool:    true,
		privUint:    42,
		privFloat:   3.14,
		privComplex: complex(1, 2),
		privString:  "secret",
	}
	result, err := cast.ToE[map[string]any](src, cast.Op{cast.PRIVATE, true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["PubName"] != "test" {
		t.Errorf("expected PubName=test, got %v", result["PubName"])
	}
	for _, field := range []string{"privBool", "privUint", "privFloat", "privComplex", "privString"} {
		if result[field] == nil {
			t.Errorf("expected private field %q to be present in result", field)
		}
	}
}

// TestMapFromStructUnextractableField covers the extractFieldValue default case
// (returns false) for unexported fields with non-scalar kinds (e.g. chan).
func TestMapFromStructUnextractableField(t *testing.T) {
	type withChan struct {
		Name   string
		privCh chan int //nolint:unused
	}
	src := withChan{Name: "hello", privCh: make(chan int)}

	t.Run("non-strict: unextractable field skipped", func(t *testing.T) {
		result, err := cast.ToE[map[string]any](src, cast.Op{cast.PRIVATE, true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result["Name"] != "hello" {
			t.Errorf("expected Name=hello, got %v", result["Name"])
		}
		if _, ok := result["privCh"]; ok {
			t.Error("expected privCh to be skipped (non-scalar unexported)")
		}
	})
	t.Run("strict: unextractable field returns error", func(t *testing.T) {
		_, err := cast.ToE[map[string]any](src,
			cast.Op{cast.PRIVATE, true},
			cast.Op{cast.STRICT, true},
		)
		if err == nil {
			t.Error("expected error in strict mode for non-scalar unexported field, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

// TestMapFromStructEmbeddedUnexported covers the early-continue in collectStructFields
// for unexported anonymous (embedded) fields when PRIVATE=false.
func TestMapFromStructEmbeddedUnexported(t *testing.T) {
	type inner struct{ val int } //nolint:unused
	type outer struct {
		inner
		Name string
	}
	src := outer{Name: "visible"}
	result, err := cast.ToE[map[string]any](src) // PRIVATE=false
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["Name"] != "visible" {
		t.Errorf("expected Name=visible, got %v", result["Name"])
	}
	// The unexported embedded field should not appear.
	if _, ok := result["inner"]; ok {
		t.Error("expected unexported embedded field to be skipped")
	}
}

// TestMapFromStructKeyIncompatible covers the key-cast-fail branch in collectStructFields
// when the field name cannot be cast to the target map key type.
func TestMapFromStructKeyIncompatible(t *testing.T) {
	type simple struct{ Score int }
	src := simple{Score: 5}

	t.Run("non-strict: incompatible key type skips field", func(t *testing.T) {
		// map[int]int target: field name "Score" cannot be parsed as int → skip
		result, err := cast.ToE[map[int]int](src)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected empty map (field skipped), got %v", result)
		}
	})
	t.Run("strict: incompatible key type returns error", func(t *testing.T) {
		_, err := cast.ToE[map[int]int](src, cast.Op{cast.STRICT, true})
		if err == nil {
			t.Error("expected error in strict mode for incompatible key type, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

// TestMapFromStructNestedMapValue covers the case reflect.Map branch in
// collectStructFields when the target map's value type is itself a map.
func TestMapFromStructNestedMapValue(t *testing.T) {
	type inner struct{ Val int }
	type outer struct {
		Name  string
		Inner inner
	}
	src := outer{Name: "top", Inner: inner{Val: 7}}
	result, err := cast.ToE[map[string]map[string]any](src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	innerMap, ok := result["Inner"]
	if !ok {
		t.Fatal("expected Inner key in result map")
	}
	if innerMap["Val"] != 7 {
		t.Errorf("expected Inner.Val=7, got %v", innerMap["Val"])
	}
}

// TestCollectStructFieldsEmbeddedError covers the error-return in
// collectStructFields when the anonymous-field recursion returns an error
// (line 117-119 in to.map.go). Requires strict=true so the nested error
// propagates instead of being skipped.
func TestCollectStructFieldsEmbeddedError(t *testing.T) {
	type BadInner struct{ Tags []string }
	type OuterEmbed struct {
		BadInner
		Score int
	}
	src := OuterEmbed{BadInner: BadInner{Tags: []string{"a"}}, Score: 5}
	_, err := cast.ToE[map[string]int](src, cast.Op{cast.STRICT, true})
	if err == nil {
		t.Fatal("expected error from embedded struct recursion, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

// TestCollectStructFieldsMapValTypeError covers the strict error-return when the
// target map value type is itself a map and the nested mapFromStruct call fails
// (lines 174-178 in to.map.go). The outer struct has only a nested struct field
// so that the map-valtype branch is reached before any scalar field can fail.
func TestCollectStructFieldsMapValTypeError(t *testing.T) {
	type Inner struct{ Tags []string } // Tags []string can't cast to int
	type Outer struct{ Inner Inner }   // only field — no scalar to fail first
	src := Outer{Inner: Inner{Tags: []string{"a"}}}
	_, err := cast.ToE[map[string]map[string]int](src, cast.Op{cast.STRICT, true})
	if err == nil {
		t.Fatal("expected error from map-valtype nested struct failure, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

// TestCollectStructFieldsAnyValTypeNestedError covers lines 164-168 in
// to.map.go: the empty-interface (any) valType branch where mapFromStruct on
// the nested struct returns an error under strict mode. The outer field name
// "True" casts to bool (the outer key type), reaching the any-valType branch;
// the inner struct's field "Score" can't cast to bool as a key, causing the
// nested mapFromStruct to fail with strict=true.
func TestCollectStructFieldsAnyValTypeNestedError(t *testing.T) {
	type Inner struct{ Score int }
	type Outer struct{ True Inner } // "True" parses as bool key
	src := Outer{True: Inner{Score: 5}}
	_, err := cast.ToE[map[bool]any](src, cast.Op{cast.STRICT, true})
	if err == nil {
		t.Fatal("expected error from nested mapFromStruct failure (any valType), got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

// TestMapFromSliceNilElement covers castToSliceType's nil-source branch
// (line 143-145 in util.reflect.go) via a map whose value type is []int and
// whose source contains a nil any element.
func TestMapFromSliceNilElement(t *testing.T) {
	result, err := cast.ToE[map[string][]int](map[string]any{"a": nil})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result["a"], []int{}) {
		t.Errorf("expected empty []int, got %v", result["a"])
	}
}

// TestMapFromStructNonEmptyInterfaceValType exercises collectStructFields'
// valType.NumMethod() != 0 branch: when the target map value type is a non-empty
// interface, struct fields are cast directly rather than wrapped in a nested map.
func TestMapFromStructNonEmptyInterfaceValType(t *testing.T) {
	// testStringer is defined in to.string_test.go (same package cast_test).
	type notStringer struct{ n int }
	type mixedStringers struct {
		Name    testStringer
		NoStr   notStringer // struct, does NOT implement fmt.Stringer
	}
	src := mixedStringers{Name: testStringer{"hello"}, NoStr: notStringer{42}}

	t.Run("non-strict: implementing field kept, non-implementing skipped", func(t *testing.T) {
		result, err := cast.ToE[map[string]fmt.Stringer](src)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result["Name"] == nil || result["Name"].String() != "hello" {
			t.Errorf("expected Name=hello Stringer, got %v", result["Name"])
		}
		if _, ok := result["NoStr"]; ok {
			t.Error("expected NoStr to be skipped (does not implement Stringer)")
		}
	})
	t.Run("strict: non-implementing struct field returns error", func(t *testing.T) {
		// NoStr is a struct that doesn't implement Stringer → strict hits the
		// non-empty interface error return inside the nested-struct if-block.
		_, err := cast.ToE[map[string]fmt.Stringer](src, cast.Op{cast.STRICT, true})
		if err == nil {
			t.Error("expected error in strict mode for non-Stringer struct field, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

// TestMapFromStructNestedStructToScalar exercises collectStructFields' default:
// branch in the nested-struct valType switch: a struct field cast to a scalar
// target type (string) goes through castToType/castToKind/toString.
func TestMapFromStructNestedStructToScalar(t *testing.T) {
	type point struct{ X, Y int }
	type shape struct {
		P    point
		Name string
	}
	src := shape{P: point{1, 2}, Name: "circle"}

	t.Run("nested struct cast to string succeeds", func(t *testing.T) {
		result, err := cast.ToE[map[string]string](src)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result["Name"] != "circle" {
			t.Errorf("expected Name=circle, got %v", result["Name"])
		}
		// P is a struct cast to string via JSON marshal
		if result["P"] == "" {
			t.Error("expected non-empty JSON for nested struct P")
		}
	})
	t.Run("nested struct cast to int fails in strict mode", func(t *testing.T) {
		type withNested struct {
			N     point
			Score int
		}
		_, err := cast.ToE[map[string]int](withNested{N: point{1, 2}, Score: 5}, cast.Op{cast.STRICT, true})
		if err == nil {
			t.Error("expected error: struct→int cast fails in strict mode")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("nested struct cast to int fails non-strict: field skipped", func(t *testing.T) {
		type withNested struct {
			N     point
			Score int
		}
		// non-strict: N (struct→int fails) skipped; Score included
		result, err := cast.ToE[map[string]int](withNested{N: point{1, 2}, Score: 5})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result["Score"] != 5 {
			t.Errorf("expected Score=5, got %v", result["Score"])
		}
		if _, ok := result["N"]; ok {
			t.Error("expected N to be skipped")
		}
	})
}

func TestMapInterfaceValueAssignability(t *testing.T) {
	type Stringer interface{ String() string }
	type myStringer struct{ s string }
	type notStringer struct{ n int }

	t.Run("value implements interface map value type", func(t *testing.T) {
		ms := myStringer{"hello"}
		result, err := cast.ToE[map[string]any](map[string]any{"k": ms})
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if result["k"] != ms {
			t.Errorf("expected %v, got %v", ms, result["k"])
		}
	})
	t.Run("nil value for interface map value type", func(t *testing.T) {
		result, err := cast.ToE[map[string]any](map[string]any{"k": nil})
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if result["k"] != nil {
			t.Errorf("expected nil, got %v", result["k"])
		}
	})
}

// TestMapFromSliceKeyCastError exercises the key-cast error path in
// mapFromSlice (to.map.go:213-215): when the slice index cannot be cast to the
// map key type, mapFromSlice returns an error. A struct{} key type triggers
// this because castToKind has no handler for reflect.Struct.
func TestMapFromSliceKeyCastError(t *testing.T) {
	_, err := cast.ToE[map[struct{}]int]([]int{1, 2, 3})
	if err == nil {
		t.Fatal("expected error casting slice index to struct{} key, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}
