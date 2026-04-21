package cast_test

import (
	"errors"
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
