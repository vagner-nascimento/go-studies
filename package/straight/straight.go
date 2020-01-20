package main

import "math"

/*
	* First letter UPPER CASE means PUBLIC (visible inside and outside of the package).
		This is valid for everything, variables, struct, attrib of struct, functions, etc...
	* First letter LOWER CASE or underline (_) means PRIVATE (visible only inside of the package).

	*For instance:
		var Point int //Ponto is public
		var point int //ponto is private
		var _Point int //_Point is private

	* You can't have duplicate functions in the same package

	* Go doesn't have visibility scope by files, only by the package

	*
*/

// Point represent a point in a cartesian plan // For public things, Go demands that you write a comment that explains it
type Point struct {
	x float64
	y float64
}

func cathethers(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return // Remember, named returns doesn't need to be explicit returned
}

//Distance is a method that calculates the distance between 2 points in a cartesian plan
func Distance(a, b Point) float64 {
	cx, cy := cathethers(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
