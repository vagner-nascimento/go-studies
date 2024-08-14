package main

import "fmt"

// Go doens't has ternary operator
func getResult(rate float64) string {
	//return rate >= 6 ? "Approved" : "Reproved"
	if rate >= 6 {
		return "Approved"
	}
	return "Reproved"
}

func main() {
	result := getResult(6.2)
	fmt.Println(result)
}
