package cast

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ToString casts an interface to a string type.
//
// Options:
//   - DEFAULT: string, default return value on error.
//   - JSON: bool, encode the string representation as a JSON string literal.
func ToString(from any, ops Ops) (string, error) {
	var ret string
	var ok bool

	if ops.HasDefault {
		if ret, ok = ops.DefaultVal.(string); !ok {
			return ret, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.DefaultVal)
		}
	}

	if ops.JsonEncode {
		s, sErr := ToString(from, ops.Delete(JSON))
		if sErr != nil {
			return ret, sErr
		}
		b, mErr := json.Marshal(s)
		if mErr != nil {
			return ret, fmt.Errorf("%w: JSON encoding failed", mErr)
		}
		return string(b), nil
	}

	switch val := from.(type) {
	case nil:
		return "", nil
	case string:
		return val, nil
	case []byte: // = []uint8
		return string(val), nil
	case []rune: // = []int32
		return string(val), nil
	case fmt.Stringer:
		return val.String(), nil
	case bool,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		complex64, complex128:
		return fmt.Sprintf("%v", val), nil
	default:
		// Channels have no useful JSON representation; preserve the address string.
		if reflect.TypeOf(val).Kind() == reflect.Chan {
			return fmt.Sprintf("%v", val), nil
		}
		b, err := json.Marshal(val)
		if err != nil {
			return ret, fmt.Errorf("%w: JSON encoding failed", err)
		}
		return string(b), nil
	}
}
