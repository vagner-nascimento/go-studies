package main

import "fmt"

func main() {
	val := 1

	incCopyOfVal(1)
	fmt.Println(val)

	incReferencedVal(&val) // Use & to send the reference of the variable
	fmt.Println(val)
}

func incCopyOfVal(val int) {
	val++
}

func incReferencedVal(val *int) {
	// The * is used to define a point and to access its value
	*val++
}
