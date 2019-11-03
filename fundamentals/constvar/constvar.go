package main

import (
	"fmt"
	m "math" //m references the math, which can be used as an alias into the code
)

func main() {
	const PI float64 = 3.1415
	var rain = 3.2 // float64 type inferred by copiller (avoid use lke that)

	//short way to declare and initialize a variable (preferred way)
	//  = means just attribution to a value into an existent variable
	// := creates a variable and set a initial value to it
	area := PI * m.Pow(rain, 2)
	fmt.Println("Circumference area is", area)

	//Construction blocks
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	//One line declaration
	var e, f bool = true, false //It gives true to e and false to f
	fmt.Println(e, f)

	//One line declaration with type inference (cream de la cream, simple the best ever)
	g, h, i := 2, false, "wow!" //Here occurs the same as above
	fmt.Println(g, h, i)
}
