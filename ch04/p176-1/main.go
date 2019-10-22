package main

import "fmt"

func main() {
  m := make(map[string]string)

  m["Yamada"] = "Taro"
  m["Sato"] = "Hanako"
  m["Yamada"] = "Jiro"

  fmt.Println(m)
}
