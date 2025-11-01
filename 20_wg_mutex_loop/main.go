package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	wg := sync.WaitGroup{} // remove race conditions
	wg.Add(100)
	mutex := sync.Mutex{} // remove data race

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			mutex.Lock()
			counter++
			mutex.Unlock()

		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
