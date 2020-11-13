package main

import (
	"fmt"
)

func main() {
	fmt.Println("Operator combine operands into expressions")
	// a := 1 + 4

	fmt.Println("Arithmetic Operators")
	fmt.Println("10 + 5 = ", 10+5)
	fmt.Println("10 - 5 = ", 10-5)
	fmt.Println("10 * 5 = ", 10*5)
	fmt.Println("13 / 5 = ", 13/5)
	fmt.Println("14 % 5 = ", 14%5)

	s := "Forrest " + "Gump"
	fmt.Println(s)

	fmt.Println("Relational Operators")
	fmt.Println("14 == 5 -> ", 14 == 5)
	fmt.Println("14 == 14 -> ", 14 == 14)

	fmt.Println("14 != 5 -> ", 14 != 5)
	fmt.Println("14 != 14 -> ", 14 != 14)

	fmt.Println("14 > 5 -> ", 14 > 5)
	fmt.Println("14 > 14 -> ", 14 > 14)
	fmt.Println("14 >= 14 -> ", 14 >= 14)

	fmt.Println("14 < 5 -> ", 14 < 5)
	fmt.Println("14 < 14 -> ", 14 < 14)
	fmt.Println("14 <= 14 -> ", 14 <= 14)

	fmt.Println("Logical Operators")

	fmt.Println("false && false -> ", false && false)
	fmt.Println("false && true -> ", false && true)
	fmt.Println("true && false -> ", true && false)
	fmt.Println("true && true -> ", true && true)

	fmt.Println("true && true && true -> ", true && true && true)
	fmt.Println("true && true && false -> ", true && true && false)

	fmt.Println("false || false -> ", false || false)
	fmt.Println("false || true -> ", false || true)
	fmt.Println("true || false -> ", true || false)
	fmt.Println("true || true -> ", true || true)

	fmt.Println("!true -> ", !true)
	fmt.Println("!false -> ", !false)

}
