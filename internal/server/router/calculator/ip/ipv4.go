package ip

import (
	"math/big"
	"net"
	"neutron-star-api/internal/server/router/calculator/model"
)

const IPv4Bits = 32

func ParseIPv4(ipAddr string) (*model.IPv4, error) {
	if ip, ipNet, err := net.ParseCIDR(ipAddr); err != nil {
		return nil, err
	} else {
		prefixLen, _ := ipNet.Mask.Size()
		start, end := getAddressRange(ipNet, IPv4Bits)
		return &model.IPv4{
			IP:           ip.String(),
			CIDR:         ipNet.String(),
			Start:        start.String(),
			End:          end.String(),
			MaskBits:     prefixLen,
			SubnetMask:   getSubnetMask(ipNet).String(),
			Count:        getAddressCount(ipNet),
			WildcardMask: getWildcardMask(ipNet).String(),
		}, nil
	}
}

func getSubnetMask(ipNet *net.IPNet) net.IP {
	prefixLen, _ := ipNet.Mask.Size()
	var n int64 = 1<<prefixLen - 1
	n = n << (IPv4Bits - prefixLen)
	return intToIp(big.NewInt(n), IPv4Bits)
}

func getWildcardMask(ipNet *net.IPNet) net.IP {
	prefixLen, _ := ipNet.Mask.Size()
	var n int64 = 1<<(IPv4Bits-prefixLen) - 1
	return intToIp(big.NewInt(n), IPv4Bits)
}
