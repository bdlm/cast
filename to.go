// Package cast provides generic type conversion for Go 1.21+.
//
// The two public entry points are [To] (ignores errors) and [ToE] (returns
// errors). Both accept an optional variadic [Op] list that controls conversion
// behavior; see [Flag] for available options.
//
// Supported target types are described by the [Types] constraint: all basic
// scalar types, slices of scalars, channels of scalars/slices, maps, and
// [Func] wrappers for each of those groups.
package cast

import (
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
	std_error "github.com/bdlm/std/v2/errors"
)

// To casts the value v to the given type, ignoring any errors. See the ToE
// documentation for more information.
func To[TTo Types](v any, o ...Op) TTo {
	ret, _ := ToE[TTo](v, o...)
	return ret
}

// ToE casts the value v to the given type, returning any errors.
//
// ops is an optional variadic list of [Op] values that control conversion
// behavior. If omitted, the default conversion behavior for the target type is
// used. Available options depend on the target type; see the documentation for
// the specific type conversion function for more information.
//
// Complex types have specific default behaviors, for example:
//
//   - If the target type is a channel, a buffered channel of size 1 is created
//     and the cast value `v` is sent to the channel before it is returned.
//
//   - If the target type is a slice, a slice is created. To pre-allocate
//     backing capacity set the LENGTH flag: `cast.ToE[[]int](v, cast.Op{cast.LENGTH, 10})`.
//     The source must itself be a slice or array; scalar sources are rejected.
//
//   - If the target type is a map, the source is converted into the target map
//     type. Supported sources: map (key/value types cast), struct or *struct
//     (field names become keys), slice or array (indices become keys).
//
// See the documentation for the specific type conversion function for more
// information.
func ToE[TTo Types](val any, ops ...Op) (panicTo TTo, panicErr error) {

	var err error
	var ok bool
	var retIface any
	var ret0Val TTo
	var retVal TTo

	// Don't panic.
	// Recover from any panic that may occur during type casting, converting it
	// to an error so callers can handle it gracefully.
	defer func() {
		if r := recover(); r != nil {
			panicTo = ret0Val
			switch e := r.(type) {
			case error:
				panicErr = errors.WrapE(ErrorUnableToCast, errors.Wrap(e, "failure casting %T to %T (panic)", val, ret0Val))
			default:
				panicErr = errors.WrapE(ErrorUnableToCast, errors.Errorf("failure casting %T to %T (panic): %v", val, ret0Val, e))
			}
		}
	}()
	options := parseOps(ops)

	toRef := reflect.ValueOf(new(TTo))
	to := reflect.Indirect(toRef)

	// Dereference pointer sources before dispatch so every converter sees the
	// concrete value. Follows pointer chains (**T, ***T, …) and updates val.
	//
	// Struct pointers are dereferenced only when the target is also a struct:
	// pointer-receiver interface implementations (error, Stringer) and named
	// converter types (e.g. *regexp.Regexp) expect the pointer, not the value,
	// so those are left alone. Pointer-to-interface is never dereferenced.
	if srcVal := reflect.ValueOf(val); srcVal.IsValid() && srcVal.Kind() == reflect.Pointer {
		changed := false
		targetIsStruct := to.Kind() == reflect.Struct
		for srcVal.Kind() == reflect.Pointer {
			if srcVal.IsNil() {
				val = nil
				changed = false // prevent the post-loop assignment from overwriting nil
				break
			}
			elem := srcVal.Elem()
			if elem.Kind() == reflect.Pointer {
				// Unwrap another level of indirection (**T → *T → …).
				srcVal = elem
				changed = true
				continue
			}
			if isScalarKind(elem.Kind()) || (elem.Kind() == reflect.Struct && targetIsStruct) {
				// Dereference scalars always; structs only when the target is a struct.
				srcVal = elem
				changed = true
			}
			break
		}
		if changed && srcVal.IsValid() {
			val = srcVal.Interface()
		}
	}

	// Named types have dedicated converters registered in namedConverters
	// (util.reflect.go). Check the table first so that adding a new named
	// type requires only one edit there.
	if fn, ok := namedConverters[to.Type()]; ok {
		retIface, err = fn(val, options)
	} else {
		switch to.Type().Kind() {
		// reflect.Array:        array targets are not in Types; use []T slice targets instead.
		// reflect.Invalid:
		// reflect.UnsafePointer:
		default:
			retIface = ret0Val
			if _, ok := retIface.(error); ok {
				retIface = errors.Errorf("%s", To[string](val, ops...))
			} else if _, ok := retIface.(std_error.Error); ok {
				retIface = errors.Errorf("%s", To[string](val, ops...))
			} else if _, ok := retIface.(fmt.Stringer); ok {
				retIface = errors.Errorf("%s", To[string](val, ops...))
			} else {
				return ret0Val, errors.WrapE(ErrorUnableToCast, errors.Errorf(ErrorStrUnableToCast, val, val, to.Interface()))
			}

		case reflect.Interface:
			retIface = val

		case reflect.Struct:
			retIface, err = toStruct(to, val, options)

		case reflect.Bool:
			retIface, err = toBool(val, options)
		case reflect.Chan:
			retIface, err = toChan(to, val, options)
		case reflect.Map:
			retIface, err = toMap(to, val, options)
		case reflect.Pointer:
			result, castErr := castToType(val, to.Type(), options)
			if castErr != nil {
				return ret0Val, errors.WrapE(ErrorUnableToCast, castErr)
			}
			retIface = result.Interface()
		case reflect.Slice:
			retIface, err = toSlice(to, val, options)
		case reflect.Func:
			retIface, err = toFunc[TTo](to, val, options)
		case reflect.Complex64:
			retIface, err = toComplex[complex64](val, options)
		case reflect.Complex128:
			retIface, err = toComplex[complex128](val, options)
		case reflect.Float32:
			retIface, err = toFloat[float32](val, options)
		case reflect.Float64:
			retIface, err = toFloat[float64](val, options)
		case reflect.Int:
			retIface, err = toInt[int](val, options)
		case reflect.Int8:
			retIface, err = toInt[int8](val, options)
		case reflect.Int16:
			retIface, err = toInt[int16](val, options)
		case reflect.Int32:
			retIface, err = toInt[int32](val, options)
		case reflect.Int64:
			retIface, err = toInt[int64](val, options)
		case reflect.Uint:
			retIface, err = toInt[uint](val, options)
		case reflect.Uint8:
			retIface, err = toInt[uint8](val, options)
		case reflect.Uint16:
			retIface, err = toInt[uint16](val, options)
		case reflect.Uint32:
			retIface, err = toInt[uint32](val, options)
		case reflect.Uint64:
			retIface, err = toInt[uint64](val, options)
		case reflect.Uintptr:
			retIface, err = toInt[uintptr](val, options)
		case reflect.String:
			retIface, err = toString(val, options)
		}
	}

	if retVal, ok = retIface.(TTo); !ok && retIface != nil {
		// Direct assertion failed; try a reflect-based conversion for named types
		// whose underlying kind matches (e.g. type MyInt int → int convertible to MyInt).
		rv := reflect.ValueOf(retIface)
		if rv.IsValid() && rv.Type().ConvertibleTo(to.Type()) {
			retVal, ok = rv.Convert(to.Type()).Interface().(TTo)
		}
		if !ok {
			return ret0Val, errors.WrapE(ErrorUnableToCast, errors.Errorf("unable to cast %#.10v of type %T to %T (%#.10v %T)", val, val, *new(TTo), retVal, retVal))
		}
	}

	if err != nil {
		return retVal, errors.WrapE(ErrorUnableToCast, err)
	}

	if retIface == nil {
		return ret0Val, nil
	}

	return retVal, nil
}
