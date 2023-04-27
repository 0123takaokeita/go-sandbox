package main

import (
	"fmt"
)

// 可変長引数
func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(2, 4, 5)

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	// sum(nums) だとエラーになる
	// nums... でスライスを展開して渡す
	sum(nums...)
}
