package cast_test

import (
	"fmt"
	"testing"

	"github.com/bdlm/cast"
	"github.com/stretchr/testify/assert"
)

func TestToUint(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint
		err    bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := cast.ToUintE(test.input)
		if test.err {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint64(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint64
		err    bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := cast.ToUint64E(test.input)
		if test.err {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint32(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint32
		err    bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := cast.ToUint32E(test.input)
		if test.err {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint16(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint16
		err    bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := cast.ToUint16E(test.input)
		if test.err {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint8(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint8
		err    bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := cast.ToUint8E(test.input)
		if test.err {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}
