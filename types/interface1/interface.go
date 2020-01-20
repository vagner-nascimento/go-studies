package main

import "fmt"

/*
	IMPORTANT: If the interface has more functions, it MUST be implemented on the struct
	otherwise, it will cause compilation errors
*/
type printable interface {
	toString() string
}

type person struct { //There's no way to explicit implements the interface with a word like "implements"
	name    string
	surname string
}

type product struct {
	name  string
	price float64
}

/*
	In Go, you can't implicit inform that a struct implents a interface
	it is inferred by the compiler.
	The only thing you must to is create functions with the dame signature of the interface
	then, the compiler identifies it and execute your polymorphic functions
*/
func (per person) toString() string {
	return fmt.Sprintf("%s %s", per.name, per.surname)
}

func (prod product) toString() string {
	return fmt.Sprintf("%s R$%.2f", prod.name, prod.price)
}

func print(p printable) {
	fmt.Println(p.toString())
}

func main() {
	//You can declare at this way, informing that it is of your interface type
	var thing printable = person{"Robert", "Langdon"}
	print(thing)

	//Then you can change this variable to any other struct which implements the interface
	thing = product{"Jeans Pant", 60.91}
	print(thing)

	//thing = "I can't do that" //You can't change to types that not implements the interface's functions
	per := person{"Michel", "Stuart"}
	print(per)

	prod := product{"Leather Shoes", 99.80} //You can directly declare the struct and use the polymorphic function too
	//prod = person{"Julian", "Rodriguez"}    //Yoy can't do it as well, you should to explicit declare with the interface type
	print(prod)
}
