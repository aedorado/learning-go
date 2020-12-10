package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	// 1. Main function in main package is entry point of Go Application

	// 2. Function syntax (func, naming, camelCase, pascalCase, parameters, body)

	// 3. Sample function to print stirng passed to it (function name, argument, etc)

	// 4. Can pass mutiple parameters (printNTimes)

	// 5. Shorthand way of passing arguments of same type sayHello

	// 6. Example of pass by value and pass by reference
	s := "original name"
	fmt.Println(s)
	passByValue(s)
	fmt.Println(s)

	// 7. Example of pass by reference
	// Passing large data structures for performance benefits
	fmt.Println(s)
	passByValue(s)
	fmt.Println(s)

	// 8. Variadic Parameters
	// can only have one and at the end
	sum(1, 2, 3, 4, 5)

	// 9. Return values
	fmt.Println(sumret(1, 2, 4))

	// 10. Return the local variable as a pointer
	// such a value is generated on the heap and not on the stack
	fmt.Println(*sumretp(1, 2, 4))

	// 11. Named return values
	// not done generally as can decrease readability
	fmt.Println(sumNamedRet(1, 2, 10, 20))

	// 12. Multiple return values form a function
	ans, err := divide(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ans)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}

func sum(values ...int) {
	fmt.Printf("Value = %v, Type := %T\n", values, values)
	total := 0
	for _, v := range values {
		total += v
	}
	fmt.Println(total)
}

func sumret(values ...int) int {
	fmt.Printf("Value = %v, Type := %T\n", values, values)
	sumv := 0
	for _, v := range values {
		sumv += v
	}
	return sumv
}

func sumNamedRet(values ...int) (result int) {
	fmt.Printf("Value = %v, Type := %T\n", values, values)
	for _, v := range values {
		result += v
	}
	return
}

func sumretp(values ...int) *int {
	fmt.Printf("Value = %v, Type := %T\n", values, values)
	sumv := 0
	for _, v := range values {
		sumv += v
	}
	return &sumv
}

func passByValue(s string) {
	s = "changed name"
	fmt.Println(s)
}

func passByRef(s *string) {
	*s = "changed name"
	fmt.Println(*s)
}

func passByReference(sf *string) {
	*sf = "changed in function"
	fmt.Println("Value in function: ", *sf)
}

func print(s string) {
	fmt.Println(s)
}

func printNTimes(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(s)
	}
}

func sayHello(s, name string) {
	fmt.Println(s, name)
}
