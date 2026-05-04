package cast_test

import (
	"net/url"
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_url_fromSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[*url.URL]("https://example.com")
	}
}

func BenchmarkTo_url_fromFull(b *testing.B) {
	// Full URL with path, query, and fragment exercises url.Parse more completely.
	for i := 0; i < b.N; i++ {
		_ = cast.To[*url.URL]("https://user:pass@example.com:8080/path?q=1&r=2#frag")
	}
}

func BenchmarkTo_url_passthrough(b *testing.B) {
	src, _ := url.Parse("https://example.com")
	for i := 0; i < b.N; i++ {
		_ = cast.To[*url.URL](src)
	}
}

func BenchmarkToE_url_fromInvalid(b *testing.B) {
	// url.Parse almost never errors (it's very permissive), so this hits the
	// nil-source error path instead.
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[*url.URL](nil)
	}
}
