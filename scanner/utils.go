package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/briandowns/spinner"
)

// checkTCPPort basically checks whether the TCP port in question is open or
// not. As we are working concurrently, we need to "synchronize" some
// information, so if any TCP port is found, we "instantiate" a ScanReport and
// send it to the scannerChan channel, where it will later be displayed in the
// final report.
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

// checkUDPPort basically checks whether the UDP port in question is open or
// not. As we are working concurrently, we need to "synchronize" some
// information, so if any UDP port is found, we "instantiate" a ScanReport and
// send it to the scannerChan channel, where it will later be displayed in the
// final report.
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

// isTCPPortOpened tries to establish a TCP connection to the specified port.
// If the connection is successful, we return true, otherwise, false.
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

// isUDPPortOpened sends three "messages" to the specified UDP port, if there
// is no error after the three sends, we can infer that the port is open.
// Unlike TCP, UDP does not actually establish a connection, which is why we
// have this "three retries" logic.
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

// predictTCPPortService attempts to infer the service that is running based on
// the reported port number.
func predictTCPPortService(port int) string {
	if service, ok := TCP_PORTS[port]; ok {
		return service
	}
	return UNKNOWN
}

// predictUDPPortService attempts to infer the service that is running based on
// the reported port number.
func predictUDPPortService(port int) string {
	if service, ok := UDP_PORTS[port]; ok {
		return service
	}
	return UNKNOWN
}

// timeTrack is a utility for checking the time spent on the scanning process.
func timeTrack(start time.Time, portsNumber int) {
	elapsed := time.Since(start)
	fmt.Printf("Scanning took %v to check %v port(s)\n", elapsed, portsNumber)
}

// startSpinner basically generates a "spinner" to make viewing the scanning
// process more "pleasant" for the user.
func startSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	s.Suffix = " Starting scan..."
	s.Start()
	return s
}

// stopSpinner for the execution of a "spinner" informed as a parameter.
func stopSpinner(s *spinner.Spinner) {
	s.Stop()
}
