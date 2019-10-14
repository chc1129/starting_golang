package main

import (
  "fmt"
)

func main() {
  n := 4
  switch {
  case n > 0 && n < 3:
    fmt.Println("0 < n < 3")
  case n > 3 && n < 6:
    fmt.Println("3 < n < 6")
  }
}
