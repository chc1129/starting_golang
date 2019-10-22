package main

import "fmt"

func pow(a []int) {
  /* スライスの各要素を2乗する */
  for i, v := range a {
    a[i] = v * v
  }
  return
}

func main() {
  /* 3要素のスライス */
  a := []int{1, 2, 3}
  pow(a)
  fmt.Println(a) // => "[1, 4, 9]"
}
