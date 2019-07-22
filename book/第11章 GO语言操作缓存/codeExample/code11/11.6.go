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
	errd := mc.Delete("foo")
	if errd != nil {
		fmt.Println("Delete failed:", errd.Error())
	}
	_, err := mc.Get("foo")
	if err != nil {
		fmt.Println("delete success")
	}
}
