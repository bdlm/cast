package cast

import (
	"errors"
	"fmt"
	"strconv"
)

/*
ErrUintBelowZero defines the error returned when attempting to cast a
negative value as a uint.
*/
var ErrUintBelowZero = errors.New("cannot cast negative value as uint")

/*
ToUint casts an interface to a uint type.
*/
func ToUint(i interface{}) (uint, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 0)
		if err == nil {
			return uint(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint: %s", i, err)
	case int:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint(s), nil
	case int64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint(s), nil
	case int32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint(s), nil
	case int16:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint(s), nil
	case int8:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint(s), nil
	case uint:
		return s, nil
	case uint64:
		return uint(s), nil
	case uint32:
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint8:
		return uint(s), nil
	case float64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint(s), nil
	case float32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
}

/*
ToUint64 casts an interface to a uint64 type.
*/
func ToUint64(i interface{}) (uint64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint64: %s", i, err)
	case int:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint64(s), nil
	case int64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint64(s), nil
	case int32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint64(s), nil
	case int16:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint64(s), nil
	case int8:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint64(s), nil
	case uint:
		return uint64(s), nil
	case uint64:
		return s, nil
	case uint32:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil
	case float32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint64(s), nil
	case float64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint64(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}

/*
ToUint32 casts an interface to a uint32 type.
*/
func ToUint32(i interface{}) (uint32, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 32)
		if err == nil {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint32: %s", i, err)
	case int:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint32(s), nil
	case int64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint32(s), nil
	case int32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint32(s), nil
	case int16:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint32(s), nil
	case int8:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint32(s), nil
	case uint:
		return uint32(s), nil
	case uint64:
		return uint32(s), nil
	case uint32:
		return s, nil
	case uint16:
		return uint32(s), nil
	case uint8:
		return uint32(s), nil
	case float64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint32(s), nil
	case float32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint32(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}

/*
ToUint16 casts an interface to a uint16 type.
*/
func ToUint16(i interface{}) (uint16, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 16)
		if err == nil {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint16: %s", i, err)
	case int:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint16(s), nil
	case int64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint16(s), nil
	case int32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint16(s), nil
	case int16:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint16(s), nil
	case int8:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint16(s), nil
	case uint:
		return uint16(s), nil
	case uint64:
		return uint16(s), nil
	case uint32:
		return uint16(s), nil
	case uint16:
		return s, nil
	case uint8:
		return uint16(s), nil
	case float64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint16(s), nil
	case float32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint16(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", i, i)
	}
}

/*
ToUint8 casts an interface to a uint type.
*/
func ToUint8(i interface{}) (uint8, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 8)
		if err == nil {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint8: %s", i, err)
	case int:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint8(s), nil
	case int64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint8(s), nil
	case int32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint8(s), nil
	case int16:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint8(s), nil
	case int8:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint8(s), nil
	case uint:
		return uint8(s), nil
	case uint64:
		return uint8(s), nil
	case uint32:
		return uint8(s), nil
	case uint16:
		return uint8(s), nil
	case uint8:
		return s, nil
	case float64:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint8(s), nil
	case float32:
		if s < 0 {
			return 0, ErrUintBelowZero
		}
		return uint8(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
}
