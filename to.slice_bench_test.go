package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

var (
	strSlice10  = makeStrSlice(10)
	strSlice100 = makeStrSlice(100)
	intSlice10  = makeIntSlice(10)
	intSlice100 = makeIntSlice(100)
)

func makeStrSlice(n int) []string {
	s := make([]string, n)
	for i := range s {
		s[i] = "42"
	}
	return s
}

func makeIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}

func BenchmarkTo_sliceInt_fromStrings_10(b *testing.B) {
	b.SetBytes(int64(len(strSlice10)))
	for i := 0; i < b.N; i++ {
		_ = cast.To[[]int](strSlice10)
	}
}

func BenchmarkTo_sliceInt_fromStrings_100(b *testing.B) {
	b.SetBytes(int64(len(strSlice100)))
	for i := 0; i < b.N; i++ {
		_ = cast.To[[]int](strSlice100)
	}
}

func BenchmarkTo_sliceString_fromInts_10(b *testing.B) {
	b.SetBytes(int64(len(intSlice10)))
	for i := 0; i < b.N; i++ {
		_ = cast.To[[]string](intSlice10)
	}
}

func BenchmarkTo_sliceString_fromInts_100(b *testing.B) {
	b.SetBytes(int64(len(intSlice100)))
	for i := 0; i < b.N; i++ {
		_ = cast.To[[]string](intSlice100)
	}
}

func BenchmarkTo_sliceInt_unique_10(b *testing.B) {
	src := make([]int, 10)
	for i := range src {
		src[i] = i % 5 // half are duplicates
	}
	b.SetBytes(int64(len(src)))
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[[]int](src, cast.Op{Flag: cast.UNIQUE_VALUES, Val: true})
	}
}

func BenchmarkTo_sliceInt_unique_100(b *testing.B) {
	src := make([]int, 100)
	for i := range src {
		src[i] = i % 50
	}
	b.SetBytes(int64(len(src)))
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[[]int](src, cast.Op{Flag: cast.UNIQUE_VALUES, Val: true})
	}
}

func BenchmarkTo_sliceInt_withLength(b *testing.B) {
	b.SetBytes(int64(len(strSlice10)))
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[[]int](strSlice10, cast.Op{Flag: cast.LENGTH, Val: 100})
	}
}

// NamedIntSlice and NamedStringSlice are named slice types that hit the reflect
// path in toSlice (the default: branch), as opposed to the concrete-type switch
// that handles []int, []string, etc. directly.
type NamedIntSlice []int
type NamedStringSlice []string

// BenchmarkTo_namedSliceInt_fromStrings measures the reflect path for a named
// []int slice type, to compare against the concrete-type-switch path above.
func BenchmarkTo_namedSliceInt_fromStrings_10(b *testing.B) {
	b.SetBytes(int64(len(strSlice10)))
	for i := 0; i < b.N; i++ {
		_ = cast.To[NamedIntSlice](strSlice10)
	}
}

func BenchmarkTo_namedSliceInt_fromStrings_100(b *testing.B) {
	b.SetBytes(int64(len(strSlice100)))
	for i := 0; i < b.N; i++ {
		_ = cast.To[NamedIntSlice](strSlice100)
	}
}

func BenchmarkTo_namedSliceString_fromInts_10(b *testing.B) {
	b.SetBytes(int64(len(intSlice10)))
	for i := 0; i < b.N; i++ {
		_ = cast.To[NamedStringSlice](intSlice10)
	}
}

func BenchmarkTo_namedSliceString_fromInts_100(b *testing.B) {
	b.SetBytes(int64(len(intSlice100)))
	for i := 0; i < b.N; i++ {
		_ = cast.To[NamedStringSlice](intSlice100)
	}
}
