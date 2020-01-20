package main

import "fmt"

type sportive interface {
	setTurboOn()
}

type luxury interface {
	autoPark()
}

type sportiveLuxury interface {
	sportive
	luxury
	run() // You can also add other methods here
}

type car struct{}

func (c car) setTurboOn() {
	fmt.Println("Turbo on")
}

func (c car) autoPark() {
	fmt.Println("Parking")
}

func (c car) run() {
	fmt.Println("Running")
}

func main() {
	var car1 sportiveLuxury = car{}
	car1.setTurboOn()
	car1.autoPark()
	car1.run()
}
