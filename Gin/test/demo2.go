package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	r := gin.Default()

	//随便起一个服务,5s后执行
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(200, "hello world")
	})

	srv := &http.Server{
		Addr:    "127.0.0.1:8888",
		Handler: r,
	}

	quit := make(chan os.Signal) //channel 传递信息
	signal.Notify(quit, os.Interrupt)

	// 等待/接收中断信号
	go func() {
		<-quit
		log.Println("receive interrupt signal! - 接收到中断信号")
		if err := srv.Close(); err != nil {
			log.Fatal("server close failed,", err)
		}
	}()

	//监听
	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request! - 服务器已根据请求关闭")
		} else {
			log.Fatal("server closed unexpect! - 服务器未预期关闭")
		}
	}

	log.Println("server exiting!\n 服务正在退出")
}
