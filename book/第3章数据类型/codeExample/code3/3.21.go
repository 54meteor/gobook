package main

import (
	"fmt"
)

func main() {
	s := make([]string, 1, 3)
	fmt.Println(cap(s))
	fmt.Println(len(s))
	s[0] = "0"
	s1 := make([]string, 2, 2)
	s1[0] = "1"
	s1[1] = "2"
	for i := 0; i < 5; i++ {
		s = append(s, s1...)
		fmt.Println(s)
		fmt.Println("cap:", cap(s))
		fmt.Println("len:", len(s))
	}
}
