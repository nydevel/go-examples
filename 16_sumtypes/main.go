package main

import "fmt"

// Result - sum type для результата операции
type Result interface {
	isResult() // приватный метод-маркер
}

// Success - вариант успеха
type Success struct {
	Value string
}

func (s Success) isResult() {}

// Error - вариант ошибки
type Error struct {
	Message string
}

func (e Error) isResult() {}

// Функция, возвращающая Result
func divide(a, b int) Result {
	if b == 0 {
		return Error{Message: "деление на ноль"}
	}
	return Success{Value: fmt.Sprintf("результат: %d", a/b)}
}

func main() {
	result := divide(10, 2)

	// Pattern matching через type switch
	switch r := result.(type) {
	case Success:
		fmt.Println("Успех:", r.Value)
	case Error:
		fmt.Println("Ошибка:", r.Message)
	}
}
