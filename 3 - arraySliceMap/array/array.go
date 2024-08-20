package main

import "fmt"

// Arrays are homogenius (all must have same type) and statics (don't changes their size once created)
func main() {
	var scores [3]float64
	fmt.Println(scores) // if none data is informed, it fills the array with the type's default value (zero value)

	scores[0], scores[1], scores[2] = 7.8, 4.3, 9.1 // giving values in one line
	fmt.Println(scores)

	var total float64
	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}

	avarage := total / float64(len(scores))
	fmt.Printf("Avarage: %.2f\n", avarage)
}
