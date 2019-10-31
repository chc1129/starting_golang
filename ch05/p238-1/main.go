package main

import "fmt"

type T struct {
  Id int
  Name string
}

func main() {
  t := &T{Id: 10, Name: "Taro"}
  fmt.Println(t) // => "&{10 Taro}"
}
