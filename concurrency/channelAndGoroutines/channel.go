package main

import (
	"fmt"
	"time"
)

func twoThreeFourTimes(value int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * value //In this moment, the go routine stops and wait for someone to read this value from the channel

	time.Sleep(time.Second) //This line just executes after someone reads the last message sent
	ch <- 3 * value

	time.Sleep(time.Second * 3)
	ch <- 4 * value
}

func main() {
	myChan := make(chan int)
	go twoThreeFourTimes(2, myChan)

	/*
		The call for values stops the execution and wait for the return sent for the function.
		If you do like bellow, it will wait for the 2 values before continues the main
	*/
	a, b := <-myChan, <-myChan
	fmt.Println(a, b) //Is how Go synchronize the go routines

	fmt.Println(<-myChan)
	// fmt.Println(<-myChan) //If you try to read a value from an empty channel, it will causes a dead lock error saying that all go routines are stopped
}
