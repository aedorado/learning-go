package main

import (
	"fmt"
)

func main() {
	fmt.Println("Understanding conditional statements")

	// Problem Statement:
	// Grade a student according to score as follows:
	// 0-40  -> F
	// 41-80 -> B
	// 81-90 -> A
	// 91+   -> A+

	score := 100

	switch {
	case score >= 0 && score <= 40:
		fmt.Println("F")
	case score >= 41 && score <= 80:
		fmt.Println("B")
	case score >= 81 && score <= 90:
		fmt.Println("A")
	case score >= 91:
		fmt.Println("A+")
	}

}
