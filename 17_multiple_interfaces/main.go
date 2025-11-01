package main

type Interface1 interface {
	Method1() string
}

type Interface2 interface {
	Method2() string
}

type Struct struct{}

func (s Struct) Method1() string {
	return "method1"
}

func (s Struct) Method2() string {
	return "method2"
}

func returnInterface1() Interface1 {
	return Struct{}
}

func returnInterface2() Interface2 {
	return Struct{}
}

func main() {
	var _ Interface1 = returnInterface1()
	var _ Interface2 = returnInterface2()
}
