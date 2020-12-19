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
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type Rect struct {
	length, width float64
}

func (r Rect) area() float64 {
	return r.length * r.width
}
func (r Rect) perim() float64 {
	return 2*r.length + 2*r.width
}

func measure(s Shape) {
	fmt.Println(s)
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perim())
}

type Shape interface {
	area() float64
	perim() float64
}

func main() {
	// 1. Example, implicit implementation
	fmt.Println("Hello, playground")
	c1 := Circle{radius: 1}
	r1 := Rect{length: 5, width: 2}

	allShapes := []Shape{c1, r1}
	fmt.Printf("%+v\n", allShapes)

	// 2. Type Conversion
	var c Shape = Circle{radius: 1}
	c3 := c.(Circle)
	fmt.Printf("Value = %+v, Type = %T\n", c3, c3)

	// r3 := c.(Rect)
	// fmt.Printf("Value = %+v, Type = %T\n", r3, r3)

	// 3. Avoid panicing
	r3, ok := c.(Rect)
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
