package main

import "fmt"

func main() {
	p1 := Point{2.0, 2.0}
	p2 := Point{2.0, 4.0}

	fmt.Println(cathethers(p1, p2)) //I can access private from here because they are in the same package, main
	fmt.Println(Distance(p1, p2))
}
