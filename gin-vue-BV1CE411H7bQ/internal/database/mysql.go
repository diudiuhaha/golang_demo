package database

import (
	"gin-demo/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitMysql() *gorm.DB {
	dsn := "gin_demo:gin_demo@tcp(127.0.0.1:3306)/gin_demo?charset=utf8mb4&parseTime=True&loc=Asia%2Fshanghai"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败!" + err.Error())
	}
	log.Println("[mysql] 连接mysql成功")
	// 自动创建表格:
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("创建数据库失败", err)
	}

	return db
}

func CloseMysql() {
	s, err := db.DB()
	if err != nil {
		return
	}

	err = s.Close()
	if err != nil {
		log.Println("关闭失败,", err)
		return
	}
}
