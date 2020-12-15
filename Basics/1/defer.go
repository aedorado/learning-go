package main

import (
	"fmt"
)

func main() {
	// 1. function runs form start to end
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)

	// 2. deferred statements run when the function executes its final statement
	// but before the function returns
	fmt.Println(11)
	defer fmt.Println(12)
	fmt.Println(13)

	// 3. deferred function calls are LIFO
	defer fmt.Println(21)
	defer fmt.Println(22)
	defer fmt.Println(23)

	// 4. Use? To close resources. Which is better done in reverse order

	// 5.

	// 6. Do not use deffered statements in a loop to close resources
	// as they are executed only when the function ends
	// better to close explicitly

	a := "start"
	defer fmt.Println(a) // value decided at the time the defer is called
	a = "end"
}
