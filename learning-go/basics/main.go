package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Anderson",
	}
	jimPointer := &jim
	fmt.Println(jimPointer)
	jimPointer.updateFirstName()
	jim.printPerson()

}

func (pointerToPerson *Person) updateFirstName() {
	(*pointerToPerson).firstName = "Jimmy"
}

func (p Person) printPerson() {
	fmt.Println(p.firstName, p.lastName)
}
