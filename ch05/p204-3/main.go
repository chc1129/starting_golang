package main

import "fmt"

func main() {
  p := &[3]int{1, 2, 3}

  fmt.Println(len(p))
  fmt.Println(cap(p))
  fmt.Println(p[0:2])
}
