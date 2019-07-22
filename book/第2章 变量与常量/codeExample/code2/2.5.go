package main

import (
	"fmt"
)

func main() {
	var a = 2
	var b = 3
	var t = a //t=a=1连等关系
	a = b     //a = b = 2
	b = t     //b = t = 1

	fmt.Println("a=", a, "b=", b)
}
