package main

import "fmt"

func Print[T any](value T) {
	fmt.Println(value)
}

// Constraints - ограничение на типы
func Sum[T int | int64 | float64](a, b T) T {
	return a + b
}

func main() {
	Print[int](1)
	Print(1) // type inference - сам определяет тип
}
