package main

import "fmt"

type score float64

/*
	func (s float64) isBetween(start, end float64) bool { } // You can't associate a primitive value to a function
	You should to create a personalized type with desired primitive type to do that
*/
func (s score) isBetween(start, end float64) bool {
	return float64(s) >= start && float64(s) <= end
}

func getScoreConcept(s score) string {
	switch {
	case s.isBetween(9.0, 10.0):
		return "A"
	case s.isBetween(7.0, 8.99):
		return "B"
	case s.isBetween(5.0, 7.99):
		return "C"
	case s.isBetween(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(getScoreConcept(9.8))
	fmt.Println(getScoreConcept(7.2))
	fmt.Println(getScoreConcept(6.9))
	fmt.Println(getScoreConcept(4.11))
	fmt.Println(getScoreConcept(2.1))
}
