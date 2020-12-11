package main

import (
	"fmt"
)

func func1(g func(int) int) {
	a := g(10)

	fmt.Printf("Type = %T, Value = %v\n", a, a)

}

func main() {

}

func main() {
	fmt.Println("Functions as Types & Anonymous Functions")

	var f func(i int) int = func(i int) int {
		fmt.Println("Hello Go!", i)
		return 0
	}

	func1(f)

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
