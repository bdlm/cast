<p align="center">
    <a href="https://gopherize.me/gopher/0b8aa47b088b43d10817e8a13cb115fdd87c0bcb"><img src="https://github.com/bdlm/cast/wiki/assets/images/gopher.png" width="300px"></a>
</p>

# cast

Now with Generics!

<a href="https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#mature"><img src="https://img.shields.io/badge/stability-mature-008000.svg" alt="Mature"></a> This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). This package is considered mature, you should expect package stability in <strong>Minor</strong> and <strong>Patch</strong> version releases

- **Major**: backwards incompatible package updates
- **Minor**: feature additions
- **Patch**: bug fixes, backward compatible model and function changes, etc.

**[CHANGELOG](CHANGELOG.md)**<br>

<p align="center">
    <a href="https://github.com/bdlm/cast/blob/master/CHANGELOG.md"><img src="https://img.shields.io/github/v/release/bdlm/cast" alt="Release"></a>
    <a href="https://pkg.go.dev/github.com/bdlm/cast"><img src="https://godoc.org/github.com/bdlm/cast?status.svg" alt="GoDoc"></a>
    <a href="https://travis-ci.org/bdlm/cast"><img src="https://travis-ci.org/bdlm/cast.svg?branch=master" alt="Build status"></a>
    <a href="https://codecov.io/gh/bdlm/cast"><img src="https://img.shields.io/codecov/c/github/bdlm/cast/master.svg" alt="Coverage status"></a>
    <a href="https://goreportcard.com/report/github.com/bdlm/cast"><img src="https://goreportcard.com/badge/github.com/bdlm/cast" alt="Go Report Card"></a>
    <a href="https://github.com/bdlm/cast/issues"><img src="https://img.shields.io/github/issues-raw/bdlm/cast.svg" alt="Github issues"></a>
    <a href="https://github.com/bdlm/cast/pulls"><img src="https://img.shields.io/github/issues-pr/bdlm/cast.svg" alt="Github pull requests"></a>
    <a href="https://github.com/bdlm/cast/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT"></a>
</p>

<sub>this project is inspired by [`spf13/cast`](https://github.com/spf13/cast)</sub>

## What is Cast?

Cast is a library to convert between different `go` data types in a straigntforward and predictable way.

Cast provides a simple function to easily convert a number to a string, an interface into a bool, etc. Cast does this intelligently when an obvious conversion is possible and logically when a conversion requires additional work, such as casting a map to a slice or a bool to a channel. It doesn’t make any attempts to guess what you meant, for example you can only convert a string to an int when it is a string representation of an int such as “8”. Cast is meant to simplify consumption of untyped or poorly typed data by removing all the typing boilerplate you otherwise need to write.

## Why use Cast?

When working with dynamic data in Go you often need to cast or convert the data from one type into another. Cast goes beyond just using type assertion (though it uses that when possible) to provide a very straightforward and convenient library.

If you are working with interfaces to handle things like dynamic content you need a way to convert an interface into a given type. This is the library for you.

If you are taking in data from YAML, TOML or JSON or other formats which lack full types, then Cast is the library for you.

## Usage

Cast provides `To[T any](any) T` and `ToE[T any](any) (T, error)` methods. These methods will always return the desired type `T`. While Cast will accept `any` type, not all conversions are possible, supportable, or sensible, but several useful and unique conversions are included.

***If input is provided that will not convert to a specified type, the 0 or nil value for that type will be returned***. In order to differentiate between success and the `nil` value, the `ToE` method will return both the cast value and any [errors](https://github.com/bdlm/errors) that occurred during the conversion.

Some conversions allow for flags for finer grained control. For example, you can pass a default value for impossible conversions or specify the size of the backing array when converting to a slice.


### Examples

##### channels
Casting to a channel will return a channel of the specified type with a buffer size of 1 containing the typed value.
```go
intCh := cast.To[chan int]("10")
ten = <-intCh // <-10 (int)

var strCh chan string
strCh = cast.To[chan string](10)
str := <-strCh // <-"10" (string)

boolval := <-cast.To[chan bool](1) // <-true (bool) - I have no idea why you would do that :) but it works
```

##### func
Casting to function will return a function that returns the cast value. This requires using the `cast.Func` type
```go
var intFunc cast.Func[int]
intFunc = cast.To[cast.Func[int]]("10")
fmt.Printf("%#v (%T)\n", intFunc(), intFunc()) // 10 (int)

strFunc := cast.To[cast.Func[string]](10)
fmt.Printf("%#v (%T)\n", strFunc(), strFunc()) // "10" (string)

var boolFunc cast.Func[bool]
boolFunc = cast.To[cast.Func[bool]](1)
fmt.Printf(
    "%#v (%T)\n", // true (bool) - why tho?
    cast.To[cast.Func[bool]](1)(),
    cast.To[cast.Func[bool]](1)(),
)
```

##### String
```go
strVal := cast.To[string]("Hi!")              // "Hi!"
strVal := cast.To[string](8)                  // "8"
strVal := cast.To[string](8.31)               // "8.31"
strVal := cast.To[string]([]byte("one time")) // "one time"
strVal := cast.To[string](nil)                // ""

var foo interface{} = "one more time"
intVal := cast.To[string](foo)                  // "one more time"
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
To capture any conversion errors, use the `ToE` method:
```go
intVal := cast.To[int]("Hi!")         // 0
intVal, err := cast.ToE[int]("Hi!")   // 0, unable to cast "Hi!" of type string to int
```
