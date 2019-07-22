package main

import (
	"fmt"
)

func main() {
	_, _, nickName := GetName()
	fmt.Println(nickName)
}
func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}
