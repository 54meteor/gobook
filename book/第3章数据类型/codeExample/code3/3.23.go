package main

import (
	"fmt"
)

func main() {
	s := [...]string{0: "zero", 1: "one", 2: "two", 3: "three", 5: "five", 4: "four"}
	fmt.Println(s[:2])            //输出：[zero one]
	fmt.Println(s[3:])            //输出：[three four five]
	s1 := append(s[:2], s[3:]...) //将[three four five] 连接到[zero one]后面生成新的切片
	fmt.Println(s1)               //输出：[zero one three four five]我们通过这种写法删除掉了2:"two"这个元素，由此我们实现了删除一个元素的功能
}
