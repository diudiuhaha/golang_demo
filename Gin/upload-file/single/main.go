package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 8 << 20 //8 Mib
	//r.Static("/", "./static")
	// 渲染静态文件
	r.LoadHTMLFiles("./static/index.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(200, "./static/index.html", nil)
	})

	//处理文件上传
	r.POST("/singleFileUpload", func(c *gin.Context) {
		// 绑定表单的元素
		name := c.PostForm("name")
		email := c.PostForm("email")
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err:", err.Error())
			return
		}
		fileName := filepath.Base(file.Filename) // 获取文件名称
		dst := "./static/" + fileName            //设置上传后的路径,也可以加前缀后缀啥的
		if err = c.SaveUploadedFile(file, dst); err != nil {
			c.String(400, "upload failed ,err:", err.Error())
			return
		}
		// 设置上传后的返回的字符串
		c.String(200, "file %s upload at %s successfully with fields name = %s and email = %s.", file.Filename, dst, name, email)
	})

	//运行
	err := r.Run(":8888")
	if err != nil {
		fmt.Println("upload failed!\nerr:", err)
		return
	}
}
