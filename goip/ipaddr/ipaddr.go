package ipaddr

import (
	"net"
)

// GetIpv4 get ipv4addrs for all interfaces
func GetIpv4() []string {
	var ipv4addrs []string
	//get all interfaces ip info
	interfaceaddrs, _ := net.InterfaceAddrs()
	//filter ip address
	for _, ipaddrs := range interfaceaddrs {
		ipaddrs := ipaddrs.(*net.IPNet)
		if !ipaddrs.IP.IsLoopback() && ipaddrs.IP.To4() != nil {
			ipv4addrs = append(ipv4addrs, ipaddrs.IP.String())

		}

	}
	return ipv4addrs

}
