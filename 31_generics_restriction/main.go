package main

import "fmt"

// Constraints - ограничение на типы

func Sum[T int | int64 | float64](a, b T) T {
	return a + b
}

func main() {
	s := Sum[int](1, 2)
	fmt.Println(s)
}
