package main

import "fmt"

type T0 struct {
  Name1 string
}

type T1 struct {
  T0
  Name2 string
}

type T2 struct {
  T1
  Name3 string
}

func main() {
  t := T2{T1: T1{T0: T0{Name1: "X"}, Name2: "Y"}, Name3: "Z"}
  fmt.Println(t.Name1)
  fmt.Println(t.Name2)
  fmt.Println(t.Name3)
}
