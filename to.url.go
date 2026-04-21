package cast

import (
	"net/url"

	"github.com/bdlm/errors/v2"
)

// toURL converts v to *url.URL.
//
// String values are parsed with url.Parse. A *url.URL or url.URL value is
// returned (or copied) directly.
//
// Options:
//   - DEFAULT: *url.URL, value to return on error.
func toURL(v any, ops ops) (any, error) {
	var ret any = (*url.URL)(nil)

	if ops.hasDefault {
		u, ok := ops.defaultVal.(*url.URL)
		if !ok {
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		ret = u
	}

	switch val := v.(type) {
	case nil:
		return ret, errors.Errorf(ErrorStrUnableToCast, v, v, (*url.URL)(nil))
	case *url.URL:
		return val, nil
	case url.URL:
		return &val, nil
	case string:
		u, err := url.Parse(val)
		if err != nil {
			return ret, errors.Errorf(ErrorStrUnableToCast, v, v, (*url.URL)(nil))
		}
		return u, nil
	default:
		if s, err := toString(v, ops.Delete(DEFAULT)); err == nil {
			if u, urlErr := url.Parse(s.(string)); urlErr == nil {
				return u, nil
			}
		}
	}

	return ret, errors.Errorf(ErrorStrUnableToCast, v, v, (*url.URL)(nil))
}
