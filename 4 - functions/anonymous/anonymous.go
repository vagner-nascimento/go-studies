package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function without return")
	}()

	func(a, b int) {
		fmt.Println("Anonymous function params", a, b)
	}(2, 5)

	r := func(a, b int) int {
		return a * b
	}(2, 5)

	fmt.Println("Anonymous func res", r)
}
