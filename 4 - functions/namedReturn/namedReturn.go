package main

import "fmt"

func main() {
	v1, v2 := exchange(1, 2) // Its not necessary to use the same name declared on function
	fmt.Println(v1, v2)
	first, second := exchange(8, 9) // But you can do it
	fmt.Println(first, second)
}

func exchange(p1, p2 int) (first, second int) { // Second and first are named return variables
	first = p1
	second = p2
	return // Clean return, it will return you named variables in the same order that thoose were declared
	//return second, first // You can return the variables and change the order of return as well
}
