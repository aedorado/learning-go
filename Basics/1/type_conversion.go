package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Type Conversion")

	var a int = 65
	fmt.Printf("Value = %v, Type = %T\n", a, a)

	var b float64 = float64(a)
	fmt.Printf("Value = %v, Type = %T\n", b, b)

	// T(v)
	var c int = int(b)
	fmt.Printf("Value = %v, Type = %T\n", c, c)

	var d string = strconv.Itoa(c)
	fmt.Printf("Value = %v, Type = %T\n", d, d)
}
