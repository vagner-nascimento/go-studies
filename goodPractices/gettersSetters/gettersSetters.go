package main

import (
	"encoding/json"
	"fmt"
	"github.com/vagner-nascimento/go-studies/goodPractices/gettersSetters/order"
)

func main() {
	o := order.Order{}
	// o.c :=  order.Customer{Name: "Gerald"} // Note that you can't do it, only by the Setter
	o.SetCustomer(order.Customer{Name: "Gerald"})

	// customer := o.c // Note that you can't do it, only by the Getter
	customer := o.Customer()
	b, _ := json.Marshal(customer)
	fmt.Printf("Customer from getter %s\n", string(b))

	customer.Name = "Yennefer"
	o.SetCustomer(customer)
	b, _ = json.Marshal(o.Customer())
	fmt.Printf("Changed customer name from setter %s\n", string(b))
}
