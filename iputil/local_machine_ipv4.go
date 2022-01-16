package iputil

import (
	"net"

	"github.com/dajinkuang/errors"
)

// LocalMachineIPV4 获取本机IP
func LocalMachineIPV4() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		err = errors.Wrap(err)
		return
	}
	for _, v := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := v.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				return
			}
		}
	}
	return
}
