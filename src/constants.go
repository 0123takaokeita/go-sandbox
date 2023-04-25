package main

import (
	"fmt"
	"math"
)

func main() {
	// 円周率
	fmt.Println(math.Pi)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
