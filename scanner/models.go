package scanner

import (
	"fmt"
)

type ScanReport struct {
	ServiceName string
	Port        int
	Protocol    string
}

func (sr *ScanReport) PrintReport() {
	fmt.Printf("[%v] %v => %v\n",
		sr.Protocol,
		sr.Port,
		sr.ServiceName,
	)
}
