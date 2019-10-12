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

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

var plusAlias = plus

func main() {

	fmt.Printf("%d\n", plus(1, 2))
	fmt.Printf("%d\n", plusAlias(10, 5))
	hello()

	q, r := div(19, 7)
	fmt.Printf("商=%d 剰余=%d\n", q, r)

	fmt.Printf("%T\n", func(x, y int) int { return x + y })
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3))

	f := returnFunc()
	f()

	callFunction(func() {
		fmt.Println("I'm a function")
	})

}
