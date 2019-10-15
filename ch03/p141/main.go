package main

func main() {

  var x interface{} = 3
  i := x.(int) // 変数iはint型で値は3
  f := x.(float64) // エラーが発生

}
