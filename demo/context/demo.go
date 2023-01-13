//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//// 初始的例子
////var wg sync.WaitGroup
//
//func woker() {
//	//死循环 ，每一秒进行打印一次消息
//	for {
//		fmt.Println("this is work message")
//		time.Sleep(time.Second)
//	}
//	// 如何才能实现接收外部命令实现退出
//	wg.Done()
//}
//
//func main() {
//	wg.Add(1)
//	go woker()
//	// 如何优雅的结束
//	wg.Wait()
//	fmt.Println("over!")
//
//}

package main

func main() {

}
