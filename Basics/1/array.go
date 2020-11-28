package main

import (
	"fmt"
)

func main() {

	// 1. Why do we need arrays?
	// price1 :=
	// price2 :=
	// price3 :=
	// ...

	// 2. Declaring an array
	prices := [3]int{1, 2, 3}
	fmt.Printf("Prices: %v\n", prices)

	// 3. Contiguous in memory

	// 4. Can skip size if providing initializer list
	prices = [...]int{1, 2, 3}

	// 5. Array with zero values
	var products [3]string

	// 6. Accessing and updating values in the array
	fmt.Printf("Products: %v\n", products)
	products[0] = "Paints"
	products[1] = "Colors"
	products[2] = "Brushes"
	fmt.Printf("Products: %v\n", products)

	// 7. Length of the Array
	fmt.Printf("Number of products: %v\n", len(products))

	// 8. Arrays can be of any types

	// 9. Arrays of arrays
	var matrix [2][2]int
	matrix[0] = [2]int{0, 1}
	matrix[1] = [2]int{2, 3}
	fmt.Printf("Matrix: %v\n", matrix)

	// 10. ARRAYS ARE VALUES IN GO, AND NOT REFERENCES
	a := [...]int{1, 2, 3}
	b := a
	b[0] = 9
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// 11. USE of refernces is by using pointers
	c := [...]int{1, 2, 3}
	d := &c
	d[0] = 9
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)

	// 12. Arrays have fixed size that has be knows at compile time. Slices can have variable size

	// 13. For range loops to work with arrays
	for i := 0; i < len(products); i++ {
		// fmt.Println(i, products[i])
	}

	for _, v := range products {
		fmt.Println(v)
	}
}
