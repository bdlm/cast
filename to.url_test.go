package cast_test

import (
	"errors"
	"net/url"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestToEURL(t *testing.T) {
	mustParse := func(s string) *url.URL {
		u, err := url.Parse(s)
		if err != nil {
			panic(err)
		}
		return u
	}

	cases := []struct {
		name      string
		in        any
		expectStr string // compare via URL.String()
		expectErr bool
	}{
		{name: "https URL", in: "https://example.com/path?q=1", expectStr: "https://example.com/path?q=1"},
		{name: "http URL", in: "http://user:pass@host:8080/p", expectStr: "http://user:pass@host:8080/p"},
		{name: "relative URL", in: "/foo/bar", expectStr: "/foo/bar"},
		{name: "empty string", in: "", expectStr: ""},
		{name: "*url.URL direct", in: mustParse("https://go.dev"), expectStr: "https://go.dev"},
		{name: "url.URL value", in: *mustParse("https://go.dev"), expectStr: "https://go.dev"},

		// Error cases
		{name: "nil", in: nil, expectErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[*url.URL](tc.in)
			if err != nil && !tc.expectErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectErr {
				t.Error("expected error, got nil")
			} else if err != nil && !errors.Is(err, cast.ErrorUnableToCast) {
				t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
			} else if err == nil && result.String() != tc.expectStr {
				t.Errorf("expected %q, got %q", tc.expectStr, result.String())
			}
		})
	}
}

func TestToURL(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := cast.To[*url.URL]("https://example.com")
		if result == nil || result.Host != "example.com" {
			t.Errorf("unexpected: %v", result)
		}
	})
	t.Run("error returns nil", func(t *testing.T) {
		result := cast.To[*url.URL](nil)
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})
}

func TestToEURLDefault(t *testing.T) {
	def, _ := url.Parse("https://fallback.example.com")
	result, err := cast.ToE[*url.URL](nil, cast.Op{Flag: cast.DEFAULT, Val: def})
	if err == nil {
		t.Error("expected error, got nil")
	}
	if result == nil || result.Host != def.Host {
		t.Errorf("expected default %v, got %v", def, result)
	}
}

func TestToEURLInvalidDefault(t *testing.T) {
	// A non-*url.URL DEFAULT value must cause an error even with valid input.
	_, err := cast.ToE[*url.URL]("https://example.com", cast.Op{Flag: cast.DEFAULT, Val: "wrong-type"})
	if err == nil {
		t.Error("expected error for non-*url.URL DEFAULT, got nil")
	}
	if !errors.Is(err, cast.ErrorUnableToCast) {
		t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
	}
}

func TestToEURLDefaultCase(t *testing.T) {
	// Inputs that are not nil/*url.URL/url.URL/string route through the default:
	// branch of toURL, which tries toString then url.Parse.
	t.Run("int source converts to relative URL via default branch", func(t *testing.T) {
		// toString(42) → "42", which url.Parse accepts as a relative path.
		result, err := cast.ToE[*url.URL](int(42))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil {
			t.Fatal("expected non-nil URL")
		}
		if result.String() != "42" {
			t.Errorf("expected URL path \"42\", got %q", result.String())
		}
	})
}

// TestToEURLStringerDefaultBranch validates that fmt.Stringer values reach
// toURL via the default: branch (toString → url.Parse).
func TestToEURLStringerDefaultBranch(t *testing.T) {
	t.Run("Stringer with valid URL succeeds", func(t *testing.T) {
		result, err := cast.ToE[*url.URL](testStringer{"https://stringer.example.com"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil || result.Host != "stringer.example.com" {
			t.Errorf("expected host stringer.example.com, got %v", result)
		}
	})
	// url.Parse accepts almost any string as a relative reference, so even
	// clearly-not-a-URL strings succeed. This is documented behavior.
	t.Run("Stringer with non-URL string produces relative URL (url.Parse is permissive)", func(t *testing.T) {
		result, err := cast.ToE[*url.URL](testStringer{"not a url"})
		// url.Parse does not error on this input; it returns a URL with spaces
		// or a path component. The test just asserts no panic and non-nil result.
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil {
			t.Error("expected non-nil URL for permissive url.Parse input")
		}
	})
}

func TestToEURLStructField(t *testing.T) {
	type Endpoint struct {
		URL  *url.URL
		Name string
	}
	result, err := cast.ToStructE[Endpoint](map[string]any{
		"URL":  "https://api.example.com/v1",
		"Name": "primary",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.URL == nil || result.URL.Host != "api.example.com" {
		t.Errorf("unexpected URL: %v", result.URL)
	}
	if result.Name != "primary" {
		t.Errorf("expected Name=primary, got %q", result.Name)
	}
}
