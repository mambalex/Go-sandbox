package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string) // key value

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Alex"] = "alex@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	// Declare map and kv
	emails2 := map[string]string{"Bob2": "bob@gmail.com", "Alex2": "alex@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails2)
	fmt.Println(len(emails))
	fmt.Println(emails["Alex"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
