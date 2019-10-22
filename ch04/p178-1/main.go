package main

import "fmt"

func main() {
  m := map[int][]int{
    /* 内側のスライスリテラルの[]intは省略可能 */
    1: {1},
    2: {1, 2},
    3: {1, 2, 3},
  }
  fmt.Println(m)
}
