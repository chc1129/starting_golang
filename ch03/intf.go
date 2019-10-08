package main

import (
	"fmt"
)

func main() {

	var x interface{}
	fmt.Printf("%#v\n", x)

var a interface{}
	a = 1
	fmt.Printf("%#v\n", a)
	a = 3.14
	fmt.Printf("%#v\n", a)
	a = '山'
	fmt.Printf("%#v\n", a)
	a = "文字列"
	fmt.Printf("%#v\n", a)
	a = [...]uint8{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", a)

//var i, j interface{}
//i, j = 1, 2
//k := i + j

}
