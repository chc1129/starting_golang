package main

import "fmt"

func main() {
  /* (A) */
  s := make([]int, 0, 0)
  fmt.Printf("(A) len=%d, cap=%d\n", len(s), cap(s))
  /* (B) */
  s = append(s, 1)
  fmt.Printf("(B) len=%d, cap=%d\n", len(s), cap(s))
  /* (C) */
  s = append(s, []int{2, 3, 4}...)
  fmt.Printf("(C) len=%d, cap=%d\n", len(s), cap(s))
  /* (D) */
  s = append(s, 5)
  fmt.Printf("(D) len=%d, cap=%d\n", len(s), cap(s))
  /* (E) */
  s = append(s, 6, 7, 8, 9)
  fmt.Printf("(E) len=%d, cap=%d\n", len(s), cap(s))
  s = append(s, 5)
  fmt.Printf("(D) len=%d, cap=%d\n", len(s), cap(s))
  /* (E) */
  s = append(s, 6, 7, 8, 9)
  fmt.Printf("(E) len=%d, cap=%d\n", len(s), cap(s))
}
