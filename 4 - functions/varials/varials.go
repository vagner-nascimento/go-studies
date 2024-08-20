package main

import "fmt"

func main() {
	fmt.Printf("Avarage %.2f\n", avg(7.7, 3.14, 11))
	fmt.Printf("Avarage %.2f\n", avg(7.7))
	fmt.Printf("Avarage %.2f\n", avg())
}

func avg(numbers ...float64) float64 { // At this way, you can send as many parameters as you want or none
	if len(numbers) <= 0 {
		return 0.0
	}

	total := 0.0

	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}
