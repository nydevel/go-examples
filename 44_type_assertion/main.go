package main

import "fmt"

// Извлечение значения с проверкой типа

func main() {
	var i any = "hello"

	s, ok := i.(string)

	if ok {
		fmt.Println("Это строка:", s)
	}
}
