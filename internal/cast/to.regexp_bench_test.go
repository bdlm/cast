package cast_test

import (
	"regexp"
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_regexp_fromSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[*regexp.Regexp](`\d+`)
	}
}

func BenchmarkTo_regexp_fromComplex(b *testing.B) {
	// More complex pattern exercises regexp.Compile more heavily.
	for i := 0; i < b.N; i++ {
		_ = cast.To[*regexp.Regexp](`^([a-z]+)(\d{2,4})(?:[-_]\w+)?$`)
	}
}

func BenchmarkTo_regexp_passthrough(b *testing.B) {
	src := regexp.MustCompile(`\d+`)
	for i := 0; i < b.N; i++ {
		_ = cast.To[*regexp.Regexp](src)
	}
}

func BenchmarkToE_regexp_fromInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[*regexp.Regexp](`[invalid`)
	}
}
