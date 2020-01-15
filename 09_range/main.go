package main

import "fmt"

func main() {
	ids := []int{33, 34, 1, 2, 4, 11, 3, 22}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob2": "bob@gmail.com", "Alex2": "alex@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Printf("Name: " + k + "\n")
	}
}
