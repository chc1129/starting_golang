package main

import "fmt"

func main() {
  m := map[int]string{1: "A", 2: "B", 3: "C"}
  fmt.Println(m)
  delete(m, 2)
  fmt.Println(m)
}
