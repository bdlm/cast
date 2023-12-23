package cast_test

import (
	"fmt"

	"github.com/bdlm/cast/v2"
)

func ExampleTo_int() {
	v, e := cast.ToE[int]("1")
	fmt.Printf("%#v (%T), %#v", v, v, e)
	// Output: 1 (int), <nil>
}
