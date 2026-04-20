package cast_test

import (
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
	t.Run("negative: scalar source → error", func(t *testing.T) {
		_, err := cast.ToE[tags]("not a slice")
		if err == nil {
			t.Error("expected error for scalar source, got nil")
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
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
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
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
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
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
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
