package main

import "fmt"

func main() {
  var (
    ch0 chan int
    ch1 <-chan int
    ch2 chan<- int
  )
  ch1 = ch0
  ch2 = ch0
  ch0 = ch1
  ch0 = ch2
  ch1 = ch2
  ch2 = ch1
}
