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

func ExampleToE_map_to_struct() {
	type MyStruct struct {
		X int
		Y int
		A string
		B string
	}
	yourData := map[string]string{
		"X": "3",
		"Y": "4",
		"A": "hello",
		"B": "world",
	}
	p, e := cast.ToE[MyStruct](yourData)
	fmt.Printf(
		"p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T), %v",
		p, p.X, p.X, p.Y, p.Y, p.A, p.A, p.B, p.B, e,
	)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), A=hello (string), B=world (string), <nil>
}

func ExampleToE_struct_to_struct() {
	type MyStruct struct {
		X int
		Y int
		A string
		B string
	}
	type YourStruct struct {
		X string
		Y string
		A string
		B string
	}
	yourStruct := YourStruct{
		X: "3",
		Y: "4",
		A: "hello",
		B: "world",
	}
	p, e := cast.ToE[MyStruct](yourStruct, cast.Op{cast.PRIVATE, true})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T), %v", p, p.X, p.X, p.Y, p.Y, p.A, p.A, p.B, p.B, e)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), A=hello (string), B=world (string), <nil>
}

func ExampleToE_map_to_struct_tags() {
	type MyStruct struct {
		X int    `cast:"field_x"`
		Y int    `cast:"field_y"`
		A string `cast:"field_a"`
		B string `cast:"field_b"`
	}
	yourData := map[string]string{
		"field_x": "3",
		"field_y": "4",
		"field_a": "hello",
		"field_b": "world",
	}
	p, e := cast.ToE[MyStruct](yourData)
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T), %v", p, p.X, p.X, p.Y, p.Y, p.A, p.A, p.B, p.B, e)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), A=hello (string), B=world (string), <nil>
}

func ExampleToE_struct_to_struct_private() {
	type MyStruct struct {
		X int
		Y int
		a string
		b string
	}
	type YourStruct struct {
		X string
		Y string
		a string
		b string
	}
	yourStruct := YourStruct{
		X: "3",
		Y: "4",
		a: "hello",
		b: "world",
	}
	p, e := cast.ToE[MyStruct](yourStruct, cast.Op{cast.PRIVATE, true})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), a=%v (%T), b=%v (%T), %v", p, p.X, p.X, p.Y, p.Y, p.a, p.a, p.b, p.b, e)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), a=hello (string), b=world (string), <nil>
}

func ExampleToE_struct_to_struct_json_tags() {
	type MyStruct struct {
		X int    `json:"fieldX"`
		Y int    `json:"fieldY"`
		a string `json:"fieldA"`
		b string `json:"fieldB"`
	}
	type YourStruct struct {
		fieldX string
		fieldY string
		fieldA string
		fieldB string
	}
	yourStruct := YourStruct{
		fieldX: "3",
		fieldY: "4",
		fieldA: "hello",
		fieldB: "world",
	}
	p, e := cast.ToE[MyStruct](yourStruct, cast.Op{cast.PRIVATE, true})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), a=%v (%T), b=%v (%T), %v", p, p.X, p.X, p.Y, p.Y, p.a, p.a, p.b, p.b, e)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), a=hello (string), b=world (string), <nil>
}

func ExampleToE_struct_to_map() {
	type YourStruct struct {
		X int
		Y int
		A string
		B string
	}
	yourStruct := YourStruct{
		X: 3,
		Y: 4,
		A: "hello",
		B: "world",
	}
	p, e := cast.ToE[map[string]string](yourStruct)
	fmt.Printf(
		"p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T), %v",
		p, p["X"], p["X"], p["Y"], p["Y"], p["A"], p["A"], p["B"], p["B"], e,
	)
	// Output: p=(map[string]string), X=3 (string), Y=4 (string), A=hello (string), B=world (string), <nil>
}

func ExampleToE_struct_to_map_error() {
	type YourStruct struct {
		X int
		Y int
		A string
		B string
	}
	yourStruct := YourStruct{
		X: 3,
		Y: 4,
		A: "hello",
		B: "world",
	}
	p, e := cast.ToE[map[string]int](yourStruct)
	fmt.Printf(
		"p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T), %v",
		p, p["X"], p["X"], p["Y"], p["Y"], p["A"], p["A"], p["B"], p["B"], e,
	)
	// Output: p=(map[string]int), X=3 (int), Y=4 (int), A=0 (int), B=0 (int), error
}

type Output struct {
	fieldX string `cast:"X"`
	fieldY int    `cast:"Y"`
	fieldA string `cast:"a"`
	fieldB []byte `cast:"b"`
}

func (o Output) Get(field string) any {
	return ""
}

// func ExampleToE_generic() {
// 	type Input struct {
// 		X int
// 		Y int
// 		a string
// 		b string
// 	}

// 	input := Input{
// 		X: 3,
// 		Y: 4,
// 		a: "hello",
// 		b: "world",
// 	}
// 	data := []Output{}
// 	output, e := cast.ToE[Output](input, cast.Op{cast.PRIVATE, true})
// 	data = append(data, output)

// 	fmt.Printf("data=(%T), fieldX=%v (%T), fieldY=%v (%T), fieldA=%v (%T), fieldB=%v (%T), %v", data[0], data[0].fieldX, data[0].fieldX, data[0].fieldY, data[0].fieldY, data[0].fieldA, data[0].fieldA, data[0].fieldB, data[0].fieldB, e)
// 	// Output: data=(cast_test.Output), fieldX=3 (string), fieldY=4 (int), fieldA=hello (string), fieldB=world (string), <nil>
// }
