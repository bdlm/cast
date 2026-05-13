package cast_test

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/bdlm/cast/v2"
)

type testCase struct {
	in        any
	expect    any
	err       error
	expectErr bool
}

type testCases map[string][]testCase

func TestSimpleTypes(t *testing.T) {
	for name, cases := range simpleCases {
		switch name {
		case "string":
			testSimpleCases[string](t, cases)
		case "bool":
			testSimpleCases[bool](t, cases)
		case "byte":
			testSimpleCases[byte](t, cases)
		case "rune":
			testSimpleCases[rune](t, cases)
		case "int":
			testSimpleCases[int](t, cases)
		case "int8":
			testSimpleCases[int8](t, cases)
		case "int16":
			testSimpleCases[int16](t, cases)
		case "int32":
			testSimpleCases[int32](t, cases)
		case "int64":
			testSimpleCases[int64](t, cases)
		case "uint":
			testSimpleCases[uint](t, cases)
		case "uint8":
			testSimpleCases[uint8](t, cases)
		case "uint16":
			testSimpleCases[uint16](t, cases)
		case "uint32":
			testSimpleCases[uint32](t, cases)
		case "uint64":
			testSimpleCases[uint64](t, cases)
		case "uintptr":
			testSimpleCases[uintptr](t, cases)
		case "float32":
			testSimpleCases[float32](t, cases)
		case "float64":
			testSimpleCases[float64](t, cases)
		case "complex64":
			testSimpleCases[complex64](t, cases)
		case "complex128":
			testSimpleCases[complex128](t, cases)
		}
	}
}

func testSimpleCases[TTo any](t *testing.T, cases []testCase) {
	var typ TTo
	name := fmt.Sprintf("%T", typ)

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
			actual, err := cast.ToE[TTo](test.in)
			testInfo := fmt.Sprintf(`
case: ToE[%s]
input: %v (%T)
expect error: %v; actual error: %v
expected result: %v (%T); actual result: %v (%T)
test: %#v
			`,
				name,
				test.in,
				test.in,
				test.expectErr,
				err,
				test.expect,
				test.expect,
				actual,
				actual,
				test,
			)

			if err != nil && !test.expectErr {
				t.Error("1. expected nil, got error", testInfo)
			} else if err == nil && test.expectErr {
				t.Error("2. expected error, got nil", testInfo)
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Error("3. expected cast.Error, got different error type", testInfo)
			} else if nil == err && !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("4. expected %v to equal %v %s", test.expect, actual, testInfo)
			}
		})
	}
}

// concError is a concrete struct type implementing error. When used as TTo,
// to.Type().Kind() == reflect.Struct → ToE's default: branch, which detects
// the error interface and builds an error-string result.
type concError struct{}

func (e concError) Error() string { return "concError" }

// concStringer is a concrete struct type implementing fmt.Stringer, similarly
// exercising the fmt.Stringer sub-branch of ToE's default: case.
type concStringer struct{}

func (s concStringer) String() string { return "concStringer" }

