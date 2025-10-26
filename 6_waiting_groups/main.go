package main

import (
	"sync"
)

func killCounter(wg *sync.WaitGroup) {
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	killCounter(&wg) // copy link to wg, not wg himself
	wg.Wait()
}
