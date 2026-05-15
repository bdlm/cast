package cast

import (
	"fmt"
	"net/url"
)

// toURL converts v to *url.URL.
//
// String values are parsed with url.Parse. A *url.URL or url.URL value is
// returned (or copied) directly.
//
// Options:
//   - DEFAULT: *url.URL, value to return on error.
func toURL(v any, ops Ops) (any, error) {
	var ret any = (*url.URL)(nil)

	if ops.HasDefault {
		u, ok := ops.DefaultVal.(*url.URL)
		if !ok {
			return ret, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.DefaultVal)
		}
		ret = u
	}

	switch val := v.(type) {
	case nil:
		return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, (*url.URL)(nil))
	case *url.URL:
		return val, nil
	case url.URL:
		return &val, nil
	case string:
		u, err := url.Parse(val)
		if err != nil {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, (*url.URL)(nil))
		}
		return u, nil
	default:
		if s, err := ToString(v, ops.Delete(DEFAULT)); err == nil {
			if u, urlErr := url.Parse(s); urlErr == nil {
				return u, nil
			}
		}
	}

	return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, (*url.URL)(nil))
}
