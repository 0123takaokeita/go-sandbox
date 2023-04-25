package main

import (
	"fmt"
)

func main() {
	i := 5
	// break がなくて良い
	for i >= 0 {
		switch i {
		case 1:
			fmt.Println("i is 1")
		case 2:
			fmt.Println("i is 2")
		case 3:
			fmt.Println("i is 3")
		case 4:
			fmt.Println("i is 4")
		case 5:
			fmt.Println("i is 5")
		default:
			fmt.Println("i is not 1, 2, 3, 4 or 5")
		}
		// increment decrement 使えるのか〜
		i--
	}
}

func separator() {
	fmt.Println("============")
}
