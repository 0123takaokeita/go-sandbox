package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// * をつけるとpointer type を使用できる。
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	// この段階では、i は 1
	fmt.Println("initial:", i)

  // zeroval は、i の値を変更しない。
  // 関数の中で変数に0を入れているのでcopyされているだけ。
	zeroval(i)
	fmt.Println("zeroval:", i)

  // zeroptr は、i の値を変更する。
  // 関数の中で変数のメモリアドレスを指定しているので、
  // そのメモリアドレスに0を入れると0になる。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

  // &i は、i のメモリアドレスを示す。
	fmt.Println("pointer:", &i)
}
