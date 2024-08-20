package main

import "fmt"

func main() {
	func1()
	func2("Str 1", "Str 2")

	ret1, ret2 := func3(), func4("Str 3", "Str 4") // You can call multiple functions in the same line, separated by comma
	fmt.Printf("%s\n%s", ret1, ret2)

	func5()               // You can complete ignore a return of a function
	ret3, ret4 := func6() // But if you take, you should to take all (usually)
	fmt.Println(ret3, ret4)

	ret5, _, _ := func7() // You can ignore paramters as much as you need and take only those you'll really use by using underline '_' symbol
	fmt.Println(ret5)
}

func func1() {
	fmt.Println("func1")
}

func func2(p1 string, p2 string) {
	fmt.Printf("func2 - p1: %s p2: %s\n", p1, p2)
}

func func3() string {
	return "func3 - is mandatory return a string, otherwise, it will not compile"
}

func func4(p1, p2 string) string { // You can group parameters by type
	return fmt.Sprintf("func4 - p1: %s p2: %s", p1, p2)
}

// Returning multiple values
func func5() (string, string) {
	return "return 1", "return 2"
}

func func6() (string, float64) {
	return "return 1", 75954.91
}

func func7() (string, float64, bool) {
	return "return 1", 354.82, true
}
