package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// function that is executing in a known context (kc -> any type)
// recieves a copy of the context ie. the Person struct
func (g Person) details() {
	fmt.Println(g.name, g.age)
}

func main() {
	fmt.Println("Methods")
	g := Person{
		age:  25,
		name: "James",
	}
	g.greet()
}

// value reciever
func (g Person) greetVal() {
	g.name = "does not affect original"
	fmt.Println(g.age, g.name)
}

// pointer
func (g *Person) greetPtr() {
	g.name = "does affect original"
	fmt.Println(g.age, g.name)
}
