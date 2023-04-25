package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // index が欲しい場合は range の前に変数を置く
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // Ruby の each みたいに key, value が取れる
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // key だけ取りたい場合は _ で省略できる
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // 文字列の場合は index, rune が取れる
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
