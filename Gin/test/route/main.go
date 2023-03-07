package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//简单的路由组1
	v1 := r.Group("/v1")
	{
		v1.GET("/submit", func(c *gin.Context) {
			c.String(200, "/v1/submit")
		})

		demo := v1.Group("demo")
		{
			demo.GET("/demo1", func(c *gin.Context) {
				c.String(200, "/v1/demo/demo1")
			})
		}
	}
	//简单的路由组2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(200, "/v2/login")
		})
		v2.GET("/submit", func(c *gin.Context) {
			c.String(200, "/v2/submit")
		})
	}

	r.Run(":8888")
}
