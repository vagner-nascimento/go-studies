package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2) //both chare the same internal array from pos 1 until 9

	s1[0] = 7
	fmt.Println(s1, s2) // both have now 7 at firs pos of array

	s2[5] = 9
	fmt.Println(s1, s2)

	s2[9] = 42
	fmt.Println(s1, s2)

	// s2[10] = 66 //it throws this error "panic: runtime error: index out of range"
	s1 = append(s1, 66)
	fmt.Println(s1, s2)
}
