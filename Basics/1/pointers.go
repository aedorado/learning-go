package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	// 1. Normal way of working
	var a int = 42
	fmt.Printf("Value = %v, Type = %T\n", a, a)

	var b int = a
	fmt.Printf("Value = %v, Type = %T\n", b, b)

	b = 21
	fmt.Printf("Value = %v, Type = %T\n", a, a)
	fmt.Printf("Value = %v, Type = %T\n\n", b, b)

	// 2. Declaring p as pointer

	var p *int = &a
	fmt.Printf("Value = %v, Type = %T\n", a, a)
	fmt.Printf("Value = %v, Type = %T\n", p, p)
	fmt.Printf("Value = %v, Type = %T\n\n", &a, &a)

	// 3. Dereferncing operator

	fmt.Printf("Value = %v, Type = %T\n", *p, *p)

	// 4. Update a and dereference it
	a = 50
	fmt.Printf("Value = %v, Type = %T\n", a, a)
	fmt.Printf("Value = %v, Type = %T\n\n", *p, *p)

	*p = 90
	fmt.Printf("Value = %v, Type = %T\n", a, a)
	fmt.Printf("Value = %v, Type = %T\n\n", *p, *p)

	// 5. Go does not allow pointer arithmetic as it leads to complex code
	// Use unsafe package for pointer aritmetic. Actually don't use it.

	// 6. Can directly create a variable
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Printf("Value = %v, Type = %T\n\n", ms, ms)

	// 7. new function
	ms = new(myStruct)
	fmt.Printf("Value = %v, Type = %T\n\n", ms, ms)
	(*ms).foo = 42
	fmt.Printf("Value = %v, Type = %T\n\n", ms, ms)
	// Syntactic Sugar
	ms.foo = 57
	fmt.Printf("Value = %v, Type = %T\n\n", ms, ms)

	// 8. Zero value of a pointer
	var q *int
	fmt.Printf("Value = %v, Type = %T\n\n", q, q)

}

type myStruct struct {
	foo int
}
