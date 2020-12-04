package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	// 1. What is a map?
	// Map of Employee Salary
	// one key maps to one value
	// literal sysntax
	// what are keys and values?
	population := map[string]int{}

	empSal := map[string]int{
		"Nami":    10000,
		"Lufy":    20000,
		"Zoro":    5000,
		"Brook":   45000,
		"Chooper": 1000,
	}

	fmt.Printf("Value = %v, Type = %T\n", population, population)
	fmt.Printf("Value = %v, Type = %T\n", empSal, empSal)

	// 2. Map Keys should be able to be tested for equality
	// Can be: bool, int, float, struct, array, string, pointer, interfaces, channels
	// Cannot be: Slices, Maps, Functions

	// m := map[[]int]string{}
	m := map[[3]int]string{}
	fmt.Printf("Value = %v, Type = %T\n", m, m)

	// 3. Make function
	// when we dont know what values will be in the map at the time of creating it
	map1 := make(map[string]int)
	fmt.Printf("Value = %v, Type = %T\n", map1, map1)

	// 4. Accesing the values form a map
	fmt.Printf("Value = %v, Type = %T\n", empSal["Nami"], empSal["Nami"])

	// 5. Updating values
	empSal["Frankie"] = 40000
	fmt.Printf("Value = %v, Type = %T\n", empSal["Nami"], empSal["Nami"])
	fmt.Printf("Value = %v, Type = %T\n", empSal, empSal)

	// 6. deleting keys
	delete(empSal, "Brook")
	fmt.Printf("Value = %v, Type = %T\n", empSal, empSal)

	// 7. Accessing keys not present in the map
	fmt.Printf("Value = %v, Type = %T\n", empSal["Brook"], empSal["Brook"])
	fmt.Printf("Value = %v, Type = %T\n", empSal["rook"], empSal["rook"])

	// 8. Optional Comma Ok syntax
	empSalValue, ok := empSal["Nami"]
	fmt.Printf("%v %v\n", empSalValue, ok)
	empSalValue, ok = empSal["Mami"]
	fmt.Printf("%v %v\n", empSalValue, ok)

	// 9. Check only for presence
	_, ok = empSal["Nami"]
	fmt.Printf("%v %v\n", empSalValue, ok)

	// 10. length
	fmt.Printf("Length = %v\n", len(empSal))

	// 11. Maps are passed by reference
	mapRef := empSal
	fmt.Printf("Value = %v, Type = %T\n", mapRef, mapRef)
	delete(mapRef, "Lufy")
	fmt.Printf("Value = %v, Type = %T\n", mapRef, mapRef)
	fmt.Printf("Value = %v, Type = %T\n", empSal, empSal)

}
