package main

import "fmt"

type Point struct {
  X, Y int
}

/* Point型のポインタを受け取るように変更 */
func swap(p *Point) {
  x, y := p.Y, p.X
  p.X = x
  p.Y = y
}

func main() {
  p := Point{X: 1, Y: 2}
  fmt.Println(p.X, p.Y)
  swap(&p) /* Point型のポインタを渡す */
  fmt.Println(p.X, p.Y)
}
