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
	v3, e3 := cast.ToStructE[withSlice](map[string]any{"Tags": "hello"})
	fmt.Printf("struct []string field from scalar \"hello\": %v, err=%v\n", v3.Tags, e3)

	// 3. Map field in struct — does castToType handle reflect.Map?
	type withMap struct{ Labels map[string]string }
	v4, e4 := cast.ToStructE[withMap](map[string]any{"Labels": map[string]string{"k": "v"}})
	fmt.Printf("struct map field: %v, err=%v\n", v4.Labels, e4)

	// 4. []byte struct field from string
	type withBytes struct{ Data []byte }
	v5, e5 := cast.ToStructE[withBytes](map[string]any{"Data": "world"})
	fmt.Printf("struct []byte field from string: %v, err=%v\n", v5.Data, e5)

	// 5. []byte struct field from []byte directly
	v6, e6 := cast.ToStructE[withBytes](map[string]any{"Data": []byte("world")})
	fmt.Printf("struct []byte field from []byte: %v, err=%v\n", v6.Data, e6)
}
