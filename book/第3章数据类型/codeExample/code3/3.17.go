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
	fmt.Println(&p) //输出：0xc42000c028
	//定义一个指向指针的指针
	var ptr **int
	//将指针ptr指向指针p的地址
	ptr = &p
	fmt.Println(ptr)   //输出：指针ptr的地址为0xc42000c028
	fmt.Println(*ptr)  //输出：0xc420014100
	fmt.Println(**ptr) //输出：10
}
