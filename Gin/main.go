package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_demo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	// 连接成功
	fmt.Println(db, "成功")

	err = db.AutoMigrate(&Article{})
	if err != nil {
		return
	}

	r := gin.Default()

	r.POST("/article", createArticles)                    //新增文章
	r.GET("/article", getArticles)                        //获取列表
	r.GET("/article/:id", getArticlesById)                //获取详情
	r.PUT("/article/:id", updateArticles)                 //修改文章
	r.DELETE("/article/:id", delArticlesById)             //删除文章
	r.DELETE("/batchDeleteArticles", BatchDeleteArticles) //删除文章

	panic(r.Run(":8888"))
}

// Article 模型定义
type Article struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 新增
func createArticles(c *gin.Context) {
	var article Article
	err := c.BindJSON(&article)
	if err != nil {
		log.Println(err)
		return
	}

	result := db.Create(&article)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	c.JSON(200, gin.H{
		"code":   "200",
		"status": "ok",
	})
}

func getArticles(c *gin.Context) {
	var article []Article
	db.Find(&article)
	c.JSON(200, gin.H{
		"code":   "200",
		"status": "ok",
		"msg":    article,
	})

}

func getArticlesById(c *gin.Context) {
	id := c.Params.ByName("id")
	var article Article
	if err := db.Where("id", id).First(&article).Error; err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, gin.H{
			"code":   "200",
			"status": "ok",
			"msg":    article,
		})
	}
}

func updateArticles(c *gin.Context) {
	id := c.Params.ByName("id")
	var article Article
	if err := db.First(&article, id).Error; err != nil {
		c.JSON(200, gin.H{
			"code":   "404",
			"status": "err",
			"msg":    "文章查不到",
		})

		log.Println(err)
		return
	}

	var updateArticle Article
	if err := c.ShouldBindJSON(&updateArticle); err != nil {
		c.JSON(200, gin.H{
			"code":   "500",
			"status": "err",
			"msg":    "参数错误",
		})
		return
	}

	db.Model(&article).Updates(updateArticle)
	c.JSON(200, gin.H{
		"code":   "200",
		"status": "ok",
		"msg":    "成功",
	})

}

func delArticlesById(c *gin.Context) {
	id := c.Params.ByName("id")
	var article Article
	if err := db.First(&article, id).Error; err != nil {
		c.JSON(200, gin.H{
			"code":   "404",
			"status": "err",
			"msg":    "文章查不到",
		})

		log.Println(err)
		return
	}

	db.Delete(&article)
	c.JSON(200, gin.H{
		"code":   "200",
		"status": "ok",
		"msg":    "成功",
	})
}

// BatchDeleteRequest 批量删除请求结构体
type BatchDeleteRequest struct {
	ArticleIDs []int `json:"ids"`
}

// BatchDeleteArticles 批量删除文章
func BatchDeleteArticles(c *gin.Context) {
	// 从请求中获取要删除的文章ID列表
	var req BatchDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//var article Article
	//db.Where("ID in (?)", ids).Find(&article) //全部删除
	//db.Delete(&article)

	// 查询并逐个删除文章
	for _, id := range req.ArticleIDs {
		db.Delete(&Article{}, id)
	}

	// 返回删除成功的消息
	c.JSON(200, gin.H{"message": "Articles deleted successfully"})

}
