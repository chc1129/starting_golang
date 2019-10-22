package main

import "fmt"

func main() {
  m := map[int]string{1: "A", 2: "B", 3: "C"}

  s := m[1]
  fmt.Println(s)
  s = m[9]
  fmt.Println(s)
}
