package main

import (
	"fmt"
	)

func main() {
	s := "Taro" + " " + "Yamada"
	fmt.Printf("%v\n", s)

var a int
a = 4 + 2
fmt.Printf("%d\n", a)
a = 14 - 7
fmt.Printf("%d\n", a)
a = 3 * 7
fmt.Printf("%d\n", a)
a = 12 / 4
fmt.Printf("%d\n", a)
a = 5 % 2
fmt.Printf("%d\n", a)

var n int
n = 165 & 155
fmt.Printf("%d\n", n)
n = 197 | 169
fmt.Printf("%d\n", n)
n = 92 ^ 137
fmt.Printf("%d\n", n)
n = 108 &^ 13
fmt.Printf("%d\n", n)
n = 1 &^ 1
fmt.Printf("%d\n", n)
n = 99 &^ 99
fmt.Printf("%d\n", n)
n = 255 &^ 255
fmt.Printf("%d\n", n)

var m uint8
m = ^uint8(13)
fmt.Printf("%d\n", m)
var l int8
l = ^int8(13)
fmt.Printf("%d\n", l)

var h int
h = 1 << 1
fmt.Printf("%d\n", h)
h = 1 >> 1
fmt.Printf("%d\n", h)

var i int
i += 5
fmt.Printf("%d\n", i)
i *= 10
fmt.Printf("%d\n", i)
//i &= x
//fmt.Printf("%d\n", i)
i <<= 3
fmt.Printf("%d\n", i)

var j string
j += "の解説"
fmt.Printf("%s\n", j)

}
