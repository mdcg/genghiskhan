package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func checkTCPPort(addr string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	if isTCPPortOpened(addr, port) {
		serviceName := predictTCPPortService(port)
		scan := ScanReport{
			ServiceName: serviceName,
			Port:        port,
			Protocol:    "TCP",
		}
		fmt.Println(scan)
	}
}

func isTCPPortOpened(addr string, port int) bool {
	host := fmt.Sprintf("%v:%v", addr, port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		return false
	}
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func predictTCPPortService(port int) string {
	if service, ok := TCP_PORTS[port]; ok {
		return service
	}
	return UNKNOWN
}
