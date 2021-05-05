package cast

import (
	"strings"
	"time"

	"github.com/bdlm/errors/v2"
)

// ToDuration casts an interface to a time.Duration type, discarding any
// errors.
func ToDuration(i interface{}) time.Duration {
	ret, _ := ToDurationE(i)
	return ret
}

// ToDurationE casts an interface to a time.Duration type.
func ToDurationE(i interface{}) (time.Duration, error) {
	var d time.Duration
	var err error
	i = indirect(i)

	switch s := i.(type) {
	case time.Duration:
		return s, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		v := ToInt64(s)
		d = time.Duration(v)
		return d, nil
	case float32, float64:
		v := ToFloat64(s)
		d = time.Duration(v)
		return d, nil
	case string:
		if strings.ContainsAny(s, "nsuµmh") {
			d, err = time.ParseDuration(s)
		} else {
			d, err = time.ParseDuration(s + "ns")
		}
		return d, err
	default:
		return d, errors.Errorf("unable to cast %#v of type %T to Duration", i, i)
	}
}
