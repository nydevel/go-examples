package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Асинхронный счетчик с блокировкой через atomic - быстрее чем через Mutex, но поддерживает только простые типы

func main() {
	var counter int32
	wg := sync.WaitGroup{} // remove race conditions
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
