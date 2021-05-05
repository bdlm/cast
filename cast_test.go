package cast_test

import (
	"testing"
	"time"

	"github.com/bdlm/cast"

	"github.com/stretchr/testify/assert"
)

const (
	TestDuration = "100000000000000000000000"
)

func TestCast(t *testing.T) {
	var val interface{}
	var err error

	tests := []struct {
		data     interface{}
		to       interface{}
		expected interface{}
	}{
		{
			map[string]string{"1": "false", "3": "false", "5": "false", "6": "true", "7": "false", "8": "true"},
			map[int]bool{},
			map[int]bool{1: false, 3: false, 5: false, 6: true, 7: false, 8: true},
		}, {
			map[int]bool{1: false, 3: false, 5: false, 6: true, 7: false, 8: true},
			map[string]string{},
			map[string]string{"1": "false", "3": "false", "5": "false", "6": "true", "7": "false", "8": "true"},
		}, {
			map[interface{}]interface{}{"1": "100", 3: 100, 5.0: 0, int8(6): 0.1, float32(7): "10000000000", 8: -1},
			map[int64]time.Duration{},
			map[int64]time.Duration{1: 100, 3: 100, 5: 0, 6: 0, 7: 10000000000, 8: -1},
		}, {
			map[int64]time.Duration{1: 100, 3: 100, 5: 0, 6: 0, 7: 10000000000, 8: -1},
			map[string]string{},
			map[string]string{"1": "100ns", "3": "100ns", "5": "0s", "6": "0s", "7": "10s", "8": "-1ns"},
		},
		//{
		//	map[interface{}]interface{}{"1": "100", 3: 100, 5.0: 0, int8(6): 0.1, float32(7): "10000000000", 8: -1},
		//	map[int64]time.Time{},
		//	map[int64]time.Time{1: time.Now()},
		//},
	}

	for _, test := range tests {
		val, err = cast.CastE(test.data, test.to)
		assert.NoError(t, err)
		t.Logf("expected: %v, actual: %v", test.expected, val)
		assert.Equal(t, test.expected, val)
	}
}
