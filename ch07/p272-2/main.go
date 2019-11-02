package main

import (
  "os"
  "log"
)

func main() {
  _, err := os.Open("foo")
  if err != nil {
    log.Fatal(err)
  }
}
