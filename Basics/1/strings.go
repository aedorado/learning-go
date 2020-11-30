package main

import (
	"fmt"
)

func main() {
	// 8. Strings
	var s string = "this is a string"
	// treated like an array (collection or array)
	fmt.Printf("Value = %v, Type = %T\n", s, s)
	fmt.Printf("Value = %v, Type = %T\n", s[0], s[0])
	// s[2] = 'o'
	// concat using +

	// convert them into a collection/slice of bytes
	b := []byte(s)
	fmt.Printf("Value = %v, Type = %T\n", b, b)

	// 9. Rune (utf-32 characters)
	// Type alias for int32
	var r rune = 'a'
	fmt.Printf("Value = %v, Type = %T\n", r, r)
}
