package main

import (
	"fmt"
	"reflect"
)

// VARIABLES

func main() {
	// variable decleration
	var name string = "MICROSOFT"
	var price float64
	var quant int

	// variable assignment
	price, quant = 1000.23, 5

	// shorthand decleration
	profit := 100

	var loss bool

	totalCost := price * float64(quant)

	fmt.Println(reflect.TypeOf(quant))
	fmt.Println(reflect.TypeOf(float64(quant)))

	fmt.Println(name)
	fmt.Println(price)
	fmt.Println(quant)
	fmt.Println(profit)
	fmt.Println(loss)
	fmt.Println(totalCost)
}
