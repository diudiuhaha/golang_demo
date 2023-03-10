package main

import (
	"gin-demo/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	//gorm
	database.InitMysql()
	defer database.CloseMysql()
	//gin
	r := gin.Default()
	DefaultRoutes(r)

	panic(r.Run(":8888"))
}
