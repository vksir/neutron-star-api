package ip

import (
	"math/big"
	"net"
)

func getAddressCount(ipNet *net.IPNet) int64 {
	prefixLen, bits := ipNet.Mask.Size()
	return 1 << (bits - prefixLen)
}

func getAddressRange(ipNet *net.IPNet, bits int) (net.IP, net.IP) {
	start := ipNet.IP
	count := getAddressCount(ipNet)
	startIpInt := ipToInt(start)
	endIpInt := startIpInt.Add(startIpInt, big.NewInt(count-1))
	end := intToIp(endIpInt, bits)
	return start, end
}

func ipToInt(ip net.IP) *big.Int {
	ipInt := &big.Int{}
	ipInt.SetBytes(ip)
	return ipInt
}

func intToIp(ipInt *big.Int, bits int) net.IP {
	ipBytes := ipInt.Bytes()
	ip := make([]byte, bits/8)
	for i := 1; i <= len(ipBytes); i++ {
		ip[len(ip)-i] = ipBytes[len(ipBytes)-i]
	}
	return ip
}
