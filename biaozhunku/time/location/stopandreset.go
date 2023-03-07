package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now() //开始时间
	fmt.Println(startTime)
	//现在的时间与开始的时间相隔多少
	timer := time.AfterFunc(1*time.Second, func() { //2）2s后执行AfterFunc
		fmt.Println("after func callback,elaspe:", time.Now().Sub(startTime), time.Now())
	})

	//time.Sleep(1 * time.Second) // 休眠一秒
	time.Sleep(3 * time.Second)
	//Reset 在 Timer 还未触发时返回 true；触发了或 Stop 了，返回 false

	if timer.Reset(3 * time.Second) { // 1）第1s时，未触发3s重置，返回true
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop", "nowTime:", time.Now())
	}

	time.Sleep(20 * time.Second) // 防止程序AfterFunc还没执行就停掉了

}
