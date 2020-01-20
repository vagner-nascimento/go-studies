package main

import "fmt"

type sportive interface {
	setTurboOn()
}

type ferrari struct {
	model           string
	turboOn         bool
	currentVelocity int
}

func (f *ferrari) setTurboOn() {
	f.turboOn = true
}

func setTurboOnInterface(c sportive) {
	c.setTurboOn()
}

func main() {
	car1 := ferrari{"F40", false, 0} // Directly from the class, you don't need to initialize as a memory reference with & symbol
	car1.setTurboOn()

	var car2 sportive = &ferrari{"F250", false, 0} // Through the interface, you must to initialize as a memory reference, otherwise, throws compile time error
	setTurboOnInterface(car2)                      // It just works when you initialize through the interface

	fmt.Println(car1, car2)
}
