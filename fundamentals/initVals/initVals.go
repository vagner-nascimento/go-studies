package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("Initital values\n%v: %v\n%v: %v\n%v: %v\n%v: %q\n%v: %v\n",
		reflect.TypeOf(a), a,
		reflect.TypeOf(b), b,
		reflect.TypeOf(c), c,
		reflect.TypeOf(d), d,
		reflect.TypeOf(e), e)
}
