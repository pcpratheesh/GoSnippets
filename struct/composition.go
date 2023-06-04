package main

import (
	"fmt"
)

type PersonalInformation struct {
	Name    string
	Age     int
	Address Address
}
type Student struct {
	StudentID string
	PersonalInformation
}

type Teacher struct {
	Department string
	Teaching   string
	PersonalInformation
}

type Address struct {
	Street  string
	City    string
	Country string
}

func main() {
	student := Student{
		StudentID: "#0001",
		PersonalInformation: PersonalInformation{
			Name: "John Doe",
			Age:  30,
			Address: Address{
				Street:  "123 Main St",
				City:    "New York",
				Country: "USA",
			},
		},
	}

	fmt.Println("ID:", student.StudentID)
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Address:", student.Address)

	teacher := Teacher{
		Department: "Computer",
		Teaching:   "Golang Programming",
		PersonalInformation: PersonalInformation{
			Name: "Else Stark",
			Age:  30,
			Address: Address{
				Street:  "123 Main St",
				City:    "New York",
				Country: "USA",
			},
		},
	}
	fmt.Println("Department:", teacher.Department)
	fmt.Println("Teaching:", teacher.Teaching)
	fmt.Println("Name:", teacher.Name)
	fmt.Println("Age:", teacher.Age)
	fmt.Println("Address:", teacher.Address)
}
