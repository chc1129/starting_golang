package main

import "fmt"

func main() {
  a := [3]string{"Apple", "Banana", "Cherry"}
  p := &a
  fmt.Println(a[1])
  fmt.Println(p[1])
  p[2] = "Grape"
  fmt.Println(a[2])
  fmt.Println(p[2])
}
