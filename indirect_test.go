package cast_test

import (
	"testing"

	"github.com/bdlm/cast"
	"github.com/stretchr/testify/assert"
)

func TestIndirectPointers(t *testing.T) {
	x := 13
	y := &x
	z := &y

	yInt := cast.ToInt(y)
	zInt := cast.ToInt(z)

	assert.Equal(t, yInt, 13)
	assert.Equal(t, zInt, 13)
}
