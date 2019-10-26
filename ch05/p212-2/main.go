package main

import "fmt"

type Point struct {
  X, Y int
}

func main() {
  pt := Point{X: 1, Y: 2}
  fmt.Println(pt.X)
  fmt.Println(pt.Y)
}
