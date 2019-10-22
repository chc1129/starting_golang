package main

import "fmt"

func main() {
  /* 要素5,容量5のスライス*/
  s1 := make([]int, 5)
  fmt.Println(len(s1)) // => "5"
  fmt.Println(cap(s1)) // => "5"

  /* 要素5,容量10のスライス */
  s2 := make([]int, 5, 10)
  fmt.Println(len(s2)) // => "5"
  fmt.Println(cap(s2)) // => "10"
}
