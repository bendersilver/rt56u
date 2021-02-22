package handler

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func toString(b []byte) string {
	return strings.TrimSpace(string(b))
}

// ParseHeader -
func ParseHeader(r io.Reader, bufLen int) (method, path, host string, err error) {
	var buf = make([]byte, bufLen)
	var i, line int
	var sep []int
	for {
		if bufLen <= i {
			err = fmt.Errorf("malformed HTTP")
			return
		}
		if _, err = r.Read(buf[i : i+1]); err != nil {
			return
		}
		// end header
		if i == 1 && buf[1] == '\n' {
			break
		}
		if buf[i] == ' ' {
			sep = append(sep, i)
		}
		if buf[i] == '\n' {
			line++
			switch line {
			case 1:
				if len(sep) == 2 {
					method = toString(buf[:sep[0]])
					path = toString(buf[sep[0]:sep[1]])
				} else {
					err = fmt.Errorf("malformed HTTP")
					return
				}
			case 2:
				if len(sep) == 1 {
					host = toString(buf[sep[0]:i])
				} else {
					err = fmt.Errorf("malformed MIME header")
					return
				}
			}
			i = 0
			sep = nil
		} else {
			i++
		}
	}
	return
}

// defaultFilteredNetworks net.IPNets that are loopback, private, link local, default unicast
// based on https://github.com/letsencrypt/boulder/blob/master/bdns/dns.go
// https://github.com/wader/filtertransport/blob/master/filter.go
var privateNet = []net.IPNet{
	mustParseCIDR("10.0.0.0/8"),     // RFC1918
	mustParseCIDR("172.16.0.0/12"),  // private
	mustParseCIDR("192.168.0.0/16"), // private
	mustParseCIDR("127.0.0.0/8"),    // RFC5735
	mustParseCIDR("169.254.0.0/16"), // RFC3927
	mustParseCIDR("192.0.0.0/24"),   // RFC 5736
	mustParseCIDR("192.0.2.0/24"),   // RFC 5737
	mustParseCIDR("192.88.99.0/24"), // RFC 3068
	mustParseCIDR("192.18.0.0/15"),  // RFC 2544
	mustParseCIDR("::/128"),         // RFC 4291: Unspecified Address
	mustParseCIDR("::1/128"),        // RFC 4291: Loopback Address

}

// mustParseCIDR parses string into net.IPNet
func mustParseCIDR(s string) net.IPNet {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		panic(err)
	}
	return *ipnet
}

// IsPrivate true if any of the ipnets contains ip
func IsPrivate(addr string) bool {
	h, _, err := net.SplitHostPort(addr)
	if err != nil {
		return false
	}
	ip := net.ParseIP(h)
	for _, ipnet := range privateNet {
		if ipnet.Contains(ip) {
			return true
		}
	}
	return false
}
