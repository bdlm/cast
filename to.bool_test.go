package cast_test

import (
	"reflect"
	"testing"

	"github.com/bdlm/cast/v2"
	"github.com/bdlm/errors/v2"
)

// ToBool casts an interface to a bool type.
func TestBoolToBool(t *testing.T) {
	a, err := cast.ToE[bool](true)
	if err != nil || a != true {
		t.Errorf("ToE[bool](true) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](true) failed"))
	}
	a, err = cast.ToE[bool](false)
	if err != nil || a != false {
		t.Errorf("ToE[bool](false) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](false) failed"))
	}
}

func TestByteToBool(t *testing.T) {
	a, err := cast.ToE[bool](byte(1))
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](byte(1)) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](byte(1)) failed"))
	}
	a, err = cast.ToE[bool](byte(0))
	if err != nil || a != false || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](byte(0)) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](byte(0)) failed"))
	}
}

func TestComplexToBool(t *testing.T) {
	a, err := cast.ToE[bool](complex(1, 0))
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](complex(1, 0)) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](complex(1, 0)) failed"))
	}
	a, err = cast.ToE[bool](complex(1, 1))
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](complex(1, 1)) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](complex(1, 1)) failed"))
	}
	a, err = cast.ToE[bool](complex(0, 0))
	if err != nil || a != false || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](complex(0, 0)) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](complex(0, 0)) failed"))
	}
	a, err = cast.ToE[bool](complex(-1, 0))
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](complex(-1, 0)) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](complex(-1, 0)) failed"))
	}
	a, err = cast.ToE[bool](complex(-1, -1))
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](complex(-1, -1)) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](complex(-1, -1)) failed"))
	}
}

func TestFloatToBool(t *testing.T) {
	a, err := cast.ToE[bool](1.0)
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](1.0) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](1.0) failed"))
	}
	a, err = cast.ToE[bool](1.1)
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](1.1) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](1.1) failed"))
	}
	a, err = cast.ToE[bool](0.0)
	if err != nil || a != false || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](0.0) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](0.0) failed"))
	}
	a, err = cast.ToE[bool](-1.0)
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](-1.0) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](-1.0) failed"))
	}
	a, err = cast.ToE[bool](-1.1)
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](-1.1) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](-1.1) failed"))
	}
}

func TestIntToBool(t *testing.T) {
	a, err := cast.ToE[bool](1)
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](1) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](1) failed"))
	}
	a, err = cast.ToE[bool](0)
	if err != nil || a != false || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](0) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](0) failed"))
	}
	a, err = cast.ToE[bool](-1)
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool](-1) failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool](-1) failed"))
	}
}

func TestStringToBool(t *testing.T) {
	a, err := cast.ToE[bool]("1")
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool]('1') failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool]('1') failed"))
	}
	a, err = cast.ToE[bool]("0")
	if err != nil || a != false || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool]('0') failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool]('0') failed"))
	}
	a, err = cast.ToE[bool]("-1")
	if err != nil || a != true || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool]('-1') failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool]('-1') failed"))
	}
	a, err = cast.ToE[bool]("Hi!")
	if err == nil || a != false || reflect.TypeOf(a).Kind() != reflect.Bool {
		t.Errorf("ToE[bool]('-1') failed: %#v, %#v", err, errors.Wrap(err, "ToE[bool]('-1') failed"))
	}
}
