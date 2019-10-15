package main

func main() {

  switch v := x.(type) {
  case bool:
    fmt.Println("bool:", v)
  case int:
    fmt.Println(v * v)
  case string:
    fmt.Println(v)
  default:
    fmt.Printf("%#v\n", v)
}
