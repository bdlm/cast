package cast_test

import (
	"fmt"
	"math/big"
	"net"
	"net/url"
	"regexp"
	"time"

	"github.com/bdlm/cast/v2"
)

// ── bool ─────────────────────────────────────────────────────────────────────

func ExampleTo_bool() {
	fmt.Println(cast.To[bool](1))
	fmt.Println(cast.To[bool](0))
	// Output:
	// true
	// false
}

func ExampleToE_boolErr() {
	_, e := cast.ToE[bool]("maybe")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

// ── string ───────────────────────────────────────────────────────────────────

func ExampleTo_string() {
	v := cast.To[string](1.234)
	fmt.Printf("%#v (%T)", v, v)
	// Output: "1.234" (string)
}

func ExampleToE_string() {
	// A map containing a channel cannot be JSON-marshalled — string conversion fails.
	_, e := cast.ToE[string](map[string]chan int{"a": make(chan int)})
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_stringFromBytes() {
	v := cast.To[string]([]byte("hello"))
	fmt.Printf("%v", v)
	// Output: hello
}

func ExampleTo_stringJson() {
	v := cast.To[string](`hello "world"`, cast.Op{cast.JSON, true})
	fmt.Printf("%v", v)
	// Output: "hello \"world\""
}

// ── int ──────────────────────────────────────────────────────────────────────

func ExampleTo_int() {
	v := cast.To[int]("1")
	fmt.Printf("%#v (%T)", v, v)
	// Output: 1 (int)
}

func ExampleToE_int() {
	_, e := cast.ToE[int]("not-a-number")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

// ── uint ─────────────────────────────────────────────────────────────────────

func ExampleTo_uintAbs() {
	v := cast.To[uint]("-1", cast.Op{cast.ABS, true})
	fmt.Printf("%v (%T)", v, v)
	// Output: 1 (uint)
}

func ExampleToE_uintErr() {
	v, e := cast.ToE[uint]("-1")
	fmt.Printf("%v (%T), %v", v, v, e)
	// Output: 0 (uint), cannot cast signed value to unsigned integer: unable to cast "-1" of type string to uint
}

// ── float ────────────────────────────────────────────────────────────────────

func ExampleTo_float64() {
	v := cast.To[float64]("1.234")
	fmt.Printf("%#v (%T)", v, v)
	// Output: 1.234 (float64)
}

func ExampleToE_float64() {
	_, e := cast.ToE[float64]("not-a-float")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

// ── complex ──────────────────────────────────────────────────────────────────

func ExampleTo_complex128() {
	v := cast.To[complex128](3.14)
	fmt.Printf("%v", v)
	// Output: (3.14+0i)
}

func ExampleToE_complex128() {
	_, e := cast.ToE[complex128]("not-a-number")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_complex64() {
	v := cast.To[complex64](float32(1.5))
	fmt.Printf("%v", v)
	// Output: (1.5+0i)
}

// ── slice ────────────────────────────────────────────────────────────────────

func ExampleTo_slice() {
	v := cast.To[[]int]([]string{"1", "2", "3"})
	fmt.Printf("%v (%T)", v, v)
	// Output: [1 2 3] ([]int)
}

func ExampleToE_sliceErr() {
	// Elements that cannot convert to the target type cause an error.
	_, e := cast.ToE[[]int]([]string{"a", "b"})
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_sliceUniqueValues() {
	v := cast.To[[]int]([]int{1, 2, 1, 3}, cast.Op{cast.UNIQUE_VALUES, true})
	fmt.Printf("%v (%T)", v, v)
	// Output: [1 2 3] ([]int)
}

// ── chan ─────────────────────────────────────────────────────────────────────

func ExampleTo_chan() {
	ch := cast.To[chan int]("10")
	v := <-ch
	fmt.Printf("%v (%T)", v, v)
	// Output: 10 (int)
}

func ExampleToE_chanErr() {
	_, e := cast.ToE[chan int]("not-a-number")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_chanLength() {
	ch := cast.To[chan int](10, cast.Op{cast.LENGTH, 5})
	v := <-ch
	fmt.Printf("%v (cap %d)", v, cap(ch))
	// Output: 10 (cap 5)
}

func ExampleTo_chanSlice() {
	ch := cast.To[chan []int]([]string{"1", "2", "3"})
	fmt.Printf("%v (%T)", <-ch, ch)
	// Output: [1 2 3] (chan []int)
}

// ── Func ─────────────────────────────────────────────────────────────────────

func ExampleTo_func() {
	f := cast.To[cast.Func[int]]("10")
	fmt.Printf("%v (%T)", f(), f())
	// Output: 10 (int)
}

func ExampleToE_funcErr() {
	_, e := cast.ToE[cast.Func[int]]("not-a-number")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_funcChan() {
	f := cast.To[cast.Func[chan int]](42)
	fmt.Printf("%v", <-f())
	// Output: 42
}

// ── map ──────────────────────────────────────────────────────────────────────

func ExampleTo_map() {
	m := cast.To[map[string]int](map[string]string{"a": "1", "b": "2"})
	fmt.Printf("a=%v b=%v (%T)", m["a"], m["b"], m)
	// Output: a=1 b=2 (map[string]int)
}

func ExampleToE_mapErr() {
	// A scalar source cannot be converted to a map.
	_, e := cast.ToE[map[string]int](42)
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleToE_mapDuplicateKeyError() {
	// After casting, "1" and "01" both become the integer key 1 — a duplicate.
	m, e := cast.ToE[map[int]string](
		map[string]string{"1": "one", "01": "also-one"},
		cast.Op{cast.DUPLICATE_KEY_ERROR, true},
	)
	fmt.Printf("m=%v, err=%v", m, e != nil)
	// Output: m=map[], err=true
}

func ExampleTo_mapFromMap() {
	m := cast.To[map[string]int](map[string]string{"a": "1"})
	fmt.Printf("%v (%T)", m["a"], m["a"])
	// Output: 1 (int)
}

func ExampleTo_mapFromStruct() {
	type Point struct{ X, Y int }
	m := cast.To[map[string]any](Point{X: 3, Y: 4})
	fmt.Printf("X=%v Y=%v", m["X"], m["Y"])
	// Output: X=3 Y=4
}

func ExampleTo_mapFromPrivateStruct() {
	// PRIVATE includes unexported fields as map keys.
	type Point struct{ x, y int }
	m := cast.To[map[string]any](Point{x: 3, y: 4}, cast.Op{cast.PRIVATE, true})
	fmt.Printf("x=%v y=%v", m["x"], m["y"])
	// Output: x=3 y=4
}

func ExampleTo_mapFromSlice() {
	m := cast.To[map[int]string]([]string{"a", "b", "c"})
	fmt.Printf("%v (%T)", m[0], m[0])
	// Output: a (string)
}

func ExampleTo_mapFromJSONString() {
	m := cast.To[map[string]int](`{"a":1,"b":2}`)
	fmt.Printf("a=%v b=%v", m["a"], m["b"])
	// Output: a=1 b=2
}

// ── struct ───────────────────────────────────────────────────────────────────

func ExampleTo_struct() {
	type Point struct{ X, Y int }
	p := cast.To[Point](map[string]any{"X": "3", "Y": "4"})
	fmt.Printf("X=%v Y=%v (%T)", p.X, p.Y, p)
	// Output: X=3 Y=4 (cast_test.Point)
}

func ExampleToE_structStrict() {
	type Point struct{ X, Y int }
	// "Z" has no matching field in Point — STRICT turns that into an error.
	_, e := cast.ToE[Point](
		map[string]any{"X": 1, "Y": 2, "Z": 3},
		cast.Op{cast.STRICT, true},
	)
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_structNested() {
	type Address struct{ City, Country string }
	type Person struct {
		Name    string
		Age     int
		Address Address
	}
	m := map[string]any{
		"Name": "Alice",
		"Age":  "30",
		"Address": map[string]any{
			"City":    "Portland",
			"Country": "US",
		},
	}
	p := cast.To[Person](m)
	fmt.Printf("Name=%v Age=%v City=%v Country=%v", p.Name, p.Age, p.Address.City, p.Address.Country)
	// Output: Name=Alice Age=30 City=Portland Country=US
}

func ExampleTo_mapToStruct() {
	type MyStruct struct {
		X int
		Y int
		A string
		B string
	}
	p := cast.To[MyStruct](map[string]string{
		"X": "3", "Y": "4", "A": "hello", "B": "world",
	})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T)",
		p, p.X, p.X, p.Y, p.Y, p.A, p.A, p.B, p.B)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), A=hello (string), B=world (string)
}

func ExampleTo_mapToPrivateStruct() {
	// PRIVATE allows map values to be written into unexported struct fields.
	type Point struct{ x, y int }
	p := cast.To[Point](map[string]any{"x": 3, "y": 4}, cast.Op{cast.PRIVATE, true})
	fmt.Printf("x=%v y=%v (%T)", p.x, p.y, p)
	// Output: x=3 y=4 (cast_test.Point)
}

func ExampleTo_structToStruct() {
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
	p := cast.To[MyStruct](YourStruct{X: "3", Y: "4", A: "hello", B: "world"})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T)",
		p, p.X, p.X, p.Y, p.Y, p.A, p.A, p.B, p.B)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), A=hello (string), B=world (string)
}

func ExampleTo_mapToStructTags() {
	type MyStruct struct {
		X int    `cast:"field_x"`
		Y int    `cast:"field_y"`
		A string `cast:"field_a"`
		B string `cast:"field_b"`
	}
	p := cast.To[MyStruct](map[string]string{
		"field_x": "3", "field_y": "4", "field_a": "hello", "field_b": "world",
	})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T)",
		p, p.X, p.X, p.Y, p.Y, p.A, p.A, p.B, p.B)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), A=hello (string), B=world (string)
}

func ExampleTo_structToStructPrivate() {
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
	p := cast.To[MyStruct](YourStruct{X: "3", Y: "4", a: "hello", b: "world"}, cast.Op{cast.PRIVATE, true})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), a=%v (%T), b=%v (%T)",
		p, p.X, p.X, p.Y, p.Y, p.a, p.a, p.b, p.b)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), a=hello (string), b=world (string)
}

