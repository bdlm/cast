package cast_test

import (
	"errors"
	"math/big"
	"net"
	"testing"

	"github.com/bdlm/cast/v2"
)

func TestToENetIP(t *testing.T) {
	cases := []struct {
		name      string
		in        any
		expect    net.IP
		expectErr bool
	}{
		// String parsing
		{name: "IPv4 string", in: "192.168.1.1", expect: net.ParseIP("192.168.1.1")},
		{name: "IPv6 string", in: "::1", expect: net.ParseIP("::1")},
		{name: "IPv4-mapped IPv6", in: "::ffff:192.0.2.1", expect: net.ParseIP("::ffff:192.0.2.1")},

		// net.IP passthrough
		{name: "net.IP direct", in: net.ParseIP("10.0.0.1"), expect: net.ParseIP("10.0.0.1")},

		// []byte
		{name: "[]byte IPv4", in: []byte{192, 168, 0, 1}, expect: net.IP{192, 168, 0, 1}},
		{name: "[]byte IPv6", in: func() []byte {
			ip := net.ParseIP("::1").To16()
			return []byte(ip)
		}(), expect: net.ParseIP("::1").To16()},
		{name: "[]byte as string", in: []byte("127.0.0.1"), expect: net.ParseIP("127.0.0.1")},

		// uint32 → packed IPv4
		{name: "uint32", in: uint32(0xC0A80101), expect: net.IPv4(192, 168, 1, 1)},

		// Error cases
		{name: "nil", in: nil, expectErr: true},
		{name: "invalid string", in: "not-an-ip", expectErr: true},
		{name: "[]byte wrong length", in: []byte{1, 2, 3}, expectErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[net.IP](tc.in)
			if err != nil && !tc.expectErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectErr {
				t.Error("expected error, got nil")
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Errorf("expected cast.Error, got %T: %v", err, err)
			} else if err == nil && !result.Equal(tc.expect) {
				t.Errorf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

func TestToNetIP(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result := cast.To[net.IP]("10.0.0.1")
		if !result.Equal(net.ParseIP("10.0.0.1")) {
			t.Errorf("unexpected: %v", result)
		}
	})
	t.Run("error returns nil", func(t *testing.T) {
		result := cast.To[net.IP]("bad")
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})
}

func TestToENetIPDefault(t *testing.T) {
	def := net.ParseIP("1.2.3.4")
	result, err := cast.ToE[net.IP]("bad", cast.Op{Flag: cast.DEFAULT, Val: def})
	if err == nil {
		t.Error("expected error, got nil")
	}
	if !result.Equal(def) {
		t.Errorf("expected default %v, got %v", def, result)
	}
}

func TestToENetIPInvalidDefault(t *testing.T) {
	// A non-net.IP DEFAULT value must cause an error even when the input is valid.
	_, err := cast.ToE[net.IP]("192.168.1.1", cast.Op{Flag: cast.DEFAULT, Val: "wrong-type"})
	if err == nil {
		t.Error("expected error for non-net.IP DEFAULT, got nil")
	}
	if !errors.Is(err, cast.Error) {
		t.Errorf("expected cast.Error, got %T: %v", err, err)
	}
}

func TestToENetIPDefaultCase(t *testing.T) {
	// Inputs that are not nil/net.IP/string/[]byte/uint32 route through the
	// default: branch of toNetIP, which tries toString then net.ParseIP.
	t.Run("stringer with valid IP succeeds", func(t *testing.T) {
		result, err := cast.ToE[net.IP](testStringer{"10.0.0.1"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !result.Equal(net.ParseIP("10.0.0.1")) {
			t.Errorf("expected 10.0.0.1, got %v", result)
		}
	})
	t.Run("int source (non-IP string) returns error", func(t *testing.T) {
		_, err := cast.ToE[net.IP](int(42))
		if err == nil {
			t.Error("expected error for int→net.IP (not a valid IP string), got nil")
		}
		if !errors.Is(err, cast.Error) {
			t.Errorf("expected cast.Error, got %T: %v", err, err)
		}
	})
}

func TestToENetIPStructField(t *testing.T) {
	type Server struct {
		Addr net.IP
		Port int
	}
	result, err := cast.ToStructE[Server](map[string]any{
		"Addr": "192.168.1.100",
		"Port": "8080",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.Addr.Equal(net.ParseIP("192.168.1.100")) {
		t.Errorf("expected 192.168.1.100, got %v", result.Addr)
	}
	if result.Port != 8080 {
		t.Errorf("expected Port=8080, got %d", result.Port)
	}
}

// TestToENetIPNewSources validates the int32, *big.Int, big.Int, *big.Float,
// and big.Float source paths added in v2.1.1.
func TestToENetIPNewSources(t *testing.T) {
	// Helper for constructing a 16-byte IPv6 address with specific trailing bytes.
	ipv6 := func(trailing ...byte) net.IP {
		ip := make(net.IP, net.IPv6len)
		copy(ip[net.IPv6len-len(trailing):], trailing)
		return ip
	}

	cases := []struct {
		name      string
		in        any
		expect    net.IP
		expectErr bool
	}{
		// int32 → big-endian IPv4 (positive int32 values only; 0x7F000001 = 127.0.0.1)
		{name: "int32 0x01020304 → 1.2.3.4", in: int32(0x01020304), expect: net.IPv4(1, 2, 3, 4)},
		{name: "int32 0x7F000001 → 127.0.0.1", in: int32(0x7F000001), expect: net.IPv4(127, 0, 0, 1)},
		{name: "int32 0 → 0.0.0.0", in: int32(0), expect: net.IPv4(0, 0, 0, 0)},
		{name: "int32 negative → error", in: int32(-1), expectErr: true},

		// *big.Int → 16-byte IPv6 (big-endian, zero-padded)
		{name: "*big.Int zero → all-zeros IPv6", in: big.NewInt(0), expect: ipv6()},
		{name: "*big.Int 1 → ::1-like", in: big.NewInt(1), expect: ipv6(0x01)},
		{name: "*big.Int 0xC0A80101 → IPv6 with IPv4 in last 4 bytes",
			in:     new(big.Int).SetUint64(0xC0A80101),
			expect: ipv6(0xC0, 0xA8, 0x01, 0x01),
		},
		{name: "*big.Int negative → error", in: big.NewInt(-1), expectErr: true},
		{name: "*big.Int too large (>128 bits) → error",
			in:        new(big.Int).Lsh(big.NewInt(1), 128),
			expectErr: true,
		},
		{name: "*big.Int nil → error", in: (*big.Int)(nil), expectErr: true},

		// big.Int value (non-pointer)
		{name: "big.Int value zero → all-zeros IPv6", in: *big.NewInt(0), expect: ipv6()},
		{name: "big.Int value negative → error", in: *big.NewInt(-1), expectErr: true},

		// *big.Float → truncated to *big.Int then to IPv6
		{name: "*big.Float 0 → all-zeros IPv6", in: new(big.Float).SetFloat64(0), expect: ipv6()},
		{name: "*big.Float 1.9 truncated to 1", in: new(big.Float).SetFloat64(1.9), expect: ipv6(0x01)},
		{name: "*big.Float nil → error", in: (*big.Float)(nil), expectErr: true},
		{name: "*big.Float +Inf → error", in: new(big.Float).SetInf(false), expectErr: true},
		{name: "*big.Float -Inf → error", in: new(big.Float).SetInf(true), expectErr: true},

		// big.Float value (non-pointer)
		{name: "big.Float value 0 → all-zeros IPv6", in: *new(big.Float).SetFloat64(0), expect: ipv6()},
		{name: "big.Float value +Inf → error", in: func() big.Float {
			var f big.Float
			f.SetInf(false)
			return f
		}(), expectErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := cast.ToE[net.IP](tc.in)
			if err != nil && !tc.expectErr {
				t.Errorf("unexpected error: %v", err)
			} else if err == nil && tc.expectErr {
				t.Error("expected error, got nil")
			} else if err != nil && !errors.Is(err, cast.Error) {
				t.Errorf("expected cast.Error, got %T: %v", err, err)
			} else if err == nil && !result.Equal(tc.expect) {
				t.Errorf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}
