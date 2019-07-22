package main

import (
	"fmt"
)

func main() {
	const ( // iota被重设为0
		a = iota // c0 == 0
		b        // c1 == 1
		c        // c2 == 2
	)

	fmt.Println(a, b, c)

}