func ExampleTo_structToStructJsonTags() {
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
	p := cast.To[MyStruct](YourStruct{fieldX: "3", fieldY: "4", fieldA: "hello", fieldB: "world"}, cast.Op{cast.PRIVATE, true})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), a=%v (%T), b=%v (%T)",
		p, p.X, p.X, p.Y, p.Y, p.a, p.a, p.b, p.b)
	// Output: p=(cast_test.MyStruct), X=3 (int), Y=4 (int), a=hello (string), b=world (string)
}

func ExampleTo_structToMap() {
	type YourStruct struct {
		X int
		Y int
		A string
		B string
	}
	p := cast.To[map[string]string](YourStruct{X: 3, Y: 4, A: "hello", B: "world"})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T)",
		p, p["X"], p["X"], p["Y"], p["Y"], p["A"], p["A"], p["B"], p["B"])
	// Output: p=(map[string]string), X=3 (string), Y=4 (string), A=hello (string), B=world (string)
}

func ExampleTo_structToMapNonStrict() {
	type YourStruct struct {
		X int
		Y int
		A string
		B string
	}
	// Non-strict: string fields A and B cannot convert to int and silently become 0.
	p := cast.To[map[string]int](YourStruct{X: 3, Y: 4, A: "hello", B: "world"})
	fmt.Printf("p=(%T), X=%v (%T), Y=%v (%T), A=%v (%T), B=%v (%T)",
		p, p["X"], p["X"], p["Y"], p["Y"], p["A"], p["A"], p["B"], p["B"])
	// Output: p=(map[string]int), X=3 (int), Y=4 (int), A=0 (int), B=0 (int)
}

