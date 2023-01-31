package main

import (
	"context"
	"fmt"
	"time"
)

func handle(c context.Context, duration time.Duration) {
	select {
	case <-c.Done():
		fmt.Println("handle", c.Err()) //返回结束的原因
	case <-time.After(duration):
		/*
			是本次监听动作的超时时间
			意思就说，只有在本次select 操作中会有效
			再次select 又会重新开始计时，
		*/
		fmt.Println("process request with", duration) //持续时间
	}
}

func main() {
	// 创建一个过期时间为1s的上下文
	// context.WithTimeout 返回一个context结构体和一个取消的函数
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() //关闭

	//并向上下文传入handle函数，该方法会使用500ms时间处理传入的请求
	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

}
