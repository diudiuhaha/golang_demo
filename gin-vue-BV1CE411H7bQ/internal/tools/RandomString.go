package tools

import (
	"math/rand"
	"time"
)

// 随机字符串函数,首字母大写表示可以被其他引用
func RandomString(n int) string {
	var letters = []byte("abcdABCD123")
	//创建一个字符串切片，长度位n
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))] //返回 0-n之间的数
	}
	return string(result)
}
