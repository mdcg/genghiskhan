package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func checkTCPPort(addr string, port int, wg *sync.WaitGroup, scannerChan chan<- ScanReport) {
	defer wg.Done()
	if isTCPPortOpened(addr, port) {
		serviceName := predictTCPPortService(port)
		scan := ScanReport{
			ServiceName: serviceName,
			Port:        port,
			Protocol:    "TCP",
		}
		scannerChan <- scan
	}
}

func checkUDPPort(addr string, port int, wg *sync.WaitGroup, scannerChan chan<- ScanReport) {
	defer wg.Done()
	if isUDPPortOpened(addr, port) {
		serviceName := predictUDPPortService(port)
		scan := ScanReport{
			ServiceName: serviceName,
			Port:        port,
			Protocol:    "UDP",
		}
		scannerChan <- scan
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

func isUDPPortOpened(addr string, port int) bool {
	host := fmt.Sprintf("%v:%v", addr, port)
	udpAddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		return false
	}
	conn, err := net.DialTimeout("udp", udpAddr.String(), 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	error_count := 0
	for i := 0; i < 3; i++ {
		buf := []byte("0")
		_, err := conn.Write(buf)
		if err != nil {
			error_count++
		}
	}

	if error_count > 0 {
		return false
	}
	return true
}

func predictTCPPortService(port int) string {
	if service, ok := TCP_PORTS[port]; ok {
		return service
	}
	return UNKNOWN
}

func predictUDPPortService(port int) string {
	if service, ok := UDP_PORTS[port]; ok {
		return service
	}
	return UNKNOWN
}

func timeTrack(start time.Time, portsNumber int) {
	elapsed := time.Since(start)
	fmt.Printf("Scanning took %v to check %v port(s)\n", elapsed, portsNumber)
}
