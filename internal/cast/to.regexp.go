package cast

import (
	"fmt"
	"regexp"
)

// toRegexp converts v to *regexp.Regexp.
//
// String values are compiled with regexp.Compile. A *regexp.Regexp value is
// returned directly.
//
// Options:
//   - DEFAULT: *regexp.Regexp, value to return on error.
func toRegexp(v any, ops Ops) (any, error) {
	var ret any = (*regexp.Regexp)(nil)

	if ops.HasDefault {
		r, ok := ops.DefaultVal.(*regexp.Regexp)
		if !ok {
			return ret, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.DefaultVal)
		}
		ret = r
	}

	switch val := v.(type) {
	case nil:
		return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, (*regexp.Regexp)(nil))
	case *regexp.Regexp:
		return val, nil
	case string:
		r, err := regexp.Compile(val)
		if err != nil {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, (*regexp.Regexp)(nil))
		}
		return r, nil
	default:
		if s, err := ToString(v, ops.Delete(DEFAULT)); err == nil {
			if r, reErr := regexp.Compile(s); reErr == nil {
				return r, nil
			}
		}
	}

	return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, (*regexp.Regexp)(nil))
}
