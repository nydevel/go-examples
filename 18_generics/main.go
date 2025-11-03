package main

import "fmt"

// Generics

func Print[T any](value T) {
	fmt.Println(value)
}

func main() {
	Print[int](1)
	Print(1) // type inference - сам определяет тип
}
