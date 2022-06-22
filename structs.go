package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
	salary    float32
}

func structs() {

	var emp1 = Employee{
		firstName: "Rohan",
		lastName:  "Maity",
		id:        2563,
		salary:    10000.00,
	}

	fmt.Printf("%T\n", emp1)
}
