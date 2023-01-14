package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("static/*/*")
	r.GET("/api/posts", func(c *gin.Context) {
		// 被折磨死了，下面的网页文件的地址是基于上面的相对地址！！！！！！！！！！！
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "<a href='https://www.baidu.com'>百度</a>",
		})
	})

	r.GET("/api/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
