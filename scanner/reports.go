package scanner

import (
	"fmt"
	"sort"
)

// generateScanReport basically it will be receiving the data from the channel
// scannerChan, and the moment this channel is closed, a slice will be returned
// containing all the information received.
func generateScanReport(scannerChan <-chan ScanReport) []ScanReport {
	scans := make([]ScanReport, 0)
	for scan := range scannerChan {
		scans = append(scans, scan)
	}
	return scans
}

// FormatScanReport facilitates the display of scan data. It is basically here
// that we make our infamous "report".
func FormatScanReport(scans []ScanReport) {
	if len(scans) > 0 {
		scans = sortScanResults(scans)
		for i := 0; i < len(scans); i++ {
			scans[i].PrintReport()
		}
	} else {
		fmt.Println("No port is open.")
	}
}

// sortScanResults sorts the scan results based on the protocol. As some
// goroutines can finish before others, regardless of the "firing order",
// especially when we are working with "fullscan", it is interesting for
// visibility purposes that they are ordered, separating the report of TCP port
// scans from UDP ports.
func sortScanResults(scans []ScanReport) []ScanReport {
	sort.Slice(scans, func(i, j int) bool {
		return scans[i].Protocol < scans[j].Protocol
	})
	return scans
}
