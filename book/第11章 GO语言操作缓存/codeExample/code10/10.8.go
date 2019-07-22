package main

import (
	"fmt"
	//	"reflect"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	start := time.Now().UnixNano() //获取开始时间戳
	for i := 0; i < 100000; i++ {
		c.Send("SET", "foo"+string(i), "bar"+string(i))
	}

	end := time.Now().UnixNano() //获取结束时间戳
	fmt.Println(end - start)     //输出141727169纳秒=0.1417272秒
}
