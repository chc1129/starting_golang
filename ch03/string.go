package main

import (
	"fmt"
)

func main() {
	s := "Goの文字列"
	fmt.Printf("%v\n", s) // => "Goの文字列"

	s = "\n"
	fmt.Printf("%v", s)
	s = ""
	fmt.Printf("%v", s)
	s = "Hello, world!\n"
	fmt.Printf("%v", s)
	s = "日本語"
	fmt.Printf("%v\n", s)
	s = "\u65e5本\U00008a9e"
	fmt.Printf("%v\n", s)
	s = "\xff\u00FF"
	fmt.Printf("%v\n", s)
/*
s = "\uD800\n"
	fmt.Printf("%v\n", s)
	s = "\U00110000\n"
	fmt.Printf("%v\n", s)
*/

}
