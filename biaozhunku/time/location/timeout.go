package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go func() { //起一个协程，每2秒取channel
		time.Sleep(2 * time.Second)
		fmt.Println(<-c)
	}()
	//主流程挂了 协程也没了
	select {
	case c <- 1:
		fmt.Println("channel...")
	case <-time.After(3 * time.Second):
		close(c)
		fmt.Println("timeout")
	}

}
