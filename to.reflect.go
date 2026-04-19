package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

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
		r, err := ToE[bool](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Int:
		r, err := ToE[int](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Int8:
		r, err := ToE[int8](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Int16:
		r, err := ToE[int16](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Int32:
		r, err := ToE[int32](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Int64:
		r, err := ToE[int64](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Uint:
		r, err := ToE[uint](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Uint8:
		r, err := ToE[uint8](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Uint16:
		r, err := ToE[uint16](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Uint32:
		r, err := ToE[uint32](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Uint64:
		r, err := ToE[uint64](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Uintptr:
		r, err := ToE[uintptr](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Float32:
		r, err := ToE[float32](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Float64:
		r, err := ToE[float64](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Complex64:
		r, err := ToE[complex64](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.Complex128:
		r, err := ToE[complex128](v, ops.List()...)
		return reflect.ValueOf(r), err
	case reflect.String:
		r, err := ToE[string](v, ops.List()...)
		return reflect.ValueOf(r), err
	}
	return reflect.Value{}, errors.Errorf("unsupported kind %v", kind)
}

// castToType casts v to the type t and returns the result as a reflect.Value.
func castToType(v any, t reflect.Type, ops ops) (reflect.Value, error) {
	switch t.Kind() {
	case reflect.Interface:
		if v == nil {
			return reflect.Zero(t), nil
		}
		src := reflect.ValueOf(v)
		if !src.Type().AssignableTo(t) {
			return reflect.Value{}, errors.Errorf(ErrorStrUnableToCast, v, v, t)
		}
		return src, nil
	case reflect.Slice:
		return castToSliceType(v, t, ops)
	case reflect.Func:
		// Only zero-arg, one-return functions are supported (matches Func[T]).
		if t.NumIn() != 0 || t.NumOut() != 1 {
			return reflect.Value{}, errors.Errorf("unsupported func type %v", t)
		}
		retVal, err := castToType(v, t.Out(0), ops)
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
				return reflect.Value{}, errors.Errorf(ErrorInvalidOption, "LENGTH", ops.lengthVal)
			}
			if s < 1 {
				return reflect.Value{}, errors.Errorf("invalid channel buffer size %d", s)
			}
			size = s
		}
		elem, err := castToType(v, t.Elem(), ops.Delete(LENGTH))
		if err != nil {
			return reflect.Value{}, err
		}
		ch := reflect.MakeChan(t, size)
		ch.Send(elem)
		return ch, nil
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
		return reflect.Value{}, errors.Errorf("cannot convert %v to %v", result.Type(), t)
	}
}

func castToSliceType(v any, t reflect.Type, ops ops) (reflect.Value, error) {
	srcVal := reflect.ValueOf(v)
	if !srcVal.IsValid() {
		return reflect.MakeSlice(t, 0, 0), nil
	}
	switch srcVal.Kind() {
	case reflect.Slice, reflect.Array:
		result := reflect.MakeSlice(t, srcVal.Len(), srcVal.Len())
		for i := 0; i < srcVal.Len(); i++ {
			elem, err := castToType(srcVal.Index(i).Interface(), t.Elem(), ops)
			if err != nil {
				return reflect.Value{}, err
			}
			result.Index(i).Set(elem)
		}
		return result, nil
	default:
		elem, err := castToType(v, t.Elem(), ops)
		if err != nil {
			return reflect.Value{}, err
		}
		result := reflect.MakeSlice(t, 1, 1)
		result.Index(0).Set(elem)
		return result, nil
	}
}

