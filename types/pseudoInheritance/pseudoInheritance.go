package main

import "fmt"

type car struct {
	name         string
	currentSpeed int
}

// THIS IS CALLED COMPOSITION
type ferrari struct {
	car     //Unamed fields, in this way it take all attributes of car and put inside ferrari. It simulates a inheritance
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.currentSpeed = 0
	f.turboOn = false

	fmt.Printf("The %s ferrari's is running at %v Km/h and turbo is on? %v\n", f.name, f.currentSpeed, f.turboOn)

	// f2 := ferrari{name: "F250", currentSpeed: 320, turboOn: true} // That is impossible
	c := car{name: "F250", currentSpeed: 320} // You should to set values separated or do like this, create a car and pass to the ferrari
	f1 := ferrari{c, true}
	fmt.Printf("The %s ferrari's is running at %v Km/h and turbo is on? %v\n", f1.name, f1.currentSpeed, f1.turboOn)

	fmt.Println(f, f1)
	/*
	 The objects printed has this structure: {{F40 0} false} note that the car's value are inside {}
	 	which means that is another data struct
	*/
}
