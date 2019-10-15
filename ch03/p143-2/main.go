package main

func main() {

  var x interface{} = 7

  switch v := x.(type) {
  case int:
    fmt.Println(v * v) // => "49"
  }
}
