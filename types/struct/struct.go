package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// Method: function with receiver (the owner of function)
func (p product) getPriceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var p product
	p = product{
		name:     "Pen",
		price:    1.79,
		discount: 0.05,
	}
	fmt.Println(p, p.getPriceWithDiscount())

	p1 := product{"Notebook", 1989.90, 0.10}
	fmt.Println(p1, p1.getPriceWithDiscount())
}
