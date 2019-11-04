package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" line.")

	fmt.Println("New")
	fmt.Println(" line.")

	x := 3.141516
	//fmt.Println("The value of x is " + x) //Gives error on compile
	sx := fmt.Sprint(x) //Returns a formatted string	
	fmt.Println("The value of x is " + sx)
	fmt.Println("The value of x is", x)       //This already convert and format the desired value into a string
	fmt.Printf("The value of x is %f", x)     //The whole float
	fmt.Printf("\nThe value of x is %.2f", x) //The float with only 2 decimal fields

	a, b, c, d := 1, 1.9999, false, "Wow"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d) // BE AWARE: %.xf (%.1f) rounds the value / %t means boolean
	fmt.Printf("\n%v %v %v %v", a, b, c, d)         // %v = print value as it appears, the most generic format
}
