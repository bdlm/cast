package cast_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

func TestSliceToBoolSlice(t *testing.T) {
	testSliceCases[[]bool](t, []testCase{
		{in: []bool{true, false}, expect: []bool{true, false}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []bool{false, true}, err: nil, expectErr: false},
		{in: []bool{}, expect: []bool{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []bool{true, true, true, false, true}, err: nil, expectErr: false},
		{in: []bool{}, expect: []bool{}, err: nil, expectErr: false},
		// Scalars wrap as single-element slices.
		{in: 1, expect: []bool{true}, err: nil, expectErr: false},
		{in: "1", expect: []bool{true}, err: nil, expectErr: false},
		{in: -1, expect: []bool{true}, err: nil, expectErr: false},
		{in: "-1", expect: []bool{true}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []bool{true}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []bool{true}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []bool{true}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []bool{true}, err: nil, expectErr: false},
	})
}

func TestSliceToIntSlice(t *testing.T) {
	testSliceCases[[]int](t, []testCase{
		{in: []bool{true, false}, expect: []int{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		// Scalars wrap as single-element slices.
		{in: 1, expect: []int{1}, err: nil, expectErr: false},
		{in: "1", expect: []int{1}, err: nil, expectErr: false},
		{in: -1, expect: []int{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []int{-1}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []int{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]int8](t, []testCase{
		{in: []bool{true, false}, expect: []int8{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int8{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int8{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int8{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []int8{1}, err: nil, expectErr: false},
		{in: "1", expect: []int8{1}, err: nil, expectErr: false},
		{in: -1, expect: []int8{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []int8{-1}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int8{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int8{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]int16](t, []testCase{
		{in: []bool{true, false}, expect: []int16{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int16{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int16{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int16{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []int16{1}, err: nil, expectErr: false},
		{in: "1", expect: []int16{1}, err: nil, expectErr: false},
		{in: -1, expect: []int16{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []int16{-1}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int16{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int16{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]int32](t, []testCase{
		{in: []bool{true, false}, expect: []int32{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int32{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int32{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []int32{1}, err: nil, expectErr: false},
		// string → []int32 uses rune conversion: '1' = codepoint 49
		{in: "1", expect: []int32{49}, err: nil, expectErr: false},
		{in: -1, expect: []int32{-1}, err: nil, expectErr: false},
		// string → []int32 uses rune conversion: '-' = 45, '1' = 49
		{in: "-1", expect: []int32{45, 49}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int32{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int32{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]int64](t, []testCase{
		{in: []bool{true, false}, expect: []int64{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []int64{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []int64{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []int64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []int64{1}, err: nil, expectErr: false},
		{in: "1", expect: []int64{1}, err: nil, expectErr: false},
		{in: -1, expect: []int64{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []int64{-1}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []int64{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []int64{1}, err: nil, expectErr: false},
	})
}

func TestSliceToUintSlice(t *testing.T) {
	testSliceCases[[]uint](t, []testCase{
		{in: []bool{true, false}, expect: []uint{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint{1}, err: nil, expectErr: false},
		{in: "1", expect: []uint{1}, err: nil, expectErr: false},
		{in: -1, expect: []uint{}, err: nil, expectErr: true},  // negative signed → unsigned errors
		{in: "-1", expect: []uint{}, err: nil, expectErr: true}, // negative string → unsigned errors
		{in: []float32{1.1}, expect: []uint{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]uint8](t, []testCase{
		{in: []bool{true, false}, expect: []uint8{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint8{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint8{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint8{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint8{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint8{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint8{1}, err: nil, expectErr: false},
		// string → []uint8 uses byte conversion: '1' = 49
		{in: "1", expect: []uint8{49}, err: nil, expectErr: false},
		{in: -1, expect: []uint8{}, err: nil, expectErr: true},          // negative int → unsigned errors
		{in: "-1", expect: []uint8{45, 49}, err: nil, expectErr: false}, // string → bytes: '-'=45 '1'=49
		{in: []float32{1.1}, expect: []uint8{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint8{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint8{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint8{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]uint16](t, []testCase{
		{in: []bool{true, false}, expect: []uint16{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint16{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint16{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint16{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint16{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint16{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint16{1}, err: nil, expectErr: false},
		{in: "1", expect: []uint16{1}, err: nil, expectErr: false},
		{in: -1, expect: []uint16{}, err: nil, expectErr: true},
		{in: "-1", expect: []uint16{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint16{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint16{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint16{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint16{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]uint32](t, []testCase{
		{in: []bool{true, false}, expect: []uint32{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint32{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint32{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint32{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint32{1}, err: nil, expectErr: false},
		{in: "1", expect: []uint32{1}, err: nil, expectErr: false},
		{in: -1, expect: []uint32{}, err: nil, expectErr: true},
		{in: "-1", expect: []uint32{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint32{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint32{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint32{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint32{1}, err: nil, expectErr: false},
	})
	testSliceCases[[]uint64](t, []testCase{
		{in: []bool{true, false}, expect: []uint64{1, 0}, err: nil, expectErr: false},
		{in: []bool{false, true}, expect: []uint64{0, 1}, err: nil, expectErr: false},
		{in: []bool{}, expect: []uint64{}, err: nil, expectErr: false},
		{in: []int32{1, 30, 10, 0, 1}, expect: []uint64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []int32{-1, 30, 10, 0, 1}, expect: []uint64{}, err: nil, expectErr: true},
		{in: []uint32{1, 30, 10, 0, 1}, expect: []uint64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: 1, expect: []uint64{1}, err: nil, expectErr: false},
		{in: "1", expect: []uint64{1}, err: nil, expectErr: false},
		{in: -1, expect: []uint64{}, err: nil, expectErr: true},
		{in: "-1", expect: []uint64{}, err: nil, expectErr: true},
		{in: []float32{1.1}, expect: []uint64{1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []uint64{1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []uint64{1}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []uint64{1}, err: nil, expectErr: false},
	})
}

func TestSliceToFloatSlice(t *testing.T) {
	testSliceCases[[]float32](t, []testCase{
		{in: []float32{1, 30, 10, 0, 1}, expect: []float32{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []float32{-1}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []float32{-1}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []float32{-1.9}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []float32{-1.9}, err: nil, expectErr: false},
		{in: []float32{}, expect: []float32{}, err: nil, expectErr: false},
		// Scalars wrap as single-element slices.
		{in: 1, expect: []float32{1}, err: nil, expectErr: false},
		{in: "1", expect: []float32{1}, err: nil, expectErr: false},
		{in: -1, expect: []float32{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []float32{-1}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []float32{1.1}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []float32{1.1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []float32{1.9}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []float32{1.9}, err: nil, expectErr: false},
	})
	testSliceCases[[]float64](t, []testCase{
		{in: []float64{1, 30, 10, 0, 1}, expect: []float64{1, 30, 10, 0, 1}, err: nil, expectErr: false},
		// float32 source: widened to float64 via float64(float32(x)), not float64(x)
		{in: []float32{-1}, expect: []float64{-1}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []float64{-1}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []float64{float64(float32(-1.9))}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []float64{-1.9}, err: nil, expectErr: false},
		{in: []float64{}, expect: []float64{}, err: nil, expectErr: false},
		// Scalars wrap as single-element slices.
		{in: 1, expect: []float64{1}, err: nil, expectErr: false},
		{in: "1", expect: []float64{1}, err: nil, expectErr: false},
		{in: -1, expect: []float64{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []float64{-1}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []float64{float64(float32(1.1))}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []float64{1.1}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []float64{float64(float32(1.9))}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []float64{1.9}, err: nil, expectErr: false},
	})
}

func TestSliceToComplexSlice(t *testing.T) {
	testSliceCases[[]complex64](t, []testCase{
		{in: []float64{1, 30, 10, 0, 1}, expect: []complex64{complex64(1), complex64(30), complex64(10), complex64(0), complex64(1)}, err: nil, expectErr: false},
		{in: []float32{-1}, expect: []complex64{complex64(-1)}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []complex64{complex64(-1)}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []complex64{complex64(-1.9)}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []complex64{complex64(-1.9)}, err: nil, expectErr: false},
		{in: []complex64{}, expect: []complex64{}, err: nil, expectErr: false},
		// Scalars wrap as single-element slices.
		{in: 1, expect: []complex64{1}, err: nil, expectErr: false},
		{in: "1", expect: []complex64{1}, err: nil, expectErr: false},
		{in: -1, expect: []complex64{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []complex64{-1}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []complex64{complex64(1.1)}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []complex64{complex64(1.1)}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []complex64{complex64(1.9)}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []complex64{complex64(1.9)}, err: nil, expectErr: false},
	})
	testSliceCases[[]complex128](t, []testCase{
		{in: []float64{1, 30, 10, 0, 1}, expect: []complex128{complex128(1), complex128(30), complex128(10), complex128(0), complex128(1)}, err: nil, expectErr: false},
		// float32 source: real part is float64(float32(x)), not float64(x)
		{in: []float32{-1}, expect: []complex128{complex128(-1)}, err: nil, expectErr: false},
		{in: []string{"-1"}, expect: []complex128{complex128(-1)}, err: nil, expectErr: false},
		{in: []float32{-1.9}, expect: []complex128{complex(float64(float32(-1.9)), 0)}, err: nil, expectErr: false},
		{in: []string{"-1.9"}, expect: []complex128{complex128(-1.9)}, err: nil, expectErr: false},
		{in: []complex128{}, expect: []complex128{}, err: nil, expectErr: false},
		// Scalars wrap as single-element slices.
		{in: 1, expect: []complex128{1}, err: nil, expectErr: false},
		{in: "1", expect: []complex128{1}, err: nil, expectErr: false},
		{in: -1, expect: []complex128{-1}, err: nil, expectErr: false},
		{in: "-1", expect: []complex128{-1}, err: nil, expectErr: false},
		{in: []float32{1.1}, expect: []complex128{complex(float64(float32(1.1)), 0)}, err: nil, expectErr: false},
		{in: []string{"1.1"}, expect: []complex128{complex128(1.1)}, err: nil, expectErr: false},
		{in: []float32{1.9}, expect: []complex128{complex(float64(float32(1.9)), 0)}, err: nil, expectErr: false},
		{in: []string{"1.9"}, expect: []complex128{complex128(1.9)}, err: nil, expectErr: false},
	})
}

func TestUniqueValuesSlice(t *testing.T) {
	t.Run("int dedup", func(t *testing.T) {
		actual, err := cast.ToE[[]int]([]int{1, 2, 2, 3, 1}, cast.Op{Flag: cast.UNIQUE_VALUES, Val: true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expect := []int{1, 2, 3}
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("expected %v, got %v", expect, actual)
		}
	})
	t.Run("string dedup", func(t *testing.T) {
		actual, err := cast.ToE[[]string]([]string{"a", "b", "a", "c"}, cast.Op{Flag: cast.UNIQUE_VALUES, Val: true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expect := []string{"a", "b", "c"}
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("expected %v, got %v", expect, actual)
		}
	})
	t.Run("no duplicates unchanged", func(t *testing.T) {
		actual, err := cast.ToE[[]int]([]int{1, 2, 3}, cast.Op{Flag: cast.UNIQUE_VALUES, Val: true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expect := []int{1, 2, 3}
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("expected %v, got %v", expect, actual)
		}
	})
	t.Run("order preserved", func(t *testing.T) {
		actual, err := cast.ToE[[]int]([]int{3, 1, 2, 1, 3}, cast.Op{Flag: cast.UNIQUE_VALUES, Val: true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expect := []int{3, 1, 2}
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("expected %v, got %v", expect, actual)
		}
	})
}

func TestSliceUniqueValuesNonComparable(t *testing.T) {
	t.Run("[]any containing slices does not panic", func(t *testing.T) {
		in := []any{[]int{1, 2}, []int{3, 4}, []int{1, 2}}
		result, err := cast.ToE[[]any](in, cast.Op{cast.UNIQUE_VALUES, true})
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 unique elements, got %d: %v", len(result), result)
		}
	})
	t.Run("[]any containing maps does not panic", func(t *testing.T) {
		in := []any{map[string]int{"a": 1}, map[string]int{"b": 2}, map[string]int{"a": 1}}
		result, err := cast.ToE[[]any](in, cast.Op{cast.UNIQUE_VALUES, true})
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 unique elements, got %d: %v", len(result), result)
		}
	})
	t.Run("[]any mixing comparable and non-comparable", func(t *testing.T) {
		in := []any{1, []int{2, 3}, 1, []int{2, 3}, "x"}
		result, err := cast.ToE[[]any](in, cast.Op{cast.UNIQUE_VALUES, true})
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if len(result) != 3 {
			t.Errorf("expected 3 unique elements, got %d: %v", len(result), result)
		}
	})
	t.Run("[]any containing nil", func(t *testing.T) {
		in := []any{nil, 1, nil}
		result, err := cast.ToE[[]any](in, cast.Op{cast.UNIQUE_VALUES, true})
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if len(result) != 2 {
			t.Errorf("expected 2 unique elements, got %d: %v", len(result), result)
		}
	})
}

func TestSliceToStringSlice(t *testing.T) {
	testSliceCases[[]string](t, []testCase{
		{in: []int{1, 2, 3}, expect: []string{"1", "2", "3"}, err: nil, expectErr: false},
		{in: []bool{true, false}, expect: []string{"true", "false"}, err: nil, expectErr: false},
		{in: []float64{1.5, -1.5}, expect: []string{"1.5", "-1.5"}, err: nil, expectErr: false},
		{in: []string{"a", "b"}, expect: []string{"a", "b"}, err: nil, expectErr: false},
		{in: []string{}, expect: []string{}, err: nil, expectErr: false},
		// Scalars wrap as single-element slices.
		{in: 1, expect: []string{"1"}, err: nil, expectErr: false},
		{in: "hello", expect: []string{"hello"}, err: nil, expectErr: false},
	})
}

func TestSliceToAnySlice(t *testing.T) {
	testSliceCases[[]any](t, []testCase{
		{in: []int{1, 2, 3}, expect: []any{1, 2, 3}, err: nil, expectErr: false},
		{in: []string{"a", "b"}, expect: []any{"a", "b"}, err: nil, expectErr: false},
		{in: []bool{true, false}, expect: []any{true, false}, err: nil, expectErr: false},
		{in: []any{1, "a", true}, expect: []any{1, "a", true}, err: nil, expectErr: false},
		{in: []any{}, expect: []any{}, err: nil, expectErr: false},
		// Scalars wrap as single-element slices.
		{in: 1, expect: []any{1}, err: nil, expectErr: false},
	})
}

func TestSliceToUintptrSlice(t *testing.T) {
	testSliceCases[[]uintptr](t, []testCase{
		{in: []int{1, 2, 3}, expect: []uintptr{1, 2, 3}, err: nil, expectErr: false},
		{in: []string{"1", "2"}, expect: []uintptr{1, 2}, err: nil, expectErr: false},
		{in: []int{-1}, expect: []uintptr{}, err: nil, expectErr: true},
		// Scalars wrap as single-element slices.
		{in: 1, expect: []uintptr{1}, err: nil, expectErr: false},
	})
}

func TestSliceLengthErrors(t *testing.T) {
	t.Run("invalid LENGTH string errors", func(t *testing.T) {
		_, err := cast.ToE[[]int]([]string{"1", "2"}, cast.Op{cast.LENGTH, "bad"})
		if err == nil {
			t.Fatal("expected error for invalid LENGTH value, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("negative LENGTH errors", func(t *testing.T) {
		_, err := cast.ToE[[]int]([]string{"1", "2"}, cast.Op{cast.LENGTH, -1})
		if err == nil {
			t.Fatal("expected error for negative LENGTH, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

func TestSliceInvalidDefault(t *testing.T) {
	t.Run("string DEFAULT for []int", func(t *testing.T) {
		_, err := cast.ToE[[]int]([]string{"1", "2"}, cast.Op{cast.DEFAULT, "not a slice"})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("int DEFAULT for []string", func(t *testing.T) {
		_, err := cast.ToE[[]string]([]int{1, 2}, cast.Op{cast.DEFAULT, 42})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]string DEFAULT for []int", func(t *testing.T) {
		_, err := cast.ToE[[]int]([]string{"1", "2"}, cast.Op{cast.DEFAULT, []string{"fallback"}})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("compatible DEFAULT for []int does not error", func(t *testing.T) {
		_, err := cast.ToE[[]int]([]string{"1", "2"}, cast.Op{cast.DEFAULT, []int{-1}})
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})
}

// TestToSliceElementErrors covers the error-return inside toSlice's per-type
// loop for each typed slice case. Each subtest passes a slice with an
// unconvertible element so the element-cast fails mid-loop.
func TestToSliceElementErrors(t *testing.T) {
	bad := []string{"bad"}

	t.Run("[]bool element error", func(t *testing.T) {
		_, err := cast.ToE[[]bool](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]complex64 element error", func(t *testing.T) {
		_, err := cast.ToE[[]complex64](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]complex128 element error", func(t *testing.T) {
		_, err := cast.ToE[[]complex128](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]float32 element error", func(t *testing.T) {
		_, err := cast.ToE[[]float32](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]float64 element error", func(t *testing.T) {
		_, err := cast.ToE[[]float64](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]int element error", func(t *testing.T) {
		_, err := cast.ToE[[]int](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]int8 element error", func(t *testing.T) {
		_, err := cast.ToE[[]int8](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]int16 element error", func(t *testing.T) {
		_, err := cast.ToE[[]int16](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]int32 element error", func(t *testing.T) {
		_, err := cast.ToE[[]int32](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
	t.Run("[]int64 element error", func(t *testing.T) {
		_, err := cast.ToE[[]int64](bad)
		if err == nil || !errors.Is(err, cast.Error) {
			t.Fatalf("expected cast.Error, got %v", err)
		}
	})
}

// TestToSliceNamedTypeUniqueVals covers the dedupeSliceVal call inside toSlice's
// default/named-type branch.
func TestToSliceNamedTypeUniqueVals(t *testing.T) {
	result, err := cast.ToE[tags]([]string{"a", "b", "a", "c", "b"}, cast.Op{cast.UNIQUE_VALUES, true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, tags{"a", "b", "c"}) {
		t.Errorf("expected [a b c], got %v", result)
	}
}

// TestNilToSlice verifies that a nil source produces a single-element slice
// containing the zero value of the element type.
func TestNilToSlice(t *testing.T) {
	testSliceCases[[]int](t, []testCase{
		{in: nil, expect: []int{0}},
	})
	testSliceCases[[]string](t, []testCase{
		{in: nil, expect: []string{""}},
	})
	testSliceCases[[]bool](t, []testCase{
		{in: nil, expect: []bool{false}},
	})
}

// TestMapToSlice verifies that a map source produces a slice of the map's
// values (order is non-deterministic, so we check length and use a single-entry
// map where possible).
func TestMapToSlice(t *testing.T) {
	t.Run("map[string]int → []int single entry", func(t *testing.T) {
		result, err := cast.ToE[[]int](map[string]int{"a": 42})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 || result[0] != 42 {
			t.Errorf("expected [42], got %v", result)
		}
	})
	t.Run("map[string]string → []string", func(t *testing.T) {
		result, err := cast.ToE[[]string](map[string]string{"x": "hello"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 || result[0] != "hello" {
			t.Errorf("expected [hello], got %v", result)
		}
	})
	t.Run("map[string]string → []int conversion", func(t *testing.T) {
		result, err := cast.ToE[[]int](map[string]string{"n": "7"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 || result[0] != 7 {
			t.Errorf("expected [7], got %v", result)
		}
	})
	t.Run("map value unconvertible → error", func(t *testing.T) {
		_, err := cast.ToE[[]int](map[string]string{"n": "bad"})
		if err == nil {
			t.Error("expected error for unconvertible map value, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

// TestStructToSlice verifies that a plain struct source iterates exported field
// values into a slice, while named struct types (e.g. time.Time) use scalar wrap.
func TestStructToSlice(t *testing.T) {
	type Point struct{ X, Y int }
	t.Run("plain struct exported fields → []int", func(t *testing.T) {
		result, err := cast.ToE[[]int](Point{X: 3, Y: 7})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Fatalf("expected 2 elements, got %d: %v", len(result), result)
		}
		// Field order matches struct declaration order.
		if result[0] != 3 || result[1] != 7 {
			t.Errorf("expected [3 7], got %v", result)
		}
	})
	t.Run("plain struct field iteration → []string", func(t *testing.T) {
		// General structs use field iteration first. Each exported int field is
		// cast to string via toString, giving ["1", "2"].
		result, err := cast.ToE[[]string](Point{X: 1, Y: 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []string{"1", "2"}) {
			t.Errorf("expected [1 2], got %v", result)
		}
	})
	t.Run("time.Time scalar wrap → []string single element", func(t *testing.T) {
		// time.Time is a named struct type that has a scalar string representation
		// via fmt.Stringer; it should produce a 1-element []string, not iterate
		// its many unexported fields.
		ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		result, err := cast.ToE[[]string](ts)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d: %v", len(result), result)
		}
		if result[0] == "" {
			t.Errorf("expected non-empty time string, got empty")
		}
	})
	t.Run("time.Time scalar wrap → []int64 Unix seconds", func(t *testing.T) {
		epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		result, err := cast.ToE[[]int64](epoch)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 || result[0] != 0 {
			t.Errorf("expected [0], got %v", result)
		}
	})
}

// TestDecodeOptionSlice verifies the DECODE flag: when DECODE="JSON" the source
// string is parsed as a JSON array or object before element conversion, bypassing
// the default scalar-wrap path that would otherwise win for []string targets.
func TestDecodeOptionSlice(t *testing.T) {
	t.Run(`DECODE=JSON: ["a","b"] → []string`, func(t *testing.T) {
		result, err := cast.ToE[[]string](`["a","b"]`, cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []string{"a", "b"}) {
			t.Errorf(`expected ["a" "b"], got %v`, result)
		}
	})
	t.Run(`DECODE=json (lowercase): ["a","b"] → []string`, func(t *testing.T) {
		result, err := cast.ToE[[]string](`["a","b"]`, cast.Op{cast.DECODE, "json"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []string{"a", "b"}) {
			t.Errorf(`expected ["a" "b"], got %v`, result)
		}
	})
	t.Run(`DECODE=JSON: [1,2,3] → []int`, func(t *testing.T) {
		result, err := cast.ToE[[]int](`[1,2,3]`, cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []int{1, 2, 3}) {
			t.Errorf("expected [1 2 3], got %v", result)
		}
	})
	t.Run("DECODE=JSON with scalar → single-element slice", func(t *testing.T) {
		// "42" is valid JSON that decodes to float64(42), which is then scalar-wrapped.
		result, err := cast.ToE[[]int]("42", cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []int{42}) {
			t.Errorf("expected [42], got %v", result)
		}
	})
	t.Run("DECODE=JSON with invalid JSON → error", func(t *testing.T) {
		_, err := cast.ToE[[]int]("not json", cast.Op{cast.DECODE, "JSON"})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
	t.Run("DECODE=JSON from error source", func(t *testing.T) {
		type myErr struct{ s string }
		result, err := cast.ToE[[]string](fmt.Errorf(`["x","y"]`), cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []string{"x", "y"}) {
			t.Errorf("expected [x y], got %v", result)
		}
	})
	t.Run("DECODE=JSON from plain struct Stringer source", func(t *testing.T) {
		// testStringer is a plain user struct that implements fmt.Stringer.
		// It is NOT a named scalar type, so the normalization block converts
		// it to its string representation, and DECODE=JSON then applies.
		result, err := cast.ToE[[]int](testStringer{`[10,20]`}, cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []int{10, 20}) {
			t.Errorf("expected [10 20], got %v", result)
		}
	})
	t.Run("DECODE=JSON from named scalar struct (time.Time) is ignored", func(t *testing.T) {
		// time.Time implements fmt.Stringer, but it IS a named scalar type
		// registered in namedConverters. The normalization block skips it so
		// it reaches the struct dispatch → scalar wrap path unchanged.
		// DECODE=JSON has no effect: the result is a single-element slice
		// containing the converted time value, not a JSON-decoded array.
		ts := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		result, err := cast.ToE[[]int64](ts, cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// Unix epoch → 0 seconds, scalar-wrapped as [0].
		if !reflect.DeepEqual(result, []int64{0}) {
			t.Errorf("expected [0] (scalar wrap), got %v", result)
		}
	})
	t.Run("DECODE=JSON from plain struct without Stringer is ignored", func(t *testing.T) {
		// A plain struct that does not implement fmt.Stringer is never
		// matched by the normalization block, so DECODE has no effect.
		// The struct reaches the struct dispatch path where sliceFromStruct
		// iterates exported fields normally.
		type Point struct{ X, Y int }
		result, err := cast.ToE[[]int](Point{X: 3, Y: 7}, cast.Op{cast.DECODE, "JSON"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []int{3, 7}) {
			t.Errorf("expected [3 7] (field iteration), got %v", result)
		}
	})
	t.Run("without DECODE: string still scalar-wraps for []string", func(t *testing.T) {
		// Without DECODE, a JSON array string is scalar-wrapped for []string targets.
		result, err := cast.ToE[[]string](`["a","b"]`)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 || result[0] != `["a","b"]` {
			t.Errorf("expected scalar wrap [%q], got %v", `["a","b"]`, result)
		}
	})
}

// TestJSONStringToSlice verifies that a JSON array string is decoded as a last
// resort when scalar wrap fails for string sources.
func TestJSONStringToSlice(t *testing.T) {
	t.Run(`"[1,2,3]" → []int`, func(t *testing.T) {
		result, err := cast.ToE[[]int](`[1,2,3]`)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []int{1, 2, 3}) {
			t.Errorf("expected [1 2 3], got %v", result)
		}
	})
	t.Run(`["a","b"] string → []string scalar wrap (JSON path not needed)`, func(t *testing.T) {
		// For []string target, sliceFromSingle succeeds immediately because
		// a string IS a string — the JSON path is only reached when scalar wrap
		// fails. So the whole JSON string becomes a single element.
		result, err := cast.ToE[[]string](`["a","b"]`)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 1 || result[0] != `["a","b"]` {
			t.Errorf(`expected [["a","b"]] (scalar wrap), got %v`, result)
		}
	})
	t.Run("non-JSON string uses scalar wrap, not JSON", func(t *testing.T) {
		// "42" is not a JSON array/object so looksLikeCollection returns false;
		// the string is cast as a scalar → []int{42}.
		result, err := cast.ToE[[]int]("42")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, []int{42}) {
			t.Errorf("expected [42], got %v", result)
		}
	})
	t.Run("invalid JSON array string → error", func(t *testing.T) {
		_, err := cast.ToE[[]int](`[1,"bad",3]`)
		if err == nil {
			t.Error("expected error for unconvertible JSON element, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

func testSliceCases[TTo any](t *testing.T, cases []testCase) {
	var typ TTo
	name := fmt.Sprintf("%T", typ)

	for k, test := range cases {
		t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
			actual, err := cast.ToE[TTo](test.in)
			testInfo := fmt.Sprintf(`
case #%d: ToE[%s]
input: %v (%T)
expect error: %v; actual error: % #+v
expected result: %v (%T); actual result: %v (%T)
test: %#v
			`,
				k,
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
			if nil != err && !test.expectErr {
				t.Error("1. expected nil, got error", testInfo)
			} else if nil == err && test.expectErr {
				t.Error("2. expected error, got nil", testInfo)
			} else if nil != err && !errors.Is(err, cast.Error) {
				t.Error("3. expected cast.Error, got different error type", testInfo)
			} else if nil == err && !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("4. expected %v (%T) to equal %v (%T)\n%s", test.expect, test.expect, actual, actual, testInfo)
			}
		})
	}
}
