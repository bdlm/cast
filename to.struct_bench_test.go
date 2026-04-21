package cast_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/bdlm/cast/v2"
)

// ── shared types ──────────────────────────────────────────────────────────────

type benchFlat struct {
	Name   string
	Age    int
	Score  float64
	Active bool
}

type benchAddress struct {
	Street string
	City   string
	Zip    string
}

type benchPerson struct {
	Name    string
	Age     int
	Address benchAddress
	Tags    []string
}

type benchOrder struct {
	ID       int
	Product  string
	Quantity int
	Price    float64
	Buyer    benchAddress
	Notes    []string
}

type benchLeaf struct {
	Value int
	Label string
}

type benchMid struct {
	Leaf  benchLeaf
	Count int
	Name  string
}

type benchDeep struct {
	Mid    benchMid
	Title  string
	Active bool
}

type benchRich struct {
	Name     string
	Score    float64
	Birthday time.Time
	Balance  *big.Int
}

type benchLarge struct {
	F1, F2, F3, F4, F5  string
	F6, F7, F8, F9, F10 int
	G1, G2, G3, G4, G5  float64
	H1, H2, H3, H4, H5  bool
}

type benchEmbedBase struct {
	ID   int
	Role string
}

type benchEmbedded struct {
	benchEmbedBase
	Name string
}

type benchEmbeddedPtr struct {
	*benchEmbedBase
	Name string
}

type benchProduct struct {
	Name  string
	Price float64
	Tags  []string
}

type benchCart struct {
	Owner    benchFlat
	Items    []benchProduct
	Subtotal float64
}

// ── shared source data ────────────────────────────────────────────────────────

var (
	benchFlatMap = map[string]any{
		"Name": "Alice", "Age": "30", "Score": "9.5", "Active": "true",
	}
	benchFlatStringMap = map[string]string{
		"Name": "Alice", "Age": "30", "Score": "9.5", "Active": "true",
	}
	benchFlatSrc = benchFlat{Name: "Alice", Age: 30, Score: 9.5, Active: true}

	benchPersonMap = map[string]any{
		"Name": "Bob",
		"Age":  "25",
		"Address": map[string]any{
			"Street": "123 Main St", "City": "Springfield", "Zip": "12345",
		},
		"Tags": []any{"admin", "user", "moderator"},
	}

	benchOrderMap = map[string]any{
		"ID": "42", "Product": "Widget", "Quantity": "5", "Price": "19.99",
		"Buyer": map[string]any{
			"Street": "456 Elm St", "City": "Shelbyville", "Zip": "67890",
		},
		"Notes": []any{"rush", "gift-wrap"},
	}

	benchDeepMap = map[string]any{
		"Title":  "report",
		"Active": "true",
		"Mid": map[string]any{
			"Name":  "mid-level",
			"Count": "7",
			"Leaf":  map[string]any{"Value": "42", "Label": "leaf"},
		},
	}

	benchRichMap = map[string]any{
		"Name":     "Carol",
		"Score":    "98.6",
		"Birthday": "1990-06-15",
		"Balance":  "999999999999999999999",
	}

	benchLargeMap = map[string]any{
		"F1": "a", "F2": "b", "F3": "c", "F4": "d", "F5": "e",
		"F6": "1", "F7": "2", "F8": "3", "F9": "4", "F10": "5",
		"G1": "1.1", "G2": "2.2", "G3": "3.3", "G4": "4.4", "G5": "5.5",
		"H1": "true", "H2": "false", "H3": "true", "H4": "false", "H5": "true",
	}

	benchEmbeddedMap    = map[string]any{"ID": "7", "Role": "admin", "Name": "Dave"}
	benchEmbeddedPtrMap = map[string]any{"ID": "8", "Role": "user", "Name": "Eve"}

	benchCartMap = map[string]any{
		"Subtotal": "59.97",
		"Owner":    map[string]any{"Name": "Frank", "Age": "40", "Score": "7.0", "Active": "true"},
		"Items": []any{
			map[string]any{"Name": "Widget", "Price": "9.99", "Tags": []any{"sale"}},
			map[string]any{"Name": "Gadget", "Price": "24.99", "Tags": []any{"new", "featured"}},
			map[string]any{"Name": "Doohickey", "Price": "24.99", "Tags": []any{"clearance"}},
		},
	}
)

