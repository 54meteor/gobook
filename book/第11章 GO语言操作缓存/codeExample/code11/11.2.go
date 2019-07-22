package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New("127.0.0.1:11211")
	err := mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value1")})
	fmt.Println("set is success")
}
