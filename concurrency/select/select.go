package main

import (
	"fmt"
	"time"

	"github.com/vagner-nascimento/gostudieshtml"
)

// getFaster return the title of the first site that respond
func getFaster(url1, url2, url3 string) string {
	ch1 := gostudieshtml.GetTitles(url1)
	ch2 := gostudieshtml.GetTitles(url2)
	ch3 := gostudieshtml.GetTitles(url3)

	// Concurrency's specifc control structure
	select {
	case title1 := <-ch1: // Get the value of channel 1 into "title1" variable
		return title1
	case title2 := <-ch2:
		return title2
	case title3 := <-ch3:
		return title3
	case <-time.After(1000 * time.Millisecond): //If none respond in this timeout, it will return this message
		return "All lost!"
		// default: // In this case, the default will be always called
		// 	return "None response!"
	}
}

func main() {
	winner := getFaster(
		"https://www.youtube.com",
		"https://www.stackoverflow.com",
		"https://www.microsoft.com",
	)

	fmt.Println("Winner: ", winner)
}
