package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Panic")

	// 1. There are no exceptions in go
	// exceptional situations are considered normal
	// Opening a file that does not exist, is reasonable, not excpetional
	// errors are returned to handle such situations

	// 2. Sometimes a program cannot figure out what to do and cannot continue to run
	// This situation is called panicing in Go

	// a, b := 1, 0
	// fmt.Println(a / b)

	// 3. Panic situation in our program
	// panic("cannot continue the program")

	// 4. Very common pattern in Go
	// ParseInt doesn't have an opinon that this is a panicable situation or not
	// it just tries to execute somethign and returns error if something goes wrong
	// So mostly library functions wont panic, they will retunr errors
	// We as devs have to decide whether our application can continue even in this situation or not

	s := "Learning Go"
	// s = "100"
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		//	panic(err.Error())
	}
	fmt.Println(n)

	// 5. Panics are executed after deferred statements
	// fmt.Println("1 Start")
	// defer fmt.Println("2 Deffered")
	// panic("cannot continue the program")
	// fmt.Println("3 Post Panic")

}
