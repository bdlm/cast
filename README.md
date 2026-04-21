# cast

<p align="center">
    <a href="https://gopherize.me/gopher/0b8aa47b088b43d10817e8a13cb115fdd87c0bcb"><img src="https://github.com/bdlm/cast/wiki/assets/images/gopher.png" width="300px"></a>
</p>

<p align="center">Now with Generics!</p>

<p align="center">
    <a href="https://github.com/bdlm/cast/actions/workflows/go.yml"><img src="https://github.com/bdlm/cast/actions/workflows/go.yml/badge.svg"></a>
    <a href="https://github.com/bdlm/cast/blob/master/CHANGELOG.md"><img src="https://img.shields.io/github/v/release/bdlm/cast" alt="Release"></a>
    <a href="https://pkg.go.dev/github.com/bdlm/cast/v2"><img src="https://godoc.org/github.com/bdlm/cast/v2?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/bdlm/cast"><img src="https://goreportcard.com/badge/github.com/bdlm/cast/v2" alt="Go Report Card"></a>
    <a href="https://github.com/bdlm/cast/issues"><img src="https://img.shields.io/github/issues-raw/bdlm/cast.svg" alt="Github issues"></a>
    <a href="https://github.com/bdlm/cast/pulls"><img src="https://img.shields.io/github/issues-pr/bdlm/cast.svg" alt="Github pull requests"></a>
    <a href="https://github.com/bdlm/cast/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT"></a>
</p>

**[CHANGELOG](CHANGELOG.md)** - This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). You should expect API stability in <strong>Minor</strong> and <strong>Patch</strong> version releases

<a href="https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#mature"><img src="https://img.shields.io/badge/stability-mature-008000.svg" alt="Mature"></a> Code has proven satisfactory and is in wide production use, cleanup of the underlying code may cause some minor changes. Backwards-compatibility is guaranteed.

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

### Supported Conversions

`cast.To[T](v, opts...)` · `cast.ToE[T](v, opts...)` · Named types work everywhere: `type Celsius float32` → `cast.ToE[Celsius]("98.6")`

**Legend:**\
&nbsp;&nbsp;`✓` always succeeds\
&nbsp;&nbsp;`~` succeeds for valid input\
&nbsp;&nbsp;`✗` always errors

| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Target →<br>↓ Source | `bool` | `int*` | `uint*` | `float*` · `complex*` | `string` | `[]T` | `map[K]V` | `chan T` · `Func[T]` | `any` |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| `bool`                | ✓   | ✓   | ✓   | ✓    | ✓    | ✗   | ✗    | ✓   | ✓ |
| `int*` (signed)       | ✓   | ✓   | ~¹  | ✓    | ✓    | ✗   | ✗    | ✓   | ✓ |
| `uint*` (unsigned)    | ✓   | ✓   | ✓   | ✓    | ✓    | ✗   | ✗    | ✓   | ✓ |
| `float*` · `complex*` | ✓   | ~²³ | ~¹² | ✓    | ✓    | ✗   | ✗    | ✓   | ✓ |
| `string`              | ~⁴  | ~⁵  | ~⁵  | ~⁵   | ✓    | ✗   | ✗    | ~⁵  | ✓ |
| `[]byte`              | ✗   | ✗   | ✗   | ✗    | ✓⁶   | ✗   | ✗    | ✗   | ✓ |
| `[]T` · `[N]T`        | ✗   | ✗   | ✗   | ✗    | ~⁷   | ✓   | ✓⁸   | ✓   | ✓ |
| `map[K]V`             | ✗   | ✗   | ✗   | ✗    | ~⁷   | ✗   | ✓    | ✗   | ✓ |
| `struct`              | ✗   | ✗   | ✗   | ✗    | ~⁷   | ✗   | ✓⁹   | ✗   | ✓ |
| `nil`                 | ✓   | ✓   | ✓   | ✓    | ✓    | ✗   | ✗    | ✗   | ✓ |
| `error` · `Stringer`  | ~¹⁰ | ~¹⁰ | ~¹⁰ | ~¹⁰  | ✓¹⁰  | ✗   | ✗    | ~¹⁰ | ✓ |
| `any` / interface     | ✓   | ✓   | ✓   | ✓    | ✓    | ✓   | ✓    | ✓   | ✓ |

`int*` = int · int8 · int16 · int32 · int64
`uint*` = uint · uint8 · uint16 · uint32 · uint64 · uintptr
`float*` = float32 · float64  ·  `complex*` = complex64 · complex128

---

**Notes**

