All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

- **Major**: backwards incompatible package updates
- **Minor**: feature additions
- **Patch**: bug fixes, backward compatible model and function changes, etc.


## v2.1.0 - 2026-04-17

### Added

#### Map target (`map[K]V`)
- `toMap` fully implemented; three source kinds are supported:
  - **map → map**: keys and values are individually cast to the target types;
    duplicate key detection is opt-in via the new `DUPLICATE_KEY_ERROR` flag.
  - **struct / *struct → map**: exported field names become keys; anonymous
    (embedded) struct fields are promoted to the top level; nested structs
    recurse into `map[K]any` or `map[K]map[...]` when the value type allows.
    Unexported fields are included when `PRIVATE=true`.
  - **slice / array → map**: element indices become keys, keys and values are cast to
    the target value types.
- New `Op` flags:
  - `DUPLICATE_KEY_ERROR` (`bool`, default `false`) — error on duplicate key
    after casting (map→map only).
  - `PRIVATE` (`bool`, default `false`) — include unexported struct fields
    (struct→map only).
  - `STRICT` (`bool`, default `false`) — return an error instead of silently
    skipping unconvertible fields (struct→map only).

#### Extended channel targets
`toChan` now handles composite element types in addition to scalars:
- `chan []T` — channel of slices
- `chan Func[T]` — channel of closures
- `chan chan T` — channel of channels
- `chan [N]T` — channel of fixed-size arrays (reflection path; N known only at
  runtime)

#### Extended func targets
`toFunc` now handles composite return types in addition to scalars:
- `Func[[]T]` — closure returning a slice
- `Func[[N]T]` — closure returning a fixed-size array (reflection path)
- `Func[chan T]` — closure returning a channel
- `Func[chan []T]` — closure returning a channel of slices
- `Func[chan Func[T]]` — closure returning a channel of closures
- `Func[chan chan T]` — closure returning a channel of channels
- `Func[chan [N]T]` — closure returning a channel of arrays (reflection path)

#### Reflection helpers (`to.reflect.go`)
New internal package-level functions shared by `toMap`, `toChan`, and `toFunc`:
- `castToKind` — cast any value to a scalar `reflect.Kind`
- `castToType` — cast any value to an arbitrary `reflect.Type`
  (handles interface, slice, and scalar targets)
- `castToSliceType` — element-wise cast to a slice type
- `castToArray` — cast a slice/array source to a fixed-size array type
- `makeArrayChan` — build and pre-load a `chan [N]T` at runtime
- `makeArrayFunc` — build a `func() [N]T` closure at runtime
- `makeArrayChanFunc` — build a `func() chan [N]T` closure at runtime

#### Type system (`to.type.go`)
- Added `~chan chan Tbase` to the `Tchan` constraint, enabling `chan chan T`
  as a concrete target type.
- Added `PRIVATE` and `STRICT` flag constants (see Map section above).

### Changed

#### `ToE` dispatch (`to.go`)
- Uncommented and wired the `reflect.Map` case to `toMap`.
- `reflect.Array` and `reflect.Slice` cases now both accept array sources
  (previously only slice sources were accepted for slice targets).
- Replaced incorrect goroutine-based panic recovery with `defer/recover`.
- Removed an unnecessary intermediate `reflect.Value` (`from`) that wrapped
  `val`; internal helpers now receive `val` directly, eliminating a layer of
  reflection indirection and fixing type-display in error messages.

#### `toChan` refactor (`to.chan.go`)
- Replaced per-type inline make+send blocks with a shared `makeChan[T]`
  generic helper (cast, make, send, return).

#### `toFunc` refactor (`to.func.go`)
- Replaced per-type inline logic with calls to `makeFunc[T]` and the new
  array/chan reflection helpers.

### Fixed
- **`error` / `fmt.Stringer` interface target documentation**: the README
  incorrectly claimed that any source value could be cast to `error` by
  converting it to a string message. The actual behavior — and the behavior
  validated by `to.interface_test.go` — is that values are accepted only when
  they already implement the target interface, and are returned as-is.
  Documentation now accurately describes pass-through semantics.

### Tests
- `to.map_test.go` — 21 new tests covering map→map (including
  `DUPLICATE_KEY_ERROR`), struct→map (including `PRIVATE`, `STRICT`, nested
  struct recursion, embedded field promotion, pointer source), and
  slice/array→map.
- `to.interface_test.go` — 11 new tests for `error`, `std_error.Error`, and
  `fmt.Stringer` interface targets, documenting pass-through semantics.
- `examples_test.go` — 9 new runnable examples: `To[[]int]`, `ToE[[]int]`
  with `UNIQUE_VALUES`, `To[chan int]`, `ToE[chan int]` with `LENGTH`,
  `To[Func[int]]`, `ToE[map[string]int]` from map/struct/slice, and
  `ToE[string]` with `JSON`.

### Documentation (`README.md`)
- Corrected function signatures: `To[T Types](v any, o ...Op) T` and
  `ToE[T Types](v any, o ...Op) (T, error)` (previously showed wrong `any`
  constraint and omitted the variadic `Op` parameter).
- Added Options section with a full flag reference table (all 8 flags, their
  types, defaults, applicable targets, and descriptions).
- Added examples for slice, map, channel, func, string with `JSON`, int with
  `ABS`, and `error`/`fmt.Stringer` interface targets.
- Corrected the conversion matrix: `slice→map` and `map→map` now show `y`;
  new `struct` source row added.
- Removed incorrect claim that "any source value can be cast to `error`".
- Fixed several prose typos (straightforward, loosely, incoming, maintained,
  measurable).


## v2.0.5 - 2025-11-12
* []byte to string conversion bugfix


## v2.0.4 - 2025-11-03
* Remove debug code
* Upgrade to go v1.24
* Cleanup non-constant format strings


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
