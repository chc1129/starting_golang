package main

import (
  "./foo"
  "fmt"
)

func main() {
  t := foo.NewI()
  fmt.Println( t.Method1() ) // OK
  // t.method2() // コンパイルエラー
}
