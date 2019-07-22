package main

import (
	"fmt"
)

func main() {
	var v int64
	v1 := 10 //系统自动推断数据类型为int
	v = v1   //此处报错
	fmt.Println(v)
}
