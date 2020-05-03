package cast_test

import (
	"testing"

	"github.com/bdlm/cast"
	"github.com/stretchr/testify/assert"
)

func TestStringerToString(t *testing.T) {
	var x foo
	x.val = "bar"
	xStr := cast.ToString(x)
	assert.Equal(t, "bar", xStr)
}

func TestErrorToString(t *testing.T) {
	var x fu
	x.val = "bar"
	xStr := cast.ToString(x)
	assert.Equal(t, "bar", xStr)
}

type foo struct {
	val string
}

func (x foo) String() string {
	return x.val
}

type fu struct {
	val string
}

func (x fu) Error() string {
	return x.val
}
