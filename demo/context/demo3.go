package main

import (
	"context"
	"fmt"
)

// 定义一个方法，传入context.Context类型的参数，输出一个整型（从channel中取值）
func gen(context1 context.Context) <-chan int {
	chan1 := make(chan int) //声明一个管道并分配内存
	n := 1
	//起一个协程，匿名函数
	go func() {
		for {
			select {
			case <-context1.Done(): // Context Done(）返回的是一个管道
				fmt.Println("context返回channel的值", <-context1.Done())
				return //结束协程？？
			case chan1 <- n:
				fmt.Println("n:", n)
				n++
			default:
			}
		}
	}()
	return chan1
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 当我们取完需要的整数后调用cancel，相当于向ctx里面添加值
	defer cancel() //var cancel context.CancelFunc
	//遍历chanel
	for n := range gen(ctx) {
		if n == 5 {
			break
		}
	}
}
