package main

import (
	"fmt"
	// "math"
)

const a int = 76

func main() {

	// 1. Constants use camelcase
	// 2. Exported constants begin with an upper case

	const GO_CONSTANT int = 10

	// 3. How to declare a constant
	const goCon1 int = 42
	fmt.Printf("Value = %v, Type = %T\n", goCon1, goCon1)

	// 4. Value assigned to a constant cannot be changed
	// goCon1 = 54

	// 5. Constant should be assignable at compile time!
	// const goCon2 float64 = math.Sqrt(5)

	// 6. Constants can be shadowed
	// const a float64 = 5.4
	fmt.Printf("Value = %v, Type = %T\n", a, a)

	// 7. Constants can be used just as varibales in most cases
	// const b int16 = 1
	var c int64 = 2
	// fmt.Printf("Value = %v, Type=%T \n", b + c, b + c)

	// 8. Untyped Constants
	const b = 54
	fmt.Printf("Value = %v, Type=%T \n", b+c, b+c)
}
