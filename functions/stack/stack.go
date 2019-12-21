package main

import "runtime/debug"

func main() {
	func1()
}

func func1() {
	func2()
}

func func2() {
	func3()
}

func func3() {
	debug.PrintStack()
}
