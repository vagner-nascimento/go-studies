package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		before ";" the variable i is created and initialized with a value
		this variable is visible only inside if or else or "else if" of the same sentence
	*/
	if i := aletoryNumber(); i > 5 { // it works also to switch
		fmt.Println("Win, i:", i)
	} else if i > 3 {
		fmt.Println("Almost, i:", i)
	} else {
		fmt.Println("Lose, i:", i)
	}

	//fmt.Println(i) // the variable "i" doesn't exists out of if sentence
}

func aletoryNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}
