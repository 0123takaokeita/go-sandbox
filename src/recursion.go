package main

import (
	"fmt"
)

// 再帰関数
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

  // 再帰関数を変数に代入
	var fib func(n int) int

  // フィボナッチ数列
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
