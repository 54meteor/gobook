package main

import (
	"fmt"
)

func main() {
	//创建一个预留长度为10，初始原素个数为5的数组切片
	s := make([]string, 5, 10)
	fmt.Println(s)      //输出:[     ]
	fmt.Println(cap(s)) //输出：10
	fmt.Println(len(s)) //输出：5
	s = append(s, "1")  //增加一个元素
	s[0] = "zero"       //修改第1个元素的值
	s[1] = "one"        //修改第2个元素的值
	s[7] = "seven"      //这里执行时错误：index out of range
	fmt.Println(s)      //[zero one    1]
	fmt.Println(cap(s)) //输出：10
	fmt.Println(len(s)) //输出：6
}
