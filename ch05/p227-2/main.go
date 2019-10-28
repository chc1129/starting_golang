package main

import "fmt"

type Point struct{ X, Y int }

/* *Point型のレシーバー */
func (p *Point) Set(x, y int) {
  p.X = x
  p.Y = y
}

func main() {
  /* 変数p1はPoint型 */
  p1 := Point{}
  p1.Set(1, 2)
  fmt.Println(p1.X, p1.Y)
  /* 変数p2は*Point型 */
  p2 := &Point{}
  p2.Set(1, 2)
  fmt.Println(p2.X, p2.Y)
}
