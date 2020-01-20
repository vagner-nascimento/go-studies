package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	surname string
}

func (p person) getFullName() string {
	return fmt.Sprintf("%s %s", p.name, p.surname)
}

// Note that to alter the values you should to use a pointer to pass the receiver
// Different to the example at (functions/pointer/pointer.go), you don't need to use * to alter the pointer values
// If you pass the receiver without * it will not alter the struct values where you called it
func (p *person) setFullName(fullName string) {
	p.name = ""
	p.surname = ""
	parts := strings.Split(fullName, " ")
	for i, part := range parts {
		if i == 0 {
			p.name = part
		} else if len(p.surname) > 1 {
			p.surname += fmt.Sprintf(" %s", part)
		} else {
			p.surname = part
		}
	}
}

func main() {
	p1 := person{"Petter", "Simon"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Jonny")
	fmt.Println(p1.getFullName())

	p1.setFullName("Mary Juana")
	fmt.Println(p1.getFullName())

	p1.setFullName("Catherine Valdez Gregory")
	fmt.Println(p1.getFullName())
}
