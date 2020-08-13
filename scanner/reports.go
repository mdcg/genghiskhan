package scanner

import "fmt"

func generateScanReport(scannerChan <-chan ScanReport) []ScanReport {
	scans := make([]ScanReport, 0)
	for scan := range scannerChan {
		scans = append(scans, scan)
	}
	return scans
}

func FormatScanReport(scans []ScanReport) {
	if len(scans) > 0 {
		fmt.Println("The following ports are open:")
		for i := 0; i < len(scans); i++ {
			scans[i].PrintReport()
		}
	} else {
		fmt.Println("No port is open.")
	}
}
