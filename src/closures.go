package main
import (
  "fmt"
)

// この関数は intSeq() が返す関数を返す
func intSeq() func () int {
  i := 0
  return func () int {
    i++
    return i
  }
}

func main() {
  // intSeq() が返す関数を nextInt に代入
  // incrementしかしない関数がnextIntになる
  nextInt := intSeq()

  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  newInts := intSeq()
  fmt.Println(newInts())
}
