package main

import "fmt"

func main() {
  type T0 int
  type T1 int

  t0 := T0(5)
  i0 := int(t0)
  fmt.Println(t0)
  fmt.Println(i0)

  t1 := T1(8)
  i1 := int(t1)
  fmt.Println(t1)
  fmt.Println(i1)

  // t0 = t1   // コンパイルエラー
}
