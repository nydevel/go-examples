package main

func main() {
	mp := make(map[string]int)
	mp["a"] = 1
	mp["b"] = 2
	mp["c"] = 3

	// v := &mp["a"] // так нельзя потому что может перестроиться хештаблица
	// fmt.Println(v)
}
