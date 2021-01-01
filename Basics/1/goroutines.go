package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 1. Spin off a green thread
	// run the hello function in the green thread
	go hello()

	// most languages use OS threads
	// this means that they have
	// an individual function call stack dedicated
	// to the execution of the code of the thread
	// 1 MB ran, take time to setup
	// leads to thread pools as creation and destruction of threads is difficult

	// Go used thread model similar to Earlang - Green threads
	// Instead of massive heavy overhead threads
	// we will create an abstraction of a thread called a go routine
	// Go runtime has a scheduler that maps goroutines onto OS threads
	// And keeps on assigning the goroutine
	// some amounts of processing time on those CPU threads
	// we dont interact with low level threads, but high level go routines

	// Advantages
	// Goroutines can be started with very very small stack spaces
	// And as they can be reallocated very quickly
	// they can be easily created or destroyed
	// In a go application we can have 10000s of go routines and have not issues

	// Threads in other langs with 1MB overhead cannot scale in this manner

	// 2. Run this program, no output
	// Main function is executing in a goroutune
	// Main spawns a goroutine
	// but the application exits as soon as the application runs
	// hello function had no time to print its message

	// 3. Put a sleep

	// 4. Anonymous functions
	// Go can form closures
	msg := "hello"
	// go func() {
	// 	fmt.Println(msg)
	// }()
	// // 5. Reassign the variable, goroutine prints the updated value
	// // The go scheduler will not interrupt the main thread generally
	// // the value in the next line is assigned before the goroutine starts
	// msg = "goodbye"

	// 6. Can add argument to ananymous function
	msg = "hello"
	go func(msg1 string) {
		fmt.Println(msg1)
	}(msg) // value passed in while invoking message
	msg = "goodbye"

	time.Sleep(100 * time.Millisecond)

	// Number of CPU threads available
	// By default equal to the number of cores on the machine
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}

func hello() {
	fmt.Println("Hello form hello function!")
}
