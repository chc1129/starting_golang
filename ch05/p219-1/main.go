package main

import "fmt"

func showStruct(s struct{ X, Y int}) {
  fmt.Println(s)
}

func main() {
  s := struct{ X, Y int}{X: 1, Y: 2}
  showStruct(s)
}
