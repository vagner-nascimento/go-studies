package main

import "fmt"

func main() {
	// var approved map[int]string //maps MUST be initialized
	approved := make(map[int]string) // map[{index type}]{value type} | make initializes a map
	approved[1] = "Mary"
	approved[2] = "Petter"
	approved[3] = "Ann"
	fmt.Println(approved)

	//iteration over index and value
	for cpf, name := range approved {
		fmt.Printf("%s (ID: %d)\n", name, cpf)
	}

	//getting a value straight from the index
	fmt.Println(approved[2])

	delete(approved, 3)
	fmt.Println(approved[3])
	fmt.Println(approved)
}