func makeProductMaps(n int) []any {
	out := make([]any, n)
	for i := range out {
		out[i] = map[string]any{
			"Name":  "Item",
			"Price": "9.99",
			"Tags":  []any{"tag1", "tag2"},
		}
	}
	return out
}

// ── flat struct ───────────────────────────────────────────────────────────────

// Baseline: map[string]any → flat struct. All values are strings that require
// casting, exercising the full normalise→hydrate→castToType path.
func BenchmarkToStructE_flat_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchFlat](benchFlatMap)
	}
}

// map[string]string source: exercises key-cast + normaliseToStringMap.
func BenchmarkToStructE_flat_fromStringMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchFlat](benchFlatStringMap)
	}
}

// Struct-to-struct: source already has the right types; measures reflection
// overhead with no string parsing.
func BenchmarkToStructE_flat_fromStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchFlat](benchFlatSrc)
	}
}

// STRICT mode adds an unused-key scan after hydration.
func BenchmarkToStructE_flat_strict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchFlat](benchFlatMap, cast.Op{Flag: cast.STRICT, Val: true})
	}
}

// STRICT mode struct→struct: exercises checkUnusedSourceFields path.
func BenchmarkToStructE_flat_fromStruct_strict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchFlat](benchFlatSrc, cast.Op{Flag: cast.STRICT, Val: true})
	}
}

// Via the generic ToE entry point instead of ToStructE.
func BenchmarkToE_flat_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[benchFlat](benchFlatMap)
	}
}

// ── nested struct ─────────────────────────────────────────────────────────────

// One level of nesting: benchPerson contains benchAddress and []string.
func BenchmarkToStructE_nested_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchPerson](benchPersonMap)
	}
}

// Two fields that themselves contain a nested struct (benchOrder.Buyer).
func BenchmarkToStructE_nestedWithSlice_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchOrder](benchOrderMap)
	}
}

// ── deeply nested struct ──────────────────────────────────────────────────────

// Three levels deep: benchDeep → benchMid → benchLeaf.
func BenchmarkToStructE_deep_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchDeep](benchDeepMap)
	}
}

// ── special field types ───────────────────────────────────────────────────────

// time.Time and *big.Int fields stress the namedConverters path.
func BenchmarkToStructE_richTypes_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchRich](benchRichMap)
	}
}

// ── large struct ──────────────────────────────────────────────────────────────

// 20 fields of mixed types: measures per-field reflection overhead at scale.
func BenchmarkToStructE_large_fromMap(b *testing.B) {
	b.SetBytes(int64(len(benchLargeMap)))
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchLarge](benchLargeMap)
	}
}

// ── embedded (anonymous) struct fields ───────────────────────────────────────

// Unexported embedded struct: promoted fields must be found via reflection.
func BenchmarkToStructE_embedded_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchEmbedded](benchEmbeddedMap)
	}
}

// Nil pointer-to-struct embedded field: allocation + recursion path.
func BenchmarkToStructE_embeddedPtr_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchEmbeddedPtr](benchEmbeddedPtrMap)
	}
}

// ── slice of structs ──────────────────────────────────────────────────────────

// A struct containing a []benchProduct — exercises castToSliceType with struct
// elements, which calls castToStructType for each item.
func BenchmarkToStructE_sliceOfStructs_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchCart](benchCartMap)
	}
}

func BenchmarkToStructE_sliceOfStructs_10(b *testing.B) {
	src := map[string]any{
		"Subtotal": "0",
		"Owner":    benchFlatMap,
		"Items":    makeProductMaps(10),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchCart](src)
	}
}

func BenchmarkToStructE_sliceOfStructs_100(b *testing.B) {
	src := map[string]any{
		"Subtotal": "0",
		"Owner":    benchFlatMap,
		"Items":    makeProductMaps(100),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToStructE[benchCart](src)
	}
}
