package main

import "fmt"

func main() {
	empByLetter := map[string]map[string]float64{
		"A": {
			"Ana":    15456.54,
			"Amaret": 8657.98,
		},
		"J": {
			"John": 10457.78,
		},
		"P": {
			"Petter": 2547.87,
			"Paul":   225487.11,
			"Pamela": 54847.58,
		},
	}

	fmt.Println(empByLetter)

	// delete(empByLetter, "A")

	for letter, employees := range empByLetter {
		fmt.Println("Employees with letter", letter)

		for emp, salary := range employees {
			fmt.Printf("Name: %v, Salary: %.2f\n", emp, salary)
		}

		fmt.Println()
	}
}
