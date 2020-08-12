package main

import (
	"github.com/mdcg/go-port-scanner/scanner"
)

func main() {
	scanner.TCPScanner("localhost", 10000)
}
