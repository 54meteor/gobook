package main

import (
	"fmt"
)

func main() {
	//定义一个整型变量
	var a int = 10
	//定义一个指针变量
	var p *int
	//指针p指向变量a的地址
	p = &a
	fmt.Println(a)  //输出：10
	fmt.Println(&a) //输出：0xc420014100
	fmt.Println(p)  //输出：0xc420014100
	fmt.Println(*p) //输出：10
	*p = 100
	fmt.Println(a)       //输出：100
	fmt.Println(&a)      //输出：0xc420014100
	fmt.Println(p)       //输出：0xc420014100
	fmt.Println(*p)      //输出：100
	fmt.Println(*p == a) //输出：true
}
