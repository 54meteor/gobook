package main

import (
	"fmt"
)

func main() {
	var s = "this is String"
	ch := s[1]
	fmt.Printf("%c", ch)
	s[1] = "k" //cannot assign to s[1]
}
