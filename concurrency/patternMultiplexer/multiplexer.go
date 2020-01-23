package main

import (
	"fmt"

	"github.com/vagner-nascimento/gostudieshtml"
)

func forward(from <-chan string, to chan string) {
	for { //  This infinite for will block when don't have more messages
		to <- <-from //Send a message from a channel to the other channel
	}
}

// It mix the message from two channels into other channel
func multiplexer(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)

	go forward(ch1, ch)
	go forward(ch2, ch)

	return ch
}

func main() {
	ch := multiplexer(
		gostudieshtml.GetTitles("https://www.cod3r.com.br", "https://www.google.com"),
		gostudieshtml.GetTitles("https://stackoverflow.com/", "https://www.youtube.com"),
	)

	fmt.Println("Firsts:", <-ch, "|", <-ch)
	fmt.Println("Seconds:", <-ch, "|", <-ch)
}
