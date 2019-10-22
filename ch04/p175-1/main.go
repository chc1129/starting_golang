package main

import "fmt"

func main() {
  /* A */
  a := [3]int{1, 2, 3}
  s := a[:]
  fmt.Println(s)
  fmt.Println(len(s))
  fmt.Println(cap(s))
  /* B */
  s = append(s, 4)
  fmt.Println(s)
  fmt.Println(a)
  fmt.Println(len(s))
  fmt.Println(cap(s))
  a[0] = 9
  fmt.Println(s)
  fmt.Println(a)
}
