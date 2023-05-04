package main

import "fmt"

type base struct {
    num int
}

// baseのメソッド
func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

// baseを含む構造体
type container struct {
    base
    str string
}

func main() {

    // containerの初期化
    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }

    // containerの表示
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    // baseにアクセスしてもnumが表示される。
    fmt.Println("also num:", co.base.num)

    // baseのメソッドも使用できる。
    fmt.Println("describe:", co.describe())

    type describer interface {
        describe() string
    }

    // coがbasewを実装しているんので、descriverも実装していると言える。
    var d describer = co
    fmt.Println("describer:", d.describe())
}
