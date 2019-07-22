package main

import (
	"fmt"
)

func main() {
	//年龄变量
	var age int = 19
	//姓名变量
	var name string = "张三"

	var (
		age2  int    = 18
		name2 string = "李四"
	)
	fmt.Println(age, name, age2, name2)
}
