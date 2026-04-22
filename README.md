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

## Table of Contents

- [What is Cast?](#what-is-cast)
- [Usage](#usage)
  - [Supported Conversions](#supported-conversions)
    - [Named-type targets](#named-type-targets)
  - [Options](#options)
  - [Examples](#examples)
    - [Scalars](#scalars)
    - [Slices](#slices)
    - [Maps](#maps)
    - [Structs](#structs)
      - [Field key resolution](#field-key-resolution)
      - [Unexported fields (PRIVATE)](#unexported-fields-private)
    - [Channels and Funcs](#channels-and-funcs)
    - [Named types](#named-types)
    - [Interface targets](#interface-targets)
    - [Error handling](#error-handling)

## What is Cast?

`cast` is a library to easily convert between different data types in a straightforward and predictable way. It provides a generic function to easily convert both simple types (number to a string, interface to a bool, etc.) and complex types (slice to map, any to func() any, any to chan any, etc.). Cast does this intelligently when an obvious conversion is possible and logically when a conversion requires a predictable measurable process.

A concrete example: casting `"6.789"` to `int` yields `6`, not `7`. Cast converts to `float64` first, then calls `math.Floor()` — because there is no integer that is _almost_ `7`, but there _is_ a `6` that can be contained within the original value.

The primary use-case is consuming untyped or loosely typed data from external sources (YAML, TOML, JSON, API responses) without repetitive type-assertion boilerplate. [More about `cast`](ABOUT.md).

## Usage

```go
// To returns the cast value; errors are silently dropped (zero value returned on failure).
func To[T Types](v any, o ...Op) T

// ToE returns the cast value and any error.
func ToE[T Types](v any, o ...Op) (T, error)

// ToStruct and ToStructE hydrate any struct type T from a map or struct source.
// T is constrained to any (not Types), so arbitrary struct types are accepted.
func ToStruct[T any](v any, o ...Op) T
func ToStructE[T any](v any, o ...Op) (T, error)
```

`T` in `To`/`ToE` is constrained to `cast.Types`, which covers all scalar types (`bool`, `int*`, `uint*`, `float*`, `complex*`, `string`), `[]T` slices, `map[K]V` maps, `chan T` channels, `cast.Func[T]` closures, and stdlib named types (`time.Time`, `time.Duration`, `net.IP`, `*url.URL`, `*regexp.Regexp`, `*big.Int`, `*big.Float`). Named types with a matching underlying kind also satisfy `Types` (e.g. `type Celsius float32`).

**On error**, `To` returns the zero value for `T`; `ToE` returns the zero value plus the error. Use `ToE` when you need to distinguish a successful zero-value cast from a failed one.

### Supported Conversions

`cast.To[T](v, opts...)` · `cast.ToE[T](v, opts...)` · Named types work everywhere: `type Celsius float32` → `cast.ToE[Celsius]("98.6")`\
Struct hydration: `cast.ToStruct[T](v, opts...)` · `cast.ToStructE[T](v, opts...)` — `T` must be a struct type.

**Legend:**\
&nbsp;&nbsp;`✓` always succeeds\
&nbsp;&nbsp;`~` succeeds for valid input\
&nbsp;&nbsp;`✗` always errors

| Source | `bool` | `int*` | `uint*` | `float*` · `complex*` | `string` | `[]T` | `map[K]V` | `struct` | `chan T` · `Func[T]` | `any` |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| `bool`                | ✓   | ✓   | ✓   | ✓    | ✓    | ✗   | ✗    | ✗    | ✓   | ✓ |
| `int*` (signed)       | ✓   | ✓   | ~¹  | ✓    | ✓    | ✗   | ✗    | ✗    | ✓   | ✓ |
| `uint*` (unsigned)    | ✓   | ✓   | ✓   | ✓    | ✓    | ✗   | ✗    | ✗    | ✓   | ✓ |
| `float*` · `complex*` | ✓   | ~²³ | ~¹² | ✓    | ✓    | ✗   | ✗    | ✗    | ✓   | ✓ |
| `string`              | ~⁴  | ~⁵  | ~⁵  | ~⁵   | ✓    | ✗   | ✗    | ✗    | ~⁵  | ✓ |
| `[]byte`              | ✗   | ✗   | ✗   | ✗    | ✓⁶   | ✗   | ✗    | ✗    | ✗   | ✓ |
| `[]T` · `[N]T`        | ✗   | ✗   | ✗   | ✗    | ~⁷   | ✓   | ✓⁸   | ✗    | ✓   | ✓ |
| `map[K]V`             | ✗   | ✗   | ✗   | ✗    | ~⁷   | ✗   | ✓    | ~¹¹  | ✗   | ✓ |
| `struct`              | ✗   | ✗   | ✗   | ✗    | ~⁷   | ✗   | ✓⁹   | ~¹¹  | ✗   | ✓ |
| `nil`                 | ✓   | ✓   | ✓   | ✓    | ✓    | ✗   | ✗    | ✗    | ✗   | ✓ |
| `error` · `Stringer`  | ~¹⁰ | ~¹⁰ | ~¹⁰ | ~¹⁰  | ✓¹⁰  | ✗   | ✗    | ✗    | ~¹⁰ | ✓ |
| `any` / interface     | ✓   | ✓   | ✓   | ✓    | ✓    | ✓   | ✓    | ✓    | ✓   | ✓ |
| Named types†          | ~   | ~   | ~   | ~    | ✓    | ✗   | ✗    | ✗    | ~   | ✓ |

`int*` = int · int8 · int16 · int32 · int64
`uint*` = uint · uint8 · uint16 · uint32 · uint64 · uintptr
`float*` = float32 · float64  ·  `complex*` = complex64 · complex128
† Named types: `time.Time` · `time.Duration` · `net.IP` · `*url.URL` · `*regexp.Regexp` · `*big.Int` · `*big.Float` — see [Named-type targets](#named-type-targets) for casting *to* these types; all implement `fmt.Stringer`, so casting *from* them to scalar targets follows the string-parse path.

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
¹⁰ `error`/`Stringer` → any target: calls `.Error()` or `.String()` to obtain the string representation, then parses it the same way a plain `string` source would be. Succeeds whenever the string value would succeed for that target type.\
¹¹ `map`/`struct` → `struct`: source map keys (or source struct exported field names) are matched case-sensitively to the target struct's exported fields. Fields with no matching source key retain their zero value; `STRICT` promotes mismatches and unmatched keys to errors. Anonymous (embedded) fields are promoted on both sides. `ToStruct[T]` / `ToStructE[T]` accept any struct type for `T` (`T any`, not constrained to `Types`).

**`chan T` / `Func[T]`** wrap any value that can be cast to T — `source → chan T` succeeds whenever `source → T` succeeds. `chan T` returns a buffered channel (size 1) pre-loaded with the value; `Func[T]` returns a `func() T` closure. T may itself be `[]T`, `Func[T]`, or `chan T` (nesting supported).

**Interface targets** (`error`, `fmt.Stringer`) — source must already implement the interface; the value is returned as-is. Values that do not implement the target interface always error.

#### Named-type targets

`time.Time`, `time.Duration`, `net.IP`, `*url.URL`, `*regexp.Regexp`, `*big.Int`, and `*big.Float` are first-class cast targets via `To[T]` / `ToE[T]`. All converters accept any source type that stringifies to a valid representation (via `fmt.Stringer` or `json.Marshal`). All support the `DEFAULT` op. All are also supported as struct field types during struct hydration.

| Source | `time.Time` | `time.Duration` | `net.IP` | `*url.URL` | `*regexp.Regexp` | `*big.Int` | `*big.Float` |
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| same named type              | ✓   | ✓   | ✓   | ✓   | ✓   | ✓ᵃ  | ✓ᵃ  |
| `*time.Time`                 | ~ᵇ  | ✗   | ✗   | ✗   | ✗   | ✗   | ✗   |
| `url.URL` (value, not ptr)   | ✗   | ✗   | ✗   | ✓   | ✗   | ✗   | ✗   |
| `big.Int` · `big.Float` (values) | ✗ | ✗ | ✗  | ✗   | ✗   | ✓ᶜ  | ✓   |
| `*big.Float` → `*big.Int`    | ✗   | ✗   | ✗   | ✗   | ✗   | ~ᶜ  | —   |
| `*big.Int` → `*big.Float`    | ✗   | ✗   | ✗   | ✗   | ✗   | —   | ✓   |
| `int*` · `uint*`             | ✓ᵈ  | ✓   | ✗   | ✗   | ✗   | ✓   | ✓   |
| `float*`                     | ✓ᵉ  | ✓   | ✗   | ✗   | ✗   | ~ᶜ  | ✓   |
| `string`                     | ~ᶠ  | ~ᵍ  | ~   | ~   | ~   | ~ʰ  | ~   |
| `[]byte`                     | ~ᶠ  | ✗   | ~ⁱ  | ✗   | ✗   | ✗   | ✗   |
| `uint32`                     | ✗   | ✗   | ✓ⁱ  | ✗   | ✗   | ✓   | ✓   |
| `nil`                        | ✗   | ✗   | ✗   | ✗   | ✗   | ✗   | ✗   |
| `any` / interface            | ✓   | ✓   | ✓   | ✓   | ✓   | ✓   | ✓   |

ᵃ Returns an independent copy — `new(big.Int).Set(src)` / `new(big.Float).Copy(src)`.\
ᵇ `*time.Time` → `time.Time`: dereferenced when non-nil; nil pointer errors.\
ᶜ `*big.Float` / `big.Float` / `float*` → `*big.Int`: truncated toward zero. NaN/±Inf float inputs error.\
ᵈ `int*` / `uint*` → `time.Time`: treated as Unix nanoseconds — `time.Unix(0, n).UTC()`.\
ᵉ `float*` → `time.Time`: treated as Unix seconds; fractional seconds preserved via nanosecond conversion.\
ᶠ `string` / `[]byte` → `time.Time`: 19 formats tried in order — RFC3339Nano, RFC3339, DateTime, RFC1123Z, RFC1123, RFC822Z, RFC822, DateOnly, then Layout/ANSIC/UnixDate/RubyDate/RFC850/Kitchen/Stamp/StampMilli/StampMicro/StampNano/TimeOnly.\
ᵍ `string` → `time.Duration`: via `time.ParseDuration`; accepts units "ns", "µs"/"us", "ms", "s", "m", "h" and combinations (e.g. "1h30m45s").\
ʰ `string` → `*big.Int`: base auto-detected from prefix — `0x` hex, `0o` octal, `0b` binary, decimal otherwise.\
ⁱ `uint32` → `net.IP`: packed IPv4 in host byte order. `[]byte` of exactly 4 or 16 bytes: direct copy into `net.IP`; other lengths parsed as string.

### Options

Both functions accept optional `Op` values that control conversion behavior. Each `Op` carries a `Flag` constant and a value. Multiple options may be passed:

```go
result := cast.To[float32](val, cast.Op{cast.DEFAULT, float32(3.14)})

items := []any{1, "two", true, 1}
result, err := cast.ToE[[]string](items, cast.Op{cast.UNIQUE_VALUES, true})
// result = []string{"1", "two", "true"}  (duplicate 1 removed)
```

**Available flags**

| Flag | Applies to | Effect |
|:---|:---|:---|
| `DEFAULT` | all targets | Return this value on error instead of the zero value |
| `ABS` | `uint*` targets | Use absolute value for negative signed inputs instead of erroring |
| `LENGTH` | `[]T`, `chan T` | Pre-allocate slice capacity or set channel buffer size (≥ 1 for chan) |
| `UNIQUE_VALUES` | `[]T` | Deduplicate after conversion, preserving first-seen order |
| `JSON` | `string` | JSON-encode the resulting string (adds quotes and escaping) |
| `PRIVATE` | map from struct; struct from map/struct | Include unexported fields when reading a struct source or hydrating a struct target |
| `STRICT` | map from struct; struct from map/struct | Error instead of silently skipping unconvertible fields or unmatched keys |
| `DUPLICATE_KEY_ERROR` | map from map | Error when two source keys cast to the same target key |

### Examples

#### Scalars

##### Bool

Only specific string values are accepted; `"yes"`, `"on"`, and similar are **not** supported.

```go
cast.To[bool](1)       // true
cast.To[bool](0)       // false
cast.To[bool]("true")  // true  — also "TRUE", "True", "t", "T", "1"
cast.To[bool]("false") // false — also "FALSE", "False", "f", "F", "0"
cast.To[bool]("yes")   // false — "yes"/"no"/"on"/"off" are not accepted
cast.To[bool](nil)     // false
```

##### Int and Float

Float-to-int truncates toward zero via `math.Floor`; it does not round.

```go
cast.To[int](8.31)   // 8   — Floor, not round
cast.To[int]("8.51") // 8   — string → float64 → Floor
cast.To[int](true)   // 1
cast.To[int](nil)    // 0

// Negative → unsigned: errors by default; ABS takes the absolute value instead
cast.To[uint](-5)                               // 0 (error silently dropped)
cast.To[uint](-5, cast.Op{cast.ABS, true})      // 5
```

##### String

```go
cast.To[string](8)                  // "8"
cast.To[string](8.31)               // "8.31"
cast.To[string]([]byte("hi"))       // "hi" — string(b) directly, not element-wise
cast.To[string](true)               // "true"
cast.To[string](nil)                // ""

// JSON-encode the result (adds quotes and escaping)
jsonStr, _ := cast.ToE[string](`hello "world"`, cast.Op{cast.JSON, true})
// jsonStr = `"hello \"world\""`
```

#### Slices

The source must be a slice, array, or map; scalar sources always error.

```go
cast.To[[]int]([]string{"1", "2", "3"})           // []int{1, 2, 3}
cast.To[[]string]([]int{1, 2, 3})                 // []string{"1", "2", "3"}
cast.To[[]bool]([]int{1, 0, 1})                   // []bool{true, false, true}

// Map source: map values become slice elements; iteration order is undefined
cast.To[[]string](map[string]int{"a": 1, "b": 2}) // []string{"1", "2"} (order varies)

// UNIQUE_VALUES: deduplicate after conversion, preserving first-seen order
cast.ToE[[]int]([]int{1, 2, 1, 3}, cast.Op{cast.UNIQUE_VALUES, true})
// []int{1, 2, 3}

// LENGTH: pre-allocate backing capacity
cast.ToE[[]int]([]string{"1", "2"}, cast.Op{cast.LENGTH, 100})
```

#### Maps

```go
// map → map: keys and values are individually cast to the target types
cast.ToE[map[string]int](map[string]string{"a": "1", "b": "2"})
// map[string]int{"a": 1, "b": 2}

// struct → map: exported field names become keys; embedded structs are inlined
type Point struct{ X, Y int }
cast.ToE[map[string]any](Point{X: 3, Y: 4})
// map[string]any{"X": 3, "Y": 4}

// slice/array → map: element indices become keys
cast.ToE[map[int]string]([]string{"a", "b", "c"})
// map[int]string{0: "a", 1: "b", 2: "c"}

// Options
cast.ToE[map[string]any](myStruct,
    cast.Op{cast.PRIVATE, true},             // include unexported fields
    cast.Op{cast.STRICT, true},              // error on unconvertible fields
    cast.Op{cast.DUPLICATE_KEY_ERROR, true}, // error on duplicate keys (map→map)
)
```

#### Structs

`ToStruct[T]` / `ToStructE[T]` hydrate a struct from a map or another struct. Field matching is **case-sensitive** on exported field names. Fields with no matching source key retain their zero value.

```go
type Point struct {
    X int
    Y int
}

// From a map: values are cast to each field's type
point, _ := cast.ToStructE[Point](map[string]any{"X": 3, "Y": "4"})
// Point{X: 3, Y: 4}

// From another struct: matched by exported field name; extra source fields are ignored
type Src struct{ X, Y, Z int }
pointFromStruct, _ := cast.ToStructE[Point](Src{X: 10, Y: 20, Z: 30})
// Point{X: 10, Y: 20}  — Z ignored, no matching field

// STRICT: error when source has keys with no matching target field
_, strictErr := cast.ToStructE[Point](
    map[string]any{"X": 1, "Y": 2, "Z": 3},
    cast.Op{cast.STRICT, true},
)
// strictErr != nil — "Z" has no matching field in Point

// Also accessible via the standard To[T] / ToE[T] entry points
pointFromMap := cast.To[Point](map[string]any{"X": 5, "Y": 6})
// Point{X: 5, Y: 6}
```

##### Field key resolution

The lookup key for each target field follows this priority: a `cast:` struct tag first, then the name portion of a `json:` tag (options like `omitempty` are stripped), then the bare field name. A tag value of `"-"` skips the field entirely. The same resolution applies when the source is another struct.

```go
type Config struct {
    Host string `cast:"host"`           // matched by key "host"
    Port int    `json:"port,omitempty"` // matched by key "port"
    Skip string `cast:"-"`             // never populated from any source
}
cfg, _ := cast.ToStructE[Config](map[string]any{"host": "localhost", "port": 8080})
// Config{Host: "localhost", Port: 8080}

// Struct source: source field keys follow the same tag priority
type Src struct {
    Host string `cast:"host"`
    Port int    `json:"port"`
}
cfg2, _ := cast.ToStructE[Config](Src{Host: "db", Port: 5432})
// Config{Host: "db", Port: 5432}
```

##### Unexported fields (`PRIVATE`)

Unexported fields are skipped in both the source and target by default. Set `PRIVATE` to read unexported source fields and hydrate unexported target fields.

```go
type connConfig struct {
    host string
    port int
}

// From a map
conn, _ := cast.ToStructE[connConfig](
    map[string]any{"host": "localhost", "port": 8080},
    cast.Op{cast.PRIVATE, true},
)
// connConfig{host: "localhost", port: 8080}

// From a struct with unexported source fields
type Src struct{ host string; port int }
conn2, _ := cast.ToStructE[connConfig](Src{host: "db", port: 5432}, cast.Op{cast.PRIVATE, true})
// connConfig{host: "db", port: 5432}
```

#### Channels and Funcs

`chan T` returns a buffered channel (buffer size 1) pre-loaded with the cast value. `cast.Func[T]` returns a `func() T` closure.

```go
// Channel
ten := <-cast.To[chan int]("10")  // 10

// Custom buffer size (≥ 1)
intCh, _ := cast.ToE[chan int](42, cast.Op{cast.LENGTH, 10})

// Func: cast.Func[T] is a named type (type Func[T] func() T);
// a named type is required because Go generics do not accept plain function literals
// as type parameters.
intFunc := cast.To[cast.Func[int]]("10")
fmt.Println(intFunc()) // 10

// Nested composite types are supported
sliceCh    := cast.To[chan []int]([]int{1, 2, 3}) // chan []int
nestedFunc := cast.To[cast.Func[chan int]](42)    // func() chan int
```

#### Named types

`time.Time`, `time.Duration`, `net.IP`, `*url.URL`, `*regexp.Regexp`, `*big.Int`, and `*big.Float` are direct cast targets via `To[T]` / `ToE[T]`. See [Named-type targets](#named-type-targets) for the full source compatibility matrix.

```go
// time.Time — string (19 formats tried), int/uint (Unix ns), float (Unix s)
rfcTime, _    := cast.ToE[time.Time]("2024-04-22T12:00:00Z")    // RFC3339
dateTime, _   := cast.ToE[time.Time]("2024-04-22")              // DateOnly
unixNsTime    := cast.To[time.Time](int64(1713787200000000000)) // Unix nanoseconds
unixSecTime   := cast.To[time.Time](float64(1713787200.5))      // Unix seconds (fractional)

// time.Duration — time.ParseDuration syntax, or int/float as nanoseconds
duration, _  := cast.ToE[time.Duration]("1h30m45s")
fiveSeconds  := cast.To[time.Duration](int64(5000000000)) // 5s (as ns)

// net.IP — string (IPv4 or IPv6), uint32 (packed IPv4), []byte (4 or 16 bytes)
localIP, _  := cast.ToE[net.IP]("192.168.1.1")
packedIP    := cast.To[net.IP](uint32(0xC0A80101)) // 192.168.1.1 packed

// *url.URL
pageURL, _ := cast.ToE[*url.URL]("https://example.com/path?q=1")

// *regexp.Regexp
fooPattern, _ := cast.ToE[*regexp.Regexp](`^foo\d+$`)

// *big.Int — string auto-detects base: 0x hex, 0o octal, 0b binary, decimal
hexInt, _    := cast.ToE[*big.Int]("0xFF")      // 255
binaryInt, _ := cast.ToE[*big.Int]("0b1010")    // 10
bigFromInt   := cast.To[*big.Int](int64(12345))

// *big.Float — arbitrary precision; float64 sources are limited to float64 precision
bigPi, _ := cast.ToE[*big.Float]("3.14159265358979323846264338327")
```

#### Interface targets

`error`, `fmt.Stringer`, and `github.com/bdlm/std/v2/errors.Error` can be cast **targets**, but only when the source already implements the interface. The value is returned as-is; no string parsing occurs.

```go
myErr := fmt.Errorf("something failed")
cast.To[error](myErr)      // returns myErr unchanged
cast.ToE[error](myErr)     // (myErr, nil)
cast.ToE[error](42)        // (nil, error) — int does not implement error
cast.ToE[fmt.Stringer](42) // (nil, error) — int does not implement fmt.Stringer
```

#### Error handling

`To` drops errors and returns the zero value. `ToE` returns both the value and the error. Use `DEFAULT` to substitute a custom fallback on error.

```go
cast.To[int]("Hi!")                 // 0  — error silently dropped
result, err := cast.ToE[int]("Hi!") // 0, error: unable to cast "Hi!" of type string to int

// DEFAULT: return a custom value instead of zero on error
result := cast.To[int]("Hi!", cast.Op{cast.DEFAULT, -1})
// result = -1
```
