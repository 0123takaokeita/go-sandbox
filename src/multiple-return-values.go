package main

import (
	"fmt"
)

// Go では複数の戻り値を返すことができる
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a, b)

  // _ で不要な値を捨てることができる
	_, c := vals()
	fmt.Println(c)
}
