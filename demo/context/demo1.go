// package main
//
// import (
//
//	"fmt"
//	"sync"
//	"time"
//
// )
//
// //实现demo的两个功能点
//
// var (
//
//	wg sync.WaitGroup
//	//
//	exitWork bool
//
// )
//
//	func worker() {
//		for {
//			fmt.Println("this is work message!")
//			time.Sleep(time.Second)
//			if exitWork {
//				break
//			}
//		}
//		wg.Done()
//	}
//
//	func main() {
//		wg.Add(1)
//		go worker()
//		time.Sleep(time.Second * 3) //延时3s再给全局变量重新赋值
//		exitWork = true
//		wg.Wait()
//		fmt.Println("结束！")
//	}
package main

func main() {

}
