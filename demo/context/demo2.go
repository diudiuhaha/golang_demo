package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(exitWork chan struct{}) {
LOOP:
	for {
		fmt.Println("this is work message")
		time.Sleep(time.Second)
		select {
		case <-exitWork:
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	var exitWork = make(chan struct{})
	wg.Add(1)
	go worker(exitWork)
	time.Sleep(time.Second * 3)
	exitWork <- struct{}{}
	wg.Wait()
	fmt.Println("结束")
}
