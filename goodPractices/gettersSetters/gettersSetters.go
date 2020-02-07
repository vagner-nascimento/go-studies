package main

import (
	"encoding/json"
	"fmt"
)

type customer struct {
	Name string
}

type order struct {
	c customer
}

// Getters should have only the type name
func (o *order) Customer() customer {
	return o.c
}

func (o *order) SetCustomer(c customer)  {
	o.c = c
}

func main() {
	o := order{}
	o.SetCustomer(customer{Name:"Gerald"})

	b, _ := json.Marshal(o.Customer())
	fmt.Printf("Customer from getter %s\n", string(b))
}