package main

import (
	"fmt"
)

func main() {
	const ( // iota被重设为0
		a = iota // a == 0
		b = iota // b == 1
		c = iota // c == 2
	)
	const (
		x         = iota * 12 // x == 0
		y float64 = iota * 12 // y == 12.0
		z         = iota * 12 // z == 24
	)
	const j = iota // j == 0 (因为iota主要遇到const就会被重设为0)
	const k = iota //k == 0 (同上)
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
	fmt.Println(j, k)
}
