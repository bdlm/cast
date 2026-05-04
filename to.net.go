package cast

import (
	"fmt"
	"math/big"
	"net"
)

// toNetIP converts v to net.IP.
//
// String values are parsed with net.ParseIP and accept both IPv4
// ("192.168.1.1") and IPv6 ("::1") representations.
// []byte values of length net.IPv4len (4) or net.IPv6len (16) are copied
// directly into a net.IP; other lengths are attempted as string parsing.
// uint32 values are treated as a packed IPv4 address in host byte order.
//
// Options:
//   - DEFAULT: net.IP, value to return on error.
func toNetIP(v any, ops ops) (any, error) {
	var ret any = net.IP(nil)

	if ops.hasDefault {
		ip, ok := ops.defaultVal.(net.IP)
		if !ok {
			return ret, fmt.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		ret = ip
	}

	switch val := v.(type) {
	case nil:
		return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
	case net.IP:
		return val, nil
	case string:
		ip := net.ParseIP(val)
		if ip == nil {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
		}
		return ip, nil
	case []byte:
		if len(val) == net.IPv4len || len(val) == net.IPv6len {
			ip := make(net.IP, len(val))
			copy(ip, val)
			return ip, nil
		}
		ip := net.ParseIP(string(val))
		if ip == nil {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
		}
		return ip, nil
	case int32:
		if val < 0 {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
		}
		u32 := uint32(val)
		return net.IPv4(byte(u32>>24), byte(u32>>16), byte(u32>>8), byte(u32)), nil
	case uint32:
		return net.IPv4(byte(val>>24), byte(val>>16), byte(val>>8), byte(val)), nil
	case *big.Int:
		if val == nil {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
		}
		return bigIntToIPv6(val, v, ret)
	case big.Int:
		return bigIntToIPv6(&val, v, ret)
	case *big.Float:
		if val == nil || val.IsInf() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
		}
		i, _ := val.Int(nil)
		return toNetIP(i, ops)
	case big.Float:
		if val.IsInf() {
			return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
		}
		i, _ := val.Int(nil)
		return toNetIP(i, ops)
	default:
		if s, err := toString(v, ops.Delete(DEFAULT)); err == nil {
			if ip := net.ParseIP(s); ip != nil {
				return ip, nil
			}
		}
	}

	return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
}

// bigIntToIPv6 converts a *big.Int to a 16-byte IPv6 net.IP. The value must
// be non-negative and representable in 16 bytes (i.e. fit within 128 bits).
func bigIntToIPv6(val *big.Int, v any, ret any) (any, error) {
	if val.Sign() < 0 {
		return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
	}
	b := val.Bytes() // big-endian, no leading zeros
	if len(b) > net.IPv6len {
		return ret, fmt.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
	}
	ip := make(net.IP, net.IPv6len)
	copy(ip[net.IPv6len-len(b):], b)
	return ip, nil
}
