package main

import "fmt"

type Person struct {
  Id int
  Name string
  Area string
}

func main() {
  /* 変数pは*Person型 */
  p := new(Person)
  fmt.Println(p.Id)
  fmt.Println(p.Name)
  fmt.Println(p.Area)
}
