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
	mc.Add(&memcache.Item{Key: "foo3", Value: []byte("bluegogo1")})
	it, err := mc.Get("foo3")
	if err != nil {
		fmt.Println("Add failed")
	} else {
		if string(it.Key) == "foo3" {
			fmt.Println("Add value is ", string(it.Value))
		} else {
			fmt.Println("Get failed")
		}
	}
}
