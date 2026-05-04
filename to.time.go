package cast

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

// toTime converts v to time.Time.
//
// Integer and unsigned-integer values are treated as Unix seconds.
// Floating-point values are treated as Unix seconds with fractional second
// precision preserved in the nanosecond field.
// *big.Int and big.Int values are treated as Unix seconds (int64 range only).
// *big.Float and big.Float values are treated as Unix seconds with fractional
// second precision (converted via float64).
// String and []byte values are tried against each format in timeFormats.
// fmt.Stringer values are converted via String() and then parsed.
//
// Options:
//   - DEFAULT: time.Time, value to return on error.
func toTime(v any, ops ops) (any, error) {
	var ret any = time.Time{}

	if ops.hasDefault {
		defaultVal, ok := ops.defaultVal.(time.Time)
		if !ok {
			return ret, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		ret = defaultVal
	}

	switch val := v.(type) {
	case nil:
		return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
	case time.Time:
		return val, nil
	case *time.Time:
		if val == nil {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		return *val, nil
	case string:
		t, ok := parseTimeString(val, ops.formatVal)
		if !ok {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		return t, nil
	case []byte:
		t, ok := parseTimeString(string(val), ops.formatVal)
		if !ok {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
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
		secs := float64(val)
		floor := math.Floor(secs)
		return time.Unix(int64(floor), int64((secs-floor)*1e9)).UTC(), nil
	case float64:
		floor := math.Floor(val)
		return time.Unix(int64(floor), int64((val-floor)*1e9)).UTC(), nil
	case *big.Int:
		if val == nil || !val.IsInt64() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		return time.Unix(0, val.Int64()).UTC(), nil
	case big.Int:
		if !val.IsInt64() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		return time.Unix(0, val.Int64()).UTC(), nil
	case *big.Float:
		if val == nil || val.IsInf() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		f64, _ := val.Float64()
		if math.IsInf(f64, 0) || math.IsNaN(f64) {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		floor := math.Floor(f64)
		return time.Unix(int64(floor), int64((f64-floor)*1e9)).UTC(), nil
	case big.Float:
		if val.IsInf() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		f64, _ := val.Float64()
		if math.IsInf(f64, 0) || math.IsNaN(f64) {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
		}
		floor := math.Floor(f64)
		return time.Unix(int64(floor), int64((f64-floor)*1e9)).UTC(), nil
	default:
		if s, err := toString(v, ops.Delete(DEFAULT)); err == nil {
			return toTime(s, ops)
		}
	}

	return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, time.Time{})
}

// timeFormats is the ordered list of formats tried when parsing a string into
// time.Time. Formats with time zone information are tried first; date-only last.
var timeFormats = []string{
	// Commonly used formats — tz-aware first, then date-only.
	time.RFC3339,
	time.RFC3339Nano,
	time.DateOnly, // "2006-01-02"
	time.DateTime, // "2006-01-02 15:04:05"
	time.RFC1123Z, // RFC1123 with numeric zone
	time.RFC1123,
	time.RFC822Z,
	time.RFC822,

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
func parseTimeString(str string, format string) (time.Time, bool) {
	if format == "" {
		for _, stdFormat := range timeFormats {
			if parsed, err := time.Parse(stdFormat, str); err == nil {
				return parsed, true
			}
		}
	}
	if parsed, err := time.Parse(format, str); err == nil {
		return parsed, true
	}
	return time.Time{}, false
}
