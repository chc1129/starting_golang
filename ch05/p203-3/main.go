package main

import "fmt"

func pow(p *[3]int) {
  i := 0
  for i < 3 {
    /* 各要素を累乗する*/
    (*p)[i] = (*p)[i] * (*p)[i]
    i++
  }
}

func main() {
  /* 変数pは*[3]int型 */
  p := &[3]int{1, 2, 3}
  pow(p)
  fmt.Println(p)
}
