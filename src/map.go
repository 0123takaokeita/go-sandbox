package main

import "fmt"

func main() {

   // Ruby でいう Hash の定義
   // key value の型を予め指定している。
    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    m["t1"] = 11
    fmt.Println(m)

    v1 := m["k1"]
    fmt.Println("v1:", v1)

    v3 := m["k3"]
    fmt.Println("v3:", v3)

    fmt.Println("len:", len(m))

    // 削除はdeleteを使うのか
    delete(m, "k2")
    fmt.Println("map:", m)

    // _ は Ruby と同じ考え方だ。
    _, prs := m["k2"]
    // key が見つからなければ false が返る。
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}
