package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int { // (num1, num2 int)   works as well
	return num1 + num2
}
func main() {
	fmt.Println(greeting("Alex"))
	fmt.Println(getSum(3, 4))
}
