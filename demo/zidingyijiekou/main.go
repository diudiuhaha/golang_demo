package main

import "fmt"

func main() {
	//创建一个新接口,名字不能是一个变量
	type Weekday int
	const ( //常量，一定要首字母大写
		Monday Weekday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Sunday)

}
