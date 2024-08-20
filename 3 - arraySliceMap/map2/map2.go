package main

import "fmt"

func main() {
	//literal way to initialize maps:
	emplAndSalaries := map[string]float64{
		"Jose John":      11000.12,
		"Gabriela Silva": 15214.40,
		"Petter Jr":      1200.00, //This last comma is necessary. Instead you put "}" in its place, like bellow
		// "Pedro Jr":    1200.00}
	}

	emplAndSalaries["Rafael Son"] = 1350.0 //It creates this value into map
	delete(emplAndSalaries, "Inexistent")  //It don't throws errors

	fmt.Println(emplAndSalaries)

	//the first value will be always the key of map, and the second the value
	for key, value := range emplAndSalaries {
		fmt.Printf("Name: %v, Salary: %.2f\n", key, value)
	}
}
