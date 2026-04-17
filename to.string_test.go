package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestStringJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       any
		expectJSON  string
		expectPlain string
	}{
		{"string", "hello", `"hello"`, "hello"},
		{"string with quotes", `say "hi"`, `"say \"hi\""`, `say "hi"`},
		{"int", 42, `"42"`, "42"},
		{"bool true", true, `"true"`, "true"},
		{"bool false", false, `"false"`, "false"},
		{"float", 3.14, `"3.14"`, "3.14"},
		{"nil", nil, `""`, ""},
		{"empty string", "", `""`, ""},
	}

	for _, test := range tests {
		t.Run(test.name+"/json=true", func(t *testing.T) {
			actual, err := cast.ToE[string](test.input, cast.Op{Flag: cast.JSON, Val: true})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if actual != test.expectJSON {
				t.Errorf("expected %q, got %q", test.expectJSON, actual)
			}
		})
		t.Run(test.name+"/json=false", func(t *testing.T) {
			actual, err := cast.ToE[string](test.input, cast.Op{Flag: cast.JSON, Val: false})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if actual != test.expectPlain {
				t.Errorf("expected %q, got %q", test.expectPlain, actual)
			}
		})
	}
}
