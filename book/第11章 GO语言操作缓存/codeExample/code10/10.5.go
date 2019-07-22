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

	redigo.Do("SADD", "tt", "1")
	redigo.Do("SADD", "tt", "2")
	redigo.Do("SADD", "tt", "3")
	redigo.Do("SADD", "tt", "1")

	values, _ := redis.Values(redigo.Do("SMEMBERS", "tt"))

	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
	//输出：
	//	1
	//	2
	//	3
}
