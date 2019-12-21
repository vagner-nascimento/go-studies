package main

import "fmt"

func main() {
	fmt.Println(getApprovedValue(6000))
	fmt.Println(getApprovedValue(5000))
	fmt.Println(getApprovedValue(4000))
	deferWithoutReturn()
}

// This is very used to close external resources, like a db connection, before leaves the function
func getApprovedValue(number int) int {
	defer fmt.Println("defer function called") // defer delays the execution of this function to last act before it is completed

	if number > 5000 {
		fmt.Println("Hig value") // The defer function will execute before the return
		return 5000
	}

	fmt.Println("Low value") // Here occurs the same
	return number
}

func deferWithoutReturn() {
	defer fmt.Println("defer without return function called")
	fmt.Println("deferWithoutReturn")
}
