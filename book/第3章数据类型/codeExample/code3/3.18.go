package main

func main() {
	//定义一个数组
	arr := [...]string{0: "zero", 1: "one", 2: "two", 3: "three", 5: "five", 4: "four"}
	//通过数组生成一个切片
	s := arr[:5]
	//给切片追加数据
	s = append(s, "six")
	fmt.Println(s) //[zero one two three four six]
	arr[0] = "111"
	fmt.Println(s) //[111 one two three four six]
}
