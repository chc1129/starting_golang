package main

import "fmt"

type Point struct {
  X, Y int
}

func swap(p Point) {
  /* フィールドX,Yの値を入れ替える */
  x, y := p.Y, p.X
  p.X = x
  p.Y = y
}

func main() {
  p := Point{X: 1, Y: 2}
  fmt.Println(p.X, p.Y)
  swap(p) /* 構造体は値渡しされる */
  fmt.Println(p.X, p.Y)
}
