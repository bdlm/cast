package cast

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

/*
jsonStringToObject attempts to unmarshall a string as JSON into the object
passed as pointer.
*/
func jsonStringToObject(s string, v interface{}) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}

/*
indirect returns the value, after dereferencing as many times as necessary
to reach the base type (or nil).

From html/template/content.go
Copyright 2011 The Go Authors. All rights reserved.
*/
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

/*
indirectToStringerOrError returns the value, after dereferencing as many
times as necessary to reach the base type (or nil) or an implementation of
fmt.Stringer or error.

From html/template/content.go
Copyright 2011 The Go Authors. All rights reserved.
*/
func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

/*
parseDateWith
*/
func parseDateWith(s string, dates []string) (time.Time, error) {
	var t time.Time
	var err error
	for _, dateType := range dates {
		if t, err = time.Parse(dateType, s); err == nil {
			return t, err
		}
	}
	return t, fmt.Errorf("unable to parse date: %s", s)
}
