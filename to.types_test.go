package cast_test

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
)

// Named types with an underlying basic kind route through the reflect.Convert
// path in ToE: the internal function produces the base type (e.g. int), the
// direct type assertion to TTo fails, and reflect.Convert bridges the gap.

type myInt int
type celsius float32
type myBool bool
type myString string
type tags []string
type myInts []int
type myChan chan int

// textUnmarshalStruct implements encoding.TextUnmarshaler for testing the
// castToStructType TextUnmarshaler path.
type textUnmarshalStruct struct{ Val string }

func (t *textUnmarshalStruct) UnmarshalText(b []byte) error {
	t.Val = string(b)
	return nil
}

var _ encoding.TextUnmarshaler = (*textUnmarshalStruct)(nil)

type textUnmarshalStructSlice []textUnmarshalStruct

func TestNamedIntType(t *testing.T) {
	t.Run("positive: int literal → myInt", func(t *testing.T) {
		result, err := cast.ToE[myInt](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != myInt(42) {
			t.Errorf("expected myInt(42), got %v", result)
		}
	})
	t.Run("positive: string → myInt", func(t *testing.T) {
		result, err := cast.ToE[myInt]("7")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != myInt(7) {
			t.Errorf("expected myInt(7), got %v", result)
		}
	})
	t.Run("negative: invalid string → error", func(t *testing.T) {
		_, err := cast.ToE[myInt]("bad")
		if err == nil {
			t.Error("expected error for invalid input, got nil")
		}
	})
}

func TestNamedFloatType(t *testing.T) {
	t.Run("positive: string → celsius", func(t *testing.T) {
		result, err := cast.ToE[celsius]("98.6")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != celsius(float32(98.6)) {
			t.Errorf("expected celsius(98.6), got %v", result)
		}
	})
	t.Run("negative: invalid string → error", func(t *testing.T) {
		_, err := cast.ToE[celsius]("hot")
		if err == nil {
			t.Error("expected error for invalid input, got nil")
		}
	})
}

func TestNamedBoolType(t *testing.T) {
	t.Run("positive: int 1 → myBool true", func(t *testing.T) {
		result, err := cast.ToE[myBool](1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != myBool(true) {
			t.Errorf("expected myBool(true), got %v", result)
		}
	})
	t.Run("negative: unparseable string → error", func(t *testing.T) {
		_, err := cast.ToE[myBool]("maybe")
		if err == nil {
			t.Error("expected error for unparseable string, got nil")
		}
	})
}

func TestNamedStringType(t *testing.T) {
	t.Run("positive: int → myString", func(t *testing.T) {
		result, err := cast.ToE[myString](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != myString("42") {
			t.Errorf("expected myString(\"42\"), got %v", result)
		}
	})
	t.Run("positive: nil → myString empty", func(t *testing.T) {
		result, err := cast.ToE[myString](nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != myString("") {
			t.Errorf("expected myString(\"\"), got %v", result)
		}
	})
}

func TestNamedSliceType(t *testing.T) {
	t.Run("positive: []string → tags", func(t *testing.T) {
		result, err := cast.ToE[tags]([]string{"a", "b", "c"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expect := tags{"a", "b", "c"}
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("expected %v, got %v", expect, result)
		}
	})
	t.Run("positive: []int → tags (via string conversion)", func(t *testing.T) {
		result, err := cast.ToE[tags]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expect := tags{"1", "2"}
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("expected %v, got %v", expect, result)
		}
	})
	t.Run("scalar string wraps as single-element slice", func(t *testing.T) {
		result, err := cast.ToE[tags]("not a slice")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expect := tags{"not a slice"}
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("expected %v, got %v", expect, result)
		}
	})
}

// TestNamedIntSliceType exercises toSlice's "default:" branch, which calls
// castToType → castToKind for each element instead of the explicit type arms.
func TestNamedIntSliceType(t *testing.T) {
	t.Run("positive: []string → myInts", func(t *testing.T) {
		result, err := cast.ToE[myInts]([]string{"1", "2", "3"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expect := myInts{1, 2, 3}
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("expected %v, got %v", expect, result)
		}
	})
	t.Run("negative: unconvertible element errors", func(t *testing.T) {
		_, err := cast.ToE[myInts]([]string{"1", "bad", "3"})
		if err == nil {
			t.Error("expected error for bad element, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// TestMapWithSliceValues exercises castToSliceType via castToType's reflect.Slice
// branch, reached when a map's value type is itself a slice.
func TestMapWithSliceValues(t *testing.T) {
	t.Run("map[string][]int from map[string][]string", func(t *testing.T) {
		src := map[string][]string{"a": {"1", "2"}, "b": {"3"}}
		result, err := cast.ToE[map[string][]int](src)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result["a"], []int{1, 2}) {
			t.Errorf("expected [1 2], got %v", result["a"])
		}
		if !reflect.DeepEqual(result["b"], []int{3}) {
			t.Errorf("expected [3], got %v", result["b"])
		}
	})
	t.Run("negative: unconvertible slice element errors", func(t *testing.T) {
		src := map[string][]string{"a": {"bad"}}
		_, err := cast.ToE[map[string][]int](src)
		if err == nil {
			t.Error("expected error for bad element, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// TestMapInterfaceValueNonAssignable exercises castToType's reflect.Interface branch
// when the source value does not implement the target interface type.
func TestMapInterfaceValueNonAssignable(t *testing.T) {
	// int does not implement fmt.Stringer; castToType returns an error.
	src := map[string]int{"a": 1}
	_, err := cast.ToE[map[string]fmt.Stringer](src)
	if err == nil {
		t.Error("expected error: int does not implement fmt.Stringer")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

// TestCastToKindAllScalars covers each arm of castToKind by routing through
// mapFromMap → castToType (default branch) → castToKind for each scalar kind.
func TestCastToKindAllScalars(t *testing.T) {
	src := map[string]int{"a": 5}

	t.Run("bool", func(t *testing.T) {
		r, err := cast.ToE[map[string]bool](map[string]int{"a": 1, "b": 0})
		if err != nil {
			t.Fatal(err)
		}
		if !r["a"] || r["b"] {
			t.Errorf("unexpected %v", r)
		}
	})
	t.Run("int8", func(t *testing.T) {
		r, err := cast.ToE[map[string]int8](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("int16", func(t *testing.T) {
		r, err := cast.ToE[map[string]int16](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("int32", func(t *testing.T) {
		r, err := cast.ToE[map[string]int32](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("int64", func(t *testing.T) {
		r, err := cast.ToE[map[string]int64](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("uint", func(t *testing.T) {
		r, err := cast.ToE[map[string]uint](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("uint8", func(t *testing.T) {
		r, err := cast.ToE[map[string]uint8](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("uint16", func(t *testing.T) {
		r, err := cast.ToE[map[string]uint16](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("uint32", func(t *testing.T) {
		r, err := cast.ToE[map[string]uint32](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("uint64", func(t *testing.T) {
		r, err := cast.ToE[map[string]uint64](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("uintptr", func(t *testing.T) {
		r, err := cast.ToE[map[string]uintptr](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("float32", func(t *testing.T) {
		r, err := cast.ToE[map[string]float32](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("float64", func(t *testing.T) {
		r, err := cast.ToE[map[string]float64](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("complex64", func(t *testing.T) {
		r, err := cast.ToE[map[string]complex64](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("complex128", func(t *testing.T) {
		r, err := cast.ToE[map[string]complex128](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != 5 {
			t.Errorf("expected 5, got %v", r["a"])
		}
	})
	t.Run("string", func(t *testing.T) {
		r, err := cast.ToE[map[string]string](src)
		if err != nil {
			t.Fatal(err)
		}
		if r["a"] != "5" {
			t.Errorf("expected \"5\", got %v", r["a"])
		}
	})
	t.Run("unsupported kind returns error", func(t *testing.T) {
		// struct{} as map value type → castToKind Struct → default error arm
		_, err := cast.ToE[map[string]struct{}](map[string]int{"a": 1})
		if err == nil {
			t.Error("expected error for unsupported kind, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// TestCastToTypeFuncBranch exercises castToType's Func branch by using a map
// whose value type is cast.Func[int], forcing castToType to build a closure.
func TestCastToTypeFuncBranch(t *testing.T) {
	r, err := cast.ToE[map[string]cast.Func[int]](map[string]int{"a": 42})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	fn := r["a"]
	if fn == nil {
		t.Fatal("expected non-nil Func")
	}
	if fn() != 42 {
		t.Errorf("expected 42, got %v", fn())
	}
}

// TestCastToTypeChanBranch exercises castToType's Chan branch including the
// LENGTH option, the invalid-length-string error, and the size<1 error.
func TestCastToTypeChanBranch(t *testing.T) {
	t.Run("with LENGTH option", func(t *testing.T) {
		r, err := cast.ToE[map[string]chan int](map[string]int{"a": 42}, cast.Op{cast.LENGTH, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		ch := r["a"]
		if cap(ch) != 2 {
			t.Errorf("expected cap 2, got %v", cap(ch))
		}
		val := <-ch
		if val != 42 {
			t.Errorf("expected 42, got %v", val)
		}
	})
	t.Run("invalid LENGTH string errors", func(t *testing.T) {
		_, err := cast.ToE[map[string]chan int](map[string]int{"a": 42}, cast.Op{cast.LENGTH, "bad"})
		if err == nil {
			t.Fatal("expected error for bad LENGTH, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
	t.Run("LENGTH zero errors (size < 1)", func(t *testing.T) {
		_, err := cast.ToE[map[string]chan int](map[string]int{"a": 42}, cast.Op{cast.LENGTH, 0})
		if err == nil {
			t.Fatal("expected error for LENGTH=0, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// TestCastToSliceTypeDefaultBranch exercises castToSliceType's default branch
// (scalar source → single-element slice) via a map whose value type is []int.
func TestCastToSliceTypeDefaultBranch(t *testing.T) {
	t.Run("scalar source produces single-element slice", func(t *testing.T) {
		r, err := cast.ToE[map[string][]int](map[string]int{"a": 42})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(r["a"], []int{42}) {
			t.Errorf("expected [42], got %v", r["a"])
		}
	})
	t.Run("unconvertible scalar source errors", func(t *testing.T) {
		_, err := cast.ToE[map[string][]int](map[string]string{"a": "bad"})
		if err == nil {
			t.Fatal("expected error for bad scalar source, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// TestCastToTypeInterfaceNil exercises castToType's Interface branch with a nil
// source value, which returns reflect.Zero of the interface type.
func TestCastToTypeInterfaceNil(t *testing.T) {
	r, err := cast.ToE[map[string]fmt.Stringer](map[string]interface{}{"a": nil})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r["a"] != nil {
		t.Errorf("expected nil Stringer value, got %v", r["a"])
	}
}

// TestCastToTypeFuncUnsupportedType covers castToType's Func branch when the
// func type has NumIn != 0 or NumOut != 1 — line 94-96 in util.reflect.go.
// map[string]func() has a function value with NumOut==0, triggering the error.
func TestCastToTypeFuncUnsupportedType(t *testing.T) {
	_, err := cast.ToE[map[string]func()](map[string]string{"a": "b"})
	if err == nil {
		t.Fatal("expected error for unsupported func type, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

// TestCastToTypeFuncElementError covers castToType's Func branch when the
// element (return-type) cast fails — line 98-100 in util.reflect.go.
func TestCastToTypeFuncElementError(t *testing.T) {
	_, err := cast.ToE[map[string]cast.Func[int]](map[string]string{"a": "bad"})
	if err == nil {
		t.Fatal("expected error for uncastable Func element, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

// TestCastToTypeChanElementError covers castToType's Chan branch when the
// element cast fails — line 120-122 in util.reflect.go.
func TestCastToTypeChanElementError(t *testing.T) {
	_, err := cast.ToE[map[string]chan int](map[string]string{"a": "bad"})
	if err == nil {
		t.Fatal("expected error for uncastable Chan element, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

// TestCastToTypeConvertibleTo covers the reflect.ConvertibleTo path in
// castToType's default branch (lines 134-136 in util.reflect.go) when
// castToKind returns a base type (int) that must be converted to a named type
// (myInt) via reflect.Convert.
func TestCastToTypeConvertibleTo(t *testing.T) {
	r, err := cast.ToE[map[string]myInt](map[string]int{"a": 42})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r["a"] != myInt(42) {
		t.Errorf("expected myInt(42), got %v", r["a"])
	}
}

// TestToEDefaultElseBranch covers to.go lines 91-93: the else branch in the
// default case of ToE's kind switch, reached when TTo is a struct type that
// doesn't implement error, std_error.Error, or fmt.Stringer.
func TestToEDefaultElseBranch(t *testing.T) {
	type plain struct{ X int }
	_, err := cast.ToE[plain]("anything")
	if err == nil {
		t.Fatal("expected error for unsupported struct TTo, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

// TestToFloatDefaultSuccess covers to.float.go line 80: the happy-path return
// in toFloat's default case, triggered by a named float type (celsius) whose
// fmt.Sprintf representation is a parseable float string.
func TestToFloatDefaultSuccess(t *testing.T) {
	result, err := cast.ToE[float64](celsius(3.14))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result < 3.13 || result > 3.15 {
		t.Errorf("expected ~3.14, got %v", result)
	}
}

// TestToBoolDefaultError covers to.bool.go line 76: the final error return in
// toBool's default case, triggered by a struct type whose Sprintf representation
// cannot be parsed as a bool.
func TestToBoolDefaultError(t *testing.T) {
	type plain struct{ X int }
	_, err := cast.ToE[bool](plain{X: 1})
	if err == nil {
		t.Fatal("expected error for unparseable bool source, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

// TestToComplexDefaultError covers to.complex.go line 59: the final error
// return in toComplex, triggered by a struct type whose Sprintf representation
// cannot be parsed as a complex number.
func TestToComplexDefaultError(t *testing.T) {
	type plain struct{ X int }
	_, err := cast.ToE[complex64](plain{X: 1})
	if err == nil {
		t.Fatal("expected error for unparseable complex source, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

func TestNamedChanType(t *testing.T) {
	t.Run("positive: int → myChan", func(t *testing.T) {
		result, err := cast.ToE[myChan](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// result should be a myChan (not plain chan int)
		if reflect.TypeOf(result) != reflect.TypeOf(myChan(nil)) {
			t.Errorf("expected type myChan, got %T", result)
		}
		val := <-result
		if val != 42 {
			t.Errorf("expected 42, got %v", val)
		}
	})
	t.Run("negative: invalid source → error", func(t *testing.T) {
		_, err := cast.ToE[myChan]("bad")
		if err == nil {
			t.Error("expected error for invalid input, got nil")
		}
	})
}

// TestCastToStructTypeTextUnmarshaler exercises castToStructType's
// encoding.TextUnmarshaler path (util.reflect.go:222-228). The path is
// reachable only through castToType's reflect.Struct branch, not through the
// top-level toStruct dispatcher. A Container struct with a textUnmarshalStruct
// field forces castToType to be called for that field, which calls
// castToStructType, which tries the TextUnmarshaler before falling back to toStruct.
func TestCastToStructTypeTextUnmarshaler(t *testing.T) {
	type Container struct {
		Field textUnmarshalStruct
		Name  string
	}
	result, err := cast.ToStructE[Container](map[string]any{
		"Field": "parsed-value",
		"Name":  "test",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Field.Val != "parsed-value" {
		t.Errorf("expected Val=parsed-value, got %q", result.Field.Val)
	}
	if result.Name != "test" {
		t.Errorf("expected Name=test, got %q", result.Name)
	}
}

// TestSliceNonScalarElementPath exercises toSlice's default: case with a
// struct element type (isScalarKind=false), routing each element through
// castToType → castToStructType → TextUnmarshaler instead of castToKind.
func TestSliceNonScalarElementPath(t *testing.T) {
	t.Run("positive: []string → textUnmarshalStructSlice", func(t *testing.T) {
		src := []string{"foo", "bar", "baz"}
		result, err := cast.ToE[textUnmarshalStructSlice](src)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 3 {
			t.Fatalf("expected len 3, got %d", len(result))
		}
		for i, want := range []string{"foo", "bar", "baz"} {
			if result[i].Val != want {
				t.Errorf("result[%d].Val: expected %q, got %q", i, want, result[i].Val)
			}
		}
	})
}
