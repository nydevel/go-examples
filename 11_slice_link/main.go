package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
}
