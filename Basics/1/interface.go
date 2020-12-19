package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

type Shape interface {
	area() float64
}

func main() {
	// 1. Example, implicit implementation
	fmt.Println("Hello, playground")
	c1 := Circle{radius: 1}
	c2 := Circle{radius: 2}
	r1 := Rectangle{length: 5, width: 2}
	r2 := Rectangle{length: 2, width: 3}
	circles := []Circle{c1, c2}
	rectangles := []Rectangle{r1, r2}
	allShapes := []Shape{c1, c2, r1, r2}
	fmt.Println(totalArea(circles, rectangles))
	fmt.Println(totalAreaI(allShapes))

	// 2. Type Conversion
	var c Shape = Circle{radius: 1}
	c3 := c.(Circle)
	fmt.Printf("Value = %+v, Type = %T\n", c3, c3)

	// r3 := c.(Rectangle)
	// fmt.Printf("Value = %+v, Type = %T\n", r3, r3)

	// 3. Avoid panicing
	r3, ok := c.(Rectangle)
	if ok {
		fmt.Printf("Value = %+v, Type = %T\n", r3, r3)
	} else {
		fmt.Printf("Cannot convert\n")
	}

	// 4 - Empty Interface - Interface with no methods
	// everything can be casted to an empty interface
	// used in case of multiple things which are not type compatible
	// To use such an variable, need to convert it to appropriate type

	var c4 interface{} = Circle{radius: 1}
	r4, ok := c4.(Circle)
	if ok {
		fmt.Printf("Value = %+v, Type = %T\n", r4, r4)
	} else {
		fmt.Printf("Cannot convert\n")
	}
}

func totalArea(circles []Circle, rectangles []Rectangle) float64 {
	var total float64
	for _, c := range circles {
		total += c.area()
	}
	for _, r := range rectangles {
		total += r.area()
	}
	return total
}

func totalAreaI(shapes []Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
