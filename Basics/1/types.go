package main

import (
	"fmt"
)

func main() {
	fmt.Println("Primitive Data Types in Go: Boolean, Numeric & Text Types")

	// 1. Boolean represents two states true or false
	fmt.Printf("Value = %v, Type = %T\n", true, true)
	fmt.Printf("Value = %v, Type = %T\n", false, false)

	// Integers
	// 2. Signed Integers
	// int - unspecified size, based on platform implementation
	fmt.Printf("Value = %v, Type = %T\n", 42, 42)

	// 3. Integers of specific sizes
	// int8, int16, int32, int64
	// int8		8 bits		-128 to 127
	// int16	16 bits		-32768 to 32767
	// int32	32 bits		-2147483648 to 2147483647
	// int64	64 bits		-2^63 to (2^63)-1
	// Use math Big Package for larger numbers: https://golang.org/pkg/math/big/

	// 4. Unsigned Integers
	// uint8	8 bits		0 to 255
	// uint16	16 bits		0 to 65535
	// uint32	32 bits		0 to 4294967295

	// 5. Can't add values of different types

	// 6. To store decimal numbers:
	// Floating point numbers IEEE754 standard float32, float64
	// float32	32 bits	 +-1.18E-38 to +-3.4E36
	// float64	64 bits	 +-2.23E-308 to +-1.8E308
	fmt.Printf("Value = %v, Type = %T\n", 3.1415, 3.1415)
	fmt.Printf("Value = %v, Type = %T\n", 3.712e2, 3.712e2)
	fmt.Printf("Value = %v, Type = %T\n", 1.1e3, 1.1e3)

	// 7. Complex type (complex64, complex128)
	var n complex64 = 1 + 2i
	fmt.Printf("Value = %v, Type = %T\n", n, n)
	// n = 3i
	fmt.Printf("Value = %v, Type = %T\n", n, n)
	fmt.Printf("Value = %v, Type = %T\n", real(n), real(n))
	fmt.Printf("Value = %v, Type = %T\n", imag(n), imag(n))

	var m complex128 = complex(3, 4)
	fmt.Printf("Value = %v, Type = %T\n", m, m)

}
