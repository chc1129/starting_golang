package main

import "fmt"

func main() {
  /* 値がスライスのマップ */
  ms := map[int][]string{
    1: {"A", "B", "C"},
  }
  /* 値がマップのマップ */
  mm := map[int]map[int]string{
    1: {1: "Apple", 2: "Banana", 3: "Cherry"},
  }

  fmt.Println(ms)
  fmt.Println(mm)
}
