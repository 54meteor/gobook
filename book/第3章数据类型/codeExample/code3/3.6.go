package main

import (
	"fmt"
)

func main() {
	var f float32 = 123453.1415926
	fmt.Printf("%8.2f", f) // 123453.14
	fmt.Printf("%3.1f", f) // 123453.1
	fmt.Printf("%3.4f", f) // 123453.1406
}
