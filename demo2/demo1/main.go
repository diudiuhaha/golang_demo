package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// yaml
	// 拼接
	r.GET("yaml1", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"Name": "Tome",
			"Age":  22,
		})
	})
	// 结构体
	var stu3 struct {
	}

	r.Run(":8080")

}
