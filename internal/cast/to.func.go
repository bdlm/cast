package cast

import (
	"fmt"
	"reflect"
)

// ToFunc returns a function value of the type described by to. The function
// returns the cast value when called. The result is returned as any wrapping
// a raw `func() T` closure; the caller (public ToE) converts to the named
// Func[T] type via reflect.Convert when the underlying type matches.
//
// Options:
//   - DEFAULT: Func[T], default return value on error.
func ToFunc(to reflect.Value, from any, ops Ops) (any, error) {
	var defaultValue any

	if ops.HasDefault {
		dv := reflect.ValueOf(ops.DefaultVal)
		if !dv.IsValid() || !dv.Type().AssignableTo(to.Type()) {
			return defaultValue, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.DefaultVal)
		}
		defaultValue = ops.DefaultVal
		// Prevent DEFAULT from being passed to element casts.
		ops = ops.Delete(DEFAULT)
	}

	if to.Type().NumOut() < 1 {
		return defaultValue, fmt.Errorf(ErrorStrUnableToCast, from, from, to.Interface())
	}

	funcType := to.Type()
	retType := funcType.Out(0)

	// Cast the source value to the function's return type via reflection.
	retVal, err := CastToType(from, retType, ops.Global())
	if err != nil {
		return defaultValue, fmt.Errorf("%w: "+ErrorStrErrorCastingFunc, err, from, defaultValue)
	}

	// Build the func via reflection so the result type matches to.Type() exactly
	// when funcType is a named type like cast.Func[int].
	fn := reflect.MakeFunc(funcType, func(_ []reflect.Value) []reflect.Value {
		return []reflect.Value{retVal}
	})
	return fn.Interface(), nil
}
