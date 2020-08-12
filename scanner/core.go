package scanner

import "sync"

func TCPScanner(addr string, portsNumber int) {
	var wg sync.WaitGroup
	for port := 0; port < portsNumber; port++ {
		wg.Add(1)
		go checkTCPPort(addr, port, &wg)
	}
	wg.Wait()
}
