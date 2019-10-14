package main

import (
  "fmt"
)

func main() {

  fruits := [3]string{"Apple", "Banana", "Cherry"}
  /* rangeを伴うfor */
  for i, s := range fruits {
    // i:文字列配列のインデックス(0, 1, 2, ...)
    // s:インデックスに対応した文字列値
    fmt.Printf("fruits[%d]=%s\n", i, s)
  }
}
