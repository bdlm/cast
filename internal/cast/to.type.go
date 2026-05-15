// Package cast (internal) contains the private implementation of the cast
// package. It is importable only by `github.com/bdlm/cast/v2`.
package cast

import (
	"fmt"
	"strings"
)

// Sentinel errors and reusable format strings used throughout the package.
var (
	// ErrorUnableToCast is a sentinel error returned when a value cannot be
	// cast to the requested type. It is used as a fallback error message in
	// conversion operations that do not provide more specific error details.
	ErrorUnableToCast = fmt.Errorf("unable to cast value")

	// ErrorSignedToUnsigned is returned when a negative value is cast to an
	// unsigned integer type and the ABS flag is not set.
	ErrorSignedToUnsigned = fmt.Errorf("cannot cast signed value to unsigned integer")

	// ErrorInvalidOption is a format string (not an error value) used to build
	// error messages when an Op flag carries an unexpected value type.
	ErrorInvalidOption = "invalid %s value '%v'"

	// ErrorStrErrorCastingFunc is a format string (not an error value) used
	// when an element cast fails inside a Func[T] closure generator.
	ErrorStrErrorCastingFunc = "error casting %T to %T during function generation"

	// ErrorStrUnableToCast is a format string (not an error value) used when
	// a value cannot be converted to the requested target type.
	ErrorStrUnableToCast = "unable to cast %#.10v of type %T to %T"
)

// Flag is the key type for conversion options passed to the public To and ToE
// functions.
type Flag int

// Op is a single key/value option passed to the public To or ToE function.
type Op struct {
	Flag Flag
	Val  any
}

// Available option flags. Flags fall into two categories:
//
// Global flags propagate through the full conversion tree and are preserved by
// [Ops.Global]. They apply at every level of a nested conversion:
//   - ABS, FORMAT, JSON, LENGTH, PRIVATE, STRICT, UNIQUE_VALUES
//
// Practical consequences of global scope:
//   - ABS=true on []uint applies absolute-value logic to every element.
//   - JSON=true on []string JSON-encodes every string element.
//   - LENGTH=N on chan []int sets both the channel buffer and inner slice capacity.
//   - UNIQUE_VALUES=true on [][]int deduplicates each inner []int independently.
//   - FORMAT on []time.Time applies the custom format to every element parse.
//   - PRIVATE and STRICT apply to nested struct fields at all levels.
//
// Local flags apply only to the conversion they are passed to and are stripped
// by [Ops.Global]. Container converters read their own local flags, then pass
// [Ops.Global] to element-level casts so the local flags do not bleed into
// nested conversions where they carry no meaning or the wrong type:
//   - DECODE, DEFAULT, DUPLICATE_KEY_ERROR
const (
	// DEFAULT: target type T, LOCAL.
	// Value to return on error instead of the zero value for T.
	// Val must be the exact same type as T — a wrong-typed DEFAULT causes an
	// error immediately, before the input is inspected, even for inputs that
	// would otherwise convert successfully.
	// Not passed to nested element-level casts.
	DEFAULT Flag = iota

	// ABS: bool, GLOBAL.
	// When true and the source is negative and the target is uint*, the absolute
	// value is used instead of erroring. Applies to signed integer sources
	// (int, int8, int16, int32, int64), float sources (float32, float64), and
	// numeric string sources that parse to a negative float64.
	ABS

	// DECODE: string, LOCAL.
	// Decode the string source before conversion; only applies to string, error,
	// and fmt.Stringer sources. Only "json" / "JSON" is supported.
	// For scalar targets (bool, int*, uint*, float*, complex*) it fires as a
	// fallback after normal parsing fails (e.g. `"\"42\""` → 42).
	// For []T targets it fires first, forcing JSON decode before element conversion.
	// Not passed to nested casts.
	DECODE

	// DUPLICATE_KEY_ERROR: bool, LOCAL.
	// Error when two source keys cast to the same target key (map→map only).
	// Has no effect on slice→map or struct→map. Not meaningful in nested casts.
	DUPLICATE_KEY_ERROR

	// FORMAT: string, GLOBAL.
	// Custom time.Parse layout for time.Time string/[]byte parsing.
	// When set, the default 19-format trial is skipped entirely and only this
	// format is tried (e.g. "2006/01/02"). Propagates to nested time.Time casts.
	// Note: when FORMAT is not set, the empty string "" is a valid source that
	// returns the zero time.Time{} (time.Parse("", "") succeeds).
	FORMAT

	// JSON: bool, GLOBAL.
	// JSON-encode the resulting string: the plain string representation of the
	// source is wrapped in JSON quotes with proper escaping. Propagates globally,
	// so it affects string elements inside []string, map[K]string, etc.
	JSON

	// LENGTH: int, GLOBAL.
	// Initial capacity for []T targets (0 allowed) or buffer size for chan T
	// targets (must be >= 1). Propagates globally — for chan []int, one LENGTH
	// value sets both the channel buffer size and the inner slice capacity.
	LENGTH

	// PRIVATE: bool, GLOBAL.
	// Include unexported struct fields when reading a struct source (struct→map)
	// or hydrating a struct target (map/struct→struct). Propagates to nested
	// struct conversions at all levels.
	PRIVATE

	// STRICT: bool, GLOBAL.
	// Error instead of silently skipping unconvertible fields or unmatched keys
	// in struct/map conversions. Propagates to nested struct conversions.
	STRICT

	// UNIQUE_VALUES: bool, GLOBAL.
	// Deduplicate slice elements after conversion, preserving first-seen order.
	// Comparable elements use map lookup (O(1)); non-comparable elements
	// (e.g. inner slices) fall back to reflect.DeepEqual. Propagates globally,
	// so every level of a nested slice type is deduplicated independently.
	UNIQUE_VALUES
)

