package cast

import (
	"fmt"
	"strings"
	"time"
)

/*
ToDuration casts an interface to a time.Duration type.
*/
func ToDuration(i interface{}) (d time.Duration, err error) {
	i = indirect(i)

	switch s := i.(type) {
	case time.Duration:
		return s, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		v, _ := ToInt64(s)
		d = time.Duration(v)
		return
	case float32, float64:
		v, _ := ToFloat64(s)
		d = time.Duration(v)
		return
	case string:
		if strings.ContainsAny(s, "nsuµmh") {
			d, err = time.ParseDuration(s)
		} else {
			d, err = time.ParseDuration(s + "ns")
		}
		return
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to Duration", i, i)
		return
	}
}
