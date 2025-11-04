package main

import "fmt"

// struct{} занимает 0 байт - это лучше чем bool для проверочных значений

func main() {
	set := map[string]struct{}{
		"one": struct{}{},
		"two": struct{}{},
	}

	if _, exist := set["one"]; exist {
		fmt.Println("one exists in map")
	}
}
