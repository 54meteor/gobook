package main

import (
	"fmt"
)

func main() {
	//创建一个map,key为string类型，value为string类型
	m := make(map[string]string)
	//给map增加值
	m["username"] = "admin"
	m["sex"] = "man"
	m["age"] = "20"
	fmt.Println(m) //输出：map[username:admin sex:man age:20]
	//删除键值
	delete(m, "age")
	fmt.Println(m) //输出：map[username:admin sex:man]
	//查询键值是否存在
	value, ok := m["username"]
	if ok {
		fmt.Println(value) //输出：value的结果，此处为admin
	} else {
		fmt.Println("nil") //输出：如果ok返回false，则输出nil
	}
}
