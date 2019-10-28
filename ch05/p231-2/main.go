package main

import "fmt"

type User struct {
  Id int
  Name string
}

func main() {
  /* キーが構造体型のマップ */
  m1 := map[User]string{
    {Id: 1, Name: "Taro"}: "Tokyo",
    {Id: 2, Name: "Jiro"}: "Osaka",
  }

  /* 値が構造体型のマップ */
  m2 := map[int]User{
    1: {Id: 1, Name: "Taro"},
    2: {Id: 2, Name: "Taro"},
  }

  fmt.Println(m1)
  fmt.Println(m2)
}
