package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(typeOf(2.3))
	fmt.Println(typeOf(1))
	fmt.Println(typeOf("str"))
	fmt.Println(typeOf(func() {}))
	fmt.Println(typeOf(time.Now()))
}

func typeOf(v interface{}) string { // interface could be anything, an int, string, bool, whatever
	switch v.(type) { // this is the sintax to take type of an interface (its WEIRD)
	case int:
		return "integer"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "unknown"
	}
}
