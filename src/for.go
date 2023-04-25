package main

import (
  "fmt"
)

func main() {
  i := 3
  for i <= 10 {
    fmt.Println(i)
    i = i + 1
  }

  separator()
  for j:= 0; j <= 10; j++ {
    fmt.Println(j)
  }
}

func separator() {
  fmt.Println("============")
}
