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
	fmt.Printf("%v [%v] => %v\n",
		sr.Port,
		sr.Protocol,
		sr.ServiceName,
	)
}
