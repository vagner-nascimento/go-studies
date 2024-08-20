package main

import "fmt"

func main() {
	x := 20
	fmt.Println(x)

	printX := closure()
	printX() // It will print the X declared inside the closure function
}

// All functions in go are closures
// A closure always remember and use the variables declared inside its scope (lexis contex)
func closure() func() {
	x := 10

	functionInClosure := func() {
		fmt.Println(x)
	}

	return functionInClosure
}
