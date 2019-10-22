package main

import "fmt"

func main() {
  a := [5]int{1, 2, 3, 4, 5}
  s := a[0:2]
  fmt.Println(s)      // s == [1, 2]
  a[1] = 0
  fmt.Println(len(s)) // == 2
  fmt.Println(cap(s)) // == 5
  fmt.Println(s)      // s == [1, 0]
}
