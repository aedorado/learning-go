package main

import "fmt"

func main() {

	// Looping constructs are used to repeat a set of statements
	// 1. One looping construct i.e. for
	// 2. No while and do while loops

	for i := 0; i < 5; i++ {
		fmt.Println("Loops! Please subscribe :)", i)
	}

	j := 0
	for j < 5 {
		fmt.Println("Loops! Please subscribe :)", j)
		j++
	}

	for {
		fmt.Println("Loops! Please subscribe :) infinite")
	}

}
