package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1 //blocking operation
	fmt.Println("After sent to the channel")
}

func main() {
	c := make(chan int)
	go routine(c)

	fmt.Println(<-c) //blocking operation
	fmt.Println("Value was read")
	fmt.Println(<-c) //It causes deadlock because there's no more possibility to some function sent the value to the channel
	fmt.Println("It will never be printed")
}
