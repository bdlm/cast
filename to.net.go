package cast

import (
	"net"

	"github.com/bdlm/errors/v2"
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
			return ret, errors.Errorf(ErrorInvalidOption, "DEFAULT", ops.defaultVal)
		}
		ret = ip
	}

	switch val := v.(type) {
	case nil:
		return ret, errors.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
	case net.IP:
		return val, nil
	case string:
		ip := net.ParseIP(val)
		if ip == nil {
			return ret, errors.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
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
			return ret, errors.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
		}
		return ip, nil
	case uint32:
		return net.IPv4(byte(val>>24), byte(val>>16), byte(val>>8), byte(val)), nil
	default:
		if s, err := toString(v, ops.Delete(DEFAULT)); err == nil {
			if ip := net.ParseIP(s.(string)); ip != nil {
				return ip, nil
			}
		}
	}

	return ret, errors.Errorf(ErrorStrUnableToCast, v, v, net.IP(nil))
}
