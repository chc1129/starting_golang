package main

import "fmt"

func main() {
  a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  s1 := a[2:4] // s1 == [3, 4]
  fmt.Println(len(s1)) // == 2
  fmt.Println(cap(s1)) // == 8

  s2 := a[2:4:4] // s2 == [3, 4]
  fmt.Println(len(s2)) // == 2
  fmt.Println(cap(s2)) // == 2

  s3 := a[2:4:6] // s3 == [3, 4]
  fmt.Println(len(s3)) // == 2
  fmt.Println(cap(s3)) // == 4
}
