package cast

import (
	"time"

	"github.com/bdlm/errors/v2"
)

// toTime converts v to time.Time.
//
// Integer and unsigned-integer values are treated as Unix nanoseconds.
// Floating-point values are treated as Unix seconds (fractional seconds are
// preserved by converting to nanoseconds before calling time.Unix).
// String and []byte values are tried against each format in timeFormats.
// fmt.Stringer values are converted via String() and then parsed.
//
// Options:
//   - DEFAULT: time.Time, value to return on error.
func toTime(v any, ops ops) (any, error) {
	var ret any = time.Time{}

	if ops.hasDefault {
		t, ok := ops.defaultVal.(time.Time)
		if !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		ret = t
	}

	switch val := v.(type) {
	case nil:
		return ret, errors.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
	case time.Time:
		return val, nil
	case *time.Time:
		if val == nil {
			return ret, errors.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		return *val, nil
	case string:
		t, ok := parseTimeString(val)
		if !ok {
			return ret, errors.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		return t, nil
	case []byte:
		t, ok := parseTimeString(string(val))
		if !ok {
			return ret, errors.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		return t, nil
	case int:
		return time.Unix(0, int64(val)).UTC(), nil
	case int8:
		return time.Unix(0, int64(val)).UTC(), nil
	case int16:
		return time.Unix(0, int64(val)).UTC(), nil
	case int32:
		return time.Unix(0, int64(val)).UTC(), nil
	case int64:
		return time.Unix(0, val).UTC(), nil
	case uint:
		return time.Unix(0, int64(val)).UTC(), nil
	case uint8:
		return time.Unix(0, int64(val)).UTC(), nil
	case uint16:
		return time.Unix(0, int64(val)).UTC(), nil
	case uint32:
		return time.Unix(0, int64(val)).UTC(), nil
	case uint64:
		return time.Unix(0, int64(val)).UTC(), nil
	case uintptr:
		return time.Unix(0, int64(val)).UTC(), nil
	case float32:
		return time.Unix(0, int64(float64(val)*float64(time.Second))).UTC(), nil
	case float64:
		return time.Unix(0, int64(val*float64(time.Second))).UTC(), nil
	default:
		if s, err := toString(v, ops.Delete(DEFAULT)); err == nil {
			if t, ok := parseTimeString(s.(string)); ok {
				return t, nil
			}
		}
	}

	return ret, errors.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
}

// timeFormats is the ordered list of formats tried when parsing a string into
// time.Time. Formats with time zone information are tried first; date-only last.
var timeFormats = []string{
	// Commonly used formats — tz-aware first, then date-only.
	time.RFC3339Nano,
	time.RFC3339,
	time.DateTime,  // "2006-01-02 15:04:05"
	time.RFC1123Z,  // RFC1123 with numeric zone
	time.RFC1123,
	time.RFC822Z,
	time.RFC822,
	time.DateOnly, // "2006-01-02"

	// Other formats defined in the time package, in no particular order.
	time.Layout, // The reference time, in numerical order.
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC850,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	time.TimeOnly, // "15:04:05"
}

// parseTimeString tries each entry in timeFormats and returns the first match.
func parseTimeString(s string) (time.Time, bool) {
	for _, f := range timeFormats {
		if parsed, err := time.Parse(f, s); err == nil {
			return parsed, true
		}
	}
	return time.Time{}, false
}
