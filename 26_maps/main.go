package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5
	m["f"] = 6
	m["g"] = 7
	m["h"] = 8
	m["i"] = 9

	var m2 map[string]int
	m2["a"] = 1
	fmt.Println(m2)
}
