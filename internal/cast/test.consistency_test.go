package cast_test

import (
	"fmt"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestProbeInconsistencies(t *testing.T) {
	// 1. Direct path: does ToE[[]byte] accept a string source?
	v1, e1 := cast.ToE[[]byte]("world")
	fmt.Printf("ToE[[]byte](\"world\"): %v, err=%v\n", v1, e1)

	// 2. Scalar source to slice via struct field (castToSliceType scalar wrapping)
	type withSlice struct{ Tags []string }
	v3, e3 := cast.ToE[withSlice](map[string]any{"Tags": "hello"})
	fmt.Printf("struct []string field from scalar \"hello\": %v, err=%v\n", v3.Tags, e3)

	// 3. Map field in struct — does castToType handle reflect.Map?
	type withMap struct{ Labels map[string]string }
	v4, e4 := cast.ToE[withMap](map[string]any{"Labels": map[string]string{"k": "v"}})
	fmt.Printf("struct map field: %v, err=%v\n", v4.Labels, e4)

	// 4. []byte struct field from string
	type withBytes struct{ Data []byte }
	v5, e5 := cast.ToE[withBytes](map[string]any{"Data": "world"})
	fmt.Printf("struct []byte field from string: %v, err=%v\n", v5.Data, e5)

	// 5. []byte struct field from []byte directly
	v6, e6 := cast.ToE[withBytes](map[string]any{"Data": []byte("world")})
	fmt.Printf("struct []byte field from []byte: %v, err=%v\n", v6.Data, e6)
}

func TestProbeInconsistencies2(t *testing.T) {
	// Scalar source → []string via direct ToE vs struct field
	v1, e1 := cast.ToE[[]string]("hello")
	fmt.Printf("ToE[[]string](\"hello\"): %v, err=%v\n", v1, e1)

	type withSlice struct{ Tags []string }
	v2, e2 := cast.ToE[withSlice](map[string]any{"Tags": "hello"})
	fmt.Printf("struct []string field from \"hello\": %v, err=%v\n", v2.Tags, e2)

	// Nested map as a map value: map[string]map[string]string from map
	v3, e3 := cast.ToE[map[string]map[string]string](
		map[string]map[string]string{"outer": {"k": "v"}},
	)
	fmt.Printf("map[string]map[string]string identity: %v, err=%v\n", v3, e3)

	// Map as a struct field from a map source (key type mismatch but compatible)
	type withMapField struct{ Labels map[string]string }
	v4, e4 := cast.ToE[withMapField](
		map[string]any{"Labels": map[string]string{"k": "v"}},
	)
	fmt.Printf("struct map[string]string field: %v, err=%v\n", v4.Labels, e4)

	// Same but STRICT: should error
	_, e5 := cast.ToE[withMapField](
		map[string]any{"Labels": map[string]string{"k": "v"}},
		cast.Op{cast.STRICT, true},
	)
	fmt.Printf("struct map field strict err: %v\n", e5)
}
