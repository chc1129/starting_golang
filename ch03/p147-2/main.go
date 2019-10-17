package main

import "fmt"

func runDefer() {
  /* deferに登録された式は関数の終了時に評価される */
  defer fmt.Println("defer")
  fmt.Println("done")
}

func main() {
  runDefer()
}
