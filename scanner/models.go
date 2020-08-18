package scanner

import (
	"fmt"
)

// ScanReport is the model that is used to facilitate the scan report
// generation process.
type ScanReport struct {
	ServiceName string
	Port        int
	Protocol    string
}

// PrintReport is a method of ScanReport to facilitate the display of data
// for the report.
func (sr *ScanReport) PrintReport() {
	fmt.Printf("[%v] %v => %v\n",
		sr.Protocol,
		sr.Port,
		sr.ServiceName,
	)
}
