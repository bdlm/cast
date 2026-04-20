package cast

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/bdlm/errors/v2"
)

// toString casts an interface to a string type.
//
// Options:
//   - DEFAULT: string, default return value on error.
//   - JSON: bool, encode the string representation as a JSON string literal.
func toString(from any, ops ops) (any, error) {
	var ret any
	var ok bool

	if ops.hasDefault {
		if ret, ok = ops.defaultVal.(string); !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
	}

	if ops.jsonEncode {
		s, sErr := toString(from, ops.Delete(JSON))
		if sErr != nil {
			return ret, sErr
		}
		b, mErr := json.Marshal(s)
		if mErr != nil {
			return ret, errors.Wrap(mErr, "JSON encoding failed")
		}
		return string(b), nil
	}

	switch val := from.(type) {
	case nil:
		return "", nil
	case string:
		return val, nil
	case []byte:
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
			return ret, errors.Wrap(err, "JSON encoding failed")
		}
		return string(b), nil
	}
}
