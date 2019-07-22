package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")
	c.Flush()
	s, err := c.Receive() // reply from SET
	fmt.Println(s)        //输出:OK
	v, err := c.Receive() // reply from GET
	fmt.Printf("%s", v)   //输出：bar
}
