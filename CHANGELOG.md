All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

- **Major**: backwards incompatible package updates
- **Minor**: feature additions
- **Patch**: bug fixes, backward compatible model and function changes, etc.

## v1.2.0 - 2022-06-??
This is a full library rewrite for go v1.18 to take advantage of [generic functions and types](https://go.dev/doc/tutorial/generics).
#### Removed
- All existing exported cast functions have been removed (`ToString(any) string`, `ToStringE(any) (string, error)`, etc.)

#### Added
- All previous exported cast functions have replaced with a single generic function (and it's `error` counterpart):
  ```go
  To[T any]() T

  ToE[T any]() (T, error)
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
