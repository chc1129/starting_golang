package main

import (
	"fmt"
)

func main() {
	r := '松'
	fmt.Printf("%v\n", r)
	r = 'a'
	fmt.Printf("%v\n", r)
	r = 'ä'
	fmt.Printf("%v\n", r)
	r = '本'
	fmt.Printf("%v\n", r)
	r = '\t'
	fmt.Printf("%v\n", r)
	r = '\000'
	fmt.Printf("%v\n", r)
	r = '\007'
	fmt.Printf("%v\n", r)
	r = '\377'
	fmt.Printf("%v\n", r)
	r = '\x07'
	fmt.Printf("%v\n", r)
	r = '\xff'
	fmt.Printf("%v\n", r)
	r = '\u12e4'
	fmt.Printf("%v\n", r)
	r = '\U00101234'
	fmt.Printf("%v\n", r)
/*
r = 'aa'
	fmt.Printf("%v\n", r)
	r = '\xa'
	fmt.Printf("%v\n", r)
	r = '\0'
	fmt.Printf("%v\n", r)
	r = '\uDFFF'
	fmt.Printf("%v\n", r)
	r = '\U00110000'
	fmt.Printf("%v\n", r)
*/

}
