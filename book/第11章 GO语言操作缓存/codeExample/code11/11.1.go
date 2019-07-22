package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache" //此行代码文中后面的例子不会给出，请注意
)

func main() {
	mc := memcache.New("127.0.0.1:11211")
	if mc == nil {
		fmt.Println("memcache New failed")
		return
	}
	fmt.Println("memcache New success")
}
