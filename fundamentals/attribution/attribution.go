package main

import "fmt"

func main() {
	i := 3 // Type inference
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 3
	fmt.Println(i)

	x, y := 1, 2 // Double attribution
	fmt.Println(x, y)

	x, y = y, x // Double attribution to invert values
	fmt.Println(x, y)
}
