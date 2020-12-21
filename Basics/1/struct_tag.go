package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 1. Tags add description to the fields of a struct
// for example form data
type Animal struct {
	Name  string `required max:"100"`   // 2. format backticks with space delimited key-value pairs
	Color string `tag can be anything!` // 3. value of tags is arbitrary
}

type T struct {
	F1 string `json:"filed1"`
	F2 string `json:"filed2,omitempty"`
	F3 string `json:"filed3,omitempty"`
	F4 string `json:"-"`
}

func main() {
	a := Animal{Color: "white"}

	// 4. Getting tag values using reflect package
	typ := reflect.TypeOf(a)
	field, _ := typ.FieldByName("Name")
	fmt.Printf("Type = %T, Value = %+v\n", field.Tag, field.Tag)
	field, _ = typ.FieldByName("Color")
	fmt.Printf("Type = %T, Value = %+v\n", field.Tag, field.Tag)

	// 5. Json usecase
	t := T{"v1", "", "v3", "v4"}
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

}
