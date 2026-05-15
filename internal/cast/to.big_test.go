package cast_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/bdlm/cast/v2"
)

// ── *big.Int ──────────────────────────────────────────────────────────────────

func TestToEBigInt(t *testing.T) {
	cases := []struct {
		name      string
		in        any
		expectStr string
		expectErr bool
	}{
		// String parsing (base auto-detect)
		{name: "decimal string", in: "12345678901234567890", expectStr: "12345678901234567890"},
		{name: "hex string", in: "0xff", expectStr: "255"},
		{name: "octal string", in: "0o77", expectStr: "63"},
		{name: "negative string", in: "-42", expectStr: "-42"},
		{name: "zero string", in: "0", expectStr: "0"},

		// *big.Int passthrough (copied)
		{name: "*big.Int direct", in: big.NewInt(999), expectStr: "999"},

		// big.Int value
		{name: "big.Int value", in: *big.NewInt(42), expectStr: "42"},

		// Integer sources
		{name: "int", in: int(100), expectStr: "100"},
		{name: "int64", in: int64(-9999999999999999), expectStr: "-9999999999999999"},
		{name: "uint64", in: uint64(18446744073709551615), expectStr: "18446744073709551615"},

		// Float sources (truncated toward zero)
		{name: "float64 truncates", in: float64(3.9), expectStr: "3"},
		{name: "float64 negative truncates", in: float64(-3.9), expectStr: "-3"},
		{name: "*big.Float", in: new(big.Float).SetFloat64(7.8), expectStr: "7"},

		// Error cases
		{name: "nil", in: nil, expectErr: true},
		{name: "invalid string", in: "not-a-number", expectErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[*big.Int](tc.in)
			if err != nil && !tc.expectErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectErr {
				t.Error("expected error, got nil")
			} else if err != nil && !errors.Is(err, cast.ErrorUnableToCast) {
				t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
			} else if err == nil && result.String() != tc.expectStr {
				t.Errorf("expected %s, got %s", tc.expectStr, result.String())
			}
		})
	}
}

func TestToBigInt(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := cast.To[*big.Int]("999999999999999999999")
		if result == nil || result.String() != "999999999999999999999" {
			t.Errorf("unexpected: %v", result)
		}
	})
	t.Run("error returns nil", func(t *testing.T) {
		result := cast.To[*big.Int]("bad")
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})
}

func TestToEBigIntDefault(t *testing.T) {
	def := big.NewInt(-1)
	result, err := cast.ToE[*big.Int]("bad", cast.Op{Flag: cast.DEFAULT, Val: def})
	if err == nil {
		t.Error("expected error, got nil")
	}
	if result == nil || result.Cmp(def) != 0 {
		t.Errorf("expected default %v, got %v", def, result)
	}
}

func TestToEBigIntStructField(t *testing.T) {
	type Account struct {
		Balance *big.Int
		Name    string
	}
	result, err := cast.ToE[Account](map[string]any{
		"Balance": "999999999999999999999",
		"Name":    "whale",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Balance == nil || result.Balance.String() != "999999999999999999999" {
		t.Errorf("unexpected Balance: %v", result.Balance)
	}
}

// ── *big.Float ────────────────────────────────────────────────────────────────

func TestToEBigFloat(t *testing.T) {
	eqStr := func(f *big.Float, s string) bool {
		return f != nil && f.Text('f', -1) == s
	}

	cases := []struct {
		name      string
		in        any
		expectStr string
		expectErr bool
	}{
		// String parsing
		{name: "decimal string", in: "3.14159265358979323846", expectStr: "3.1415926535897932385"},
		{name: "integer string", in: "42", expectStr: "42"},
		{name: "negative string", in: "-1.5", expectStr: "-1.5"},
		{name: "zero string", in: "0", expectStr: "0"},

		// *big.Float passthrough (copied)
		{name: "*big.Float direct", in: func() *big.Float { f, _ := new(big.Float).SetString("2.718"); return f }(), expectStr: "2.718"},

		// *big.Int source
		{name: "*big.Int", in: big.NewInt(100), expectStr: "100"},

		// Integer sources
		{name: "int", in: int(7), expectStr: "7"},
		{name: "int64", in: int64(-500), expectStr: "-500"},
		{name: "uint64", in: uint64(1000), expectStr: "1000"},

		// Float sources
		{name: "float64", in: float64(1.5), expectStr: "1.5"},
		{name: "float32", in: float32(2.5), expectStr: "2.5"},

		// Error cases
		{name: "nil", in: nil, expectErr: true},
		{name: "invalid string", in: "not-a-float", expectErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[*big.Float](tc.in)
			if err != nil && !tc.expectErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectErr {
				t.Error("expected error, got nil")
			} else if err != nil && !errors.Is(err, cast.ErrorUnableToCast) {
				t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
			} else if err == nil && !eqStr(result, tc.expectStr) {
				t.Errorf("expected %s, got %s", tc.expectStr, result.Text('f', -1))
			}
		})
	}
}

func TestToBigFloat(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := cast.To[*big.Float]("3.14")
		if result == nil {
			t.Fatal("expected non-nil")
		}
		f64, _ := result.Float64()
		if f64 < 3.13 || f64 > 3.15 {
			t.Errorf("unexpected: %v", result)
		}
	})
	t.Run("error returns nil", func(t *testing.T) {
		result := cast.To[*big.Float]("bad")
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})
}

