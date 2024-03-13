package cast_test

import (
	"fmt"

	"github.com/bdlm/cast/v2"
)

func ExampleToE_string() {
	v, e := cast.ToE[string](float64(1.0))
	fmt.Printf("%#v (%T), %v", v, v, e)
	// Output: "1" (string), <nil>
}
func ExampleToE_error_with_default() {
	v, e := cast.ToE[int]("Hi!", cast.Op{cast.DEFAULT, 10})
	fmt.Printf("%#v (%T), %v", v, v, e)
	// Output: 10 (int), unable to cast "Hi!" of type string to int
}

func ExampleTo_string() {
	v := cast.To[string](1.234)
	fmt.Printf("%#v (%T)", v, v)
	// Output: "1.234" (string)
}

func ExampleTo_int() {
	v := cast.To[int]("1")
	fmt.Printf("%#v (%T)", v, v)
	// Output: 1 (int)
}

func ExampleToE_int() {
	v, e := cast.ToE[int]("1")
	fmt.Printf("%#v (%T), %v", v, v, e)
	// Output: 1 (int), <nil>
}

func ExampleToE_uint_err() {
	v, e := cast.ToE[uint]("-1")
	fmt.Printf("%v (%T), %v", v, v, e)
	// Output: 0 (uint), unable to cast "-1" of type string to uint
}

func ExampleToE_uint_abs() {
	v, e := cast.ToE[uint]("-1", cast.Op{cast.ABS, true})
	fmt.Printf("%v (%T), %v", v, v, e)
	// Output: 1 (uint), <nil>
}

func ExampleTo_float64() {
	v := cast.To[float64]("1.234")
	fmt.Printf("%#v (%T)", v, v)
	// Output: 1.234 (float64)
}

func ExampleToE_float64() {
	v, e := cast.ToE[float64]("1")
	fmt.Printf("%#v (%T), %v", v, v, e)
	// Output: 1 (float64), <nil>
}
