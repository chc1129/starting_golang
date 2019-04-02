package main

import (
	"fmt"
)

func main() {
	// fmt.Println
	fmt.Println("fmt.Println")
	// fmt.Printf
	fmt.Printf("fmt.Printf num=%d\n", 5)
	fmt.Printf("10進数=%d, 2進数=%b, 8進数=%o, 16進数=%x", 17, 17, 17, 17)
	fmt.Printf("数値=%v, 文字列=%v, 配列=%v\n", 5, "Golang", [...]int{1,2,3})
	fmt.Printf("数値=%#v, 文字列=%#v, 配列=%#v\n", 5, "Golang", [...]int{1,2,3})
	fmt.Printf("数値=%T, 文字列=%T, 配列=%T\n", 5, "Golang", [...]int{1,2,3})

	//print
	print("print!")
	println("println!")
}
