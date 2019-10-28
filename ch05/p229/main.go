package main

import (
  "./foo"
  "fmt"
)

func main() {
  t := &foo.T{}
  t.Method1()
  fmt.Println(  t.Field1 )
//  t.method2() // コンパイルエラー
//  t.fiedl2    // コンパイルエラー
}
