package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string
	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	// Declare and assign
	fruitArr2 := [2]string{"Blueberry", "Banana"}

	fmt.Println(fruitArr, fruitArr[1], fruitArr2)

	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:2])
}
