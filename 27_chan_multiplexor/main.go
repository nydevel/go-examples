package main

import (
	"fmt"
	"sync"
)

// Мультиплексер каналов - n каналов будут направлены в один, выходной канал будет отдан сразу

func multiplexer(chans ...chan int) chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))

	for _, ch := range chans {
		go func(c chan int) {
			defer wg.Done()

			for val := range ch {
				out <- val
			}
		}(ch)
	}

	// Если не обернуть в горутину, то функция будет заблокирована и пока не отработает WaitGroup, канал не будет отдан
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() { a <- 1; a <- 2; close(a) }()
	go func() { b <- 3; close(b) }()

	for v := range multiplexer(a, b) {
		fmt.Println(v) // вывод: 1 2 3 (порядок не гарантирован)
	}
}
