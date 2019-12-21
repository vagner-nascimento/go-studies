package main

import "fmt"

// Refactored factorial recursive function
func main() {
	// result := factorial(-1) // Now is impossible to send negative numbers because uint accepts only number without signal
	result := factorial(5)
	fmt.Println(result)
}

func factorial(n uint) uint {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
