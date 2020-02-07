package main

import "fmt"

func main() {
	// declaring an array with inference, it uses the number of given elements to define the size
	numbers := [...]int{1, 2, 3, 4, 5}

	/*
		in this way, for takes the index and the value at this index
		looks like .every() of JS (ECMA6)
	*/
	for index, value := range numbers {
		fmt.Printf("index %d, value %d\n", index, value)
	}

	// ignoring index, looks like forEach
	for _, value := range numbers {
		fmt.Printf("Getting only value %d\n",value)
	}

	// if you need only the first value (key), you can omit the second one
	for key := range numbers{
		fmt.Printf("Getting only key %d\n", key)
	}
}
