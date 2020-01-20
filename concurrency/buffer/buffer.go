package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	ch <- 1
	fmt.Println("Executed 1!")
	ch <- 2
	fmt.Println("Executed 2!")
	ch <- 3
	fmt.Println("Executed 3!")
	ch <- 4
	fmt.Println("Executed 4!")
	ch <- 5
	fmt.Println("Executed 5!")
	ch <- 6
	fmt.Println("Executed 6!")
}

func main() {
	//When you use a buffer, the routine can send without need to be read until the buffer is full (in this case, 3 integers)
	c := make(chan int, 3)
	go routine(c)

	time.Sleep(time.Second)
	fmt.Println(<-c) //once it is read, the routine function put another value into the channel

	time.Sleep(time.Second)
	fmt.Println(<-c) //and here again

	time.Sleep(time.Second)
	fmt.Println(<-c) //and again

	//If the main terminates and you don't read all buffer, that's not a problem, none error will be throwed
}
