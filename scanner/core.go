package scanner

import (
	"sync"
)

func TCPScanner(addr string, portsNumber int) []ScanReport {
	var wg sync.WaitGroup
	scannerChan := make(chan ScanReport, portsNumber)
	for port := 0; port < portsNumber; port++ {
		wg.Add(1)
		go checkTCPPort(addr, port, &wg, scannerChan)
	}
	wg.Wait()
	close(scannerChan)
	return generateScanReport(scannerChan)
}
