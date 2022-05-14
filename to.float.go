package cast

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bdlm/errors/v2"
	"golang.org/x/exp/constraints"
)

// toFloat casts an interface to a float type.
func toFloat[TTo constraints.Float](from reflect.Value) (TTo, error) {
	to := reflect.Indirect(reflect.ValueOf(new(TTo)))

	switch s := from.Interface().(type) {
	case bool:
		if s {
			return TTo(1), nil
		}
		return TTo(0), nil
	case float64:
		return TTo(s), nil
	case float32:
		return TTo(s), nil
	case int:
		return TTo(s), nil
	case int64:
		return TTo(s), nil
	case int32:
		return TTo(s), nil
	case int16:
		return TTo(s), nil
	case int8:
		return TTo(s), nil
	case nil:
		return TTo(0), nil
	case fmt.Stringer:
		var e, err error
		var val interface{}
		var bitSize = 64
		str := s.String()

		if to.Type().Kind() == reflect.Float32 {
			bitSize = 32
		}

		if val, e = strconv.ParseFloat(str, bitSize); e != nil {
			err = e
			str = strings.ReplaceAll(
				strings.Split(
					strings.Trim(str, "\r\n\t "),
					".",
				)[0],
				",", "",
			)
			if val, e = strconv.ParseFloat(str, bitSize); e != nil {
				err = errors.WrapE(err, e)
				_, e := strconv.ParseComplex(str, bitSize)
				if e != nil {
					err = errors.WrapE(err, e)
				} else {
					err = errors.Wrap(err, "unable to cast complex to int")
					val = float64(0)
				}
			} else {
				err = nil
			}
		}
		if to.Type().Kind() == reflect.Float32 {
			val = float32(val.(float64))
		}
		return val.(TTo), err
	case string:
		var e, err error
		var val interface{}
		var bitSize = 64
		str := s

		if to.Type().Kind() == reflect.Float32 {
			bitSize = 32
		}

		if val, e = strconv.ParseFloat(str, bitSize); e != nil {
			err = e
			str = strings.ReplaceAll(
				strings.Split(
					strings.Trim(str, "\r\n\t "),
					".",
				)[0],
				",", "",
			)
			if val, e = strconv.ParseFloat(str, bitSize); e != nil {
				err = errors.WrapE(err, e)
				_, e := strconv.ParseComplex(str, bitSize)
				if e != nil {
					err = errors.WrapE(err, e)
				} else {
					err = errors.Wrap(err, "unable to cast %T to %T", from.Interface(), to.Interface())
					val = float64(0)
				}
			} else {
				err = nil
			}
		}
		if to.Type().Kind() == reflect.Float32 {
			val = float32(val.(float64))
		}
		return val.(TTo), err
	case uint:
		return TTo(s), nil
	case uint64:
		return TTo(s), nil
	case uint32:
		return TTo(s), nil
	case uint16:
		return TTo(s), nil
	case uint8:
		return TTo(s), nil
	}

	ret, err := toFloat[TTo](reflect.ValueOf(fmt.Sprintf("%v", from.Interface())))
	if nil != err {
		return 0, errors.Wrap(err, "unable to cast %T to %T", from.Interface(), to.Interface())
	}
	return ret, nil
}
