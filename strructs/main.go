package main

import (
	"fmt"
)

type contactInfo struct {
	email       string
	phoneNumber int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	// var alex person
	// fmt.Printf("%+v", alex)

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:       "alexanderson1@mail.com",
			phoneNumber: +123456789,
		},
	}
	// & : Get memory address of value variable is pointing to we can omit it as well
	// (&alex).updateName("Alec")
	alex.updateName("Alec")
	alex.print()

}

// * get value memory address is pointing to
func (p *person) updateName(newfirstName string) {
	(*p).firstName = newfirstName
}

func (p person) print() {
	fmt.Println(p)
}
