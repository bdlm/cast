package cast_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestToERegexp(t *testing.T) {
	cases := []struct {
		name      string
		in        any
		pattern   string // expected regexp pattern string
		expectErr bool
	}{
		{name: "simple pattern", in: `\d+`, pattern: `\d+`},
		{name: "anchored", in: `^foo.*bar$`, pattern: `^foo.*bar$`},
		{name: "empty pattern", in: "", pattern: ""},
		{name: "*regexp.Regexp direct", in: regexp.MustCompile(`[a-z]+`), pattern: `[a-z]+`},

		// Error cases
		{name: "nil", in: nil, expectErr: true},
		{name: "invalid pattern", in: `[unclosed`, expectErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[*regexp.Regexp](tc.in)
			if err != nil && !tc.expectErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectErr {
				t.Error("expected error, got nil")
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Errorf("expected cast.Error, got %T: %v", err, err)
			} else if err == nil && result.String() != tc.pattern {
				t.Errorf("expected pattern %q, got %q", tc.pattern, result.String())
			}
		})
	}
}

func TestToRegexp(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := cast.To[*regexp.Regexp](`\w+`)
		if result == nil || result.String() != `\w+` {
			t.Errorf("unexpected: %v", result)
		}
	})
	t.Run("error returns nil", func(t *testing.T) {
		result := cast.To[*regexp.Regexp](nil)
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})
}

func TestToERegexpDefault(t *testing.T) {
	def := regexp.MustCompile(`.*`)
	result, err := cast.ToE[*regexp.Regexp](nil, cast.Op{Flag: cast.DEFAULT, Val: def})
	if err == nil {
		t.Error("expected error, got nil")
	}
	if result == nil || result.String() != def.String() {
		t.Errorf("expected default %v, got %v", def, result)
	}
}

func TestToERegexpMatch(t *testing.T) {
	// Verify the compiled regexp actually works.
	result, err := cast.ToE[*regexp.Regexp](`^\d{3}-\d{4}$`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.MatchString("555-1234") {
		t.Error("expected pattern to match 555-1234")
	}
	if result.MatchString("abc") {
		t.Error("expected pattern to not match abc")
	}
}

func TestToERegexpStructField(t *testing.T) {
	type Filter struct {
		Pattern *regexp.Regexp
		Label   string
	}
	result, err := cast.ToStructE[Filter](map[string]any{
		"Pattern": `^test_`,
		"Label":   "prefix filter",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Pattern == nil || result.Pattern.String() != `^test_` {
		t.Errorf("unexpected Pattern: %v", result.Pattern)
	}
}
