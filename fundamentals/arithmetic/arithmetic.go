package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Subtraction =>", a-b)
	fmt.Println("Division =>", a/b)
	fmt.Println("Multiplication =>", a*b)
	fmt.Println("Module % ('rest of division') =>", a%b)

	// bitwise
	// a = 11, b = 10 (bits representation)
	fmt.Println("AND =>", a&b) // = 10 = 2
	fmt.Println("OR =>", a|b)  // = 11 = 3
	fmt.Println("XOR =>", a^b) // 01 = 1

	// Math operations
	c := 3.0
	d := 2.0

	fmt.Println("math.Max => ", math.Max(float64(a), float64(b)))
	fmt.Println("math.Min => ", math.Min(c, d))
	fmt.Println("math.Pow => ", math.Pow(c, d))
}
