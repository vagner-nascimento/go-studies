package main

import "fmt"

type course struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	thing = "John"
	fmt.Println(thing)

	thing = 468
	fmt.Println(thing)

	thing = course{"Golang Course"}
	fmt.Println(thing)

	type dynamic interface{}
	var anotherThing dynamic = "Jess"
	fmt.Println(anotherThing)

	anotherThing = true
	fmt.Println(anotherThing)

	anotherThing = 3.14
	fmt.Println(anotherThing)

	anotherThing = course{name: "Functional Proggraming"}
	fmt.Println(anotherThing)
}
