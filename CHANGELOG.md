All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

- **Major**: backwards incompatible package updates
- **Minor**: feature additions
- **Patch**: bug fixes, backward compatible model and function changes, etc.

# Versions

- [v2.1.4 - 2026-05-14](#v214---2026-05-14) _(Go 1.21+)_
- [v2.1.3 - 2026-05-13](#v213---2026-05-13) _(Go 1.21+)_
- [v2.1.2 - 2026-05-13](#v212---2026-05-13) _(Go 1.21+)_
- [v2.1.1 - 2026-05-04](#v211---2026-05-04) _(Go 1.21+)_
- [v2.1.0 - 2026-04-20](#v210---2026-04-20) _(Go 1.21+)_
- [v2.0.5 - 2025-11-12](#v205---2025-11-12) _(Go 1.24+)_
- [v2.0.4 - 2025-11-03](#v204---2025-11-03) _(Go 1.24+)_
- [v2.0.3 - 2024-02-22](#v203---2024-02-22) _(Go 1.21+)_
- [v2.0.2 - 2023-12-30](#v202---2023-12-30) _(Go 1.18+)_
- [v2.0.1 - 2023-12-28](#v201---2023-12-28) _(Go 1.18+)_
- [v2.0.0 - 2023-12-23](#v200---2023-12-23) _(Go 1.18+)_
- [v1.0.2 - 2020-06-25](#v102---2020-06-25) _(Go 1.14+)_
- [v1.0.1 - 2020-06-25](#v101---2020-06-25) _(Go 1.14+)_
- [v1.0.0 - 2020-05-02](#v100---2020-05-02) _(Go 1.14+)_


# v2.1.4 - 2026-05-14
Internal package refactoring, `ToStruct`/`ToStructE` removal, documentation corrections, flag-behavior fixes, and substantially expanded test and example coverage.

### Removed

- **`ToStruct[T]`** — public API removed. Use `cast.To[T]` instead; struct hydration is dispatched automatically when `T` is a struct type.
- **`ToStructE[T]`** — public API removed. Use `cast.ToE[T]` instead. The implementation moved to `internal/cast` as `ToStructE`.
- README references to `cast.ToStruct[T]` / `cast.ToStructE[T]` replaced with `cast.To[T]` / `cast.ToE[T]`.

### Changed

#### `internal/cast/` package
Implementation moved out of the public package surface:
- `to.go` and `to.type.go` remain in `package cast` and expose `To`, `ToE`, and the public type API.
- All converter sources (`to.bool.go`, `to.int.go`, `to.float.go`, `to.complex.go`, `to.string.go`, `to.slice.go`, `to.map.go`, `to.chan.go`, `to.func.go`, `to.struct.go`, `to.time.go`, `to.duration.go`, `to.big.go`, `to.net.go`, `to.url.go`, `to.regexp.go`, `util.decode.go`, `util.reflect.go`) moved to `internal/cast/`.
- All `*_test.go` files moved to `internal/cast/` except `test.examples_test.go`, which stays in main to validate the public API surface via godoc examples.
- Non-generic public types (`Flag`, `Op`, flag constants, error vars) are defined in the internal package; main package re-exports them via Go type aliases. Generic types (`Func`, `Types`, `Tbase`, `Tslice`, `Tchan`, `Tmap`) stay in main because Go 1.21 does not support generic type aliases.
- Internal `ops` struct renamed to `Ops` with exported fields so the main package can construct it via `internal.ParseOps`.
- `makeChan`/`makeFunc` rewritten to use reflection-based casting (`CastToType`) instead of recursively calling the public `ToE`, breaking the otherwise-circular import path between main and internal.
- `ToChan` and `ToFunc` are non-generic in the internal package; main's `ToE` performs the final type assertion (with `reflect.Convert` fallback) to produce the user's named `chan T` or `Func[T]`.

#### Options documentation
- Added `FORMAT` row to the README Options table (previously absent).
- Added Type and Scope columns to the Options table; reordered rows to match the constant declaration order.
- Added a "Global flag propagation" subsection below the table with runnable examples showing `ABS`, `JSON`, `LENGTH`, `UNIQUE_VALUES`, and `FORMAT` propagating into nested casts.
- All flag constants in `internal/cast/to.type.go` now carry full multi-line doc comments stating their value type, scope, applicability, and propagation behavior.

#### Example function naming convention
Happy-path `ExampleToE_*` functions that never showed an error have been renamed to `ExampleTo_*` and switched to `cast.To`; the freed-up `ExampleToE_*` names now hold corresponding error examples. Every `ExampleTo_*` has a matching `ExampleToE_*` error example. Total example functions: 73.

### Added

#### Test coverage for the conversion table
33 new sub-tests cover previously-untested `✗` cells of the conversion table: scalar → map errors, scalar → struct errors, slice/map/struct → bool|complex errors, map → int/uint/float errors, struct → uint error, and the documented `[]byte`/`[]rune` → `map[K]V` success path. New ABS coverage added for negative `float64` sources and negative-float strings to `uint` targets.

#### Godoc example functions
24 new example functions covering previously-undocumented type categories: `time.Time` (string, int, float, `FORMAT`), `time.Duration` (string, int), `net.IP` (string, packed `uint32`), `*url.URL`, `*regexp.Regexp`, `*big.Int` (decimal, hex), `*big.Float`, `bool`, `complex64`/`complex128`, `error` and `fmt.Stringer` interface targets, `DECODE` (scalar and slice), `[]byte` → `string`, JSON-string → map, nested `chan []int` and `Func[chan int]`.

### Fixed

#### `time.Time` ↔ `*big.Int` round-trip
`time.Time → *big.Int` now uses `val.UnixNano()` (was `val.Unix()`). `*big.Int → time.Time` already used nanoseconds, so the round-trip is now lossless for times within `int64` range. Three test cases in `TestTimeToBigConversions` were updated accordingly.

#### Documentation conflicts
- Named-type table footnote ᵗ said `*big.Int → time.Time` was Unix seconds; corrected to nanoseconds. `*big.Float → time.Time` remains documented as Unix seconds with fractional precision.
- Supported Conversions table cell for `string → map[K]V` was `✗`, but JSON-object/array strings auto-decode; corrected to `~ⁿ` with a new footnote.
- Note ⁴ for `string → bool` said only `strconv.ParseBool` variants are accepted; the integer-parse fallback path was undocumented and is now described (`"-1"` → true, `"0.1"` → false because `floor(0.1) = 0`, etc.).
- Note ᵈ and the named-type example for `int*` → `time.Time` claimed Unix seconds; corrected to nanoseconds. The `toTime` docstring previously said "Unix seconds" for integer sources; fixed to "Unix nanoseconds".

#### Flag documentation
- `FORMAT` constant comment said "time/duration parsing"; only `time.Time` reads `ops.FormatVal`. Corrected.
- `PRIVATE` constant comment said "include unexported struct fields in map output" — too narrow. Corrected to describe both directions (struct→map source-field reading and map/struct→struct target hydration).
- README `ABS` row said "negative signed inputs"; `ABS` also applies to negative `float32`/`float64` sources and numeric strings that parse to negative floats. Description, example block, and IMPORTANT callout updated.

#### `parseTimeString` undocumented edge case
When `FORMAT` is not set, `parseTimeString` falls through to `time.Parse("", str)` after the 19-format loop. The empty string `""` is a valid source that returns `time.Time{}` because `time.Parse("", "")` succeeds. Now documented in both the function godoc and the README footnote ᶠ.

#### Wrong-typed `DEFAULT` errors immediately
Every converter type-asserts the `DEFAULT` value at the top, before inspecting the input. Passing a wrong type returns an error immediately even for inputs that would otherwise convert successfully. Now documented on the `Op` struct, the `DEFAULT` constant, and the README Options table.

### Documentation (README.md)
- `IMPORTANT` callout updated to describe the actual `string → bool` two-step parse behavior and to mention float-source applicability for `ABS`.
- Named-type code example updated to use a meaningful Unix-nanosecond value (`int64(1_713_787_200_000_000_000)` for 2024-04-22) instead of a value that only made sense as Unix seconds.
- The struct-hydration example block under `#### Structs` rewritten to use `cast.To` / `cast.ToE` after the removal of `ToStruct` / `ToStructE`.

# v2.1.3 - 2026-05-13
Code quality, documentation, and test improvements.

### Added

- Doc comments added to previously undocumented exported and internal symbols:
  - `Tmap` type
  - `ErrorSignedToUnsigned`, `ErrorInvalidOption`, `ErrorStrErrorCastingFunc`, `ErrorStrUnableToCast` (clarifying the string vars are format strings, not error values)
  - `castToType` — all 8 dispatch cases documented
- New godoc examples: `ExampleTo_map`, `ExampleTo_struct`, `ExampleToE_mapDuplicateKeyError`, `ExampleToE_structNested`, `ExampleToE_structStrict`, `ExampleToE_mapToPrivateStruct`.

### Changed

- Renamed the sentinel error variable `Error` to `ErrorUnableToCast`. `Error` is retained as a deprecated alias with a `// Deprecated:` godoc annotation for backward compatibility.
- Example function suffixes converted from `snake_case` to `camelCase` (e.g. `ExampleToE_mapFromMap`) so all examples appear in godoc. Suffixes containing underscores are silently dropped by godoc.

### Fixed

- `ToE` doc referenced non-existent type `Ops` (correct type is `Op`) and contained double word "the the" — corrected.
- `ToStructE` doc omitted `*struct` as a valid source type — corrected.
- `toSlice` doc was circular and meaningless — rewritten.
- `ops.Global` doc omitted `FORMAT` from the global flag list and `DECODE` from the local flag list — corrected.
- `parseOps` doc omitted `FORMAT` and `DECODE` from its description of preserved values — corrected.
- `ops` struct doc incorrectly stated all non-default flags are pre-parsed to `bool`; `FORMAT` and `DECODE` are stored as strings — corrected.
- `TestPointerDerefLoop` "pointer-to-interface" sub-test used `errors.New` whose concrete pointer type is opaque; replaced with a local `ptrReceiverError` type that has `Error()` on `*T` only, making the pointer-receiver guard explicit and self-documenting.

# v2.1.2 - 2026-05-13
Code quality, bug fixes, and internal refactoring.

### Added

#### Tag-aware field key resolution (`util.reflect.go`)
`fieldKey` now resolves a struct field's source-map key with the following priority: `cast` tag → `json` tag (name portion only) → field name. A tag value of `"-"` causes the field to be skipped during both struct hydration and struct→map conversion. This applies uniformly to `hydrateStruct`, `collectSourceFieldValues`, `collectStructFields`, and `collectExportedFields`.

#### `PRIVATE` flag for struct hydration (`to.struct.go`)
`ToStructE` / `toStruct` now support the `PRIVATE` flag. When set, unexported fields are included in both source collection and target hydration. Unexported fields are read via `extractFieldValue`; unexported target fields are set via `unsafe.Pointer` (the only mechanism available without CGo).

### Changed

#### Reflection utilities consolidated into `util.reflect.go`
Four internal helpers that had grown beyond their origin files are now defined exclusively in `util.reflect.go`:
- `fieldKey` — moved from `to.struct.go`
- `extractFieldValue` — moved from `to.map.go`
- `isNamedScalarStructType` — moved from `to.slice.go`
- `isScalarKind` — moved from `to.slice.go`

#### Pointer source dereferencing in `ToE` (`to.go`)
The pointer-unwrapping loop now also dereferences pointer-to-struct sources when the target type is a struct. This makes `*myStruct` → `myStruct` consistent with the existing `*int` → `int` behavior. Pointer-to-interface sources and struct pointers whose target is not a struct (e.g. `*regexp.Regexp` when casting to `*regexp.Regexp`, or `*errorT` when casting to `error`) are left as-is so that pointer-receiver interface implementations continue to work correctly.

### Fixed
- **Multi-level nil pointer overwrites `nil`** (`to.go`): when unwrapping a pointer chain such as `**T` where the outer pointer is non-nil but the inner `*T` is nil, `changed` was left `true`, causing the post-loop `val = srcVal.Interface()` assignment to overwrite the `val = nil` set during the nil check with a typed nil `(*T)(nil)`. Fixed by resetting `changed = false` in the nil branch.
- **Dead `err` variable in `strToFloat`** (`to.float.go`): the final `return TTo(val), err` always returned a nil `err` — the variable is cleared before reaching that line. Changed to `return TTo(val), nil`.



# v2.1.1 - 2026-05-04
Struct hydration, seven new named-type cast targets, pointer dereferencing, and reflection infrastructure improvements.

### Added

#### Struct hydration (`to.struct.go`)
Any map, struct, or `*struct` can now be cast into a user-defined struct type via `ToStruct[T]` / `ToStructE[T]`, or via the standard `To[T]` / `ToE[T]` entry points. Source map keys are matched case-sensitively to exported field names. Fields whose source value cannot be cast retain their zero value by default; the `STRICT` flag promotes mismatches and unknown keys to errors. Supported sources:
- `map[string]any` and any map whose keys are string-castable
- struct or `*struct` (exported field names become keys; anonymous/embedded fields are promoted)

Nested structs, slices of structs, and embedded (anonymous) struct fields — including those from unexported embedded types and nil embedded pointer fields — are all handled recursively.

#### Named-type cast targets
Seven standard Go types are now first-class cast targets via `To[T]` / `ToE[T]`:

| Target | File | Sources |
|---|---|---|
| `time.Time` | `to.time.go` | string (19 formats), `time.Time`, `*time.Time`, integer (Unix ns), float (Unix s) |
| `time.Duration` | `to.duration.go` | `time.Duration`, string (`time.ParseDuration`), integer/float (nanoseconds) |
| `net.IP` | `to.net.go` | `net.IP`, string, `[]byte` (4 or 16 bytes), `uint32` (packed IPv4) |
| `*url.URL` | `to.url.go` | `*url.URL`, `url.URL`, string |
| `*regexp.Regexp` | `to.regexp.go` | `*regexp.Regexp`, string |
| `*big.Int` | `to.big.go` | `*big.Int`, `big.Int`, `*big.Float`, string (base auto-detect), integer types, float types |
| `*big.Float` | `to.big.go` | `*big.Float`, `big.Float`, `*big.Int`, string, integer types, float types |

All converters support the `DEFAULT` op (return the supplied fallback on error) and a `default:` string-cast fallback path for unrecognized source types.

All new targets are also supported as struct field types during struct hydration.

#### Named-type converter table and `rawToValue` helper (`util.reflect.go`)
A single `namedConverters map[reflect.Type]func(any, ops)(any, error)` table is the authoritative registry for all named-type converters. Both `ToE` and `castToType` consult it before the generic kind dispatch, eliminating previously duplicated switch blocks. The companion `rawToValue` helper ensures that `DEFAULT` values are propagated to callers on converter failure instead of being silently dropped.

#### Extended reflection infrastructure (`util.reflect.go`)
`castToType` now handles `reflect.Struct` and `reflect.Pointer` kinds in addition to scalars, slices, funcs, and chans, enabling recursive hydration of arbitrary nested types during struct field casting.

### Fixed
- Nil `*T` anonymous (embedded) pointer fields are now allocated before recursion in `hydrateStruct`, so promoted fields are properly hydrated instead of being skipped.
- `collectExportedFields` now recurses into unexported anonymous struct types (matching `hydrateStruct` semantics), fixing a gap where exported fields within unexported embedded types were missed during struct→struct conversion.
- `DEFAULT` values supplied to `castToStructType` are now correctly propagated to the caller on error instead of being discarded.
- `toStruct` no longer allocates a zero-value struct for `defaultVal` when a `DEFAULT` op is provided.

### Tests
Full test coverage added for all new functionality: `to.struct_test.go`, `to.time_test.go`, `to.duration_test.go`, `to.net_test.go`, `to.url_test.go`, `to.regexp_test.go`, `to.big_test.go`.



# v2.1.0 - 2026-04-20
Map target implementations, extended channel targets, extended function targets, expanded type definitions, performance improvements, expanded test coverage.

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

#### Extended func targets
`toFunc` now handles composite return types in addition to scalars:
- `Func[[]T]` — closure returning a slice
- `Func[chan T]` — closure returning a channel
- `Func[chan []T]` — closure returning a channel of slices
- `Func[chan Func[T]]` — closure returning a channel of closures
- `Func[chan chan T]` — closure returning a channel of channels

#### Reflection helpers (`to.reflect.go`)
New internal package-level functions shared by `toMap`, `toChan`, and `toFunc`:
- `castToKind` — cast any value to a scalar `reflect.Kind`
- `castToType` — cast any value to an arbitrary `reflect.Type`
  (handles interface, slice, func, chan, and scalar targets)
- `castToSliceType` — element-wise cast to a named or concrete slice type

#### Type system (`to.type.go`)
- Added `~chan chan Tbase` to the `Tchan` constraint, enabling `chan chan T`
  as a concrete target type.
- Added `PRIVATE` and `STRICT` flag constants (see Map section above).

### Changed

#### Expanded version support
- Expand support for Go v1.21 to current.

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



# v2.0.5 - 2025-11-12
* []byte to string conversion bugfix



# v2.0.4 - 2025-11-03
* Remove debug code
* Upgrade to go v1.24
* Cleanup non-constant format strings



# v2.0.3 - 2024-02-22
* GitHub action definition for builds and tests
* Related bug fixes and cleanup



# v2.0.2 - 2023-12-30
* Added test coverage
* Related bug fixes and cleanup



# v2.0.1 - 2023-12-28
* Additional examples
* Improved slice support
* Various bugfixes
* Added tests



# v2.0.0 - 2023-12-23
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


# v1.0.2 - 2020-06-25
#### Added
- `ToUint64Slice`
- `ToUint64SliceE`


# v1.0.1 - 2020-06-25
#### Added
- `ToInt64Slice`
- `ToInt64SliceE`


### v1.0.0 - 2020-05-02
`v1.0.0` is the production release of the previous development work.
