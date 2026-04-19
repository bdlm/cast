# cast

<p align="center">
    <a href="https://gopherize.me/gopher/0b8aa47b088b43d10817e8a13cb115fdd87c0bcb"><img src="https://github.com/bdlm/cast/wiki/assets/images/gopher.png" width="300px"></a>
</p>

Now with Generics! [![Go](https://github.com/bdlm/cast/actions/workflows/go.yml/badge.svg)](https://github.com/bdlm/cast/actions/workflows/go.yml)

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). You should expect package stability in <strong>Minor</strong> and <strong>Patch</strong> version releases

- **Major**: backwards incompatible package updates
- **Minor**: feature additions
- **Patch**: bug fixes, backward compatible model and function changes, etc.

<a href="https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#mature"><img src="https://img.shields.io/badge/stability-mature-008000.svg" alt="Mature"></a> Code has proven satisfactory and is ready for production use, cleanup of the underlying code may cause some minor changes. Backwards-compatibility is guaranteed.

**[CHANGELOG](CHANGELOG.md)**<br>

<p align="center">
    <a href="https://github.com/bdlm/cast/blob/master/CHANGELOG.md"><img src="https://img.shields.io/github/v/release/bdlm/cast" alt="Release"></a>
    <a href="https://pkg.go.dev/github.com/bdlm/cast/v2"><img src="https://godoc.org/github.com/bdlm/cast/v2?status.svg" alt="GoDoc"></a>
    <a href="https://travis-ci.org/bdlm/cast"><img src="https://travis-ci.org/bdlm/cast.svg?branch=master" alt="Build status"></a>
    <a href="https://codecov.io/gh/bdlm/cast"><img src="https://img.shields.io/codecov/c/github/bdlm/cast/master.svg" alt="Coverage status"></a>
    <a href="https://goreportcard.com/report/github.com/bdlm/cast"><img src="https://goreportcard.com/badge/github.com/bdlm/cast" alt="Go Report Card"></a>
    <a href="https://github.com/bdlm/cast/issues"><img src="https://img.shields.io/github/issues-raw/bdlm/cast.svg" alt="Github issues"></a>
    <a href="https://github.com/bdlm/cast/pulls"><img src="https://img.shields.io/github/issues-pr/bdlm/cast.svg" alt="Github pull requests"></a>
    <a href="https://github.com/bdlm/cast/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT"></a>
</p>

<sub>This project is inspired by [`spf13/cast`](https://github.com/spf13/cast)</sub>

## What is Cast?

`cast` is a library to easily convert between different data types in a straightforward and predictable way.

`cast` provides a generic function to easily convert both simple types (number to a string, interface to a bool, etc.) and complex types (slice to map, any to func() any, any to chan any, etc.). Cast does this intelligently when an obvious conversion is possible and logically when a conversion requires a predictable measurable process, such as casting a bool to a channel or a struct to a map. It doesn’t make any assumptions about how types should be converted but follows simple predictable rules.

For example you can only cast a string to an int when it is a string representation of a number, such as `"6.789"`. In a case like this, a reliable predictable rule converts that value to `int(6)` by converting it to a `float64` and calling `math.Floor()`. The reason it does not round is because there is no integer that is almost `7`, but there __is__ a `6` which can be contained within the original `float64`.

`cast` is meant to simplify consumption of untyped or inconveniently / poorly typed data by removing all the boilerplate you would otherwise write for each use-case. [More about `cast`](ABOUT.md).

## Why use Cast?

The primary use-case for `cast` is consuming untyped or poorly/loosely typed data, especially from unpredictable data sources. This can require a lot of repetitive boilerplate for validating and then typing incoming data (string representations of numbers is incredibly common and usually useless except for printing).

`cast` goes beyond just using type assertion (though it uses that whenever possible) to provide a very straightforward and usable API. If you are working with interfaces to handle dynamic content or are taking in data from YAML, TOML or JSON or other formats which lack full types or reliable producers, `cast` can be used to get the boilerplate out of your line of sight so you can just work on your code.

## Usage

`cast` provides two generic functions:

```go
func To[T Types](v any, o ...Op) T
func ToE[T Types](v any, o ...Op) (T, error)
```

`To` returns the cast value and silently ignores errors. `ToE` returns both the cast value and any error. The type parameter `T` is constrained to `cast.Types`, which covers all scalar types, slices, maps, channels, and `cast.Func[T]`.

***If input cannot be converted to the specified type, the zero value for that type is returned***. Use `ToE` to distinguish between a successful zero-value cast and a conversion error. `ToE` returns an error describing any issue along with the cast value.

### Options

Both functions accept optional `Op` values that control conversion behavior. Each `Op` carries a `Flag` constant and a value. Multiple options may be passed:

```go
result := cast.To[float32](val, cast.Op{cast.DEFAULT, float32(3.14)})
result, err := cast.ToE[[]string](items, cast.Op{cast.UNIQUE_VALUES, true})

items := map[any]any{
    "10": 10,
    0: struct{Field1 string, Field2 int}{"struct": 10},
    struct{}{}: "struct{}{} value",
}
result, err := cast.ToE[[]string](items, cast.Op{cast.UNIQUE_VALUES, true}, cast.Op{cast.JSON, true})


```

Available flags:

| Flag | Type | Default | Applies to | Description |
|------|------|---------|------------|-------------|
| `DEFAULT` | same as target | zero value | all types | Value to return instead of zero on error |
| `ABS` | `bool` | `false` | int/uint targets | Use the absolute value when casting a negative number to an unsigned type instead of returning an error |
| `DUPLICATE_KEY_ERROR` | `bool` | `false` | map target | Return an error when two source keys cast to the same target key |
| `JSON` | `bool` | `false` | string target | Encode the result as a JSON string literal (adds surrounding quotes and escaping) |
| `LENGTH` | `int` | `1` (chan), `1` (slice) | slice/chan targets | Initial backing-array capacity for slices; buffer size for channels must be 1 or greater |
| `PRIVATE` | `bool` | `false` | map target (struct source) | Include unexported struct fields in the output map |
| `STRICT` | `bool` | `false` | map target (struct source) | Return an error instead of silently skipping fields that cannot be converted |
| `UNIQUE_VALUES` | `bool` | `false` | slice target | Deduplicate slice elements after conversion |

### Examples

##### Channels
Casting to a channel returns a channel of the specified type with a buffer of 1 pre-loaded with the cast value.
```go
intCh := cast.To[chan int]("10")
ten := <-intCh // 10 (int)

var strCh chan string
strCh = cast.To[chan string](10)
str := <-strCh // "10" (string)

boolval := <-cast.To[chan bool](1) // true (bool)

// Custom buffer size
ch, err := cast.ToE[chan int](42, cast.Op{cast.LENGTH, 10})
```

##### Func
Casting to [`cast.Func[T]`](https://github.com/bdlm/cast/blob/main/to.type.go) returns a `func() T` closure that returns the cast value. The `cast.Func[T]` named type is required because Go generics cannot use a plain function type as a type parameter.
```go
var intFunc cast.Func[int]
intFunc = cast.To[cast.Func[int]]("10")
fmt.Printf("%#v (%T)\n", intFunc(), intFunc()) // 10 (int)

strFunc := cast.To[cast.Func[string]](10)
fmt.Printf("%#v (%T)\n", strFunc(), strFunc()) // "10" (string)
```

##### Slices
Casting to a slice type converts each element of the source slice or array. The source must be a slice or array; scalar values are not accepted.
```go
ints  := cast.To[[]int]([]string{"1", "2", "3"})    // []int{1, 2, 3}
strs  := cast.To[[]string]([]int{1, 2, 3})           // []string{"1", "2", "3"}
bools := cast.To[[]bool]([]int{1, 0, 1})             // []bool{true, false, true}

// Deduplicate
uniq, err := cast.ToE[[]int]([]int{1, 2, 1, 3}, cast.Op{cast.UNIQUE_VALUES, true})
// uniq = []int{1, 2, 3}

// Pre-allocate backing capacity
big, err := cast.ToE[[]int]([]string{"1", "2"}, cast.Op{cast.LENGTH, 100})
```

##### Maps
Casting to a map type is supported from maps, structs, and slices/arrays.
```go
// map → map: keys and values are cast to the target types
m, err := cast.ToE[map[string]int](map[string]string{"a": "1", "b": "2"})
// m = map[string]int{"a": 1, "b": 2}

// struct → map: exported field names become keys
type Point struct{ X, Y int }
p, err := cast.ToE[map[string]any](Point{X: 3, Y: 4})
// p = map[string]any{"X": 3, "Y": 4}

// slice → map: element indices become keys
idx, err := cast.ToE[map[int]string]([]string{"a", "b", "c"})
// idx = map[int]string{0: "a", 1: "b", 2: "c"}

// Struct options
p2, err := cast.ToE[map[string]any](
    myStruct,
    cast.Op{cast.PRIVATE, true},  // include unexported fields
    cast.Op{cast.STRICT,  true},  // error instead of skipping bad fields
)
```

##### String
```go
cast.To[string]("Hi!")              // "Hi!" (string)
cast.To[string](8)                  // "8" (string)
cast.To[string](8.31)               // "8.31" (string)
cast.To[string]([]byte("one time")) // "one time" (string)
cast.To[string](nil)                // "" (string)

var foo interface{} = "one more time"
cast.To[string](foo)                // "one more time" (string)

// JSON-encode the string result (adds quotes and escaping)
s, err := cast.ToE[string](`hello "world"`, cast.Op{cast.JSON, true})
// s = `"hello \"world\""`
```

##### Int
```go
cast.To[int](8)      // 8 (int)
cast.To[int](8.31)   // 8 (int)  — truncates via math.Floor, does not round
cast.To[int]("8")    // 8 (int)
cast.To[int]("8.31") // 8 (int)
cast.To[int]("8.51") // 8 (int)
cast.To[int](true)   // 1 (int)
cast.To[int](false)  // 0 (int)
cast.To[int](nil)    // 0 (int)

// Negative → unsigned: use ABS to take the absolute value instead of erroring
v, err := cast.ToE[uint](score, cast.Op{cast.ABS, true})
```

##### error and fmt.Stringer targets

Values that already implement `error`, `fmt.Stringer`, or the `github.com/bdlm/std/v2/errors.Error` interface can be cast to those respective interface targets; the value is returned as-is. Values that do not implement the target interface result in an error.

```go
e := fmt.Errorf("something failed")
cast.To[error](e)           // returns e unchanged
cast.ToE[error](e)          // returns (e, nil)
cast.ToE[error](42)         // returns (nil, cast.Error) — int does not implement error
cast.ToE[fmt.Stringer](42)  // returns (nil, cast.Error) — int does not implement fmt.Stringer
```

##### Error checking
To capture any conversion errors, use the `ToE` method:
```go
cast.To[int]("Hi!")           // 0 (int)
cast.ToE[int]("Hi!")          // 0, error: unable to cast "Hi!" of type string to int
```

### Supported Conversions
| | To | `bool` | `complex64` | `complex128` | `float32` | `float64` | `int` | `int8` | `int16` | `int32` | `int64` | `uint` | `uint8` | `uint16` | `uint32` | `uint64` | `string` | `slice` | `map` | `func` | `chan` |
|:---------|-----------------:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| **From** | **`any`**        | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y |
|          | **`bool`**       | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`complex64`**  | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`complex128`** | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`float32`**    | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`float64`**    | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`int`**        | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`int8`**       | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`int16`**      | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`int32`**      | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`int64`**      | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`uint`**       | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`uint8`**      | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`uint16`**     | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`uint32`**     | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`uint64`**     | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`string`**     | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | y | n | n | y | y |
|          | **`slice`**      | n | n | n | n | n | n | n | n | n | n | n | n | n | n | n | y | y | y | y | y |
|          | **`map`**        | n | n | n | n | n | n | n | n | n | n | n | n | n | n | n | y | n | y | y | y |
|          | **`struct`**     | n | n | n | n | n | n | n | n | n | n | n | n | n | n | n | n | n | y | n | n |
|          | **`func`**       | n | n | n | n | n | n | n | n | n | n | n | n | n | n | n | y | n | n | n | n |
|          | **`chan`**       | n | n | n | n | n | n | n | n | n | n | n | n | n | n | n | y | n | n | n | n |

> **Interface targets (`error`, `fmt.Stringer`):** Casting to interface targets is supported when the source value already implements the target interface; the value is returned as-is. These are not reflected in the matrix above.

