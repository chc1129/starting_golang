package main

import "fmt"

type Person struct {
  ID uint
  name string
  部署 string
}

func main() {
  p := Person{ID: 17, name: "山田太郎", 部署: "営業部"}
  fmt.Println(p.ID)
  fmt.Println(p.name)
  fmt.Println(p.部署)
}
