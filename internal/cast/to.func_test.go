package cast_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestToFuncBool(t *testing.T) {
	testFuncCases[bool](t, []testCase{
		{in: true, expect: true, err: nil, expectErr: false},
		{in: 1, expect: true, err: nil, expectErr: false},
		{in: 0, expect: false, err: nil, expectErr: false},
		{in: "hi", expect: false, err: nil, expectErr: true},
		{in: float64(1.1), expect: true, err: nil, expectErr: false},
		{in: float64(-1.1), expect: true, err: nil, expectErr: false},
	})
}

func TestToFuncByte(t *testing.T) {
	testFuncCases[byte](t, []testCase{
		{in: "a", expect: byte(0), err: nil, expectErr: true},
		{in: byte(1), expect: byte(1), err: nil, expectErr: false},
		{in: byte(0), expect: byte(0), err: nil, expectErr: false},
		{in: "hi", expect: byte(0), err: nil, expectErr: true},
		{in: float64(1.1), expect: byte(1), err: nil, expectErr: false},
		{in: float64(-1.1), expect: byte(0), err: nil, expectErr: true},
	})
}

func TestToFuncInt(t *testing.T) {
	testFuncCases[int](t, []testCase{
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
		{in: "Hi!", expect: int(0), err: nil, expectErr: true},
	})
	testFuncCases[int8](t, []testCase{
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
		{in: "Hi!", expect: int8(0), err: nil, expectErr: true},
	})
	testFuncCases[int16](t, []testCase{
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
		{in: "Hi!", expect: int16(0), err: nil, expectErr: true},
	})
	testFuncCases[int32](t, []testCase{
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
		{in: "Hi!", expect: int32(0), err: nil, expectErr: true},
	})
	testFuncCases[int64](t, []testCase{
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
		{in: "Hi!", expect: int64(0), err: nil, expectErr: true},
	})
}

func TestToFuncUint(t *testing.T) {
	testFuncCases[uint](t, []testCase{
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
		{in: "Hi!", expect: uint(0), err: nil, expectErr: true},
	})
	testFuncCases[uint8](t, []testCase{
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
		{in: "Hi!", expect: uint8(0), err: nil, expectErr: true},
	})
	testFuncCases[uint16](t, []testCase{
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
		{in: "Hi!", expect: uint16(0), err: nil, expectErr: true},
	})
	testFuncCases[uint32](t, []testCase{
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
		{in: "Hi!", expect: uint32(0), err: nil, expectErr: true},
	})
	testFuncCases[uint64](t, []testCase{
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
		{in: "Hi!", expect: uint64(0), err: nil, expectErr: true},
	})
	testFuncCases[uintptr](t, []testCase{
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
	})
}

func TestToFuncFloat(t *testing.T) {
	testFuncCases[float32](t, []testCase{
		{in: 1, expect: float32(1), err: nil, expectErr: false},
		{in: "1", expect: float32(1), err: nil, expectErr: false},
		{in: 1.1, expect: float32(1.1), err: nil, expectErr: false},
		{in: "1.1", expect: float32(1.1), err: nil, expectErr: false},
		{in: 1.9, expect: float32(1.9), err: nil, expectErr: false},
		{in: "1.9", expect: float32(1.9), err: nil, expectErr: false},
		{in: -1, expect: float32(-1), err: nil, expectErr: false},
		{in: "-1", expect: float32(-1), err: nil, expectErr: false},
		{in: -1.9, expect: float32(-1.9), err: nil, expectErr: false},
		{in: "-1.9", expect: float32(-1.9), err: nil, expectErr: false},
		{in: "Hi!", expect: float32(0), err: nil, expectErr: true},
	})
	testFuncCases[float64](t, []testCase{
		{in: 1, expect: float64(1), err: nil, expectErr: false},
		{in: "1", expect: float64(1), err: nil, expectErr: false},
		{in: 1.1, expect: float64(1.1), err: nil, expectErr: false},
		{in: "1.1", expect: float64(1.1), err: nil, expectErr: false},
		{in: 1.9, expect: float64(1.9), err: nil, expectErr: false},
		{in: "1.9", expect: float64(1.9), err: nil, expectErr: false},
		{in: -1, expect: float64(-1), err: nil, expectErr: false},
		{in: "-1", expect: float64(-1), err: nil, expectErr: false},
		{in: -1.9, expect: float64(-1.9), err: nil, expectErr: false},
		{in: "-1.9", expect: float64(-1.9), err: nil, expectErr: false},
		{in: "Hi!", expect: float64(0), err: nil, expectErr: true},
	})
}

