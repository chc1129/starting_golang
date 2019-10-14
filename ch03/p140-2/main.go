package main

import (
  "fmt"
)

func main() {
  /* caseに定数と式が混在するswitch文はコンパイルエラー */
  switch x := 1; x {
  case 1, 2, 3:
    fmt.Println(x)
  case x > 3:
    fmt.Println(x * x)
  }
}
