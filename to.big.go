package cast

import (
	"math/big"
	"time"

	"github.com/bdlm/errors/v2"
)

// toBigInt converts v to *big.Int.
//
// String values are parsed with (*big.Int).SetString with base 0, which
// auto-detects base from the prefix: 0x for hex, 0o for octal, 0b for binary,
// and decimal otherwise. Integer values are set directly. Float values are
// truncated toward zero.
//
// Options:
//   - DEFAULT: *big.Int, value to return on error.
func toBigInt(v any, ops ops) (any, error) {
	var ret any = (*big.Int)(nil)

	if ops.hasDefault {
		i, ok := ops.defaultVal.(*big.Int)
		if !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		ret = i
	}

	switch val := v.(type) {
	case nil:
		return ret, errors.Errorf(ErrorStrUnableToCast, v, v, (*big.Int)(nil))
	case *big.Int:
		return new(big.Int).Set(val), nil
	case big.Int:
		return new(big.Int).Set(&val), nil
	case *big.Float:
		i, _ := val.Int(nil)
		return i, nil
	case big.Float:
		i, _ := val.Int(nil)
		return i, nil
	case string:
		i := new(big.Int)
		if _, ok := i.SetString(val, 0); !ok {
			return ret, errors.Errorf(ErrorStrUnableToCast, v, v, (*big.Int)(nil))
		}
		return i, nil
	case int:
		return big.NewInt(int64(val)), nil
	case int8:
		return big.NewInt(int64(val)), nil
	case int16:
		return big.NewInt(int64(val)), nil
	case int32:
		return big.NewInt(int64(val)), nil
	case int64:
		return big.NewInt(val), nil
	case uint:
		return new(big.Int).SetUint64(uint64(val)), nil
	case uint8:
		return new(big.Int).SetUint64(uint64(val)), nil
	case uint16:
		return new(big.Int).SetUint64(uint64(val)), nil
	case uint32:
		return new(big.Int).SetUint64(uint64(val)), nil
	case uint64:
		return new(big.Int).SetUint64(val), nil
	case uintptr:
		return new(big.Int).SetUint64(uint64(val)), nil
	case float32:
		i, _ := new(big.Float).SetFloat64(float64(val)).Int(nil)
		return i, nil
	case float64:
		i, _ := new(big.Float).SetFloat64(val).Int(nil)
		return i, nil
	case time.Time:
		return big.NewInt(val.Unix()), nil
	default:
		if s, err := toString(v, ops.Delete(DEFAULT)); err == nil {
			i := new(big.Int)
			if _, ok := i.SetString(s.(string), 0); ok {
				return i, nil
			}
		}
	}

	return ret, errors.Errorf(ErrorStrUnableToCast, v, v, (*big.Int)(nil))
}

// toBigFloat converts v to *big.Float.
//
// String values are parsed with (*big.Float).SetString. Integer and
// floating-point values are set directly without loss of precision beyond
// what float64 can represent (for float32/float64 sources) or exactly (for
// integer sources).
//
// Options:
//   - DEFAULT: *big.Float, value to return on error.
func toBigFloat(v any, ops ops) (any, error) {
	var ret any = (*big.Float)(nil)

	if ops.hasDefault {
		f, ok := ops.defaultVal.(*big.Float)
		if !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		ret = f
	}

	switch val := v.(type) {
	case nil:
		return ret, errors.Errorf(ErrorStrUnableToCast, v, v, (*big.Float)(nil))
	case *big.Float:
		return new(big.Float).Copy(val), nil
	case big.Float:
		return new(big.Float).Copy(&val), nil
	case *big.Int:
		return new(big.Float).SetInt(val), nil
	case big.Int:
		return new(big.Float).SetInt(&val), nil
	case string:
		f := new(big.Float)
		if _, ok := f.SetString(val); !ok {
			return ret, errors.Errorf(ErrorStrUnableToCast, v, v, (*big.Float)(nil))
		}
		return f, nil
	case int:
		return new(big.Float).SetInt64(int64(val)), nil
	case int8:
		return new(big.Float).SetInt64(int64(val)), nil
	case int16:
		return new(big.Float).SetInt64(int64(val)), nil
	case int32:
		return new(big.Float).SetInt64(int64(val)), nil
	case int64:
		return new(big.Float).SetInt64(val), nil
	case uint:
		return new(big.Float).SetUint64(uint64(val)), nil
	case uint8:
		return new(big.Float).SetUint64(uint64(val)), nil
	case uint16:
		return new(big.Float).SetUint64(uint64(val)), nil
	case uint32:
		return new(big.Float).SetUint64(uint64(val)), nil
	case uint64:
		return new(big.Float).SetUint64(val), nil
	case uintptr:
		return new(big.Float).SetUint64(uint64(val)), nil
	case float32:
		return new(big.Float).SetFloat64(float64(val)), nil
	case float64:
		return new(big.Float).SetFloat64(val), nil
	case time.Time:
		secs := float64(val.Unix()) + float64(val.Nanosecond())/1e9
		return new(big.Float).SetFloat64(secs), nil
	default:
		if s, err := toString(v, ops.Delete(DEFAULT)); err == nil {
			f := new(big.Float)
			if _, ok := f.SetString(s.(string)); ok {
				return f, nil
			}
		}
	}

	return ret, errors.Errorf(ErrorStrUnableToCast, v, v, (*big.Float)(nil))
}