¹ Negative signed or float value → `uint*`: error. Use `Op{ABS, true}` to take the absolute value instead.\
² float/complex → `int*`: truncates toward zero via `math.Floor`. `1.9 → 1`, `-1.9 → -1`. Does not round.\
³ complex → real numeric: the imaginary part is discarded; only the real component is used.\
⁴ `string` → `bool`: only `"1"/"0"/"t"/"f"/"true"/"false"` and their case variants (`"True"`, `"TRUE"`, …) return `true`.\
⁵ `string` → numeric: parsed as float64 via `strconv.ParseFloat`; non-numeric strings error. Float strings truncate when targeting `int*`.\
⁶ `[]byte` → `string`: direct `string(b)`, not element-wise — bypasses JSON encoding.\
⁷ → `string` fallback: complex types, maps, and structs stringify via `fmt.Sprintf("%v", v)`.\
⁸ `[]T`/`[N]T` → `map[K]V`: element indices (0, 1, 2 …) become map keys, cast to key type K.\
⁹ `struct` → `map[K]V`: exported field names become keys; embedded structs are inlined; nested structs recurse into nested maps when value type is `any` or `map`.\
¹⁰ `error`/`Stringer` → any target: calls `.Error()` or `.String()` to obtain the string representation, then parses it the same way a plain `string` source would be. Succeeds whenever the string value would succeed for that target type.

**`chan T` / `Func[T]`** wrap any value that can be cast to T — `source → chan T` succeeds whenever `source → T` succeeds. `chan T` returns a buffered channel (size 1) pre-loaded with the value; `Func[T]` returns a `func() T` closure. T may itself be `[]T`, `Func[T]`, or `chan T` (nesting supported).

**Interface targets** (`error`, `fmt.Stringer`) — source must already implement the interface; the value is returned as-is. Values that do not implement the target interface always error.

---

**Options**

| Flag | Applies to | Effect |
|:---|:---|:---|
| `DEFAULT` | all targets | Return this value on error instead of the zero value |
| `ABS` | `uint*` targets | Use absolute value for negative signed inputs instead of erroring |
| `LENGTH` | `[]T`, `chan T` | Pre-allocate slice capacity or set channel buffer size (≥ 1 for chan) |
| `UNIQUE_VALUES` | `[]T` | Deduplicate after conversion, preserving first-seen order |
| `JSON` | `string` | JSON-encode the resulting string (adds quotes and escaping) |
| `PRIVATE` | map from struct | Include unexported scalar fields |
| `STRICT` | map from struct | Error instead of silently skipping unconvertible fields |
| `DUPLICATE_KEY_ERROR` | map from map | Error when two source keys cast to the same target key |

### Options

Both functions accept optional `Op` values that control conversion behavior. Each `Op` carries a `Flag` constant and a value. Multiple options may be passed:

```go
result := cast.To[float32](val, cast.Op{cast.DEFAULT, float32(3.14)})

items := []any{1, "two", true, 1}
result, err := cast.ToE[[]string](items, cast.Op{cast.UNIQUE_VALUES, true})
// result = []string{"1", "two", "true"}  (duplicate 1 removed)
```

Available flags:

| Flag | Type | Default | Applies to | Description |
|------|------|---------|------------|-------------|
| `DEFAULT` | same as target | zero value | all types | Value to return instead of zero on error |
| `ABS` | `bool` | `false` | int/uint targets | Use the absolute value when casting a negative number to an unsigned type instead of returning an error |
| `DUPLICATE_KEY_ERROR` | `bool` | `false` | map target | Return an error when two source keys cast to the same target key |
| `JSON` | `bool` | `false` | string target | Encode the result as a JSON string literal (adds surrounding quotes and escaping) |
| `LENGTH` | `int` | `1` (chan), `1` (slice) | slice/chan targets | Initial backing-array capacity for slices; buffer size for channels must be 1 or greater |
| `PRIVATE` | `bool` | `false` | map target (struct source) | Include unexported scalar struct fields; unexported non-scalar fields are skipped (or error with `STRICT`) |
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
Casting to a slice type converts each element of the source slice or array. The source must be a slice or array or map; scalar values are not accepted.
```go
ints  := cast.To[[]int]([]string{"1", "2", "3"})     // []int{1, 2, 3}
strs  := cast.To[[]string]([]int{1, 2, 3})           // []string{"1", "2", "3"}
bools := cast.To[[]bool]([]int{1, 0, 1})             // []bool{true, false, true}
maps  := cast.To[[]string](map[string]int{           // []string{"1", "2", "2"}
    "a": 1, "b": 2, "c": 2,
})

// Deduplicate
uniq, err := cast.ToE[[]int]([]int{1, 2, 1, 3}, cast.Op{cast.UNIQUE_VALUES, true})
// uniq := []int{1, 2, 3}

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
    cast.Op{cast.PRIVATE, true},  // include unexported scalar fields
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
cast.ToE[error](42)         // returns (nil, errors.Error) — int does not implement error
cast.ToE[fmt.Stringer](42)  // returns (nil, errors.Error) — int does not implement fmt.Stringer
```

##### Error checking
To capture any conversion errors, use the `ToE` method:
```go
cast.To[int]("Hi!")           // 0 (int)
cast.ToE[int]("Hi!")          // 0 (int), error: unable to cast "Hi!" of type string to int
```
