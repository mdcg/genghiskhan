package scanner

import (
	"fmt"
	"net"
)

func isTCPPortOpened(addr string, port int) bool {
	host := fmt.Sprintf("%s:%d", addr, port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		return false
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func isUDPPortOpened(addr string, port int) bool {
	host := fmt.Sprintf("%s:%d", addr, port)
	udpAddr, err := net.ResolveUDPAddr("udp4", host)
	if err != nil {
		return false
	}
	conn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func getTCPPort(port int) string {
	if service, ok := TCP_PORTS[port]; ok {
		return service
	}
	return UNKNOWN
}

func getUDPPort(port int) string {
	if service, ok := UDP_PORTS[port]; ok {
		return service
	}
	return UNKNOWN
}
