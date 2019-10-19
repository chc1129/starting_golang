package main

import "fmt"

func main() {
  a := make([]float64, 3)
  fmt.Println(a)    // => "[0, 0, 0]"
  a[0] = 3.14
  fmt.Println(a)    // => "[3,14, 0, 0]"
  a[1] = 6.28
  fmt.Println(a)    // => "[3.14, 6.28, 0]"
  fmt.Println(a[0]) // => "3.14"
//  fmt.Println(a[4]) // => ランタイムパニック
}
