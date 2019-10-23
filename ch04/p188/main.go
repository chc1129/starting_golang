package main

import "fmt"

func main() {
  ch := make(chan string, 3)
  ch <- "Apple"
  fmt.Println(len(ch))
  ch <- "Banana"
  ch <- "Chelly"
  fmt.Println(len(ch))
}
