package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

// 全局变量，用户结构
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
	Telephone string `gorm:"varchar(110);not null" json:"telephone"`
	Password  string `gorm:"size:255;not null" json:"password"`
}

func main() {
	//gorm
	dsn := "gin_demo:gin_demo@tcp(127.0.0.1:3306)/gin_demo?charset=utf8mb4&parseTime=True&loc=Asia%2Fshanghai"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("[mysql] 连接mysql成功")
	// 自动创建表格:
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
		return
	}

	// 连接后立刻关闭
	DB, _ := db.DB()
	defer log.Println("[mysql] 断开连接")
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(DB)

	//gin
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		//手机号长度验证：
		if len(telephone) != 11 {
			log.Println("[err] 手机号长度必须为11位")
			return
		}
		// 手机号是否在数据库中存在
		if isTelephone(db, telephone) {
			log.Println("[err] 手机号已经存在")
			return
		}

		//密码验证：
		if len(password) < 8 {
			log.Println("[err] 密码必须为8位")
			return
		}

		// 随机生成用户名
		if len(name) == 0 {
			name = RandomString(5)
		}

		//存入数据库
		createNewUser := User{ //结构体实例化
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		result := db.Create(&createNewUser)
		log.Printf("[mysql] 创建成功，影响了 %d 行数据", result.RowsAffected)

		c.JSON(200, gin.H{
			"message":   "注册成功",
			"name":      name,
			"telephone": telephone,
			"password":  password,
		})

	})

	panic(r.Run(":8888"))
}

// 随机字符串函数,首字母大写表示可以被其他引用
func RandomString(n int) string {
	var letters = []byte("abcdABCD123")
	//创建一个字符串切片，长度位n
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))] //返回 0-n之间的数
	}
	return string(result)
}

// 判断手机号是否存在
func isTelephone(db *gorm.DB, telephone string) bool {
	var user User //实例化用户结构体
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true // 如果查得出来数据，则证明是存在的
	}
	return false //否则，是不存在的
}
