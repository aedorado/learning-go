package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name     string
}

// function that is executing in a known context (kc -> any type)
// recieves a copy of the context ie. the greeter struct
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	fmt.Println("Methods")
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

// value reciever
func (g greeter) greetVal() {
	g.name = "does not affect original"
	fmt.Println(g.greeting, g.name)
}

// pointer
func (g *greeter) greetPtr() {
	g.name = "does not affect original"
	fmt.Println(g.greeting, g.name)
}
