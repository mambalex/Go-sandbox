package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeing method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Alex", lastName: "Zhang", city: "Sydney", gender: "m", age: 25}

	//Alternative
	person2 := Person{"Angela", "Lee", "Sydney", "f", 25}

	fmt.Println(person1)
	fmt.Println(person1.lastName)
	person1.age++
	fmt.Println(person1)
	person1.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person2)
	person2.getMarried("Zhang")
	fmt.Println(person2.greet())
}
