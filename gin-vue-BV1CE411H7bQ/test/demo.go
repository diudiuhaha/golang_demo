package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// 全局变量 - 人员信息 - 用于创建user表
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(110);not null;unique"`
	Password  string `gorm:"size:255;not null""`
}

func DBInit() *gorm.DB {
	//连接数据库
	//  参考：dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	user := "gin_demo"       // 用户名
	password := "gin_demo"   //密码
	host := "127.0.0.1"      //地址
	port := "3306"           //端口
	dbname := "gin_demo"     //库名字
	charset := "utf8mb4"     //字符集
	parseTime := "True"      //格式化时间
	loc := "Asia%2Fshanghai" //时区
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		user, password, host, port, dbname, charset, parseTime, loc)
	//创建连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("%s连接失败,err:%s", dbname, err)
		return
	} else {
		log.Printf("%s连接成功", dbname)
	}

	// 创建表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Printf("%s的user表创建失败，err：err:%s", dbname, err)
		return
	}
	return db
}

func main() {
	DBInit()
	// 写gin接口
	r := gin.Default()
	//注册接口
	r.POST("/register", func(c *gin.Context) {
		// 绑定表单
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		// 数据验证
		// 注册成功
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})
	})

	// 登录接口

	// 运行
	panic(r.Run(":"))

}
