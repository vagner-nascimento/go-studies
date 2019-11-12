package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		Go has one repetition structure: for
	*/

	fmt.Println("it can be used as a while looping")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("an usual for looping")
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}

	fmt.Println("infinite looping")
	for {
		fmt.Println("Forever...")
		time.Sleep(time.Second)
	}
}
