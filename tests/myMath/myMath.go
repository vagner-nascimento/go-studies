package mymath

import (
	"fmt"
	"strconv"
)

// Avarage calculate the avarage of parameters value
func Avarage(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	avarage := total / float64(len(numbers))
	roundAvarage, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avarage), 64) //64 is the size of desired float

	return roundAvarage
}

// Sum sum the values
func Sum(vallues ...int) int {
	total := 0

	for _, val := range vallues {
		total += val
	}

	return total
}
