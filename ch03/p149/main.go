package main

import (
  "fmt"
)

func main() {

  /* panic時でもdeferは実行される */
  defer fmt.Println("on defer")

  panic("runtime error!")
  fmt.Println("Hello, World!")
}
