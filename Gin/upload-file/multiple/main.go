package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()

	//渲染网页文件
	r.LoadHTMLGlob("./static/*")
	r.GET("/upload1", func(c *gin.Context) {
		c.HTML(200, "./static/index.html", nil)
	})

	// 处理上传文件
	r.POST("/multipleFileUpload", func(c *gin.Context) {
		//获取表单内容
		name := c.PostForm("name")
		email := c.PostForm("email")

		// 可以文件名列表
		//var fileList []string
		//多文件处理
		form, err := c.MultipartForm()

		if err != nil {
			//fmt.Println(err)
			c.String(http.StatusBadRequest, "get form failed,%s", err.Error())
			return
		}
		files := form.File["files"]
		for _, file := range files { //循环遍历多文件
			fileName := filepath.Base(file.Filename) //获取文件名

			//fileList = append(fileList, fileName)

			dst := "./temp/" + fileName //设置路径
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				c.String(http.StatusBadRequest, "upload file failed,%s", err.Error())
				return
			}
		}

		c.String(200, "upload successfully %d files with fields name=%s and email %s", len(files), name, email)
		//fmt.Println(fileList)
	})

	err := r.Run(":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
}
