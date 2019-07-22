package main

import (
	"fmt"
)

func main() {
	//声明一个数组
	arr := [...]string{0: "zero", 1: "one", 2: "two", 3: "three", 5: "five", 4: "four"}
	//创建一个切片
	s := make([]string, 1, 10)
	//将数组追加的方式加入到切片中
	s = append(s, arr[0:]...)
	//提取插入点之后的数据
	insertAfter := append([]string{}, s[4:]...)
	//将要插入的数据写入切片
	s = append(s[0:4], "inserted")
	//组合数据
	s = append(s, insertAfter...)
	fmt.Println(s) //输出：[ zero one two inserted three four five]
}
