package main

import (
	"fmt"
)

func main() {
	var f float32 = 1 << 24 // 1 << 24
	fmt.Println(f == f+1)   // "true"!
	fmt.Println(f)          // 1.6777216e+0.7
	fmt.Println(f + 1)      // 1.6777216e+0.7
}
