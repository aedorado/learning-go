package main

import (
	"fmt"
	"strconv"
)

var v int = 32

var (
	name      string = "My Name"
	jobTitle         = "Engineer"
	otherInfo        = "Extra Info"
)

var a int = 42

func main() {
	fmt.Println("Variables")

	// 1. Declaring variable
	var i int
	fmt.Println(i)

	// 2. Assigning a value
	i = 42

	// 3. Value of a variable can be changed
	i = 21

	// 4. Combine declration and assignment
	var j int = 42
	fmt.Println(j)

	// 5. Shorthand assignment
	k := 42
	fmt.Println(k)

	// 6. Assigning wrong type (Go is strongly and statically typed)
	var s string
	s = "characters"
	// s = 42
	fmt.Println(s)

	// 7. Declare at package level (cant use shorthand assignment)

	// 8. Variable block to clean clutter
	var name string = "My Name"
	var jobTitle string = "Engineer"
	var otherInfo string = "Extra Info"
	fmt.Println(name, jobTitle, otherInfo)

	// 9. Redeclaring varibales in the same scope (Shadowing)
	fmt.Println(a)
	var a int = 420
	fmt.Println(a)
	// var a int = 42

	// 10. Variables always have to be used

	// 11. Levels of scoping
	// local, package, global
	// > Lowercase variables scoped to package
	// > Uppercase variables exposed outside package

	// 11. Naming conventions - 1
	// concise name for loops etc
	// descriptive names for variables used for longer time
	// verbose names for package level variables
	// acronyms should be all upper case Eg. EmpID (for readability)

	// 12. Type conversion has to be explicit
	var b int = 42
	fmt.Printf("Value = %v, Type = %T\n", b, b)

	var c float64 = float64(b)
	fmt.Printf("Value = %v, Type = %T\n", b, b)

	var d int = int(c)
	fmt.Printf("Value = %v, Type = %T\n", d, d)

	// 13. Convert int to string using strconv
	var e string
	e = string(a)
	fmt.Printf("Value = %v, Type = %T\n", e, e)
	e = strconv.Itoa(a)
	fmt.Printf("Value = %v, Type = %T\n", e, e)

}
