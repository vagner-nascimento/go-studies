// In go, functions are first class citzen, it means that they are the most important think
package main

import "fmt"

func main() {
	fmt.Println(sum(2, 5))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(5, 2))
}

var sum = func(a, b int) int {
	return a + b
}
