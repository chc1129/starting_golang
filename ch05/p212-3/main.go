package main

import "fmt"

type Point struct {
  X, Y int
}

func main() {
  pt := Point{Y: 28}
  fmt.Println(pt.X)
  fmt.Println(pt.Y)
}
