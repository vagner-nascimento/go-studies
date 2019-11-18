package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // an Array
	s1 := []int{1, 2, 3}  // a Slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// from position 1 (inc) until position 3 (exclude), the element at position 3 is not taken
	s2 := a2[1:3] // Slice is not an array, it is a piece(a slice) of an array
	fmt.Println(a2, s2)

	s3 := a2[:2] // new slice of the same array, it starts at 0 position
	fmt.Println(s3)

	// Slice is structure with a pointer to the first element and an size
	// Many slice can be derived from the same array
	s4 := s2[:1] // a slice of a slice, pointing to the same array
	fmt.Println(s4)
}
