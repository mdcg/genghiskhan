package main

import "github.com/mdcg/go-port-scanner/cmd"

func main() {
	// scanner.FormatScanReport(scanner.TCPScanner("localhost", 10000))
	// scanner.FormatScanReport(scanner.UDPScanner("localhost", 10000))
	// scanner.FormatScanReport(scanner.FullScanner("localhost", 10000))
	cmd.Execute()
}
