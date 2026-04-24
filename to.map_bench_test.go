package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

type benchPoint struct {
	X, Y int
	Name string
}

var (
	benchMapSS = map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "5"}
	benchMapSI = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	benchStruct = benchPoint{X: 3, Y: 4, Name: "origin"}
)

func BenchmarkTo_map_fromMap_strStr_to_strInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[map[string]int](benchMapSS)
	}
}

func BenchmarkTo_map_fromMap_strInt_to_strStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[map[string]string](benchMapSI)
	}
}

func BenchmarkTo_map_fromStruct_to_strAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[map[string]any](benchStruct)
	}
}

func BenchmarkTo_map_fromStruct_private(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[map[string]any](benchStruct, cast.Op{Flag: cast.PRIVATE, Val: true})
	}
}

func BenchmarkTo_map_fromSlice_to_intStr(b *testing.B) {
	src := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[map[int]string](src)
	}
}

var benchLargeStrMap = makeLargeStrMap(50)

func makeLargeStrMap(n int) map[string]string {
	m := make(map[string]string, n)
	for i := 0; i < n; i++ {
		m[makeKey(i)] = "42"
	}
	return m
}

func makeKey(i int) string {
	keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	return keys[i%10] + keys[(i/10)%10]
}

func BenchmarkTo_map_fromLargeMap_50(b *testing.B) {
	b.SetBytes(int64(len(benchLargeStrMap)))
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[map[string]int](benchLargeStrMap)
	}
}

type benchNestedPoint struct {
	Inner benchPoint
	Label string
}

var benchNestedStruct = benchNestedPoint{
	Inner: benchPoint{X: 1, Y: 2, Name: "inner"},
	Label: "outer",
}

func BenchmarkTo_map_fromNestedStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[map[string]any](benchNestedStruct)
	}
}
