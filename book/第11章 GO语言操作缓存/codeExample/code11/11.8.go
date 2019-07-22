package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	//	mc := memcache.New("192.1.1.1:11211")
	mc := memcache.New("127.0.0.1:11211")
	value, err := mc.Decrement("aaa", 4)
	if err != nil {
		fmt.Println("Decrement failed", err.Error())
	} else {
		fmt.Println("after decrement the value is ", value)
	}
}
