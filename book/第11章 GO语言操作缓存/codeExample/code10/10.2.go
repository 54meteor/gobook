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
	v, err := redigo.Do("SET", "color", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v) //输出ok
	v, err = redis.String(redigo.Do("GET", "color"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v) //输出red
}
