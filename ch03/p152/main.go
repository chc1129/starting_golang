package main

import "fmt"

func main() {
  /* 無名関数によるgo */
  go func() {
    for {
      fmt.Println("sub loop")
    }
  }()
  for {
    fmt.Println("main loop")
  }
}
