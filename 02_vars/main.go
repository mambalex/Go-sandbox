package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name string = "Alex"
	// var name = "Alex"      		  // works as well
	var age int = 27 // have to use it or get errors
	// var age int32 = 27			  // works as well
	var isCool = true
	isCool = false
	// const isCool = true 			  // Can't resign

	// Shorthand
	lastname := "Zhang" // Cam resign
	size := 1.3
	email := "alexzhang@gmail.com"

	name2, email2 := "Mambalex", "mambalex@gmail.com"

	fmt.Println(name, lastname, age, isCool, email)
	fmt.Println(name2, email2)
	fmt.Printf("%T & %T\n", name, size) // the type of the value
}
