package main

import "fmt"

type Someone struct {
	Name string
}

func changeName(person *Someone) {
	// Создается новая переменная которой присваевается копия указателя на старый объект,
	// на этом все - просто новый объект в скоупе данной функции
	person = &Someone{
		Name: "Dora",
	}
}

// Работающий метод - указатель на указатель
func changeNameWorking(person **Someone) {
	*person = &Someone{
		Name: "Dora",
	}
}

func main() {
	person := &Someone{
		Name: "Nick",
	}

	fmt.Println(person.Name) // покажет Nick
	changeNameWorking(&person)
	fmt.Println(person.Name) // покажет Nick
}
