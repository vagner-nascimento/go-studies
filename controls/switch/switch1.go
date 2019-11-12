package main

import "fmt"

func main() {
	fmt.Println(scoreToConcept(9.8))
	fmt.Println(scoreToConcept(6.9))
	fmt.Println(scoreToConcept(2.1))
	fmt.Println(scoreToConcept(-1))
}

func scoreToConcept(score float64) string {
	nota := int(score)

	/*
		Go's switch don't execute one case after other by default, like Java and other languages do (which dumb)
		This is GREAT, once it rarely is used.
		If you want to do that, you should to use "fallthrough" key word
	*/
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid score"
	}
}
