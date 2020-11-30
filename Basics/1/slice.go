package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices")

	// 1. Size of an array has to be known at compile time
	a := [...]int{1, 2, 3}
	fmt.Printf("Value = %v, Type = %T\n", true, true)

	// 2. Literal syntax to Declare a slice;
	// looks like an array and works like an array
	// some exceptions
	b := []int{1, 2, 3}

	// 3. length
	// 4. capacity (slice is projection of backing/underlying array)
	fmt.Printf("Length = %v, Capacity=%v\n", len(b), cap(b))

	// 5. Unlike arrays, slices are refernce types
	// thus be careful
	a := [...]int{1, 2, 3}
	b := a
	b[0] = 9
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// 6. Creating a slice form an underlying array
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]

	// 7. All slices refer to the same underlying array
	d[0] = 100

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// 8. Make function
	a := make([]int, 3, 10) // type, length and capacity
	fmt.Println(a)
	fmt.Printf("Length = %v, Capacity=%v\n", len(b), cap(b))

	// 8. slices can vary in length unlike arrays
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length = %v, Capacity=%v\n", len(a), cap(a))
	a = append(a, 1)
	fmt.Printf("Length = %v, Capacity=%v\n", len(a), cap(a))
	// Earlier array had 0 elements
	// on appending a new array was created with increased capacity
	// this can be expensive with larger arrays so use make

	// 9. Append multiple values
	// Go keeps on doubling the capacity of the slice
	a = append(a, 2, 3, 4, 5)
	fmt.Printf("Length = %v, Capacity=%v\n", len(a), cap(a))

	// 10. Concat 2 slices in Go
	a = append(a, []int{1, 2, 3})    // doesnt work
	a = append(a, []int{1, 2, 3}...) // works

	// 11. Stack operations
	// PUSH -> append
	a := []int{1, 2, 3, 4, 5}
	// POP
	b := a[:len(a)-1]

	// 12. Remove elements from middle of slice
	b := append(a[:2], a[3:]...)

	// 13. The underlying array is not messed up
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(a)
	fmt.Println(b)

}
