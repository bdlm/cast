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

func ExampleTo_slice() {
	v := cast.To[[]int]([]string{"1", "2", "3"})
	fmt.Printf("%v (%T)", v, v)
	// Output: [1 2 3] ([]int)
}

func ExampleToE_slice_unique_values() {
	v, e := cast.ToE[[]int]([]int{1, 2, 1, 3}, cast.Op{cast.UNIQUE_VALUES, true})
	fmt.Printf("%v (%T), %v", v, v, e)
	// Output: [1 2 3] ([]int), <nil>
}

func ExampleTo_chan() {
	ch := cast.To[chan int]("10")
	v := <-ch
	fmt.Printf("%v (%T)", v, v)
	// Output: 10 (int)
}

func ExampleToE_chan_length() {
	ch, e := cast.ToE[chan int](10, cast.Op{cast.LENGTH, 5})
	v := <-ch
	fmt.Printf("%v (cap %d), %v", v, cap(ch), e)
	// Output: 10 (cap 5), <nil>
}

func ExampleTo_func() {
	f := cast.To[cast.Func[int]]("10")
	fmt.Printf("%v (%T)", f(), f())
	// Output: 10 (int)
}

func ExampleToE_map_from_map() {
	m, e := cast.ToE[map[string]int](map[string]string{"a": "1"})
	fmt.Printf("%v (%T), %v", m["a"], m["a"], e)
	// Output: 1 (int), <nil>
}

func ExampleToE_map_from_struct() {
	type Point struct{ X, Y int }
	m, e := cast.ToE[map[string]any](Point{X: 3, Y: 4})
	fmt.Printf("X=%v Y=%v, %v", m["X"], m["Y"], e)
	// Output: X=3 Y=4, <nil>
}

func ExampleToE_map_from_private_struct() {
	type Point struct{ x, y int }
	m, e := cast.ToE[map[string]any](Point{x: 3, y: 4}, cast.Op{cast.PRIVATE, true})
	fmt.Printf("x=%v y=%v, %v", m["x"], m["y"], e)
	// Output: x=3 y=4, <nil>
}

func ExampleToE_map_from_slice() {
	m, e := cast.ToE[map[int]string]([]string{"a", "b", "c"})
	fmt.Printf("%v (%T), %v", m[0], m[0], e)
	// Output: a (string), <nil>
}

func ExampleToE_string_json() {
	v, e := cast.ToE[string](`hello "world"`, cast.Op{cast.JSON, true})
	fmt.Printf("%v, %v", v, e)
	// Output: "hello \"world\"", <nil>
}
