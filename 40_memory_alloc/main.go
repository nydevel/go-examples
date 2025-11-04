package main

// Аллокация 1гб памяти

func main() {
	data := make([]byte, 1<<30)

	_ = data
}
