package calculator

import (
	"fmt"
	"net"
	"strings"
)

const IPv6Bits = 128
const IPv6Len = 8

func ParseIPv6(ipAddr string) (*IPv6, error) {
	if ip, ipNet, err := net.ParseCIDR(ipAddr); err != nil {
		return nil, err
	} else {
		prefixLen, _ := ipNet.Mask.Size()
		start, end := getAddressRange(ipNet, IPv6Bits)
		count := getAddressCount(ipNet)
		return &IPv6{
			Short: SubIPv6{
				IP:       ip.String(),
				CIDR:     ipNet.String(),
				Start:    start.String(),
				End:      end.String(),
				MaskBits: prefixLen,
				Count:    count,
			},
			Long: SubIPv6{
				IP:       getLongIPv6Str(ip.String()),
				CIDR:     getLongCIDR(ipNet.String()),
				Start:    getLongIPv6Str(start.String()),
				End:      getLongIPv6Str(end.String()),
				MaskBits: prefixLen,
				Count:    count,
			},
		}, nil
	}
}

func getLongCIDR(cidr string) string {
	cidrSlice := strings.Split(cidr, "/")
	cidrSlice[0] = getLongIPv6Str(cidrSlice[0])
	return strings.Join(cidrSlice, "/")
}

func getLongIPv6Str(ipStr string) string {
	ipSlice := strings.Split(ipStr, ":")
	for i := 0; i < len(ipSlice); i++ {
		if ipSlice[i] == "" {
			ipSlice = append(ipSlice[:i], append(make([]string, IPv6Len-len(ipSlice)+1), ipSlice[i+1:]...)...)
			break
		}
	}
	for i, str := range ipSlice {
		ipSlice[i] = fmt.Sprintf("%04s", str)
	}
	return strings.Join(ipSlice, ":")
}
