package cast_test

import (
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestComplexToInt(t *testing.T) {
	var tests = []struct {
		in     any
		expect int
		err    bool
	}{
		{complex64(1), 0, true},
		{complex64(-1), 0, true},
		{complex128(1), 0, true},
		{complex128(-1), 0, true},
	}
	// int
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == test.expect && reflect.TypeOf(got).Kind() == reflect.Int {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int8
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int8](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int8(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int8 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int16
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int16](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int16(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int16 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int32
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int32](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int32(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int32 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int64
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int64](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int64(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int64 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
}

func TestComplexToUint(t *testing.T) {
	var tests = []struct {
		in     any
		expect uint
		err    bool
	}{
		{complex64(1), 0, true},
		{complex64(-1), 0, true},
		{complex128(1), 0, true},
		{complex128(-1), 0, true},
	}
	// uint
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == test.expect && reflect.TypeOf(got).Kind() == reflect.Uint {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint8
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint8](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint8(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint8 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint16
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint16](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint16(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint16 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint32
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint32](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint32(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint32 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint64
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint64](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint64(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint64 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
}

func TestArrayToInt(t *testing.T) {
	var tests = []struct {
		in     any
		expect int
		err    bool
	}{
		{[]int{1, 2}, 0, true},
		{[]string{"1", "2"}, 0, true},
	}
	// int
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == test.expect && reflect.TypeOf(got).Kind() == reflect.Int {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int8
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int8](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int8(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int8 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int16
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int16](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int16(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int16 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int32
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int32](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int32(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int32 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int64
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int64](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int64(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int64 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
}

func TestArrayToUint(t *testing.T) {
	var tests = []struct {
		in     any
		expect uint
		err    bool
	}{
		{[]uint{1, 2}, 0, true},
		{[]string{"1", "2"}, 0, true},
	}
	// uint
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == test.expect && reflect.TypeOf(got).Kind() == reflect.Uint {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint8
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint8](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint8(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint8 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint16
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint16](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint16(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint16 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint32
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint32](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint32(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint32 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint64
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint64](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint64(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint64 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
}

func TestBoolToInt(t *testing.T) {
	var tests = []struct {
		in     any
		expect int
		err    bool
	}{
		{true, 1, false},
		{false, 0, false},
	}
	// int
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == test.expect && reflect.TypeOf(got).Kind() == reflect.Int {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int8
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int8](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int8(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int8 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int16
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int16](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int16(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int16 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int32
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int32](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int32(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int32 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int64
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int64](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int64(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int64 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
}

func TestBoolToUint(t *testing.T) {
	var tests = []struct {
		in     any
		expect uint
		err    bool
	}{
		{true, 1, false},
		{false, 0, false},
	}
	// uint
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == test.expect && reflect.TypeOf(got).Kind() == reflect.Uint {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint8
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint8](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint8(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint8 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint16
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint16](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint16(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint16 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint32
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint32](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint32(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint32 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint64
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint64](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint64(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint64 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
}

func TestStrToInt(t *testing.T) {
	var tests = []struct {
		in     any
		expect int
		err    bool
	}{
		{"1", 1, false},
		{"0", 0, false},
		{"-1", -1, false},
		{"1.1", 1, false},
		{"1.4", 1, false},
		{"1.5", 1, false},
		{"1.9", 1, false},
		{"0.0", 0, false},
		{"-1.1", -1, false},
		{"-1.4", -1, false},
		{"-1.1", -1, false},
		{"-1.1", -1, false},
		{"1,000", 1000, false},
		{"1,000,000", 1000000, false},
		{"Hi", 0, true},
	}
	// int
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == test.expect && reflect.TypeOf(got).Kind() == reflect.Int {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int8
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int8](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int8(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int8 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int16
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int16](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int16(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int16 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int32
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int32](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int32(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int32 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// int64
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[int64](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == int64(test.expect) && reflect.TypeOf(got).Kind() == reflect.Int64 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
}

func TestStrToUint(t *testing.T) {
	var tests = []struct {
		in     any
		expect uint
		err    bool
	}{
		{"1", 1, false},
		{"0", 0, false},
		{"-1", 0, true},
		{"1.1", 1, false},
		{"1.4", 1, false},
		{"1.5", 1, false},
		{"1.9", 1, false},
		{"0.0", 0, false},
		{"-1.1", 0, true},
		{"-1.4", 0, true},
		{"-1.5", 0, true},
		{"-1.9", 0, true},
		{"1,000", 1000, false},
		{"1,000,000", 1000000, false},
		{"Hi", 0, true},
	}
	// uint
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == test.expect && reflect.TypeOf(got).Kind() == reflect.Uint {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint8
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint8](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint8(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint8 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), expect error: %v, actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, test.err, got, got, err)
		}
	}
	// uint16
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint16](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint16(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint16 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint32
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint32](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint32(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint32 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
	// uint64
	for _, test := range tests {
		var success bool
		got, err := cast.ToE[uint64](test.in)
		if (nil == err && false == test.err) || (nil != err && true == test.err) {
			if got == uint64(test.expect) && reflect.TypeOf(got).Kind() == reflect.Uint64 {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %v (%T), expect: %v (%T), actual: %v (%T), error: %#+v", test.in, test.in, test.expect, test.expect, got, got, err)
		}
	}
}
