package main

import "fmt"

func main() {
  /* intの別名としてのMyInt型を定義 */
  type MyInt int

  var n1 MyInt = 5
  n2 := MyInt(7)
  fmt.Println(n1) // => "5"
  fmt.Println(n2) // => "7"
}
