package main

import "fmt"

func main() {
  m := map[int][]int{
    1: []int{1},
    2: []int{1, 2},
    3: []int{1, 2, 3},
  }
  fmt.Println(m)
}
