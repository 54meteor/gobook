package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	//	mc := memcache.New("192.1.1.1:11211")
	mc := memcache.New("127.0.0.1:11211")
	if mc == nil {
		fmt.Println("memcache New failed")
		return
	}
	mc.Set(&memcache.Item{Key: "foo2", Value: []byte("my value1")})
	mc.Replace(&memcache.Item{Key: "foo2", Value: []byte("mobike")})
	it, err := mc.Get("foo2")
	if err != nil {
		fmt.Println("Add failed")
	} else {
		if string(it.Key) == "foo2" {
			fmt.Println("Add value is ", string(it.Value))
		} else {
			fmt.Println("Get failed")
		}
	}
}
