package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	//	mc := memcache.New("192.1.1.1:11211")
	mc := memcache.New("127.0.0.1:11211")
	err := mc.Set(&memcache.Item{Key: "aaa", Value: []byte("1")})
	if err != nil {
		fmt.Println("Set failed :", err.Error())
	}
	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println("Get failed ", err.Error())
	} else {
		fmt.Println("src value is:", it.Value)
	}
	value, err := mc.Increment("aaa", 7)
	if err != nil {
		fmt.Println("Increment failed")
	} else {
		fmt.Println("after increment the value is :", value)
	}
}
