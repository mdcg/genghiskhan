package scanner

import (
	"sync"
	"time"
)

// TCPScanner starts the process of scanning TCP type ports. Basically, it will
// iterate from 1 to N (N is the number of ports that have been requested to be
// scanned), and for each of the iterations it will trigger a goroutine to
// check if the port in question is open.
func TCPScanner(addr string, portsNumber int) []ScanReport {
	defer timeTrack(time.Now(), portsNumber)
	s := startSpinner()
	var wg sync.WaitGroup
	scannerChan := make(chan ScanReport, portsNumber)
	for port := 0; port < portsNumber; port++ {
		wg.Add(1)
		go checkTCPPort(addr, port, &wg, scannerChan)
	}
	wg.Wait()
	close(scannerChan)
	stopSpinner(s)
	return generateScanReport(scannerChan)
}

// UDPScanner starts the process of scanning UDP-type ports. Basically, it will
// iterate from 1 to N (N is the number of ports that have been requested to be
// scanned), and for each of the iterations it will trigger a goroutine to
// check if the port in question is open.
func UDPScanner(addr string, portsNumber int) []ScanReport {
	defer timeTrack(time.Now(), portsNumber)
	s := startSpinner()
	var wg sync.WaitGroup
	scannerChan := make(chan ScanReport, portsNumber)
	for port := 0; port < portsNumber; port++ {
		wg.Add(1)
		go checkUDPPort(addr, port, &wg, scannerChan)
	}
	wg.Wait()
	close(scannerChan)
	stopSpinner(s)
	return generateScanReport(scannerChan)
}

// FullScanner starts the process of scanning both TCP and UDP ports.
// Basically, it will iterate from 1 to N, (N is the number of ports that have
// been requested to be scanned) and for each of the iterations, it will
// trigger two goroutines, one to verify the TCP port and one for UDP.
func FullScanner(addr string, portsNumber int) []ScanReport {
	defer timeTrack(time.Now(), portsNumber)
	s := startSpinner()
	var wg sync.WaitGroup
	scannerChan := make(chan ScanReport, portsNumber)
	for port := 0; port < portsNumber; port++ {
		wg.Add(2)
		go checkUDPPort(addr, port, &wg, scannerChan)
		go checkTCPPort(addr, port, &wg, scannerChan)
	}
	wg.Wait()
	close(scannerChan)
	stopSpinner(s)
	return generateScanReport(scannerChan)
}
