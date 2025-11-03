package main

import "fmt"

// recover — это встроенная функция, которая останавливает панику (panic) и возвращает значение, переданное в panic().

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("перехватили панику:", r)
		}
	}()
	panic("ups")
}
