All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

- **Major**: backwards incompatible package updates
- **Minor**: feature additions
- **Patch**: bug fixes, backward compatible model and function changes, etc.


## v2.0.3 - 2024-02-22
* GitHub action definition for builds and tests
* Related bug fixes and cleanup


## v2.0.2 - 2023-12-30
* Added test coverage
* Related bug fixes and cleanup


## v2.0.1 - 2023-12-28
* Additional examples
* Improved slice support
* Various bugfixes
* Added tests


## v2.0.0 - 2023-12-23
This is a full library rewrite for go v1.18+ to take advantage of [generic functions and types](https://go.dev/doc/tutorial/generics).

syntax example:
```go
intVal := cast.To[int]("8.31")      // 8 (int)
intVal := cast.To[int]("Hi!")       // 0 (int)
intVal, err := cast.ToE[int]("Hi!") // 0, unable to cast "Hi!" of type string to int (int, error)
```
#### Removed
- All existing exported cast functions have been removed (`ToString(any) string`, `ToStringE(any) (string, error)`, etc.)

#### Added
- All previous exported cast functions have replaced with a single generic function (and it's `error` counterpart):
  ```go
  To[T any](any) T

  ToE[T any](any) (T, error)
  ```


## v1.1.0 - 2020-06-27
#### Changed
- Refactoring `ToSlice*` and `ToMap` language

## v1.0.2 - 2020-06-25
#### Added
- `ToUint64Slice`
- `ToUint64SliceE`

## v1.0.1 - 2020-06-25
#### Added
- `ToInt64Slice`
- `ToInt64SliceE`

### v1.0.0 - 2020-05-02
`v1.0.0` is the production release of the previous development work.
