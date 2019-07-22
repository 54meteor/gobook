package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string
	s1 := "this is String"
	ch := s1[0]
	fmt.Println(s)                  //这里什么也不会显示
	fmt.Println(s1)                 //这里输出this is String
	fmt.Println(reflect.TypeOf(ch)) //这里输出uint8，也就是byte类型
	fmt.Println(ch)                 //这里输出ascii码116
	fmt.Printf("%c", ch)            //这里输出字母t
}
