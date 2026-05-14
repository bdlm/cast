package cast

import (
	"encoding"
	"fmt"
	"math/big"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// namedConverters maps specific reflect.Types to their dedicated converter.
// Both ToE (to.go) and castToType consult this table before the kind dispatch
// so that supporting a new named type requires only one edit here.
var namedConverters = map[reflect.Type]func(any, ops) (any, error){
	reflect.TypeOf(time.Duration(0)):      toDuration,
	reflect.TypeOf(time.Time{}):           toTime,
	reflect.TypeOf(net.IP(nil)):           toNetIP,
	reflect.TypeOf((*url.URL)(nil)):       toURL,
	reflect.TypeOf((*regexp.Regexp)(nil)): toRegexp,
	reflect.TypeOf((*big.Int)(nil)):       toBigInt,
	reflect.TypeOf((*big.Float)(nil)):     toBigFloat,
}

// namedStructTypes is a pre-computed set of struct types (and struct types
// whose pointer is registered in namedConverters) that have dedicated scalar
// converters. Using a separate set avoids an init cycle: tryDecodeJSON calls
// isNamedScalarStructType, which would otherwise access namedConverters,
// creating a cycle through the converter functions stored in that map.
// Keep in sync with namedConverters above.
var namedStructTypes = map[reflect.Type]struct{}{
	reflect.TypeOf(time.Time{}):     {},
	reflect.TypeOf(big.Int{}):       {},
	reflect.TypeOf(big.Float{}):     {},
	reflect.TypeOf(url.URL{}):       {},
	reflect.TypeOf(regexp.Regexp{}): {},
}

// rawToValue converts the (raw, err) pair returned by a named-type converter
// into the (reflect.Value, error) pair expected by castToType. On error, if
// ops carries a DEFAULT value assignable to t, it is returned alongside the
// error so the caller can propagate both.
func rawToValue(raw any, t reflect.Type, ops ops, err error) (reflect.Value, error) {
	if err == nil {
		if raw == nil {
			return reflect.Zero(t), nil
		}
		return reflect.ValueOf(raw), nil
	}
	if ops.hasDefault {
		dv := reflect.ValueOf(ops.defaultVal)
		if dv.IsValid() && dv.Type().AssignableTo(t) {
			return dv, err
		}
	}
	return reflect.Value{}, err
}

// castToKind casts v to the scalar Go type corresponding to kind and returns
// the result as a reflect.Value. It only handles scalar kinds; for slices,
// funcs, or chans use [castToType] instead.
func castToKind(v any, kind reflect.Kind, ops ops) (reflect.Value, error) {
	switch kind {
	case reflect.Interface:
		if v == nil {
			return reflect.Zero(reflect.TypeOf((*any)(nil)).Elem()), nil
		}
		return reflect.ValueOf(v), nil
	case reflect.Bool:
		r, err := toBool[bool](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Int:
		r, err := toInt[int](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Int8:
		r, err := toInt[int8](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Int16:
		r, err := toInt[int16](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Int32:
		r, err := toInt[int32](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Int64:
		r, err := toInt[int64](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Uint:
		r, err := toInt[uint](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Uint8:
		r, err := toInt[uint8](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Uint16:
		r, err := toInt[uint16](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Uint32:
		r, err := toInt[uint32](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Uint64:
		r, err := toInt[uint64](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Uintptr:
		r, err := toInt[uintptr](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Float32:
		r, err := toFloat[float32](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Float64:
		r, err := toFloat[float64](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Complex64:
		r, err := toComplex[complex64](v, ops)
		return reflect.ValueOf(r), err
	case reflect.Complex128:
		r, err := toComplex[complex128](v, ops)
		return reflect.ValueOf(r), err
	case reflect.String:
		s, err := toString(v, ops)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(s), nil
	}
	return reflect.Value{}, fmt.Errorf("unsupported kind %v", kind)
}

// castToType casts v to the type t and returns the result as a reflect.Value.
// Dispatch order:
//  1. Named types — checked in namedConverters before the kind switch.
//  2. Interface — assignability-checked and returned as-is.
//  3. Map, Slice — delegated to toMap / toSlice.
//  4. Func — requires zero-arg, one-return signature; return type is cast recursively.
//  5. Chan — element is cast recursively; a buffered channel of the target type is returned.
//  6. Pointer — element type is cast recursively; a new pointer is allocated.
//  7. Struct — delegated to castToStructType.
//  8. Scalars — delegated to castToKind with a ConvertibleTo fallback for named types.
func castToType(v any, t reflect.Type, ops ops) (reflect.Value, error) {
	// Named types have dedicated converters; consult the table before the
	// generic kind dispatch below so that adding a new named type only
	// requires a single entry in namedConverters.
	if fn, ok := namedConverters[t]; ok {
		raw, err := fn(v, ops)
		return rawToValue(raw, t, ops, err)
	}

	switch t.Kind() {
	case reflect.Interface:
		if v == nil {
			return reflect.Zero(t), nil
		}
		src := reflect.ValueOf(v)
		if !src.Type().AssignableTo(t) {
			return reflect.Value{}, fmt.Errorf(ErrorStrUnableToCast, v, v, t)
		}
		return src, nil
	case reflect.Map:
		raw, err := toMap(reflect.Zero(t), v, ops)
		return rawToValue(raw, t, ops, err)
	case reflect.Slice:
		raw, err := toSlice(reflect.Zero(t), v, ops)
		return rawToValue(raw, t, ops, err)
	case reflect.Func:
		// Only zero-arg, one-return functions are supported (matches Func[T]).
		if t.NumIn() != 0 || t.NumOut() != 1 {
			return reflect.Value{}, fmt.Errorf("unsupported func type %v", t)
		}
		retVal, err := castToType(v, t.Out(0), ops.Global())
		if err != nil {
			return reflect.Value{}, err
		}
		fn := reflect.MakeFunc(t, func(_ []reflect.Value) []reflect.Value {
			return []reflect.Value{retVal}
		})
		return fn, nil
	case reflect.Chan:
		// Use reflect.MakeChan so named channel types (type MyChan chan int) are
		// created correctly rather than a plain chan int.
		size := 1
		if ops.hasLength {
			s, sErr := ToE[int](ops.lengthVal)
			if sErr != nil {
				return reflect.Value{}, fmt.Errorf(ErrorInvalidOption, "LENGTH", ops.lengthVal)
			}
			if s < 1 {
				return reflect.Value{}, fmt.Errorf("invalid channel buffer size %d", s)
			}
			size = s
		}
		elem, err := castToType(v, t.Elem(), ops.Global())
		if err != nil {
			return reflect.Value{}, err
		}
		ch := reflect.MakeChan(t, size)
		ch.Send(elem)
		return ch, nil
	case reflect.Pointer:
		elem, err := castToType(v, t.Elem(), ops.Global())
		if err != nil {
			return reflect.Value{}, err
		}
		ptr := reflect.New(t.Elem())
		ptr.Elem().Set(elem)
		return ptr, nil
	case reflect.Struct:
		return castToStructType(v, t, ops)
	default:
		result, err := castToKind(v, t.Kind(), ops)
		if err != nil {
			return reflect.Value{}, err
		}
		if result.Type() == t {
			return result, nil
		}
		if result.Type().ConvertibleTo(t) {
			return result.Convert(t), nil
		}
		return reflect.Value{}, fmt.Errorf("cannot convert %v to %v", result.Type(), t)
	}
}

// castToStructType casts v to the struct type t. It tries, in order:
//  1. Direct assignment or convertible type (fast path for same-type values).
//  2. encoding.TextUnmarshaler (handles struct types with text parsing).
//  3. Struct hydration via toStruct (map or struct sources).
//
// time.Time is handled by namedConverters in castToType and never reaches here.
func castToStructType(v any, t reflect.Type, ops ops) (reflect.Value, error) {
	if v != nil {
		srcVal := reflect.ValueOf(v)
		if srcVal.IsValid() {
			if srcVal.Type() == t {
				return srcVal, nil
			}
			if srcVal.Type().ConvertibleTo(t) {
				return srcVal.Convert(t), nil
			}
		}
	}

	ptr := reflect.New(t)
	if tu, ok := ptr.Interface().(encoding.TextUnmarshaler); ok {
		if s, strErr := ToE[string](v); strErr == nil {
			if umErr := tu.UnmarshalText([]byte(s)); umErr == nil {
				return ptr.Elem(), nil
			}
		}
	}

	raw, err := toStruct(reflect.Zero(t), v, ops)
	return rawToValue(raw, t, ops, err)
}

// fieldKey returns the source-map key for a struct field.
// Priority: cast tag > json tag (name portion only) > field name.
// Returns ("", false) when the tag value is "-", meaning skip this field.
func fieldKey(field reflect.StructField) (string, bool) {
	if tag, ok := field.Tag.Lookup("cast"); ok {
		if tag == "-" {
			return "", false
		}
		return tag, true
	}
	if tag, ok := field.Tag.Lookup("json"); ok {
		name := strings.SplitN(tag, ",", 2)[0]
		if name == "-" {
			return "", false
		}
		if name != "" {
			return name, true
		}
	}
	return field.Name, true
}

// extractFieldValue returns a struct field's value as any, handling unexported
// fields via kind-specific reflect methods. Returns (nil, false) for unexported
// non-scalar fields that cannot be read without unsafe.
func extractFieldValue(v reflect.Value) (any, bool) {
	if v.CanInterface() {
		return v.Interface(), true
	}
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	case reflect.String:
		return v.String(), true
	default:
		return nil, false
	}
}

// isNamedScalarStructType reports whether t is a struct type that has a
// meaningful scalar representation — i.e., types in namedStructTypes (which
// mirrors the struct entries of namedConverters without creating an init cycle).
func isNamedScalarStructType(t reflect.Type) bool {
	if t.Kind() != reflect.Struct {
		return false
	}
	_, ok := namedStructTypes[t]
	return ok
}

// isScalarKind reports whether k is a scalar kind handled by castToKind.
// reflect.Interface is excluded: castToType enforces assignability checks that
// castToKind skips, so interface element types must go through castToType.
func isScalarKind(k reflect.Kind) bool {
	switch k {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String:
		return true
	}
	return false
}
