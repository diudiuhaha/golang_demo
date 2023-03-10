package controller

import (
	"gin-demo/internal/database"
	"gin-demo/internal/models"
	"gin-demo/internal/tools"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

var DB = database.InitMysql()

func Register(c *gin.Context) {
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	//手机号长度验证：
	if len(telephone) != 11 {
		c.JSON(422, gin.H{"code": "422", "message": "手机号长度必须为11位"})
		return
	}
	// 手机号是否在数据库中存在
	if isTelephone(DB, telephone) {
		c.JSON(422, gin.H{"code": "422", "message": "手机号已经存在"})
		return
	}

	//密码验证：
	if len(password) < 8 {
		c.JSON(422, gin.H{"code": "422", "message": "密码必须为8位以上"})
		return
	}

	// 随机生成用户名
	if len(name) == 0 {
		name = tools.RandomString(5)
	}

	//存入数据库
	//做密码的加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(422, gin.H{"code": "500", "message": "密码加密失败"})
		return
	}
	createNewUser := models.User{ //结构体实例化
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	result := DB.Create(&createNewUser)
	log.Printf("[mysql] 创建成功，影响了 %d 行数据", result.RowsAffected)

	c.JSON(200, gin.H{
		"code":    "200",
		"message": "注册成功",
	})

}

// 登录
func Login(c *gin.Context) {

	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	//手机号长度验证：
	if len(telephone) != 11 {
		c.JSON(422, gin.H{"code": "422", "message": "手机号长度必须为11位"})
		return
	}

	//判断手机号是否存在，这里逻辑是反过来了 ，只有手机号存在，才会继续去进行
	var user models.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 { //如果id是0则
		c.JSON(422, gin.H{"code": "422", "message": "手机号不存在！"})
		return
	}
	//判断密码是否正确

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(422, gin.H{"code": "422", "message": "密码错误！"})
		return
	}
	//发放token
	token := "test"
	//成功
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

// 判断手机号是否存在
func isTelephone(db *gorm.DB, telephone string) bool {
	var user models.User //实例化用户结构体
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true // 如果查得出来数据，则证明是存在的
	}
	return false //否则，是不存在的
}
