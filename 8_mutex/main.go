package main

import (
	"fmt"
	"sync"
)

func main() {
	value := 0

	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer func() {
				wg.Done()
				mutex.Unlock()
			}()

			mutex.Lock()
			value++
		}()
	}

	wg.Wait()

	fmt.Println(value)
}
