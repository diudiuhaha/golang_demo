package main

import "github.com/gin-gonic/gin"

func main() {
	var secrets = gin.H{ //私人数据map
		"test1": gin.H{"email": "1@test.com", "phone": "123"},
		"test2": gin.H{"email": "1@test.com", "phone": "456"},
		"test3": gin.H{"email": "1@test.com", "phone": "789"},
	}

	r := gin.Default()
	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"test1": "1", //账号密码字段
		"test2": "1",
		"test3": "1",
		"test4": "1",
	}))
	// /admin/secrets 端点
	// 触发 "localhost:8888/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secrets, ok := secrets[user]; ok {
			c.JSON(200, gin.H{
				"user":   user,
				"secret": secrets,
			})
		} else {
			c.JSON(200, gin.H{
				"user":   user,
				"secret": "NO SECRET",
			})
		}
	})

	r.Run(":8888")
}
