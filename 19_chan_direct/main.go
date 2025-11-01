package main

// Однонаправленные каналы

func receiver(ch <-chan int) int {
	v := <-ch

	return v
}

func sender(ch chan<- int) {
	ch <- 1
}

func main() {
	ch := make(chan int)
	go sender(ch)
	receiver(ch)
}
