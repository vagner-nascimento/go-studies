package main

import "fmt"

func main() {
	// approved := [...]string{"Mary", "John"} // Using ... means that go will calculate the size of the array
	approved := []string{"Mary", "John", "Mike", "Jack"} // Using [] (empty) you're creating a slice
	printApproved(approved...)                           // It doesn't works with arrays, only with slices
}

func printApproved(approved ...string) {
	fmt.Println("Approved list:")
	for i, name := range approved {
		fmt.Printf("%d) %s\n", i, name)
	}
}
