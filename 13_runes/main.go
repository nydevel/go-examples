package main

import "fmt"

func main() {
	var r = []rune("Hello")

	r[0] = '1'          // we can change rune
	fmt.Printf("%c", r) //how to print rune character
}
