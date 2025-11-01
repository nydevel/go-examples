package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	ch := make(chan int, 100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			ch <- 1
		}()
	}

	wg.Wait()
	close(ch)

	counter := 0
	for range ch {
		counter++
	}

	fmt.Println(counter)

}
