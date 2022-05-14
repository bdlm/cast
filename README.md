# cast

Now with Generics!

<a href="https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#mature"><img src="https://img.shields.io/badge/stability-mature-008000.svg" alt="Mature"></a> This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). This package is considered mature, you should expect package stability in <strong>Minor</strong> and <strong>Patch</strong> version releases

- **Major**: backwards incompatible package updates
- **Minor**: feature additions
- **Patch**: bug fixes, backward compatible model and function changes, etc.

**[CHANGELOG](CHANGELOG.md)**<br>

<a href="https://github.com/bdlm/cast/blob/master/CHANGELOG.md"><img src="https://img.shields.io/github/v/release/bdlm/cast" alt="Release"></a>
<a href="https://pkg.go.dev/github.com/bdlm/cast"><img src="https://godoc.org/github.com/bdlm/cast?status.svg" alt="GoDoc"></a>
<a href="https://travis-ci.org/bdlm/cast"><img src="https://travis-ci.org/bdlm/cast.svg?branch=master" alt="Build status"></a>
<a href="https://codecov.io/gh/bdlm/cast"><img src="https://img.shields.io/codecov/c/github/bdlm/cast/master.svg" alt="Coverage status"></a>
<a href="https://goreportcard.com/report/github.com/bdlm/cast"><img src="https://goreportcard.com/badge/github.com/bdlm/cast" alt="Go Report Card"></a>
<a href="https://github.com/bdlm/cast/issues"><img src="https://img.shields.io/github/issues-raw/bdlm/cast.svg" alt="Github issues"></a>
<a href="https://github.com/bdlm/cast/pulls"><img src="https://img.shields.io/github/issues-pr/bdlm/cast.svg" alt="Github pull requests"></a>
<a href="https://github.com/bdlm/cast/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT"></a>

`cast` is inspired by [`spf13/cast`](https://github.com/spf13/cast).

## What is Cast?

Cast is a library to convert between different go types in a consistent and straigntforward way.

Cast provides a simple function to easily convert a number to a string, an interface into a bool, etc. Cast does this intelligently when an obvious conversion is possible. It doesn’t make any attempts to guess what you meant, for example you can only convert a string to an int when it is a string representation of an int such as “8”. Cast is meant to simplify consumption of untyped or poorly typed data by removing all the typing boilerplate you otherwise need to write.

## Why use Cast?

When working with dynamic data in Go you often need to cast or convert the data from one type into another. Cast goes beyond just using type assertion (though it uses that when possible) to provide a very straightforward and convenient library.

If you are working with interfaces to handle things like dynamic content you need a way to convert an interface into a given type. This is the library for you.

If you are taking in data from YAML, TOML or JSON or other formats which lack full types, then Cast is the library for you.

## Usage

Cast provides `To[T any](any) T` and `ToE[T any](any) (T error)` methods. These methods will always return the desired type. *If input is provided that will not convert to that type, the 0 or nil value for that type will be returned*. In order to differentiate between success and the `nil` value, the `ToE` method will return both the cast value and any error that occurred during the conversion.

While Cast will accept `any` type, not all conversions are possible or supportable, but several useful and unique conversions are provided.

### Examples

##### channels
Casting to a channel will return a channel of the specified type with a buffer size of 1 containing the typed value.
```go
var intCh chan int
intCh = cast.To[chan int]("10")  // 10

var strCh chan string
strCh = cast.To[chan string](10) // "10"

var boolCh chan bool
boolCh = cast.To[chan bool](1)   // true
```

##### func
Casting to function will return a function that returns the cast value. This requires using the `cast.Func` type
```go
var intFunc cast.Func[int]
intFunc = cast.ToE[cast.Func[int]]("10")
fmt.Printf("%#v (%T)\n", intFunc(), intFunc()) // 10 (int)

var strFunc cast.Func[string]
strFunc = cast.ToE[cast.Func[string]](10)
fmt.Printf("%#v (%T)\n", strFunc(), strFunc()) // "10" (string)

var boolFunc cast.Func[bool]
boolFunc = cast.ToE[cast.Func[bool]](1)
fmt.Printf("%#v (%T)\n", boolFunc(), boolFunc()) // true (bool)

```

##### String
```go
intVal := cast.To[string]("Hi!")              // "Hi!"
intVal := cast.To[string](8)                  // "8"
intVal := cast.To[string](8.31)               // "8.31"
intVal := cast.To[string]([]byte("one time")) // "one time"
intVal := cast.To[string](nil)                // ""

var foo interface{} = "one more time"
intVal := cast.ToString(foo)                  // "one more time"
```

##### Int
```go
intVal := cast.To[int](8)           // 8
intVal := cast.To[int](8.31)        // 8
intVal := cast.To[int]("8")         // 8
intVal := cast.To[int]("8.31")      // 8
intVal := cast.To[int]("8.51")      // 8
intVal := cast.To[int](true)        // 1
intVal := cast.To[int](false)       // 0

var eight interface{} = 8
intVal := cast.To[int](eight)       // 8
intVal := cast.To[int](nil)         // 0
```

##### Error checking
```go
intVal := cast.To[int]("Hi!")         // 0
intVal, err := cast.ToE[int]("Hi!")   // 0, unable to cast "Hi!" of type string to int
```