// Ops is the internal parsed representation of conversion options. A plain
// struct is used instead of a map so that the common zero-options path
// allocates nothing and bool-flag checks are plain field reads.
//
// DefaultVal and LengthVal preserve the original Op.Val for type-checking and
// error messages. FormatVal and DecodeVal are stored as pre-parsed strings.
// All bool flags are parsed eagerly by ParseOps.
type Ops struct {
	HasDefault bool
	DefaultVal any // DEFAULT value; meaningful only when HasDefault is true

	HasLength bool
	LengthVal any // LENGTH value preserved for ToE[int] parsing and error messages

	HasFormat bool
	FormatVal string // FORMAT value preserved for time/duration parsing and error messages

	HasDecode bool
	DecodeVal string // DECODE format, normalized to lowercase (e.g. "json")

	Abs        bool
	DupKeyErr  bool
	UniqueVals bool
	JsonEncode bool
	Private    bool
	Strict     bool
}

// Global returns a copy of Ops containing only the flags that apply
// universally across all target types. Local flags (DECODE, DEFAULT,
// DUPLICATE_KEY_ERROR) are dropped; global flags (ABS, FORMAT, JSON, LENGTH,
// PRIVATE, STRICT, UNIQUE_VALUES) are retained.
//
// Container converters call this when passing Ops to element-level casts so
// that a container's own local flags do not leak into nested conversions where
// they carry no meaning or the wrong type.
func (o Ops) Global() Ops {
	return Ops{
		Abs:        o.Abs,
		HasLength:  o.HasLength,
		LengthVal:  o.LengthVal,
		HasFormat:  o.HasFormat,
		FormatVal:  o.FormatVal,
		UniqueVals: o.UniqueVals,
		JsonEncode: o.JsonEncode,
		Private:    o.Private,
		Strict:     o.Strict,
	}
}

// Delete returns a copy of Ops with the given flag cleared.
func (o Ops) Delete(key Flag) Ops {
	switch key {
	case DEFAULT:
		o.HasDefault = false
		o.DefaultVal = nil
	case ABS:
		o.Abs = false
	case DECODE:
		o.HasDecode = false
		o.DecodeVal = ""
	case DUPLICATE_KEY_ERROR:
		o.DupKeyErr = false
	case FORMAT:
		o.HasFormat = false
		o.FormatVal = ""
	case JSON:
		o.JsonEncode = false
	case LENGTH:
		o.HasLength = false
		o.LengthVal = nil
	case PRIVATE:
		o.Private = false
	case STRICT:
		o.Strict = false
	case UNIQUE_VALUES:
		o.UniqueVals = false
	}
	return o
}

// List converts Ops back to a []Op slice for passing to recursive calls
// for element-level casts inside containers.
func (o Ops) List() []Op {
	var list []Op
	if o.HasDefault {
		list = append(list, Op{DEFAULT, o.DefaultVal})
	}
	if o.HasDecode {
		list = append(list, Op{DECODE, o.DecodeVal})
	}
	if o.HasLength {
		list = append(list, Op{LENGTH, o.LengthVal})
	}
	if o.Abs {
		list = append(list, Op{ABS, true})
	}
	if o.DupKeyErr {
		list = append(list, Op{DUPLICATE_KEY_ERROR, true})
	}
	if o.HasFormat {
		list = append(list, Op{FORMAT, o.FormatVal})
	}
	if o.UniqueVals {
		list = append(list, Op{UNIQUE_VALUES, true})
	}
	if o.JsonEncode {
		list = append(list, Op{JSON, true})
	}
	if o.Private {
		list = append(list, Op{PRIVATE, true})
	}
	if o.Strict {
		list = append(list, Op{STRICT, true})
	}
	return list
}

// ParseOps collapses the public variadic []Op into the internal Ops struct
// used by all conversion functions. Bool flags are parsed eagerly; DEFAULT and
// LENGTH preserve their raw Val for type-checking at each call site; FORMAT and
// DECODE are stored as normalized strings.
func ParseOps(o []Op) Ops {
	if len(o) == 0 {
		return Ops{}
	}
	var result Ops
	for _, op := range o {
		switch op.Flag {
		case DEFAULT:
			result.HasDefault = true
			result.DefaultVal = op.Val
		case ABS:
			result.Abs, _ = op.Val.(bool)
		case DECODE:
			if s, ok := op.Val.(string); ok && s != "" {
				result.HasDecode = true
				result.DecodeVal = strings.ToLower(s)
			}
		case DUPLICATE_KEY_ERROR:
			result.DupKeyErr, _ = op.Val.(bool)
		case FORMAT:
			result.HasFormat = true
			result.FormatVal, _ = op.Val.(string)
		case JSON:
			result.JsonEncode, _ = op.Val.(bool)
		case LENGTH:
			result.HasLength = true
			result.LengthVal = op.Val
		case PRIVATE:
			result.Private, _ = op.Val.(bool)
		case STRICT:
			result.Strict, _ = op.Val.(bool)
		case UNIQUE_VALUES:
			result.UniqueVals, _ = op.Val.(bool)
		}
	}
	return result
}

// Integer, Float, and ComplexNum are internal constraints used by the
// per-kind conversion functions. They accept named types with the matching
// underlying type (e.g. type Celsius float32).
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface{ ~float32 | ~float64 }

type ComplexNum interface{ ~complex64 | ~complex128 }
