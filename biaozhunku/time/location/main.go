package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.ParseInLocation("2006/01/02 15:04:05", time.Now().Format("2006/01/02 15:04:05"), time.Local)
	fmt.Println(t) //2023-02-20 18:21:22 +0800 CST
	// 整点，向下取整
	fmt.Println(t.Truncate(1 * time.Hour)) //2023-02-20 18:00:00 +0800 CST
	// 整点，最接近
	fmt.Println(t.Round(1 * time.Hour)) //2023-02-20 18:00:00 +0800 CST
	// 整分，向下取整
	fmt.Println(t.Truncate(1 * time.Minute)) //2023-02-20 18:21:00 +0800 CST
	// 整分，最接近
	fmt.Println(t.Round(1 * time.Minute)) //2023-02-20 18:21:00 +0800 CST
	t1, _ := time.ParseInLocation("2006/01/02 15:04:05", t.Format("2006/01/02 15:04:05"), time.Local)
	fmt.Println(t1) //2023-02-20 18:21:22 +0800 CST
}
