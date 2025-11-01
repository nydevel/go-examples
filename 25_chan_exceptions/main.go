package main

import (
	"fmt"
	"time"
)

func chan1() {
	ch := make(chan int, 3) // буфер на 3 элемента

	ch <- 1 // НЕ блокируется
	ch <- 2 // НЕ блокируется
	ch <- 3 // НЕ блокируется
	ch <- 4 // БЛОКИРУЕТСЯ — буфер полон deadlock
}

// deadlock
func chan2() {
	var ch chan int // nil канал

	// ch <- 1 // БЛОКИРУЕТСЯ НАВСЕГДА
	val := <-ch

	fmt.Println(val)
}

// deadlock
func chan2_1() {
	ch := make(chan int, 1)

	val := <-ch
	fmt.Println(val)
}

// НЕ блокируется — default выполняется мгновенно
func chan3() {
	ch := make(chan int)

	select {
	case v := <-ch:
		fmt.Println("Received:", v)
	default:
		fmt.Println("No data available") // выполнится сразу
	}
}

func chan4() {
	ch := make(chan int) // небуферизированный канал

	fmt.Println("горутина-1 собирается писать…")
	go func() {
		ch <- 42                   // БЛОКИРОВКА: никто не читает
		fmt.Println("записали 42") // эта строка никогда не выполнится
	}()

	time.Sleep(100 * time.Millisecond) // даём горутине шанс запуститься
	fmt.Println("main завершается, но горутина-1 всё ещё ждёт читателя")
}

func main() {
	chan4()
}
