package main

import (
	"fmt"
)

func main() {
	//创建一个预留长度为10，初始元素个数为5的数组切片
	s := make([]string, 5, 10)
	fmt.Println(s)      //输出:[     ]
	fmt.Println(cap(s)) //输出：10
	fmt.Println(len(s)) //输出：5
	s = append(s, "1")
	s[0] = "zero"
	s[1] = "one"
	fmt.Println(s)      //[zero one    1]
	fmt.Println(cap(s)) //输出：10
	fmt.Println(len(s)) //输出：6
}
