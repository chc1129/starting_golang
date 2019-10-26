package main

import "fmt"

type T struct {
  int
  float64
  string
}

func main() {
  t := T{1, 3.14, "文字列"}
  fmt.Println(t.int)
  fmt.Println(t.float64)
  fmt.Println(t.string)
}
