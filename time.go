package cast

import (
	"fmt"
	"time"
)

// ToTime casts an interface to a time.Time type, discarding any errors.
func ToTime(i interface{}) time.Time {
	ret, _ := ToTimeE(i)
	return ret
}

// ToTimeE casts an interface to a time.Time type.
func ToTimeE(i interface{}) (time.Time, error) {
	i = indirect(i)

	switch v := i.(type) {
	case time.Time:
		return v, nil
	case string:
		return StringToDateE(v)
	case int:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	default:
		return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", i, i)
	}
}

// StringToDate attempts to parse a string into a time.Time type using a
// predefined list of formats. If no suitable format is found, an error is
// returned. Any errors are discarded
func StringToDate(s string) time.Time {
	ret, _ := StringToDateE(s)
	return ret
}

// StringToDateE attempts to parse a string into a time.Time type using a
// predefined list of formats. If no suitable format is found, an error is
// returned.
func StringToDateE(s string) (time.Time, error) {
	return parseDateWith(s, []string{
		"02 Jan 2006",
		ISO8601Date,
		ISO8601DateTime1, // iso8601 without timezone
		ISO8601DateTime2, // iso8601 without timezone
		ISO8601DateTimeOffset1,
		ISO8601DateTimeOffset2,
		ISO8601DateTimeOffset3,
		time.ANSIC,
		time.Kitchen,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RubyDate,
		time.Stamp,
		time.StampMicro,
		time.StampMilli,
		time.StampNano,
		time.UnixDate,
		TimeString, // Time.String() output
	})
}

const (
	// ISO8601Date defines the ISO-8601 date format.
	ISO8601Date = "2006-01-02"
	// ISO8601DateTime1 defines the ISO-8601 date/time format without a
	// timezone offset.
	ISO8601DateTime1 = "2006-01-02 15:04:05"
	// ISO8601DateTime2 defines the ISO-8601 date/time format without a
	// timezone offset.
	ISO8601DateTime2 = "2006-01-02T15:04:05"
	// ISO8601DateTimeOffset1 defines the ISO-8601 date/time format with a
	// timezone offset.
	ISO8601DateTimeOffset1 = "2006-01-02 15:04:05Z07:00"
	// ISO8601DateTimeOffset2 defines the ISO-8601 date/time format with a
	// timezone offset.
	ISO8601DateTimeOffset2 = "2006-01-02 15:04:05 -07:00"
	// ISO8601DateTimeOffset3 defines the ISO-8601 date/time format with a
	// timezone offset.
	ISO8601DateTimeOffset3 = "2006-01-02 15:04:05 -0700"
	// TimeString defines the Time.String() date/time format.
	TimeString = "2006-01-02 15:04:05.999999999 -0700 MST"
)
