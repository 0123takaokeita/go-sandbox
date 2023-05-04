package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

    // 省略する場合は、順番に指定する。
    fmt.Println(person{"Bob", 20})

    // k,v の形でも指定できる。
    fmt.Println(person{name: "Alice", age: 30})

    // 省略した場合は、初期値が入る。
    fmt.Println(person{name: "Fred"})

    // & でメモリアドレスを指定する。
    fmt.Println(&person{name: "Ann", age: 40})

    // newPerson で初期値を設定する。
    fmt.Println(newPerson("Jon"))

    // 構造体のフィールドにアクセスする。
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // ポインタを使用する場合も. を使用する。
    sp := &s
    fmt.Println(sp.age)

    // 再代入も可能
    sp.age = 51
    fmt.Println(sp.age)
}
