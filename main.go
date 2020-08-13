package main

import (
	"github.com/mdcg/go-port-scanner/scanner"
)

func main() {
	// scanner.FormatScanReport(scanner.TCPScanner("localhost", 10000))
	scanner.FormatScanReport(scanner.TCPScanner("localhost", 10000))
}
