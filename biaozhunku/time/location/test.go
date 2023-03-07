package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := MyTimeZone("Asia/Shanghai").Format("2016-01-02 15:04:05")
	fmt.Println(nowTime)
	//2023-02-20 14:31:33.636443 +0800 CST
	//格式化一下  20026-02-20 14:39:15

	nowTime1 := MyTimeZone("Asia/NewYork")
	fmt.Println(nowTime1)
	//err: unknown time zone Asia/NewYork
	//	0001-01-01 00:00:00 +0000 UTC

}

func MyTimeZone(tineZone string) time.Time {
	location, err := time.LoadLocation(tineZone)
	if err != nil {
		fmt.Println("err:", err)
		return time.Time{}
	}
	return time.Now().In(location)
}
