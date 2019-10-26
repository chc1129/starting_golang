package main

import "fmt"

func main() {
  /* type によるさまざまなエイリアス */
  type (
    IntPair     [2]int
    Strings     []string
    AreaMap     map[string][2]float64
    IntsChannel chan []int
  )

  pair := IntPair{1, 2}
  strs := Strings{"Apple", "Banana", "Cherry"}
  amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
  ich := make(IntsChannel)

  fmt.Println(pair)
  fmt.Println(strs)
  fmt.Println(amap)
  fmt.Println(ich)
}
