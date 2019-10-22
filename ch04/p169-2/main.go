package main

import "fmt"

func main() {
  s := []string{"Apple", "Banana", "Cherry"}

  for i := 0; i < len(s); i++ {
    fmt.Printf("[%d] => %s\n", i, s[i])
    s = append(s, "Melon") // 要素の追加
  }
}
