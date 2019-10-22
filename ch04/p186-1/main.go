package main

import "fmt"

func main() {
  ch := make(chan int)
  /* チャンネルから受信した内容を出力 */
  fmt.Println(<-ch)
}
