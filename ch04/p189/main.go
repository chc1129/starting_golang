package main

import "fmt"

func main() {
  ch := make(chan string)
  fmt.Println(cap(ch))

  ch = make(chan string, 3)
  fmt.Println(cap(ch))
}
