package main

import (
	"fmt"
	"strconv"
)

// Map мутабельны

func main() {
	m := make(map[string]int)
	m["1"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5
	m["f"] = 6
	m["g"] = 7
	m["h"] = 8
	m["i"] = 9

	for i := range 10000 {
		m[strconv.Itoa(i)] = i // 1 = "1" ...
	}

	fmt.Println(m["1"])
}
