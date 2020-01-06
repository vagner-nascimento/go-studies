package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%2 == 0 {
			return false
		}
	}
	return true
}

func primes(qtd int, ch chan int) {
	start := 2
	for i := 0; i < qtd; i++ {
		for number := start; ; number++ { //The stop clause is empty, it means that it will never stop until call break (or return?)
			if isPrime(number) {
				ch <- number
				start = number + 1
				time.Sleep(time.Millisecond * 180)
				break
			}
		}
	}
	close(ch) //Closes the channel, in this way, those who are listening (waiting values) from this channel, stops and the code continues whitout dead lock errors
}

func main() {
	ch := make(chan int, 30) // Buffer = 30 (30 values can be putted without nobody read it)
	// go primes(cap(ch), ch)   // Cap brings the buffer size of the channel
	go primes(60, ch)

	for prime := range ch {
		fmt.Printf("%d\n", prime)
	}

	fmt.Println("The end")
}
