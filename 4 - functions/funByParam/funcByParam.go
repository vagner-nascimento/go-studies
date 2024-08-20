package main

import "fmt"

func main() {
	result := exec(multiply, 3, 4)
	fmt.Println(result)
}

func multiply(a, b int) int {
	return a * b
}

func exec(function func(int, int) int, p1, p2 int) int { // You should to specify the function signature
	return function(p1, p2)
}
