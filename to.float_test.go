package cast_test

import (
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestComplexToFloat(t *testing.T) {
	var tests = []struct {
		in       any
		expect32 float32
		err32    bool
		expect64 float64
		err64    bool
	}{
		{complex64(1), 0, true, 0, true},
		{complex64(-1), 0, true, 0, true},
		{complex128(1), 0, true, 0, true},
		{complex128(-1), 0, true, 0, true},
	}
	// floats
	for _, test := range tests {
		var success32, success64 bool
		got32, err := cast.ToE[float32](test.in)
		if (nil == err && false == test.err32) || (nil != err && true == test.err32) {
			if got32 == test.expect32 && reflect.TypeOf(got32).Kind() == reflect.Float32 {
				success32 = true
			}
		}
		if !success32 {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect32, test.expect32, got32, got32, err)
		}

		got64, err := cast.ToE[float64](test.in)
		if (nil == err && false == test.err64) || (nil != err && true == test.err64) {
			if got64 == test.expect64 && reflect.TypeOf(got64).Kind() == reflect.Float64 {
				success64 = true
			}
		}
		if !success64 {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect64, test.expect64, got64, got64, err)
		}
	}
}

func TestArrayToFloat(t *testing.T) {
	var tests = []struct {
		in       any
		expect32 float32
		err32    bool
		expect64 float64
		err64    bool
	}{
		{[]int{1, 2}, 0, true, 0, true},
		{[]string{"1", "2"}, 0, true, 0, true},
	}
	// int
	for _, test := range tests {
		var success32, success64 bool
		got32, err := cast.ToE[float32](test.in)
		if (nil == err && false == test.err32) || (nil != err && true == test.err32) {
			if got32 == test.expect32 && reflect.TypeOf(got32).Kind() == reflect.Float32 {
				success32 = true
			}
		}
		if !success32 {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect32, test.expect32, got32, got32, err)
		}

		got64, err := cast.ToE[float64](test.in)
		if (nil == err && false == test.err64) || (nil != err && true == test.err64) {
			if got64 == test.expect64 && reflect.TypeOf(got64).Kind() == reflect.Float64 {
				success64 = true
			}
		}
		if !success64 {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect64, test.expect64, got64, got64, err)
		}
	}
}

func TestBoolToFloat(t *testing.T) {
	var tests = []struct {
		in       any
		expect32 float32
		err32    bool
		expect64 float64
		err64    bool
	}{
		{true, 1, false, 1, false},
		{false, 0, false, 0, false},
	}
	// int
	for _, test := range tests {
		var success32, success64 bool
		got32, err := cast.ToE[float32](test.in)
		if (nil == err && false == test.err32) || (nil != err && true == test.err32) {
			if got32 == test.expect32 && reflect.TypeOf(got32).Kind() == reflect.Float32 {
				success32 = true
			}
		}
		if !success32 {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect32, test.expect32, got32, got32, err)
		}

		got64, err := cast.ToE[float64](test.in)
		if (nil == err && false == test.err64) || (nil != err && true == test.err64) {
			if got64 == test.expect64 && reflect.TypeOf(got64).Kind() == reflect.Float64 {
				success64 = true
			}
		}
		if !success64 {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect64, test.expect64, got64, got64, err)
		}
	}
}

func TestStrToFloat(t *testing.T) {
	var tests = []struct {
		in       any
		expect32 float32
		err32    bool
		expect64 float64
		err64    bool
	}{
		{"1", 1, false, 1, false},
		{"0", 0, false, 0, false},
		{"-1", -1, false, -1, false},
		{"1.1", 1.1, false, 1.1, false},
		{"1.4", 1.4, false, 1.4, false},
		{"1.5", 1.5, false, 1.5, false},
		{"1.9", 1.9, false, 1.9, false},
		{"0.0", 0.0, false, 0.0, false},
		{"0.1", 0.1, false, 0.1, false},
		{"-1.1", -1.1, false, -1.1, false},
		{"-1.4", -1.4, false, -1.4, false},
		{"-1.1", -1.1, false, -1.1, false},
		{"-1.1", -1.1, false, -1.1, false},
		{"1,000", 1000, false, 1000, false},
		{"1,000,000", 1000000, false, 1000000, false},
		{"Hi", 0, true, 0, true},
	}
	// int
	for _, test := range tests {
		var success32, success64 bool
		got32, err := cast.ToE[float32](test.in)
		if (nil == err && false == test.err32) || (nil != err && true == test.err32) {
			if got32 == test.expect32 && reflect.TypeOf(got32).Kind() == reflect.Float32 {
				success32 = true
			}
		}
		if !success32 {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect32, test.expect32, got32, got32, err)
		}

		got64, err := cast.ToE[float64](test.in)
		if (nil == err && false == test.err64) || (nil != err && true == test.err64) {
			if got64 == test.expect64 && reflect.TypeOf(got64).Kind() == reflect.Float64 {
				success64 = true
			}
		}
		if !success64 {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect64, test.expect64, got64, got64, err)
		}
	}
}
