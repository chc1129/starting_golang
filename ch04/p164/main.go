package main

import "fmt"

func main() {
  /* 要素数3のスライス */
  s := []int{1, 2, 3}
  /* 4をスライスの末尾に追加 */
  s = append(s, 4) // s == [1, 2, 3, 4]
  fmt.Println(s)
  /* 5, 6, 7をスライスの末尾に追加 */
  s = append(s, 5, 6, 7) // s == [1, 2, 3, 4, 5, 6, 7]
  fmt.Println(s)

  s0 := []int{1, 2, 3}
  s1 := []int{4, 5, 6}
  s2 := append(s0, s1...)
  fmt.Println(s2)
}
