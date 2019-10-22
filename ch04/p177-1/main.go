package main

import "fmt"

func main() {
  m := map[int]string{
    1: "Taro",
    2: "Hanako",
    3: "Jiro", // ←カンマが必要
  }
  fmt.Println(m)
}
