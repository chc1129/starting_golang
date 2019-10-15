package main

import (
  "fmt"
)

func anything(a interface{}) {
  fmt.Println(a)
}

func main() {

  /* interface{}にはすべての型を指定できる */
  anything(1)
  anything(3.14)
  anything(4 + 5i)
  anything('海')
  anything("日本語")
  anything([...]int{1, 2, 3, 4, 5})
}
