package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis" //此行代码文中后面的例子不会给出，请注意
)

func main() {
	host := "127.0.0.1"
	port := "6379"
	protocol := "tcp"

	redis, err := redis.Dial(protocol, host+":"+port)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	fmt.Println("Connect to redis succeed")
	defer redis.Close()
}