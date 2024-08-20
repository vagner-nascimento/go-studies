package main

import "fmt"

/*
This is a Go convention to initialize things.
This is called before execute the main method when the file (or files) is called.
*/
func init() {
	fmt.Println("Init called")
}

func main() {
	fmt.Println("Main called")
}
