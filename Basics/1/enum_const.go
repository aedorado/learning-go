package main

import (
	"fmt"
)

// 1.
const a = iota

// 2.
const (
	b = iota
	c = iota
	d = iota
)

// 3.
const (
	e = iota
	f
	g
)

// 5. Using
const (
	car = iota
	bike
	truck
)

// 7. Dealing with error
// const (
// 	errorV = iota
// 	car
// 	bike
// 	truck
// )

// const (
// 	_ = iota
// 	car
// 	bike
// 	truck
// )

// 8. Arithmetic Operations on iota
// const (
// 	car = iota + 5
// 	bike
// 	truck
// )

func main() {

	// 1. iota is a counter we can use to create enumerated constants
	fmt.Printf("Value = %v, Type = %T\n\n", a, a)

	// 2. Using iota in a constant decleration block
	fmt.Printf("Value = %v, Type = %T\n", b, b)
	fmt.Printf("Value = %v, Type = %T\n", c, c)
	fmt.Printf("Value = %v, Type = %T\n", d, d)
	fmt.Printf("Value = %v, Type = %T\n\n", e, e)

	// 3. Dont assign value after the first one

	// 4. Values of iota restricted to a constant block and reset for next
	// this allows us to create related constants together, with different values
	// and another set of constants having unique values
	// but allow duplication of values across constant blocks
	fmt.Printf("Value = %v, Type = %T\n", e, e)
	fmt.Printf("Value = %v, Type = %T\n", f, f)
	fmt.Printf("Value = %v, Type = %T\n\n", g, g)

	// 5, 6 Usage and default values
	var vehicle int = bike
	fmt.Printf("%v\n", vehicle == bike)
	fmt.Printf("%v\n", vehicle == car)
	vehicle = bike
	fmt.Printf("%v\n", vehicle == car)

}
