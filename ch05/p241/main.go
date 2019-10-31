package main

import "fmt"

type T struct{ Id int }

func (t *T) GetId() int { return t.Id }

func ShowId(id interface { GetId() int } ) {
  /* 変数idは「GetId() int」 というメソッドを持つ型 */
  fmt.Println(id.GetId())
}

func main() {
  t := &T{Id: 19}
  ShowId(t) // => "19"
}
