package cast_test

import (
	"fmt"
	"testing"

	"github.com/bdlm/cast/v2"
)

// toFunc returns a function that casts an interface to the specified
// type.
func TestToFunc(t *testing.T) {
	var tests = []struct {
		in        any
		expect    any
		err       error
		expectErr bool
	}{
		{1, 1, nil, false},
		{2, 2, nil, false},
		{3, 3, nil, false},
		{4, 4, nil, false},
		{"5", 5, nil, false},
		{true, 1, nil, false},
		{1.9, 1, nil, false},
	}

	for _, test := range tests {
		var success bool
		got, err := cast.ToE[cast.Func[int]](test.in)
		var v interface{}
		fmt.Printf("got: %T, err: %#+v, ", got, err)
		if (nil == err && false == test.expectErr) || (nil != err && true == test.expectErr) {
			v = got()
			fmt.Printf("returns: %v (%T)\n", v, v)
			if v == test.expect || err == test.err {
				success = true
			}
		}
		if !success {
			t.Errorf("input: %#v (%T), expect: %v (%T), actual: %v (%T), error: %# +v", test.in, test.in, test.expect, test.expect, v, v, err)
		}
	}

	//fn, err := cast.ToE[cast.Func](1)
	//fmt.Println(err)
	//fmt.Println(fn())
	t.Errorf("stop")
}
