package cast_test

import (
	"fmt"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestProbeInconsistencies2(t *testing.T) {
	// Scalar source → []string via direct ToE vs struct field
	v1, e1 := cast.ToE[[]string]("hello")
	fmt.Printf("ToE[[]string](\"hello\"): %v, err=%v\n", v1, e1)

	type withSlice struct{ Tags []string }
	v2, e2 := cast.ToStructE[withSlice](map[string]any{"Tags": "hello"})
	fmt.Printf("struct []string field from \"hello\": %v, err=%v\n", v2.Tags, e2)

	// Nested map as a map value: map[string]map[string]string from map
	v3, e3 := cast.ToE[map[string]map[string]string](
		map[string]map[string]string{"outer": {"k": "v"}},
	)
	fmt.Printf("map[string]map[string]string identity: %v, err=%v\n", v3, e3)

	// Map as a struct field from a map source (key type mismatch but compatible)
	type withMapField struct{ Labels map[string]string }
	v4, e4 := cast.ToStructE[withMapField](
		map[string]any{"Labels": map[string]string{"k": "v"}},
	)
	fmt.Printf("struct map[string]string field: %v, err=%v\n", v4.Labels, e4)

	// Same but STRICT: should error
	_, e5 := cast.ToStructE[withMapField](
		map[string]any{"Labels": map[string]string{"k": "v"}},
		cast.Op{cast.STRICT, true},
	)
	fmt.Printf("struct map field strict err: %v\n", e5)
}
