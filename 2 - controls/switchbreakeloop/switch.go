package main

import "fmt"

func simpleBreak() {
	v := 10

	switch {
	case v < 10:
		fmt.Println(v)
	case v >= 10:
		if v == 10 {
			fmt.Println("v is 10")
			break
		}
		fmt.Println(v)
	}
}

func loopBreak() {
myLoop: // myLoop is an label to the for that can be used only inside it (and only applied to the loops)
	for i := 0; i < 100; i = i + 1 {
		switch {
		case i < 5:
			fmt.Println("is less than 5")
		case i >= 5 && i < 11:
			if i == 5 {
				fmt.Println("breaking only switch")
				break
			}
			fmt.Println("i is greater than 5")
		case i == 11:
			fmt.Println("Breaking the myLoop")
			break myLoop

		}
		fmt.Printf("Out of switch %d\n", i)
	}

	// fmt.Println(myLoop) // it gives compilation error
}

func main() {
	// break is used only to early finishes the switch
	simpleBreak()
	fmt.Println()
	// With break you can also stops a switch and a looping outside of it
	loopBreak()
}
