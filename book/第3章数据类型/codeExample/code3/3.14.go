package main

func main() {
	var a [8]int                         //定义长度为8整型数组
	var b [5]int = [5]int{0, 1, 2, 3, 4} //定义长度为5的数组并初始化
	a = b                                //编译错误：cannot use b(type [5]int) as type [8]int in assignment
}
