package main

import "fmt"

func main() {
	x := 1
	y := 2

	// Go accepts only posfix operator
	x++ // x += 1 or x = x + 1
	fmt.Println(x)

	y-- // y -=1 or y = y - 1
	fmt.Println(y)

	// fmt.Println(x == y--) // it doesn't works
	// ++x // it doesn't works
}
