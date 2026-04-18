package cast

import (
	"reflect"

	"github.com/bdlm/errors/v2"
)

// castToKind casts v to the scalar Go type corresponding to kind and returns
// the result as a reflect.Value. Used by toMap, castToArray, and toStruct.
func castToKind(v any, kind reflect.Kind, ops Ops) (reflect.Value, error) {
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
func castToType(v any, t reflect.Type, ops Ops) (reflect.Value, error) {
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

func castToSliceType(v any, t reflect.Type, ops Ops) (reflect.Value, error) {
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

// castToArray casts from (must be a slice or array) to the array type arrType,
// casting each element. Requires source length == arrType.Len().
func castToArray(arrType reflect.Type, from any, ops Ops) (reflect.Value, error) {
	src := reflect.ValueOf(from)
	if !src.IsValid() || (src.Kind() != reflect.Slice && src.Kind() != reflect.Array) {
		return reflect.Value{}, errors.Errorf(ErrorStrUnableToCast, from, from, arrType)
	}
	if src.Len() != arrType.Len() {
		return reflect.Value{}, errors.Errorf(
			"array length mismatch: source has %d elements, target needs %d",
			src.Len(), arrType.Len(),
		)
	}
	arr := reflect.New(arrType).Elem()
	elemKind := arrType.Elem().Kind()
	for i := 0; i < src.Len(); i++ {
		castVal, err := castToKind(src.Index(i).Interface(), elemKind, ops)
		if err != nil {
			return reflect.Value{}, err
		}
		arr.Index(i).Set(castVal)
	}
	return arr, nil
}

// makeArrayChan builds a chan [N]T of the given chanType (must have Kind Chan
// with Elem Kind Array), casts from into the array, and sends it on the channel.
func makeArrayChan(chanType reflect.Type, from any, size int, ops Ops) (any, error) {
	arr, err := castToArray(chanType.Elem(), from, ops)
	if err != nil {
		return nil, err
	}
	ch := reflect.MakeChan(chanType, size)
	if size == 0 {
		go ch.Send(arr)
		return ch.Interface(), nil
	}
	ch.Send(arr)
	return ch.Interface(), nil
}

// makeArrayFunc builds a func() [N]T of funcType, casting from into the array.
func makeArrayFunc(funcType reflect.Type, from any, ops Ops) (any, error) {
	arr, err := castToArray(funcType.Out(0), from, ops)
	if err != nil {
		return nil, err
	}
	fn := reflect.MakeFunc(funcType, func(_ []reflect.Value) []reflect.Value {
		return []reflect.Value{arr}
	})
	return fn.Interface(), nil
}

// makeArrayChanFunc builds a func() chan [N]T of funcType, casting from into
// the array which is sent on the channel the function returns.
func makeArrayChanFunc(funcType reflect.Type, from any, ops Ops) (any, error) {
	chanType := funcType.Out(0)
	chanVal, err := makeArrayChan(chanType, from, 1, ops)
	if err != nil {
		return nil, err
	}
	rv := reflect.ValueOf(chanVal)
	fn := reflect.MakeFunc(funcType, func(_ []reflect.Value) []reflect.Value {
		return []reflect.Value{rv}
	})
	return fn.Interface(), nil
}
