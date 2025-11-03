package main

// Нельзя изменить map если она проинициализирована как nil

func main() {
	var m2 map[string]int
	m2["a"] = 1
}
