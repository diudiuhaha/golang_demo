package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/api/upload", func(c *gin.Context) {
		// 从请求中读取文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		} else {
			// 将读取到的文件保存到xxx
			// dst := fmt.Sprintf("./%s",file.Filename) // 方法1
			dst := path.Join("./", file.Filename) // 方法2
			err := c.SaveUploadedFile(file, dst)  // 文件，目的路径
			if err != nil {
				fmt.Println(err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "upload OK",
			})
		}
	})

	r.Run(":8080")
}
