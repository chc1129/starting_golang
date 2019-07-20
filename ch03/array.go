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

}
