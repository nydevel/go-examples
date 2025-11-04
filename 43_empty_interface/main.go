package main

// Пустой интерфейс - любой тип, тоесть any (c go 1.18)
func main() {
	var _ interface{}

	var _ any

	// Any - алиас для пустого интерфейса
	type any = interface{}
}
