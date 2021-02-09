package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "xxx@email.com",
			zipCode: 123112,
		},
	}
	// &variable --> Give me the memory address of the value this variable is pointing at
	// *pointer --> Give me the value of the memory address is pointing at

	// Turn `address` into `value` with `*address`
	// Turn `vaule` into `address` with `&vaule`

	jim.updateName("Tom")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
