package main

import (
	"fmt"
)

func main() {
	var i = 9 //直接赋值
	var j int //没有直接赋值，go语言机制默认使其持有默认值0
	j = 19    //显式为上一行中的j变量赋值为19
	fmt.Println(i, j)
}