func TestToFuncComplex(t *testing.T) {
	testFuncCases[complex64](t, []testCase{
		{in: 1, expect: complex64(1), err: nil, expectErr: false},
		{in: "1", expect: complex64(1), err: nil, expectErr: false},
		{in: 1.1, expect: complex64(1.1), err: nil, expectErr: false},
		{in: "1.1", expect: complex64(1.1), err: nil, expectErr: false},
		{in: 1.9, expect: complex64(1.9), err: nil, expectErr: false},
		{in: "1.9", expect: complex64(1.9), err: nil, expectErr: false},
		{in: -1, expect: complex64(-1), err: nil, expectErr: false},
		{in: "-1", expect: complex64(-1), err: nil, expectErr: false},
		{in: -1.9, expect: complex64(-1.9), err: nil, expectErr: false},
		{in: "-1.9", expect: complex64(-1.9), err: nil, expectErr: false},
		{in: "Hi!", expect: complex64(0), err: nil, expectErr: true},
	})
	testFuncCases[complex128](t, []testCase{
		{in: 1, expect: complex128(1), err: nil, expectErr: false},
		{in: "1", expect: complex128(1), err: nil, expectErr: false},
		{in: 1.1, expect: complex128(1.1), err: nil, expectErr: false},
		{in: "1.1", expect: complex128(1.1), err: nil, expectErr: false},
		{in: 1.9, expect: complex128(1.9), err: nil, expectErr: false},
		{in: "1.9", expect: complex128(1.9), err: nil, expectErr: false},
		{in: -1, expect: complex128(-1), err: nil, expectErr: false},
		{in: "-1", expect: complex128(-1), err: nil, expectErr: false},
		{in: -1.9, expect: complex128(-1.9), err: nil, expectErr: false},
		{in: "-1.9", expect: complex128(-1.9), err: nil, expectErr: false},
		{in: "Hi!", expect: complex128(0), err: nil, expectErr: true},
	})
}

func TestToFuncString(t *testing.T) {
	testFuncCases[string](t, []testCase{
		{in: "hello", expect: "hello", err: nil, expectErr: false},
		{in: 42, expect: "42", err: nil, expectErr: false},
		{in: true, expect: "true", err: nil, expectErr: false},
		{in: false, expect: "false", err: nil, expectErr: false},
		{in: float64(1.5), expect: "1.5", err: nil, expectErr: false},
	})
}

func TestToFuncSlice(t *testing.T) {
	t.Run("Func[[]int] from []string source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]int]]([]string{"1", "2", "3"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		result := fn()
		expect := []int{1, 2, 3}
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("expected %v, got %v", expect, result)
		}
	})
	t.Run("Func[[]string] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]string]]([]int{1, 2, 3})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		result := fn()
		expect := []string{"1", "2", "3"}
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("expected %v, got %v", expect, result)
		}
	})
	t.Run("Func[[]int] from scalar source errors", func(t *testing.T) {
		_, err := cast.ToE[cast.Func[[]int]]("not a slice")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

func TestToFuncChan(t *testing.T) {
	t.Run("Func[chan int] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan int]](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		ch := fn()
		val := <-ch
		if val != 42 {
			t.Errorf("expected 42, got %v", val)
		}
	})
	t.Run("Func[chan string] from string source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan string]]("hello")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		ch := fn()
		val := <-ch
		if val != "hello" {
			t.Errorf("expected \"hello\", got %v", val)
		}
	})
	t.Run("invalid source errors", func(t *testing.T) {
		_, err := cast.ToE[cast.Func[chan int]]("bad")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
		}
	})
}

// TestToFuncSliceVariants covers additional arms of toFunc's Func[[]T] section.
func TestToFuncSliceVariants(t *testing.T) {
	t.Run("Func[[]bool] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]bool]]([]int{1, 0, 1})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []bool{true, false, true}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]float32] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]float32]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []float32{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]float64] from []string source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]float64]]([]string{"1.5", "2.5"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []float64{1.5, 2.5}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]complex64] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]complex64]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []complex64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]complex128] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]complex128]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []complex128{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]int8] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]int8]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []int8{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]int16] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]int16]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []int16{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]int32] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]int32]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []int32{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]int64] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]int64]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []int64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]uint] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]uint]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []uint{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]uint8] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]uint8]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []uint8{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]uint16] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]uint16]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []uint16{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]uint32] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]uint32]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []uint32{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]uint64] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]uint64]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []uint64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]uintptr] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]uintptr]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(fn(), []uintptr{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[[]any] from []int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[[]any]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(fn()) != 2 {
			t.Error("expected len 2")
		}
	})
}

