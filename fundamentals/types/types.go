package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//Integers
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal's type is", reflect.TypeOf(32000))

	//unsigned (only positive int): unit8, unit16, unit32, unit64
	var b byte = 255                                  //byte is an alias to unit8
	fmt.Println("Byte is type of", reflect.TypeOf(b)) // = Byte is type of uint8

	//signed (positive and negative int): int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("The maximum value of int is", i1)
	fmt.Println("The type of i1 is", reflect.TypeOf(i1))

	var i2 rune = 'a' // rune = represents a mapping of Unicode table (int32)
	fmt.Println("The type of rune is", reflect.TypeOf(i2), "and its value is", i2)

	//real numbers? float32 and float64
	var x float32 = 49.99
	fmt.Println("The type of x is", reflect.TypeOf(x), "and its value is", x)
	fmt.Println("The type literal 49.99 is", reflect.TypeOf(49.99)) //By default, a real number is float64

	//boolean
	bo := true
	fmt.Println("The type of bo is", reflect.TypeOf(bo), "and its value is", bo)
	fmt.Println("!bo =", !bo)

	//string
	s1 := "Strings are ONLY separated by double quotation"
	fmt.Println(s1 + "!")
	fmt.Println("The length of s1 is", len(s1))

	//Multiple lines string
	s2 := `Strings 
	are 
	ONLY 
	separated 
	by double quotation`
	fmt.Println(s2)
	fmt.Println("The length of s2 is", len(s2))

	//char (DOESN'T EXISTS)
	//var ch char = 'b'
	char := 'a'
	fmt.Println("The type of char is", reflect.TypeOf(char), "and its value is", char) //This a rune as well
	fmt.Println("To get the character value you must to covert to a string:", string(char))
}
