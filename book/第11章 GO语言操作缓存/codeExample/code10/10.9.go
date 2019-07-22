package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var RedisClient *redis.Pool

func init() {
	// 建立连接池
	RedisClient = &redis.Pool{
		//最大空闲连接数
		MaxIdle: 1,
		//最大激活连接数
		MaxActive: 10,
		//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
func main() {

	// 从池里获取连接
	rc := RedisClient.Get()
	v, err := rc.Do("SET", "color", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v) //输出ok
	v, err = redis.String(rc.Do("GET", "color"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v) //输出red
	// 用完后将连接放回连接池
	defer rc.Close()
}
