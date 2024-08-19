package main

import (
	"fmt"
	"time"
)

func main() {
	printHello()

	fmt.Println(scoreToConcept(9.8))
	fmt.Println(scoreToConcept(6.9))
	fmt.Println(scoreToConcept(2.1))
}

func printHello() {
	now := time.Now()

	// switch { //or...
	switch true {
	case now.Hour() < 12:
		fmt.Println("Good morning!")
	case now.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

/*
Challenge: Refactor the code bellow using switch expression:

	func scoreToConcept(n float64) string {
		if n >= 9 && n <= 10 {
		return "A"
		} else if n >= 8 && n < 9 {
		return "B"
		} else if n >= 5 && n < 8 {
		return "C"
		} else if n >= 3 && n < 5 {
		return "D"
		} else {
		return "E"
		}
	}
*/
func scoreToConcept(s float64) string {
	sInt := int(s)

	switch { // same than switch true {}
	case sInt >= 9 && sInt <= 10:
		return "A"
	case sInt >= 8 && sInt < 9:
		return "B"
	case sInt >= 5 && sInt < 8:
		return "C"
	case sInt >= 3 && sInt < 5:
		return "D"
	default:
		return "E"
	}
}
