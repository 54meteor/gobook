package main

import (
	"fmt"
)

func main() {
	var a [8]int //定义长度为8整型数组
	//定义长度为5的数组并初始化
	var b [5]int = [5]int{0, 1, 2, 3, 4}
	fmt.Println(a[0])        //输出：0
	fmt.Println(b[len(b)-1]) //输出:4
	fmt.Println(b)           //输出：[0 1 2 3 4]
	b[0] = 100               //修改b数组第0个元素的值
	fmt.Println(b[0])        //输出：100
}
