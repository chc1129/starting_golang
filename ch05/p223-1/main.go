package main

import "fmt"

type Point struct{ X, Y int }

/* *Point型のメソッドRender */
func (p *Point) Render() {
  fmt.Printf("<%d,%d>\n", p.X, p.Y)
}

func main() {
  p := &Point{X: 5, Y :12}
  /* メソッドRenderの呼び出し */
  p.Render()
}
