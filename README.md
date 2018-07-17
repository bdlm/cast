# cast

<p align="center">
	<a href="https://github.com/bdlm/cast/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT"></a>
	<a href="https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#alpha"><img src="https://img.shields.io/badge/stability-alpha-f4d03f.svg" alt="Alpha"></a>
	<a href="https://travis-ci.org/bdlm/cast"><img src="https://travis-ci.org/bdlm/cast.svg?branch=master" alt="Build status"></a>
	<a href="https://codecov.io/gh/bdlm/cast"><img src="https://img.shields.io/codecov/c/github/bdlm/cast/master.svg" alt="Coverage status"></a>
	<a href="https://goreportcard.com/report/github.com/bdlm/cast"><img src="https://goreportcard.com/badge/github.com/bdlm/cast" alt="Go Report Card"></a>
	<a href="https://github.com/bdlm/cast/issues"><img src="https://img.shields.io/github/issues-raw/bdlm/cast.svg" alt="Github issues"></a>
	<a href="https://github.com/bdlm/cast/pulls"><img src="https://img.shields.io/github/issues-pr/bdlm/cast.svg" alt="Github pull requests"></a>
	<a href="https://godoc.org/github.com/bdlm/cast"><img src="https://godoc.org/github.com/bdlm/cast?status.svg" alt="GoDoc"></a>
</p>

`cast` is forked from [`spf13/cast`](https://github.com/spf13/cast).

## What is Cast?

Cast is a library to convert between different go types in a consistent and easy way.

Cast provides simple functions to easily convert a number to a string, an interface into a bool, etc. Cast does this intelligently when an obvious conversion is possible. It doesn’t make any attempts to guess what you meant, for example you can only convert a string to an int when it is a string representation of an int such as “8”. Cast was developed for use in [Hugo](http://hugo.spf13.com), a website engine which uses YAML, TOML or JSON for meta data.

## Why use Cast?

When working with dynamic data in Go you often need to cast or convert the data from one type into another. Cast goes beyond just using type assertion (though it uses that when possible) to provide a very straightforward and convenient library.

If you are working with interfaces to handle things like dynamic content you’ll need an easy way to convert an interface into a given type. This is the library for you.

If you are taking in data from YAML, TOML or JSON or other formats which lack full types, then Cast is the library for you.

## Usage

Cast provides a handful of To_____ methods. These methods will always return the desired type. *If input is provided that will not convert to that type, the 0 or nil value for that type will be returned*. In addition, an error will be returned which tells you if conversion was successful and differentiate between success and the nil value.

### Examples

`ToString`:
```go
cast.ToString("mayonegg")         // "mayonegg"
cast.ToString(8)                  // "8"
cast.ToString(8.31)               // "8.31"
cast.ToString([]byte("one time")) // "one time"
cast.ToString(nil)                // ""

var foo interface{} = "one more time"
cast.ToString(foo)                // "one more time"
```

`ToInt`
```go
cast.ToInt(8)             // 8
cast.ToInt(8.31)          // 8
cast.ToInt("8")           // 8
cast.ToInt(true)          // 1
cast.ToInt(false)         // 0

var eight interface{} = 8
cast.ToInt(eight)         // 8
cast.ToInt(nil)           // 0
```
