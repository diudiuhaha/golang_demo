package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var g errgroup.Group

func router1() http.Handler {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/r1", func(c *gin.Context) {
		c.String(200, "test router 1")
	})

	return r
}

func router2() http.Handler {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/r2", func(c *gin.Context) {
		c.String(200, "test router 2")
	})

	return r
}

func main() {

	//  两个服务结构体
	ser1 := &http.Server{
		Addr:         ":8888",
		Handler:      router1(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	ser2 := &http.Server{
		Addr:    ":9999",
		Handler: router2(),
	}

	//goroutine启动
	g.Go(func() error {
		return ser1.ListenAndServe()
	})

	g.Go(func() error {
		return ser2.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}
