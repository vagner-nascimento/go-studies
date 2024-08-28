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

	s3 := a2[:2] // or s3 := a2[0:2], new slice of the same array, it starts at 0 position, that can be ommited
	fmt.Println(s3)

	// Slice is a structure with a pointer to the first element and an the size of the slice
	// Many slices can be derived from the same array
	s4 := s2[:1] // a slice of a slice, pointing to the same array
	fmt.Println(s2, s4)
	fmt.Println("len s2", len(s2))

	// once the slice is a pointer, if we change the value on array, it also will be changed on the slice of this array
	// in the same way, if we change the value on slice, it also will changed on the array
	fmt.Println("a2", a2, "s2", s2)
	a2[1] = 9
	fmt.Println("a2", a2, "s2", s2)
	s2[0] = 2
	fmt.Println("a2", a2, "s2", s2)

	// slice of slice has the same behavior
	// s4 is a slice from the s2 (slice of slice)
	fmt.Println("a2", a2, "s2", s2, "s4", s4)
	s4[0] = 9
	fmt.Println("a2", a2, "s2", s2, "s4", s4)
	a2[1] = 2
	fmt.Println("a2", a2, "s2", s2, "s4", s4)

	// slice passed by param
	fmt.Println("by param a2", a2, "s2", s2)
	fn(s2)
	fmt.Println("side effect a2", a2, "s2", s2)
	s2[0] = 2

	// array passed by param NO SIDE EFFECTS
	fmt.Println("by param s2", s2, "a2", a2)
	fnArr(a2)
	fmt.Println("NO side effect s2", s2, "a2", a2)
}

// with side effect
func fn(vals []int) {
	vals[0] = 9
}

// without side effects
func fnArr(vals [5]int) {
	vals[1] = 9
	fmt.Println("array chnaged into func", vals)
}
