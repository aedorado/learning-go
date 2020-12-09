package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions as first class data types in Go")

	func() {
		fmt.Println("Hello Go!")
	}()

	// 2. Wrong way
	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	// 3. Right way
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	// 4. Define function as variable
	f := func() {
		fmt.Println("Woo hoo Go")
	}
	f()
	// fmt.Printf("Type = %T, Value = %+v\n", f, f)

	// 5. Long decleration
	var f1 func() = func() {}
	f1()

	// 6. Divide function
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("Cannot divide by 0")
		}
		return a / b, nil
	}
	d, err := divide(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
