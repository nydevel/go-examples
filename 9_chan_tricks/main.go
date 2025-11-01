package main

import "fmt"

func main() {
	a := make(chan int, 1)
	for i := 0; i < 5; i++ {
		select {
		case <-a:
			fmt.Println(i)
		case a <- i:
			fmt.Println(i)
		}
	}
}
