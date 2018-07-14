package cast_test

import (
	"fmt"
	"testing"

	"github.com/bdlm/cast"
	"github.com/stretchr/testify/assert"
)

func TestToBool(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect bool
		iserr  bool
	}{
		{0, false, false},
		{nil, false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"False", false, false},
		{"f", false, false},
		{"F", false, false},
		{false, false, false},

		{"true", true, false},
		{"TRUE", true, false},
		{"True", true, false},
		{"t", true, false},
		{"T", true, false},
		{1, true, false},
		{true, true, false},
		{-1, true, false},

		// errors
		{"test", false, true},
		{testing.T{}, false, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := cast.ToBool(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func BenchmarkToBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v, _ := cast.ToBool(true)
		if !v {
			b.Fatal("ToBool returned false")
		}
	}
}
