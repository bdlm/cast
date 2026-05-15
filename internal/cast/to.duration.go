package cast

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

// toDuration converts v to time.Duration.
//
// String values are parsed using time.ParseDuration, which accepts numeric
// values followed by a unit suffix: "ns", "us" (or "µs"), "ms", "s", "m", "h".
// Multiple units may be combined: "1h30m", "2m45s".
//
// Integer, unsigned-integer, and floating-point values are converted directly
// to time.Duration (nanoseconds), matching Go's own time.Duration semantics.
//
// Options:
//   - DEFAULT: time.Duration, value to return on error.
func toDuration(v any, ops Ops) (any, error) {
	var ret any = time.Duration(0)

	if ops.HasDefault {
		d, ok := ops.DefaultVal.(time.Duration)
		if !ok {
			return ret, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.DefaultVal)
		}
		ret = d
	}

	switch val := v.(type) {
	case nil:
		return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Duration(0))
	case time.Duration:
		return val, nil
	case string:
		d, err := time.ParseDuration(val)
		if err != nil {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Duration(0))
		}
		return d, nil
	case int:
		return time.Duration(val), nil
	case int8:
		return time.Duration(val), nil
	case int16:
		return time.Duration(val), nil
	case int32:
		return time.Duration(val), nil
	case int64:
		return time.Duration(val), nil
	case uint:
		return time.Duration(val), nil
	case uint8:
		return time.Duration(val), nil
	case uint16:
		return time.Duration(val), nil
	case uint32:
		return time.Duration(val), nil
	case uint64:
		return time.Duration(val), nil
	case uintptr:
		return time.Duration(val), nil
	case float32:
		return time.Duration(int64(math.Floor(float64(val)))), nil
	case float64:
		return time.Duration(int64(math.Floor(val))), nil
	case *big.Int:
		if val == nil || !val.IsInt64() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Duration(0))
		}
		return time.Duration(val.Int64()), nil
	case big.Int:
		if !val.IsInt64() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Duration(0))
		}
		return time.Duration(val.Int64()), nil
	case *big.Float:
		if val == nil || val.IsInf() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Duration(0))
		}
		f64, _ := val.Float64()
		if math.IsInf(f64, 0) || math.IsNaN(f64) {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Duration(0))
		}
		return time.Duration(int64(math.Floor(f64))), nil
	case big.Float:
		if val.IsInf() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Duration(0))
		}
		f64, _ := val.Float64()
		if math.IsInf(f64, 0) || math.IsNaN(f64) {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Duration(0))
		}
		return time.Duration(int64(math.Floor(f64))), nil
	default:
		if s, err := ToString(v, ops.Delete(DEFAULT)); err == nil {
			if d, err := time.ParseDuration(s); err == nil {
				return d, nil
			}
		}
	}

	return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Duration(0))
}
