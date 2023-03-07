package main

import "github.com/gin-gonic/gin"

func main() {
	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	r := gin.Default()
	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"name":   name,
			"action": action,
		})
	})

	r.Run(":8888")

}