// ── interface targets ─────────────────────────────────────────────────────────

func ExampleTo_errorTarget() {
	err := fmt.Errorf("something failed")
	result := cast.To[error](err)
	fmt.Printf("%v", result)
	// Output: something failed
}

func ExampleToE_errorTarget() {
	// A source that does not implement error always fails.
	_, e := cast.ToE[error](42)
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_stringerTarget() {
	// Any source that already implements fmt.Stringer is returned as-is.
	ip := net.ParseIP("127.0.0.1")
	result := cast.To[fmt.Stringer](ip)
	fmt.Printf("%v", result)
	// Output: 127.0.0.1
}

func ExampleToE_stringerTarget() {
	// A source that does not implement fmt.Stringer always fails.
	_, e := cast.ToE[fmt.Stringer](42)
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

// ── DECODE option ─────────────────────────────────────────────────────────────

func ExampleTo_decodeScalar() {
	// DECODE=json unwraps a JSON-encoded string before converting.
	v := cast.To[int](`"42"`, cast.Op{cast.DECODE, "json"})
	fmt.Printf("%v", v)
	// Output: 42
}

func ExampleTo_decodeSlice() {
	// DECODE=json forces JSON decoding for a slice target.
	v := cast.To[[]int](`[1,2,3]`, cast.Op{cast.DECODE, "json"})
	fmt.Printf("%v", v)
	// Output: [1 2 3]
}

func ExampleToE_decodeErr() {
	// DECODE=json with invalid JSON returns an error.
	_, e := cast.ToE[[]int]("not json", cast.Op{cast.DECODE, "json"})
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

// ── DEFAULT option ─────────────────────────────────────────────────────────────

func ExampleToE_errorWithDefault() {
	v, e := cast.ToE[int]("Hi!", cast.Op{cast.DEFAULT, 10})
	fmt.Printf("%#v (%T), %v", v, v, e)
	// Output: 10 (int), strconv.ParseFloat: parsing "Hi!": invalid syntax: strconv.ParseFloat: parsing "Hi!": invalid syntax: unable to cast "Hi!" of type string to int
}

// ── time.Time ────────────────────────────────────────────────────────────────

func ExampleTo_time() {
	t := cast.To[time.Time]("2024-04-22T12:00:00Z")
	fmt.Printf("%v", t.Format(time.RFC3339))
	// Output: 2024-04-22T12:00:00Z
}

func ExampleToE_time() {
	_, e := cast.ToE[time.Time]("not-a-time")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_timeFromInt() {
	// int64 input is Unix nanoseconds.
	t := cast.To[time.Time](int64(1_000_000_000))
	fmt.Printf("%v", t.Format(time.RFC3339))
	// Output: 1970-01-01T00:00:01Z
}

func ExampleTo_timeFromFloat() {
	// float64 input is Unix seconds; fractional seconds are preserved.
	t := cast.To[time.Time](float64(1.5))
	fmt.Printf("%v", t.Format(time.RFC3339Nano))
	// Output: 1970-01-01T00:00:01.5Z
}

func ExampleTo_timeFormat() {
	// FORMAT overrides the default set of tried formats.
	t := cast.To[time.Time]("2024/04/22", cast.Op{cast.FORMAT, "2006/01/02"})
	fmt.Printf("%v", t.Format(time.DateOnly))
	// Output: 2024-04-22
}

func ExampleToE_timeFormat() {
	// With FORMAT set, standard formats are not tried as a fallback.
	// "2024-04-22" uses dashes but the custom format expects slashes.
	_, e := cast.ToE[time.Time]("2024-04-22", cast.Op{cast.FORMAT, "2006/01/02"})
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

// ── time.Duration ─────────────────────────────────────────────────────────────

func ExampleTo_duration() {
	d := cast.To[time.Duration]("1h30m")
	fmt.Printf("%v", d)
	// Output: 1h30m0s
}

func ExampleToE_duration() {
	// A bare number without a unit suffix is not a valid duration string.
	_, e := cast.ToE[time.Duration]("5")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_durationFromInt() {
	// int64 input is nanoseconds.
	d := cast.To[time.Duration](int64(5_000_000_000))
	fmt.Printf("%v", d)
	// Output: 5s
}

// ── net.IP ───────────────────────────────────────────────────────────────────

func ExampleTo_netIP() {
	ip := cast.To[net.IP]("192.168.1.1")
	fmt.Printf("%v", ip)
	// Output: 192.168.1.1
}

func ExampleToE_netIP() {
	_, e := cast.ToE[net.IP]("not-an-ip")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_netIPFromUint32() {
	// uint32 encodes a packed big-endian IPv4 address.
	ip := cast.To[net.IP](uint32(0xC0A80101))
	fmt.Printf("%v", ip)
	// Output: 192.168.1.1
}

// ── *url.URL ─────────────────────────────────────────────────────────────────

func ExampleTo_url() {
	u := cast.To[*url.URL]("https://example.com/path?q=1")
	fmt.Printf("%v %v", u.Host, u.Path)
	// Output: example.com /path
}

func ExampleToE_url() {
	// nil is the simplest source that always errors for *url.URL.
	_, e := cast.ToE[*url.URL](nil)
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

// ── *regexp.Regexp ───────────────────────────────────────────────────────────

func ExampleTo_regexp() {
	re := cast.To[*regexp.Regexp](`^\d+$`)
	fmt.Printf("%v %v", re.MatchString("42"), re.MatchString("hi"))
	// Output: true false
}

func ExampleToE_regexp() {
	_, e := cast.ToE[*regexp.Regexp]("[unclosed")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

// ── *big.Int ─────────────────────────────────────────────────────────────────

func ExampleTo_bigInt() {
	n := cast.To[*big.Int]("123456789012345678901234567890")
	fmt.Printf("%v", n)
	// Output: 123456789012345678901234567890
}

func ExampleToE_bigInt() {
	_, e := cast.ToE[*big.Int]("not-a-number")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
}

func ExampleTo_bigIntHex() {
	// Base is auto-detected: 0x=hex, 0o=octal, 0b=binary, else decimal.
	n := cast.To[*big.Int]("0xff")
	fmt.Printf("%v", n)
	// Output: 255
}

// ── *big.Float ───────────────────────────────────────────────────────────────

func ExampleTo_bigFloat() {
	f := cast.To[*big.Float]("3.14")
	fmt.Printf("%v", f.Text('f', 2))
	// Output: 3.14
}

func ExampleToE_bigFloat() {
	_, e := cast.ToE[*big.Float]("not-a-float")
	fmt.Printf("err=%v", e != nil)
	// Output: err=true
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
