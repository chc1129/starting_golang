package main

import "fmt"

func main() {
  m := map[int]string{1: "A", 2: "B", 3: "C"}
  fmt.Println(m)
  fmt.Println(len(m))
  m[4] = "D"
  m[5] = "E"
  fmt.Println(m)
  fmt.Println(len(m))
}
