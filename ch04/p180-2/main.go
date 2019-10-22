package main

import "fmt"

func main() {
  m := map[string]int{
    "Apple": 88,
    "Banana": 107,
    "Cherry": 46,
  }
  m["Grape"] = 66
  m["Lemon"] = 16
  m["Orange"] = 44
  m["Pinapple"] = 73

  for k, v := range m {
    fmt.Printf("%s => %d\n", k, v)
  }
}
