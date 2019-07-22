package main

import (
	"fmt"
)

func main() {
	var b [5]int = [5]int{0, 1, 2, 3, 4} //定义长度为5的数组并初始化
	//以下为第一种遍历方式
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
		//换行输出：
		//0
		//1
		//2
		//3
		//4

	}
	//以下为第二种遍历方式
	for key, value := range b {
		fmt.Printf("下标是：%d 值是：%d\n", key, value)
		//换行输出：
		//下标是：0 值是：0
		//下标是：1 值是：1
		//下标是：2 值是：2
		//下标是：3 值是：3
		//下标是：4 值是：4
	}
}
