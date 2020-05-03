package cast_test

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/bdlm/cast"
	"github.com/stretchr/testify/assert"
)

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float64
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
		{float32(8), 8, false},
		{float64(8.31), 8.31, false},
		{"8", 8, false},
		{true, 1, false},
		{false, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := cast.ToFloat64E(test.input)
		if test.err {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToFloat32(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float32
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
		{float32(8.31), 8.31, false},
		{float64(8.31), 8.31, false},
		{"8", 8, false},
		{true, 1, false},
		{false, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := cast.ToFloat32E(test.input)
		if test.err {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToString(t *testing.T) {
	type Key struct {
		k string
	}
	key := &Key{"foo"}

	tests := []struct {
		input  interface{}
		expect string
		err    bool
	}{
		{int(8), "8", false},
		{int8(8), "8", false},
		{int16(8), "8", false},
		{int32(8), "8", false},
		{int64(8), "8", false},
		{uint(8), "8", false},
		{uint8(8), "8", false},
		{uint16(8), "8", false},
		{uint32(8), "8", false},
		{uint64(8), "8", false},
		{float32(8.31), "8.31", false},
		{float64(8.31), "8.31", false},
		{true, "true", false},
		{false, "false", false},
		{nil, "", false},
		{[]byte("one time"), "one time", false},
		{"one more time", "one more time", false},
		{template.HTML("one time"), "one time", false},
		{template.URL("http://somehost.foo"), "http://somehost.foo", false},
		{template.JS("(1+2)"), "(1+2)", false},
		{template.CSS("a"), "a", false},
		{template.HTMLAttr("a"), "a", false},
		// errors
		{testing.T{}, "", true},
		{key, "", true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := cast.ToStringE(test.input)
		if test.err {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}
