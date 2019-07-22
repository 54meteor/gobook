package main

import (
	"fmt"
)

func main() {
	s := [...]string{0: "zero", 1: "one", 2: "two", 3: "three", 5: "five", 4: "four"}
	fmt.Println(s[0:6]) //输出：[zero one two three four five]
	fmt.Println(s[2:4]) //输出：[two three]
	fmt.Println(s[:4])  //输出：[zero one two three]
	fmt.Println(s[3:])  //输出：[three four five]
	fmt.Println(s[0:8]) //编译错误：./main.go:19:15: invalid slice index 8 (out of bounds for 6-element array)
}
