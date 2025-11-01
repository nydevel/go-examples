package main

// методы встроенной структуры становятся методами внешней структуры автоматически, но это композиция, а не наследование.
type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "Какой-то звук"
}

type Dog struct {
	Animal // встраивание
	Breed  string
}

func main() {
	dog := Dog{
		Animal: Animal{Name: "Шарик"},
		Breed:  "Овчарка",
	}

	// Оба вызова работают:
	dog.Speak()        // ✅ метод "поднялся" к Dog
	dog.Animal.Speak() // ✅ явный вызов через Animal
}
