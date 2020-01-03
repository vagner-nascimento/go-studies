package main

import (
	"fmt"
	"time"
)

func speak(name, text string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(time.Second) //Sleep for 1 second
		fmt.Printf("%s: %s (iteration %d)\n", name, text, i+1)
	}
}
func main() {
	// It runs sequentially
	// speak("Mary", "Why don't you talk to me?", 3)
	// speak("John", "I just can speak after you", 1)

	// A go routine just lives while the main function proccess is running

	// It runs concurrently, asynchronous, but prints nothing because the main function finishes before 1 second (time sleep into the speak function)
	// go speak("Mary", "Hey...", 500)
	// go speak("John", "Wow...", 500)

	// It prints the go routine proccessment while John keep talking, so, after 5 times the main finishes and Mary don't completes all your speaks
	go speak("Mary", "Got it!!!", 10)
	speak("John", "Congratulation!", 5)
}
