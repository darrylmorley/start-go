package main

import "fmt"

// Define a struct
type contactInfo struct {
	email   string
	zipCode int32
}

// Embed a struct in a struct
type person struct {
	firstName string
	lastName  string
	age       int16
	contact   contactInfo
}

func main() {
	// Declare a struct 1
	p := person{"Adam", "Anderson", 34, contactInfo{"adamanderson@gmail.com", 90210}}

	// Declare struct 2
	a := person{firstName: "John", lastName: "Johnson", age: 22}

	// Declare a struct 3
	var barry person

	// Update struct
	barry.firstName = "Barry"
	barry.lastName = "Bethel"
	barry.age = 55

	// Reference memory address for 'barry', pointer
	// barryPointer := &barry
	// barryPointer.updateName("Mark")

	// Shortcut to above works because the receiver type is *person
	barry.updateName("Mark")

	p.print()
	a.print()
	barry.print()
}

// Pointers!
// Turn a memory address into a value with *address, below
// Turn a value into a memory address with &value, above

// *person is a type description, it means we're working with a pointer to a person, it is not an operator here.
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointerToPerson is an operator that means we want to manipulate the value that the pointer is referncing.
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
