package main

import "fmt"

type item struct {
	productID int
	qtd       int
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) getTotalPrice() float64 {
	total := 0.0
	for _, item := range o.items {
		total += item.price * float64(item.qtd)
	}

	return total
}

func main() {
	o := order{
		userID: 1,
		items: []item{
			item{productID: 1, qtd: 2, price: 12.10}, //Its better initialize an struct in a declarative way to avoid confusion
			item{2, 1, 23.49},                        //This isn't clear, in a big program, where struct are declared in a place and used in others,...
			item{11, 100, 3.13},                      // ... it can turn in HUGE mess
		},
	}

	fmt.Printf("The order's total price is $%.2f\n", o.getTotalPrice())
}
