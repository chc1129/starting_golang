package main

import "fmt"

//type Point struct {
//  X int
//  Y int
//}

type Point struct {
  X, Y int
}

func main() {
  var pt Point
  fmt.Println(pt.X)
  fmt.Println(pt.Y)

  /* 構造体のフィールドへの代入 */
  pt.X = 10
  pt.Y = 8

  fmt.Println(pt.X)
  fmt.Println(pt.Y)

}