// TestToFuncChanVariants covers additional arms of toFunc's Func[chan T] section.
func TestToFuncChanVariants(t *testing.T) {
	t.Run("Func[chan any] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan any]](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != 42 {
			t.Errorf("expected 42, got %v", val)
		}
	})
	t.Run("Func[chan bool] from bool source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan bool]](true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !(<-fn()) {
			t.Error("expected true")
		}
	})
	t.Run("Func[chan float32] from string source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan float32]]("1.5")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != float32(1.5) {
			t.Errorf("expected 1.5, got %v", val)
		}
	})
	t.Run("Func[chan float64] from string source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan float64]]("3.14")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != 3.14 {
			t.Errorf("expected 3.14, got %v", val)
		}
	})
	t.Run("Func[chan complex64] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan complex64]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != complex64(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan complex128] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan complex128]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != complex128(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan int8] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan int8]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != int8(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan int16] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan int16]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != int16(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan int32] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan int32]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != int32(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan int64] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan int64]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != int64(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan uint] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan uint]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != uint(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan uint8] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan uint8]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != uint8(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan uint16] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan uint16]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != uint16(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan uint32] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan uint32]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != uint32(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan uint64] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan uint64]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != uint64(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("Func[chan uintptr] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[chan uintptr]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-fn(); val != uintptr(5) {
			t.Error("unexpected value")
		}
	})
}

// TestToFuncAny exercises the reflect.Interface arm in toFunc's Out(0) kind switch.
func TestToFuncAny(t *testing.T) {
	t.Run("Func[any] from int source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[any]](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		val := fn()
		if val != 42 {
			t.Errorf("expected 42, got %v", val)
		}
	})
	t.Run("Func[any] from string source", func(t *testing.T) {
		fn, err := cast.ToE[cast.Func[any]]("hello")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		val := fn()
		if val != "hello" {
			t.Errorf("expected \"hello\", got %v", val)
		}
	})
}

// TestToFuncDefaultValidType covers the ops.Delete(DEFAULT) line in toFunc
// (the happy path when a valid DEFAULT is provided): the type assertion
// succeeds, DEFAULT is stripped, and the cast proceeds normally.
func TestToFuncDefaultValidType(t *testing.T) {
	defaultFn := cast.Func[int](func() int { return -1 })
	fn, err := cast.ToE[cast.Func[int]](42, cast.Op{cast.DEFAULT, defaultFn})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if fn() != 42 {
		t.Errorf("expected 42, got %v", fn())
	}
}

// TestToFuncDefaultWrongType exercises toFunc's DEFAULT wrong-type error path.
func TestToFuncDefaultWrongType(t *testing.T) {
	_, err := cast.ToE[cast.Func[int]](42, cast.Op{cast.DEFAULT, "not a func"})
	if err == nil {
		t.Fatal("expected error for wrong DEFAULT type, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %v", err)
	}
}

func testFuncCases[TTo any](t *testing.T, cases []testCase) {
	var typ TTo
	name := fmt.Sprintf("%T", typ)

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
			var result TTo
			actual, err := cast.ToE[cast.Func[TTo]](test.in)
			if nil == err {
				result = actual()
			}
			testInfo := fmt.Sprintf(`
case: ToE[Func[%s]]
input: %v (%T)
expect error: %v; actual error: %v
expected result: %v (%T); actual result: %v (%T)
result fn: %#v, (%T)
test: %#v
			`,
				name,
				test.in,
				test.in,
				test.expectErr,
				err,
				test.expect,
				test.expect,
				result,
				result,
				actual,
				actual,
				test,
			)
			//fmt.Println(testInfo)
			if !test.expectErr && err != nil {
				t.Error(1, testInfo)
			} else if test.expectErr && err == nil {
				t.Error(2, testInfo)
			} else if err != nil && !errors.Is(err, cast.ErrorUnableToCast) {
				t.Error(3, testInfo)
			} else if nil == err && !reflect.DeepEqual(result, test.expect.(TTo)) {
				t.Error(4, testInfo)
			}
		})
	}
}
