package main

import (
	"fmt"
	"time"
)

func main() {
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	//it ALLWAYS compare the values, and NEVER the memory reference
	fmt.Println("d1 == d2", d1 == d2)
	fmt.Println("d1.Equal(d2)", d1.Equal(d2))

	type Person struct {
		name string
	}

	p1 := Person{"John"}
	p2 := Person{"John"}
	fmt.Println("p1 == p2", p1 == p2)
}
