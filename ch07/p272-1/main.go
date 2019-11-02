package main

import (
  "fmt"
  "os"
)

func main() {
  /* os.Exitによって破棄されるdefer */
  defer func() {
    fmt.Println("defer")
  }()

  os.Exit(0)
}
