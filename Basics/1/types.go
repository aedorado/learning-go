package main

import (
	"fmt"
	"reflect"
)

func main() {
	// fmt.Println(("Hello\"playground"))
	fmt.Println(reflect.TypeOf("Hello, playground"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(100))
	fmt.Println(reflect.TypeOf(100.45637))
}

// data types

// how a certain piece of data is intended to be used
// numbers -> + - * /
// text -> /

// rune -> 'a', 'b' .. '\n'
// text

// boolean -> true / false
