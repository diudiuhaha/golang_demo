package main

import (
	"gin-demo/internal/controller"
	"github.com/gin-gonic/gin"
)

func DefaultRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}
