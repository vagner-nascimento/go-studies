package main

import "fmt"

func main() {
	i := 1

	var p *int = nil // Only pointers can receive nil, wich means that it isn't pointing to nothing
	fmt.Println(i, p)

	p = &i                // p recieves the memory address of i
	fmt.Println(i, p, &i) // with "&" we get memory address of i, which will be the same of p
	fmt.Println(i, *p)    // with "*" we can dereference it, getting its value instead the memory address

	*p++ // Access and change the value on this address, what also changes the value of i becase both are in the same memory address
	fmt.Println(i, *p)

	i++ // It also change the value into p (*p)
	fmt.Println(i, *p)

	// p++ // Go doens't has arithimetic over pointers
}
