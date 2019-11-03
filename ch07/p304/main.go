package main

import "fmt"

type User struct {
  Id int
  Email string
}

func main() {
  fmt.Printf("%v\n", [3]int{1,2,3})
  fmt.Printf("%v\n", []string{"A","B","C"})
  fmt.Printf("%v\n", map[int]float64{1: 1.0, 2: 4.0})

  u := &User{Id: 123, Email: "mail@example.com"}
  fmt.Printf("%v\n", u)
  // => "&{123 mail@example.com}"

  // %+vは構造体のフィールドも出力する
  fmt.Printf("%+v\n", u)
  // => "&{Id: 123 Email:mail@example.com}"

  // %#vは構造体のフィールドと型についても出力する
  fmt.Printf("%#v\n", u)
  // => "&main.User{Id:123, Email:"mail@example.com}"
}
