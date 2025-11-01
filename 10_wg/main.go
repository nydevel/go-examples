package main

import (
	"fmt"
	"sync"
)

func main_bad() {
	ch := make(chan int)

	go func() {
		v := <-ch // Может еще НЕ начать ждать
		fmt.Println(v)
	}()

	// Основная горутина может дойти сюда ДО того,
	// как горутина выше начнет ждать на канале!
	select {
	case ch <- 42:
		fmt.Println("sent")
	default:
		fmt.Println("not ready") // Может выполниться это!
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var ch = make(chan int)

	go func() {
		defer wg.Done()
		v := <-ch
		fmt.Println(v)
	}()

	ch <- 42
	wg.Wait()
}
