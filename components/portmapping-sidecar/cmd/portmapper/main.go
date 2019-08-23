package main

import (
	"atakanyenel/portmapper-sidecar"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	portmapper.Start()
	wg.Wait()
}
