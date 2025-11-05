package main

import (
	"fmt"
	"reflect"
)

// Рефлексия в GO

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Vasya",
		Age:  18,
	}

	ref := reflect.TypeOf(user)

	fmt.Println(ref.Name())
}
