// This is a one line comment(c like): Executable programs starts from main package
package main

/*
	This is a block comment(c like)
	Go codes are organized in packages
	to use these ones, you should to declare their imports (one or more)
*/
import "fmt"

/*
	The main function is the entrance of a go application.
	This is the first function that the program will execute when called
*/
func main() {
	fmt.Print("First ")
	fmt.Print("Program!")
}
