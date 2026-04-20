package cast_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestToChanBool(t *testing.T) {
	testChanCases[bool](t, []testCase{
		{in: true, expect: true, err: nil, expectErr: false},
		{in: 1, expect: true, err: nil, expectErr: false},
		{in: 0, expect: false, err: nil, expectErr: false},
		{in: "hi", expect: false, err: nil, expectErr: true},
		{in: float64(1.1), expect: true, err: nil, expectErr: false},
		{in: float64(-1.1), expect: true, err: nil, expectErr: false},
	})
}

func TestToChanByte(t *testing.T) {
	testChanCases[byte](t, []testCase{
		{in: "a", expect: byte(0), err: nil, expectErr: true},
		{in: byte(1), expect: byte(1), err: nil, expectErr: false},
		{in: byte(0), expect: byte(0), err: nil, expectErr: false},
		{in: "hi", expect: byte(0), err: nil, expectErr: true},
		{in: float64(1.1), expect: byte(1), err: nil, expectErr: false},
		{in: float64(-1.1), expect: byte(0), err: nil, expectErr: true},
	})
}

func TestToChanInt(t *testing.T) {
	testChanCases[int](t, []testCase{
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
	testChanCases[int8](t, []testCase{
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
	testChanCases[int16](t, []testCase{
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
	testChanCases[int32](t, []testCase{
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
	testChanCases[int64](t, []testCase{
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

func TestToChanUint(t *testing.T) {
	testChanCases[uint](t, []testCase{
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
	testChanCases[uint8](t, []testCase{
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
	testChanCases[uint16](t, []testCase{
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
	testChanCases[uint32](t, []testCase{
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
	testChanCases[uint64](t, []testCase{
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
	testChanCases[uintptr](t, []testCase{
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

func TestToChanFloat(t *testing.T) {
	testChanCases[float32](t, []testCase{
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
	testChanCases[float64](t, []testCase{
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

func TestToChanComplex(t *testing.T) {
	testChanCases[complex64](t, []testCase{
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
	testChanCases[complex128](t, []testCase{
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

func TestToChanString(t *testing.T) {
	testChanCases[string](t, []testCase{
		{in: "hello", expect: "hello", err: nil, expectErr: false},
		{in: 42, expect: "42", err: nil, expectErr: false},
		{in: true, expect: "true", err: nil, expectErr: false},
		{in: false, expect: "false", err: nil, expectErr: false},
		{in: float64(1.5), expect: "1.5", err: nil, expectErr: false},
	})
}

func TestToChanSlice(t *testing.T) {
	t.Run("chan []int from []string source", func(t *testing.T) {
		ch, err := cast.ToE[chan []int]([]string{"1", "2", "3"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		result := <-ch
		if !reflect.DeepEqual(result, []int{1, 2, 3}) {
			t.Errorf("expected [1 2 3], got %v", result)
		}
	})
	t.Run("chan []string from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []string]([]int{1, 2, 3})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		result := <-ch
		if !reflect.DeepEqual(result, []string{"1", "2", "3"}) {
			t.Errorf(`expected ["1" "2" "3"], got %v`, result)
		}
	})
	t.Run("scalar source errors", func(t *testing.T) {
		_, err := cast.ToE[chan []int]("not a slice")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

func TestToChanNested(t *testing.T) {
	t.Run("chan chan int from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan int](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		inner := <-outer
		val := <-inner
		if val != 42 {
			t.Errorf("expected 42, got %v", val)
		}
	})
	t.Run("chan chan string from string source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan string]("hello")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		inner := <-outer
		val := <-inner
		if val != "hello" {
			t.Errorf("expected \"hello\", got %v", val)
		}
	})
	t.Run("invalid source errors", func(t *testing.T) {
		_, err := cast.ToE[chan chan int]("bad")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

func TestChanInvalidDefault(t *testing.T) {
	t.Run("string DEFAULT for chan int", func(t *testing.T) {
		_, err := cast.ToE[chan int](42, cast.Op{cast.DEFAULT, "not a chan"})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("int DEFAULT for chan string", func(t *testing.T) {
		_, err := cast.ToE[chan string]("hello", cast.Op{cast.DEFAULT, 42})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("chan string DEFAULT for chan int", func(t *testing.T) {
		_, err := cast.ToE[chan int](42, cast.Op{cast.DEFAULT, make(chan string, 1)})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
	t.Run("compatible DEFAULT for chan int does not error", func(t *testing.T) {
		_, err := cast.ToE[chan int](42, cast.Op{cast.DEFAULT, make(chan int, 1)})
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})
}

// TestToChanDefaultArm covers toChan's top-level default: case (unsupported element kind).
func TestToChanDefaultArm(t *testing.T) {
	_, err := cast.ToE[chan struct{}](42)
	if err == nil {
		t.Fatal("expected error for chan struct{} (unsupported kind), got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %v", err)
	}
}

// TestToChanAny exercises the reflect.Interface arm in toChan's element-kind switch.
func TestToChanAny(t *testing.T) {
	t.Run("chan any from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan any](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		val := <-ch
		if val != 42 {
			t.Errorf("expected 42, got %v", val)
		}
	})
	t.Run("chan any from string source", func(t *testing.T) {
		ch, err := cast.ToE[chan any]("hello")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		val := <-ch
		if val != "hello" {
			t.Errorf("expected \"hello\", got %v", val)
		}
	})
}

// TestToChanSliceVariants covers additional arms of toChan's chan []T section.
func TestToChanSliceVariants(t *testing.T) {
	t.Run("chan []any from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []any]([]int{1, 2, 3})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(<-ch) != 3 {
			t.Error("expected 3-element slice")
		}
	})
	t.Run("chan []bool from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []bool]([]int{1, 0, 1})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []bool{true, false, true}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []float32 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []float32]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []float32{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []float64 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []float64]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []float64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []complex64 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []complex64]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []complex64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []complex128 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []complex128]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []complex128{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []int8 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []int8]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []int8{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []int16 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []int16]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []int16{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []int32 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []int32]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []int32{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []int64 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []int64]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []int64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []uint from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []uint]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []uint{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []uint8 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []uint8]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []uint8{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []uint16 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []uint16]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []uint16{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []uint32 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []uint32]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []uint32{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []uint64 from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []uint64]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []uint64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan []uintptr from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan []uintptr]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(<-ch, []uintptr{1, 2}) {
			t.Error("unexpected value")
		}
	})
}

// TestToChanFuncVariants covers additional arms of toChan's chan Func[T] section.
func TestToChanFuncVariants(t *testing.T) {
	t.Run("chan Func[any] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[any]](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if fn() != 42 {
			t.Errorf("expected 42, got %v", fn())
		}
	})
	t.Run("chan Func[float32] from string source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[float32]]("1.5")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if fn() != float32(1.5) {
			t.Errorf("expected 1.5, got %v", fn())
		}
	})
	t.Run("chan Func[float64] from string source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[float64]]("3.14")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if fn() != 3.14 {
			t.Errorf("expected 3.14, got %v", fn())
		}
	})
	t.Run("chan Func[complex64] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[complex64]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if fn() != complex64(5) {
			t.Errorf("expected 5, got %v", fn())
		}
	})
	t.Run("chan Func[complex128] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[complex128]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if fn() != complex128(5) {
			t.Errorf("expected 5, got %v", fn())
		}
	})
	t.Run("chan Func[int8] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[int8]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != int8(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[int16] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[int16]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != int16(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[int32] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[int32]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != int32(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[int64] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[int64]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != int64(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[uint] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[uint]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != uint(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[uint8] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[uint8]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != uint8(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[uint16] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[uint16]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != uint16(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[uint32] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[uint32]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != uint32(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[uint64] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[uint64]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != uint64(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[uintptr] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[uintptr]](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if (<-ch)() != uintptr(5) {
			t.Error("unexpected value")
		}
	})
}

// TestToChanNestedVariants covers additional arms of toChan's chan chan T section.
func TestToChanNestedVariants(t *testing.T) {
	t.Run("chan chan bool from bool source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan bool](true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); !val {
			t.Error("expected true")
		}
	})
	t.Run("chan chan float32 from float source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan float32](float32(1.5))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != float32(1.5) {
			t.Errorf("expected 1.5, got %v", val)
		}
	})
	t.Run("chan chan float64 from float source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan float64](3.14)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != 3.14 {
			t.Errorf("expected 3.14, got %v", val)
		}
	})
	t.Run("chan chan complex64 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan complex64](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != complex64(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan complex128 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan complex128](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != complex128(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan int8 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan int8](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != int8(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan int16 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan int16](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != int16(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan int32 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan int32](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != int32(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan int64 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan int64](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != int64(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan uint from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan uint](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != uint(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan uint8 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan uint8](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != uint8(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan uint16 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan uint16](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != uint16(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan uint32 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan uint32](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != uint32(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan uint64 from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan uint64](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != uint64(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan uintptr from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan uintptr](5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != uintptr(5) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan chan any from int source", func(t *testing.T) {
		outer, err := cast.ToE[chan chan any](99)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val := <-(<-outer); val != 99 {
			t.Errorf("expected 99, got %v", val)
		}
	})
}

// TestToChanFuncSliceVariants covers the chan Func[[]T] sub-section inside toChan's
// Func element arm (slice return types for Func elements).
func TestToChanFuncSliceVariants(t *testing.T) {
	t.Run("chan Func[[]int] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]int]]([]int{1, 2, 3})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []int{1, 2, 3}) {
			t.Errorf("expected [1 2 3], got %v", fn())
		}
	})
	t.Run("chan Func[[]string] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]string]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []string{"1", "2"}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]bool] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]bool]]([]int{1, 0})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []bool{true, false}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]float32] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]float32]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []float32{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]float64] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]float64]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []float64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]complex64] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]complex64]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []complex64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]complex128] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]complex128]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []complex128{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]int8] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]int8]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []int8{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]int16] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]int16]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []int16{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]int32] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]int32]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []int32{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]int64] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]int64]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []int64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]uint] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]uint]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []uint{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]uint8] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]uint8]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []uint8{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]uint16] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]uint16]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []uint16{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]uint32] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]uint32]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []uint32{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]uint64] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]uint64]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []uint64{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]uintptr] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]uintptr]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !reflect.DeepEqual(fn(), []uintptr{1, 2}) {
			t.Error("unexpected value")
		}
	})
	t.Run("chan Func[[]any] from []int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[[]any]]([]int{1, 2})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if len(fn()) != 2 {
			t.Error("expected len 2")
		}
	})
	t.Run("invalid source errors", func(t *testing.T) {
		_, err := cast.ToE[chan cast.Func[[]int]]("bad")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

// TestToChanFuncElement exercises the reflect.Func arm in toChan's element-kind
// switch, creating channels whose element type is a Func[T].
func TestToChanFuncElement(t *testing.T) {
	t.Run("chan Func[int] from int source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[int]](42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if fn() != 42 {
			t.Errorf("expected 42, got %v", fn())
		}
	})
	t.Run("chan Func[string] from string source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[string]]("hello")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if fn() != "hello" {
			t.Errorf("expected \"hello\", got %v", fn())
		}
	})
	t.Run("chan Func[bool] from bool source", func(t *testing.T) {
		ch, err := cast.ToE[chan cast.Func[bool]](true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		fn := <-ch
		if !fn() {
			t.Errorf("expected true, got false")
		}
	})
	t.Run("invalid source errors", func(t *testing.T) {
		_, err := cast.ToE[chan cast.Func[int]]("bad")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %v", err)
		}
	})
}

func testChanCases[TTo any](t *testing.T, cases []testCase) {
	var typ TTo
	name := fmt.Sprintf("%T", typ)

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s: %v", name, test.in), func(t *testing.T) {
			var result TTo
			actual, err := cast.ToE[chan TTo](test.in)
			if nil == err {
				result = <-actual
			}
			testInfo := fmt.Sprintf(`
case: ToE[chan %s]
input: %v (%T)
expect error: %v; actual error: %v
expected result: %v (%T); actual result: %v (%T)
result chan: %#v, (%T)
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
			// fmt.Println(testInfo)
			// fmt.Printf("\n-----------------\n%v (%T): %v (%T)\n-----------------\n", result, result, test.expect, test.expect)
			if err != nil && !test.expectErr {
				t.Error("1. expected nil, got error", testInfo)
			} else if err == nil && test.expectErr {
				t.Error("2. expected error, got nil", testInfo)
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Error("3. expected cast.Error, got different error type", testInfo)
			} else if nil == err && !reflect.DeepEqual(result, test.expect.(TTo)) {
				t.Errorf("4. expected %v to equal %v %s", test.expect, actual, testInfo)
			}
		})
	}
}
