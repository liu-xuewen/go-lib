package iplib

import (
	"net"
)

// LocalIP local ip
var LocalIP string

func init() {
	LocalIP = GetLocalIP()
}

// GetLocalIP get local ip
func GetLocalIP() (ip string) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, address := range addresses {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4().String()
			}

		}
	}
	return
}
