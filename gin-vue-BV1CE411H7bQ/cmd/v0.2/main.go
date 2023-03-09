package main

import (
	"gin-demo/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	//gorm
	db := database.InitMysql()
	dbClose, _ := db.DB()
	dbClose.Close()
	//gin
	r := gin.Default()
	DefaultRoutes(r)

	panic(r.Run(":8888"))
}
