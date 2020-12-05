package main

import (
	"fmt"
)

// 2. Defining a struct data type
type Person struct {
	name    string
	age     int
	friends []string
}

func main() {
	fmt.Println("Structs")

	// 1. Gathers information related to a concept
	// can mix any type of data together as compated to other data types

	// 2. Declaring a variable
	p1 := Person{
		name:    "Person1",
		age:     15,
		friends: []string{"person2", "person3"},
	}

	fmt.Printf("Value = %+v, Type = %T\n", p1, p1)

	// 3. Accesing certain values using the dot syntx
	fmt.Printf("Value = %+v, Type = %T\n", p1.name, p1.name)
	fmt.Printf("Value = %+v, Type = %T\n", p1.age, p1.age)
	fmt.Printf("Value = %+v, Type = %T\n", p1.friends, p1.friends)

	// 4. Postional Syntax, don't use (don't tell)
	p2 := Person{
		"Person2",
		16,
		[]string{"person4", "person1"},
	}
	fmt.Printf("Value = %+v, Type = %T\n", p2, p2)

	// 5. Can ignore certain fields while creating a variable

	// 6. Naming Conventions

	// 7. Other packages can see the Person struct but cannot use the fields
	// Capitalize for exported usage

	// 8. Anonymous Struct
	// Does not have an independent name
	// first set of curly braces for struct, 2nd one is intializer
	// use for shortlived data models where formal type not needed
	// short lived data
	name := struct{ name string }{name: "xyz"}
	fmt.Printf("Value = %+v, Type = %T\n", name, name)

	// 9. Structs passed by value
	// use pointers to create reference

}
