package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/FileTest", func(c *gin.Context) {
		// FormFile返回所提供的表单键的第一个文件
		f, _ := c.FormFile("file")
		// SaveUploadedFile上传表单文件到指定的路径
		c.SaveUploadedFile(f, "./"+f.Filename)
		c.JSON(200, gin.H{
			"msg": f,
		})
	})
	r.Run(":8080")
}
