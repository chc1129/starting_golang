package main

import "fmt"

func main() {
  s1 := []int{1, 2, 3, 4, 5}
  s2 := []int{10, 11, 12, 13, 14, 15, 16}
  n := copy(s1, s2)
  fmt.Println(s1)
  fmt.Println(n)
}
