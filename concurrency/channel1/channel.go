package main

import "fmt"

func main() {
	//Channel is a type in go
	//Channels are used to share data either between go routines or between goroutines and main
	//Channels are the sync point of the go routines

	//The first argumment is the data type which will transits in the channel, the second (1) means the buffer(it will be better explained early)
	ch := make(chan int, 1)

	ch <- 1 //sending a value to the channel
	<-ch    //getting a value form the channel, you don't need to pik it if you wont

	ch <- 2
	// ch <- 3 //if you send a value without remove the other one, it throws a dead lock error
	fmt.Println(<-ch)
}
