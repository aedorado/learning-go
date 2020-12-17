package main

import (
	"fmt"
)

func main() {

	// 6. Panics are not fatal
	// they are fatal only when we panic upto the go runtime
	// in which case, the go runtime kills this application

	fmt.Println("Main: Start")
	//defer tryToRecover()
	//panic("cannot continue the program")
	iCausePanics()
	fmt.Println("Main: Post Panic")
}

func iCausePanics() {

	fmt.Println("about to panic")
	//defer tryToRecover()
	panic("func: canot continue to execute")
	fmt.Println("i just panicked")
}

func tryToRecover() {
	// recover returns nil if the application is not panicing
	// 	   returns the error that has caused the panic
	if err := recover(); err != nil {
		fmt.Println("Error: ", err)
	}
}
