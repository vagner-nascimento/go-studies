package main

import "fmt"

func main() {
	//var approved map[int]string //it gives the error "panic: assignment to entry in nil map" because maps MUST be initialized with make command
	approved := make(map[int]string) // map[{index type}]{value type} | make initializes a map
	approved[86827519090] = "Mary"
	approved[68293098082] = "Petter"
	approved[44037170043] = "Ann"
	fmt.Println(approved)

	//iteration over index and value
	for cpf, name := range approved {
		fmt.Printf("%s (cpf: %d)\n", name, cpf)
	}

	//getting a value straight from the index
	fmt.Println(approved[44037170043])

	//deleting from specific index
	delete(approved, 44037170043)
	fmt.Println(approved[44037170043]) // it prints nothing
	fmt.Println(approved)
}
