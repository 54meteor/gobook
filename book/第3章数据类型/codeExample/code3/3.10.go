package main

import (
	"fmt"
)

func main() {
	var s1 rune
	var s2 rune
	var s3 rune
	s1 = 'a'
	s2 = "a"//此处报错cannot use "a" (type string) as type rune is assignment
	s3 = 'abc'//此处报错 invalid character literal(more than one character)
	fmt.Println(s1)      //97
	fmt.Printf("%c", s1) //a
}
