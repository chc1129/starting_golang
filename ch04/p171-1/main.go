package main

import "fmt"

/* 任意個数のint型の値の合計値を返す関数 */
func sum(s ...int) int {
  n := 0
  for _, v := range s {
    n += v
  }
  fmt.Println(n)
  return n
}

func main() {
  sum(1, 2, 3)        // == 6
  sum(1, 2, 3, 4, 5)  // == 15
  sum()               // == 0
}
