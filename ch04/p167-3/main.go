package main

import "fmt"

func main() {
  a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  s1 := a[2:4]
  fmt.Println(len(s1))
  fmt.Println(cap(s1))
}
