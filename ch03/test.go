package main

import ("fmt")

func main() {
	c := 1.3 + 4.2i
	fmt.Printf("%v\n", real(c))
	fmt.Printf("%v\n", imag(c))
}
