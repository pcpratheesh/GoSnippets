package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "John Doe",
		Age:  30,
	}

	personType := reflect.TypeOf(person)
	personValue := reflect.ValueOf(person)

	for i := 0; i < personType.NumField(); i++ {
		field := personType.Field(i)
		value := personValue.Field(i)

		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}
