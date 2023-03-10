package main

import (
	"gin-demo/internal/controller"
	"github.com/gin-gonic/gin"
)

func DefaultRoutes(r *gin.Engine) *gin.Engine {
	//用户路由组
	authApi := r.Group("/api/auth")
	{
		//注册
		authApi.POST("/register", controller.Register)
		//登录
		authApi.POST("/login", controller.Login)
	}
	//
	return r
}
