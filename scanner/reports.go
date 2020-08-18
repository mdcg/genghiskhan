package scanner

import (
	"fmt"
	"sort"
)

func generateScanReport(scannerChan <-chan ScanReport) []ScanReport {
	scans := make([]ScanReport, 0)
	for scan := range scannerChan {
		scans = append(scans, scan)
	}
	return scans
}

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

func sortScanResults(scans []ScanReport) []ScanReport {
	sort.Slice(scans, func(i, j int) bool {
		return scans[i].Protocol < scans[j].Protocol
	})
	return scans
}
