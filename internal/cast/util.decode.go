package cast

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// looksLikeCollection is a cheap syntactic pre-check that reports whether s
// is likely a JSON object or array based on its first non-whitespace character.
// Used to avoid calling the more expensive unmarshalCollection for strings that
// are clearly not JSON collections.
func looksLikeCollection(s string) bool {
	s = strings.TrimSpace(s)
	return len(s) > 1 && (s[0] == '[' || s[0] == '{')
}

// unmarshalCollection attempts to decode s as a JSON array or object.
// It returns the decoded value and true only when decoding succeeds and the
// result is a slice or map; scalars (e.g. a bare "42") are rejected so that
// numeric strings don't accidentally produce single-element collections.
func unmarshalCollection(s string) (any, bool) {
	var v any
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return nil, false
	}
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return nil, false
	}
	k := rv.Kind()
	return v, k == reflect.Slice || k == reflect.Map
}

// tryDecodeJSON attempts to JSON-decode val when DECODE=json is set and val is
// a string, error, or non-named-scalar Stringer source. Returns:
//   - (decoded, true, nil): JSON decode succeeded; caller should use decoded
//   - (nil, false, nil): DECODE not set, source is not a string/error/Stringer,
//     or source is a named scalar struct — caller proceeds with original val
//   - (nil, false, err): source is a string/error/Stringer but JSON unmarshal
//     failed — caller should return an error
func tryDecodeJSON(val any, ops Ops) (any, bool, error) {
	if !ops.HasDecode || ops.DecodeVal != "json" {
		return nil, false, nil
	}
	var srcStr string
	switch v := val.(type) {
	case string:
		srcStr = v
	case error:
		srcStr = v.Error()
	case fmt.Stringer:
		t := reflect.TypeOf(val)
		if t != nil && t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t != nil && isNamedScalarStructType(t) {
			return nil, false, nil
		}
		srcStr = v.String()
	default:
		return nil, false, nil
	}
	var decoded any
	if err := json.Unmarshal([]byte(srcStr), &decoded); err != nil {
		return nil, false, fmt.Errorf("DECODE=json: unable to decode %#.10v as JSON: %v", val, err)
	}
	return decoded, true, nil
}