func TestToEBigFloatDefault(t *testing.T) {
	def := new(big.Float).SetFloat64(-1.0)
	result, err := cast.ToE[*big.Float]("bad", cast.Op{Flag: cast.DEFAULT, Val: def})
	if err == nil {
		t.Error("expected error, got nil")
	}
	if result == nil || result.Cmp(def) != 0 {
		t.Errorf("expected default %v, got %v", def, result)
	}
}

func TestToEBigFloatStructField(t *testing.T) {
	type Measurement struct {
		Value *big.Float
		Unit  string
	}
	result, err := cast.ToE[Measurement](map[string]any{
		"Value": "3.14159265358979323846",
		"Unit":  "meters",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Value == nil {
		t.Fatal("expected non-nil Value")
	}
	if result.Value.Text('f', 5) != "3.14159" {
		t.Errorf("unexpected Value: %s", result.Value.Text('f', 5))
	}
}

// ── missing value (non-pointer) sources and stringer default branches ─────────

func TestToEBigIntMissingCases(t *testing.T) {
	t.Run("big.Float value (not pointer) truncates to int", func(t *testing.T) {
		var f big.Float
		f.SetFloat64(9.9)
		result, err := cast.ToE[*big.Int](f)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil || result.Int64() != 9 {
			t.Errorf("expected 9, got %v", result)
		}
	})

	t.Run("stringer default branch success", func(t *testing.T) {
		// testStringer returns "42"; toBigInt's default: branch calls toString
		// → "42" → SetString("42", 0) succeeds.
		result, err := cast.ToE[*big.Int](testStringer{"42"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil || result.Int64() != 42 {
			t.Errorf("expected 42, got %v", result)
		}
	})

	t.Run("stringer default branch error", func(t *testing.T) {
		_, err := cast.ToE[*big.Int](testStringer{"not-a-number"})
		if err == nil {
			t.Error("expected error for invalid big.Int string, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
		}
	})
}

func TestToEBigFloatMissingCases(t *testing.T) {
	t.Run("big.Int value (not pointer) converts exactly", func(t *testing.T) {
		var i big.Int
		i.SetInt64(42)
		result, err := cast.ToE[*big.Float](i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil {
			t.Fatal("expected non-nil")
		}
		f64, _ := result.Float64()
		if f64 != 42.0 {
			t.Errorf("expected 42.0, got %v", f64)
		}
	})

	t.Run("stringer default branch success", func(t *testing.T) {
		result, err := cast.ToE[*big.Float](testStringer{"3.14"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil {
			t.Fatal("expected non-nil")
		}
		f64, _ := result.Float64()
		if f64 < 3.13 || f64 > 3.15 {
			t.Errorf("expected ~3.14, got %v", f64)
		}
	})

	t.Run("stringer default branch error", func(t *testing.T) {
		_, err := cast.ToE[*big.Float](testStringer{"not-a-float"})
		if err == nil {
			t.Error("expected error for invalid big.Float string, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
		}
	})
}

// TestDefaultBranchBigIntErrors verifies that types without a dedicated case
// in toBigInt route through the default: branch (toString → SetString), and
// that types whose string form is not a valid integer produce an error.
func TestDefaultBranchBigIntErrors(t *testing.T) {
	t.Run("bool → error (\"true\" is not a valid big.Int)", func(t *testing.T) {
		_, err := cast.ToE[*big.Int](true)
		if err == nil {
			t.Error("expected error for bool→*big.Int, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
		}
	})
	t.Run("complex128 → error (\"(1+2i)\" is not a valid big.Int)", func(t *testing.T) {
		_, err := cast.ToE[*big.Int](complex(1, 2))
		if err == nil {
			t.Error("expected error for complex→*big.Int, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
		}
	})
}

// TestDefaultBranchBigFloatErrors verifies that types without a dedicated
// case in toBigFloat route through default: (toString → SetString) and that
// non-numeric string forms produce an error.
func TestDefaultBranchBigFloatErrors(t *testing.T) {
	t.Run("bool → error (\"true\" is not a valid big.Float)", func(t *testing.T) {
		_, err := cast.ToE[*big.Float](true)
		if err == nil {
			t.Error("expected error for bool→*big.Float, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
		}
	})
	t.Run("complex128 → error (\"(1+2i)\" is not a valid big.Float)", func(t *testing.T) {
		_, err := cast.ToE[*big.Float](complex(1, 2))
		if err == nil {
			t.Error("expected error for complex→*big.Float, got nil")
		}
		if !errors.Is(err, cast.ErrorUnableToCast) {
			t.Errorf("expected cast.ErrorUnableToCast, got %T: %v", err, err)
		}
	})
}
