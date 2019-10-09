package main

import (
	"fmt"
)

func plus(x, y int) int {
	return x + y
	}

func hello() {
	fmt.Println("Hello, World!")
	return
}

func div(a, b int) (int, int) {
	q := a / b // aをbで割った商
	r := a % b // aをbで割った剰余
	return q, r
}

func main() {

	fmt.Printf("%d\n", plus(1, 2))
	hello()

	q, r := div(19, 7)
	fmt.Printf("商=%d 剰余=%d\n", q, r)

}
