package iputil

import (
	"net"
)

func GetLocalIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return net.IPv4zero
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip := ipnet.IP.To4(); ip != nil {
				return ipnet.IP
			}
		}
	}
	return net.IPv4zero
}

func GetLocalPublicIP() net.IP {
	return GetLocalWanIP()
}

func GetLocalWanIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return net.IPv4zero
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip := ipnet.IP.To4(); ip != nil && publicIP(ip) {
				return ipnet.IP
			}
		}
	}
	return net.IPv4zero
}

// pre 表示局域ip的第一个字节，10 或者172 或者192
func GetLocalLanIP(pre byte) net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return net.IPv4zero
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip := ipnet.IP.To4(); ip != nil && ip[0] == pre {
				return ipnet.IP
			}
		}
	}
	return net.IPv4zero
}

// 判断是否是公网IP
func publicIP(ip net.IP) bool {
	return ip[0] != 10 && ip[0] != 172 && ip[0] != 192
}