func TestToEDefaultCaseConcreteError(t *testing.T) {
	// concError has Kind == Struct, so it hits ToE's default: branch.
	// The branch detects it implements error and attempts an error-string cast,
	// which ultimately fails the type assertion → returns an error.
	_, err := cast.ToE[concError](42)
	if err == nil {
		t.Error("expected error for concrete error type → ToE default branch, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

func TestToEDefaultCaseConcreteStringer(t *testing.T) {
	// concStringer has Kind == Struct, hits ToE's default: branch.
	// Detects fmt.Stringer, builds string result, but final assertion to
	// concStringer fails → returns an error.
	_, err := cast.ToE[concStringer](42)
	if err == nil {
		t.Error("expected error for concrete Stringer type → ToE default branch, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

func TestTo(t *testing.T) {
	t.Run("To positive: string to int", func(t *testing.T) {
		result := cast.To[int]("42")
		if result != 42 {
			t.Errorf("expected 42, got %v", result)
		}
	})
	t.Run("To negative: invalid string returns zero value silently", func(t *testing.T) {
		result := cast.To[int]("bad")
		if result != 0 {
			t.Errorf("expected 0, got %v", result)
		}
	})
}

var simpleCases = testCases{
	"string": {
		{in: true, expect: "true", err: nil, expectErr: false},
		{in: false, expect: "false", err: nil, expectErr: false},
		{in: 1, expect: "1", err: nil, expectErr: false},
		{in: 0, expect: "0", err: nil, expectErr: false},
		{in: "hi", expect: "hi", err: nil, expectErr: false},
		{in: float64(1.1), expect: "1.1", err: nil, expectErr: false},
		{in: float64(-1.1), expect: "-1.1", err: nil, expectErr: false},
	},
	"bool": {
		{in: true, expect: true, err: nil, expectErr: false},
		{in: 1, expect: true, err: nil, expectErr: false},
		{in: 0, expect: false, err: nil, expectErr: false},
		{in: "hi", expect: false, err: nil, expectErr: true},
		{in: float64(1.1), expect: true, err: nil, expectErr: false},
		{in: float64(-1.1), expect: true, err: nil, expectErr: false},
	},
	// int
	"int": {
		{in: 1, expect: int(1), err: nil, expectErr: false},
		{in: "1", expect: int(1), err: nil, expectErr: false},
		{in: 1.1, expect: int(1), err: nil, expectErr: false},
		{in: "1.1", expect: int(1), err: nil, expectErr: false},
		{in: 1.9, expect: int(1), err: nil, expectErr: false},
		{in: "1.9", expect: int(1), err: nil, expectErr: false},
		{in: -1, expect: int(-1), err: nil, expectErr: false},
		{in: "-1", expect: int(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int(-1), err: nil, expectErr: false},
	},
	"int8": {
		{in: 1, expect: int8(1), err: nil, expectErr: false},
		{in: "1", expect: int8(1), err: nil, expectErr: false},
		{in: 1.1, expect: int8(1), err: nil, expectErr: false},
		{in: "1.1", expect: int8(1), err: nil, expectErr: false},
		{in: 1.9, expect: int8(1), err: nil, expectErr: false},
		{in: "1.9", expect: int8(1), err: nil, expectErr: false},
		{in: -1, expect: int8(-1), err: nil, expectErr: false},
		{in: "-1", expect: int8(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int8(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int8(-1), err: nil, expectErr: false},
	},
	"int16": {
		{in: 1, expect: int16(1), err: nil, expectErr: false},
		{in: "1", expect: int16(1), err: nil, expectErr: false},
		{in: 1.1, expect: int16(1), err: nil, expectErr: false},
		{in: "1.1", expect: int16(1), err: nil, expectErr: false},
		{in: 1.9, expect: int16(1), err: nil, expectErr: false},
		{in: "1.9", expect: int16(1), err: nil, expectErr: false},
		{in: -1, expect: int16(-1), err: nil, expectErr: false},
		{in: "-1", expect: int16(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int16(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int16(-1), err: nil, expectErr: false},
	},
	"int32": {
		{in: 1, expect: int32(1), err: nil, expectErr: false},
		{in: "1", expect: int32(1), err: nil, expectErr: false},
		{in: 1.1, expect: int32(1), err: nil, expectErr: false},
		{in: "1.1", expect: int32(1), err: nil, expectErr: false},
		{in: 1.9, expect: int32(1), err: nil, expectErr: false},
		{in: "1.9", expect: int32(1), err: nil, expectErr: false},
		{in: -1, expect: int32(-1), err: nil, expectErr: false},
		{in: "-1", expect: int32(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int32(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int32(-1), err: nil, expectErr: false},
	},
	"int64": {
		{in: 1, expect: int64(1), err: nil, expectErr: false},
		{in: "1", expect: int64(1), err: nil, expectErr: false},
		{in: 1.1, expect: int64(1), err: nil, expectErr: false},
		{in: "1.1", expect: int64(1), err: nil, expectErr: false},
		{in: 1.9, expect: int64(1), err: nil, expectErr: false},
		{in: "1.9", expect: int64(1), err: nil, expectErr: false},
		{in: -1, expect: int64(-1), err: nil, expectErr: false},
		{in: "-1", expect: int64(-1), err: nil, expectErr: false},
		{in: -1.9, expect: int64(-1), err: nil, expectErr: false},
		{in: "-1.9", expect: int64(-1), err: nil, expectErr: false},
	},

	// uint
	"uint": {
		{in: 1, expect: uint(1), err: nil, expectErr: false},
		{in: "1", expect: uint(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint(1), err: nil, expectErr: false},
		{in: -1, expect: uint(0), err: nil, expectErr: true},
		{in: "-1", expect: uint(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint(0), err: nil, expectErr: true},
	},
	"uint8": {
		{in: 1, expect: uint8(1), err: nil, expectErr: false},
		{in: "1", expect: uint8(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint8(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint8(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint8(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint8(1), err: nil, expectErr: false},
		{in: -1, expect: uint8(0), err: nil, expectErr: true},
		{in: "-1", expect: uint8(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint8(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint8(0), err: nil, expectErr: true},
	},
	"uint16": {
		{in: 1, expect: uint16(1), err: nil, expectErr: false},
		{in: "1", expect: uint16(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint16(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint16(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint16(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint16(1), err: nil, expectErr: false},
		{in: -1, expect: uint16(0), err: nil, expectErr: true},
		{in: "-1", expect: uint16(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint16(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint16(0), err: nil, expectErr: true},
	},
	"uint32": {
		{in: 1, expect: uint32(1), err: nil, expectErr: false},
		{in: "1", expect: uint32(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint32(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint32(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint32(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint32(1), err: nil, expectErr: false},
		{in: -1, expect: uint32(0), err: nil, expectErr: true},
		{in: "-1", expect: uint32(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint32(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint32(0), err: nil, expectErr: true},
	},
	"uint64": {
		{in: 1, expect: uint64(1), err: nil, expectErr: false},
		{in: "1", expect: uint64(1), err: nil, expectErr: false},
		{in: 1.1, expect: uint64(1), err: nil, expectErr: false},
		{in: "1.1", expect: uint64(1), err: nil, expectErr: false},
		{in: 1.9, expect: uint64(1), err: nil, expectErr: false},
		{in: "1.9", expect: uint64(1), err: nil, expectErr: false},
		{in: -1, expect: uint64(0), err: nil, expectErr: true},
		{in: "-1", expect: uint64(0), err: nil, expectErr: true},
		{in: -1.9, expect: uint64(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uint64(0), err: nil, expectErr: true},
	},
	"uintptr": {
		{in: 1, expect: uintptr(1), err: nil, expectErr: false},
		{in: "1", expect: uintptr(1), err: nil, expectErr: false},
		{in: 1.1, expect: uintptr(1), err: nil, expectErr: false},
		{in: "1.1", expect: uintptr(1), err: nil, expectErr: false},
		{in: 1.9, expect: uintptr(1), err: nil, expectErr: false},
		{in: "1.9", expect: uintptr(1), err: nil, expectErr: false},
		{in: -1, expect: uintptr(0), err: nil, expectErr: true},
		{in: "-1", expect: uintptr(0), err: nil, expectErr: true},
		{in: -1.9, expect: uintptr(0), err: nil, expectErr: true},
		{in: "-1.9", expect: uintptr(0), err: nil, expectErr: true},
	},
}

// ptrDerefStruct is a plain struct used to test pointer-to-struct dereferencing.
type ptrDerefStruct struct {
	X int
	Y string
}

// TestPointerDerefLoop covers the pointer-unwrapping logic at the top of ToE.
// Each sub-test targets a distinct branch of the loop.
func TestPointerDerefLoop(t *testing.T) {
	// ── scalars (baseline) ────────────────────────────────────────────────────

	t.Run("*int → int", func(t *testing.T) {
		v := 42
		got, err := cast.ToE[int](&v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 42 {
			t.Errorf("expected 42, got %v", got)
		}
	})

	t.Run("**int → int", func(t *testing.T) {
		v := 42
		p := &v
		got, err := cast.ToE[int](&p)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 42 {
			t.Errorf("expected 42, got %v", got)
		}
	})

	t.Run("*int nil → 0, no error", func(t *testing.T) {
		var p *int
		got, err := cast.ToE[int](p)
		if err != nil {
			t.Fatalf("unexpected error for nil *int: %v", err)
		}
		if got != 0 {
			t.Errorf("expected 0, got %v", got)
		}
	})

	// ── multi-level nil fix ───────────────────────────────────────────────────
	// Before the changed=false fix, **int where inner *int is nil produced a
	// typed nil (*int)(nil), which fell through to the default converter path
	// and returned an error instead of 0.

	t.Run("**int inner nil → 0, no error", func(t *testing.T) {
		var inner *int
		got, err := cast.ToE[int](&inner)
		if err != nil {
			t.Fatalf("unexpected error for **int with nil inner: %v", err)
		}
		if got != 0 {
			t.Errorf("expected 0, got %v", got)
		}
	})

	// ── struct dereferencing (target is struct) ───────────────────────────────

	t.Run("*struct → struct", func(t *testing.T) {
		src := ptrDerefStruct{X: 7, Y: "hello"}
		got, err := cast.ToE[ptrDerefStruct](&src)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != src {
			t.Errorf("expected %+v, got %+v", src, got)
		}
	})

	t.Run("**struct → struct", func(t *testing.T) {
		src := ptrDerefStruct{X: 7, Y: "hello"}
		p := &src
		got, err := cast.ToE[ptrDerefStruct](&p)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != src {
			t.Errorf("expected %+v, got %+v", src, got)
		}
	})

	t.Run("**struct inner nil → error", func(t *testing.T) {
		// nil source cannot be hydrated into a struct; expect an error.
		var inner *ptrDerefStruct
		_, err := cast.ToE[ptrDerefStruct](&inner)
		if err == nil {
			t.Fatal("expected error for nil struct source, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})

	// ── pointer-to-interface: must NOT dereference ────────────────────────────
	// errors.New returns *errors.errorString, which satisfies error only through
	// a pointer receiver. Dereferencing would lose the interface satisfaction.

	t.Run("pointer-to-error source not dereferenced for error target", func(t *testing.T) {
		src := errors.New("sentinel")
		got, err := cast.ToE[error](src)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != src {
			t.Errorf("expected same error value, got %v", got)
		}
	})

	// ── pointer target: struct source pointer not dereferenced ─────────────────
	// When TTo is itself a pointer (e.g. *regexp.Regexp), targetIsStruct is
	// false, so a *struct source is left as-is for the named converter.

	t.Run("*regexp.Regexp source not dereferenced for *regexp.Regexp target", func(t *testing.T) {
		re := regexp.MustCompile(`[a-z]+`)
		got, err := cast.ToE[*regexp.Regexp](re)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.String() != re.String() {
			t.Errorf("expected pattern %q, got %q", re.String(), got.String())
		}
	})
}
