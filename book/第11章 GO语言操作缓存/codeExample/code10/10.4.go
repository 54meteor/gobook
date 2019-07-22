package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	redigo, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	redigo.Do("lpush", "redlist", "qqq")
	redigo.Do("lpush", "redlist", "www")
	redigo.Do("lpush", "redlist", "eee")

	values, _ := redis.Values(redigo.Do("lrange", "redlist", "0", "100"))

	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
	//输出：
	//	eee
	//	www
	//	qqq
}
