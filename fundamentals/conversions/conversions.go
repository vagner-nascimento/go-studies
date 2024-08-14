package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) //You can't do it directly in Go because its a strong typed language

	score := 6.9
	finalScore := int(score)
	fmt.Println(finalScore) //It doesn't round values, it just removes the decimal values

	/*
		BE AWARE: It doesn't coverts the 97 to a string "97",
		its the number referenced at ASCII table, which is this case is "a"
	*/
	fmt.Println(string(97))
	fmt.Println("The right way to covert int to string is using strconv.Itoa(int val)",
		strconv.Itoa(97))
	//... it is valid with vars too
	var oo int = 97
	str := string(oo)
	fmt.Println(str) //Prints "a"

	num, _ := strconv.Atoi("123") //This function returns 2 values, the undercore "_" means that we'll ignore this result
	fmt.Println("Coverting from string to int using strconv.Atoi(string num)", num)

	//Converting string to boolean (only "true" and "1" gives true, all other string returns false)
	bol, _ := strconv.ParseBool("true")
	if bol {
		fmt.Println("Converted \"true\" to", bol)
	}

	bol, _ = strconv.ParseBool("1")
	if bol {
		fmt.Println("Converted \"1\" to", bol)
	}

	bol, _ = strconv.ParseBool("9999")
	if !bol {
		fmt.Println("Converted \"9999\" to", bol)
	}

	bol, _ = strconv.ParseBool("-44")
	if !bol {
		fmt.Println("Converted \"-44\" to", bol)
	}

	bol, _ = strconv.ParseBool("tony")
	if !bol {
		fmt.Println("Converted \"tony\" to", bol)
	}

	bol, _ = strconv.ParseBool("@")
	if !bol {
		fmt.Println("Converted \"@\" to", bol)
	}
}
