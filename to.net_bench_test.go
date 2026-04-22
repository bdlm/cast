package cast_test

import (
	"net"
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_netIP_fromStringIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[net.IP]("192.168.1.1")
	}
}

func BenchmarkTo_netIP_fromStringIPv6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cast.To[net.IP]("2001:db8::1")
	}
}

func BenchmarkTo_netIP_fromBytes(b *testing.B) {
	// []byte of IPv4len (4) copies directly without net.ParseIP
	src := []byte{192, 168, 1, 1}
	for i := 0; i < b.N; i++ {
		_ = cast.To[net.IP](src)
	}
}

func BenchmarkTo_netIP_fromUint32(b *testing.B) {
	// uint32 → packed IPv4 via net.IPv4
	for i := 0; i < b.N; i++ {
		_ = cast.To[net.IP](uint32(0xC0A80101)) // 192.168.1.1
	}
}

func BenchmarkTo_netIP_passthrough(b *testing.B) {
	src := net.ParseIP("192.168.1.1")
	for i := 0; i < b.N; i++ {
		_ = cast.To[net.IP](src)
	}
}

func BenchmarkToE_netIP_fromInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = cast.ToE[net.IP]("not-an-ip")
	}
}
