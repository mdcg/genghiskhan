package scanner

import (
	"sync"
	"time"
)

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
