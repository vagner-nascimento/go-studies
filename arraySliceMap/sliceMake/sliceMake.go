package main

import "fmt"

func main() {
	s := make([]int, 10) //make a slice with 10 postions and internal too
	fmt.Println(s, len(s), cap(s))

	s = make([]int, 10, 20) //make a slice with 10 positions and internal with 20 positions
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) //insert into s internal array the range 1 to 9
	fmt.Println(s, len(s), cap(s))

	s = append(s, 666)
	fmt.Println(s, len(s), cap(s)) // if you pass internal array's size, it creates an other array with 40 positions
}
