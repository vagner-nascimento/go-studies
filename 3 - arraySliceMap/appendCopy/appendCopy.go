package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 1, 2, 3) //append only works with slices
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	// slice1 = append(slice1, array1) //it doesn't works too
	slice2 := make([]int, 2)
	copy(slice2, slice1) //the slice2 has only 2 positions, so it copy only 2 first elements of the slice1, which has len 3
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 3)
	copy(slice3, slice2) //it copy only 2 first elements into the slice3 and fills the remaining spaces with 0
	fmt.Println(slice2, slice3)
}
