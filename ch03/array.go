package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", a) // => "[1, 2, 3, 4, 5]"

fmt.Printf("\n")

	a = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a[0]=%v\n", a[0])
	fmt.Printf("a[1]=%v\n", a[1])
	fmt.Printf("a[2]=%v\n", a[2])
	fmt.Printf("a[3]=%v\n", a[3])
	fmt.Printf("a[4]=%v\n", a[4])
//	fmt.Printf("a[5]=%v\n", a[5])
//	fmt.Printf("a[-1]=%v\n", a[-1])

	a = [5]int{}
	fmt.Printf("%v\n", a)
	a = [5]int{1, 2, 3}
	fmt.Printf("%v\n", a)
//	a = [5]int{1, 2, 3, 4, 5, 6}
//	fmt.Printf("%v\n", a)

	var b [5]int
	fmt.Printf("%v\n", b) // b == [0, 0, 0, 0]
	c := [5]int{} // b == [0, 0, 0, 0]
	fmt.Printf("%v\n", c)

	ia := [3]int{}
	fmt.Printf("%v\n", ia)
	ua := [3]uint{}
	fmt.Printf("%v\n", ua)
	ba := [3]bool{}
	fmt.Printf("%v\n", ba)
	fa := [3]float64{}
	fmt.Printf("%v\n", fa)
	ca := [3]complex128{}
	fmt.Printf("%v\n", ca)
	ra := [3]rune{}
	fmt.Printf("%v\n", ra)
	sa := [3]string{}
	fmt.Printf("%v\n", sa)

// 要素数の省略
	a1 := [...]int{1, 2, 3}
	fmt.Printf("%v\n", a1)
	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a2)
	a3 := [...]int{}
	fmt.Printf("%v\n", a3)

// 要素への代入
	d := [...]int{1, 2, 3} // a == {1, 2, 3}
	d[0] = 0
	d[2] = 0
	fmt.Printf("%v\n", d) // => "[0, 2, 0]"

// 配列型の互換性
	e1 := [3]int{1, 2, 3}
	e2 := [3]int{4, 5, 6}

	e1 = e2 // a1にa2を代入

	fmt.Printf("%v\n", e1) // => "[4, 5, 6]" a1とa2は同一の内容

	e1[0] = 0
	e1[2] = 0

	fmt.Printf("%v\n", e1) // => "[0, 5, 0]"
	fmt.Printf("%v\n", e2) // => "[4, 5, 6]"

}
