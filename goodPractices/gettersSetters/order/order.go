package order

type Customer struct {
	Name string
}

type Order struct {
	c Customer
}

// Getters should have only the type name
func (o *Order) Customer() Customer {
	return o.c
}

func (o *Order) SetCustomer(c Customer) {
	o.c = c
}
