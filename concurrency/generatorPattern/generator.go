package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Search at Youtube: Google I/O 2012 - Go Concurrency Patterns
// Genetor Pattern encapsules the use of go routines, returning a read only channel where the result will be sended while they are ready

// <-chan = A read only channel
func getTitles(urls ...string) <-chan string {
	ch := make(chan string)

	for _, url := range urls {

		go func(urlParam string) { // Calling an anonymous function in a go routine
			resp, _ := http.Get(urlParam)
			html, _ := ioutil.ReadAll(resp.Body)
			regX, _ := regexp.Compile("<title.*?>(.*?)<\\/title>")
			matches := regX.FindStringSubmatch(string(html))

			ch <- matches[1]
		}(url) // Passing each URL in URLs array
	}

	return ch
}

func main() {
	t1 := getTitles("https://www.cod3r.com.br", "https://www.google.com")
	t2 := getTitles("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("Firsts:", <-t1, "|", <-t2)
	fmt.Println("Seconds:", <-t1, "|", <-t2)
}
