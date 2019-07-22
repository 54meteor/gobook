package main

import (
	"fmt"
)

func main() {
	var v int64
	v1 := 10 //系统自动推断数据类型为int
	v = int64(v1)
	fmt.Println(v) //输出10
}
