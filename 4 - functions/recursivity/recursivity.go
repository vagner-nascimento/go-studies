package main

import "fmt"

func main() {
	result, error := factorial(5)
	if error == nil {
		fmt.Println(result)
	}

	_, err := factorial(-1)

	if err != nil {
		fmt.Println(err)
	}
}

func factorial(n int) (int, error) { // Return errors is commom in go
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid number: %d", n)
	case n == 0 || n == 1:
		return 1, nil
	default:
		lastFactorial, _ := factorial(n - 1)
		return n * lastFactorial, nil
	}
}
