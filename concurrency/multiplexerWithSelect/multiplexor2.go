package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			ch <- fmt.Sprintf("%s speaking: %d", person, i+1)
		}
	}()

	return ch
}

func merge(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for { // This infinite for will run until the go routine is alive
			select {
			case speaking := <-ch1:
				ch <- speaking
			case speaking := <-ch2:
				ch <- speaking
			}
		}
	}()

	return ch
}

func main() {
	ch := merge(speak("John"), speak("Mary"))

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch, <-ch)
	}
}
