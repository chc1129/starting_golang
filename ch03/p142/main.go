package main

func main() {

  var x interface{} = 3.14

  i, isInt := x.(int)
  f, isFloat64 := x.(float64)
  s, isString := x.(string)
}
