package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestMyTest(t *testing.T) {
	t.Run("Casting to anonymous struct", func(t *testing.T) {
		result, err := cast.ToE[struct {
			Foo string
			Bar int
		}](map[string]any{
			"Foo": "hello",
			"Bar": 42,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != struct {
			Foo string
			Bar int
		}{
			Foo: "hello",
			Bar: 42,
		} {
			t.Errorf("unexpected result: %v", result)
		}
		t.Logf("successfully cast to anonymous struct: %+v", result)
	})
}
