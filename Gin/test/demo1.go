package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//AsciiJson
	msg1 := map[string]string{
		"lange": "Go语言",
		"tag":   "<br>",
	}
	r.GET("/asciiJson", func(c *gin.Context) {
		c.AsciiJSON(200, msg1)
	})

	//jsonp
	r.GET("/jsonp-test", func(c *gin.Context) {
		msg2 := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(200, msg2)
	})

	//Multipart/Urlencoded binding
	type loginForm struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}
	r.POST("/login", func(c *gin.Context) {
		var loginForm loginForm
		// 方法1 ：显示声明
		//c.ShouldBindWith(&loginForm, binding.Form)
		// 方法2 ：自动绑定
		if c.ShouldBind(&loginForm) == nil {
			if loginForm.User == "user" && loginForm.Password == "password" {
				c.JSON(200, gin.H{
					"status": "login success!",
				})
			} else {
				c.JSON(401, gin.H{
					"status": "unauthorized!",
				})
			}
		}
	})

	// Multipart/Urlencoded 表单
	r.POST("/form_post", func(c *gin.Context) {
		//c.PostFormArray() 或 c.PostFormMap()
		message := c.PostForm("message") //获取表单中的数据，如果没有则为空
		// 获取表单中 nick 的数据，如果没有设置一个默认值
		nick := c.DefaultPostForm("nick", "test")
		c.JSON(200, gin.H{
			"status":      "posted",
			"message":     message,
			"defaultForm": nick,
		})
	})

	//secure json
	r.GET("secure_json", func(c *gin.Context) {
		// 可以使用自定义前缀
		r.SecureJsonPrefix("（*{}[],\n")
		//设置数组
		names := []string{"lena", "austin", "foo"}
		c.SecureJSON(200, names)
	})

	//////////////////////////////
	// XML/JSON/YAML/ProtoBuf 渲染
	// 结构体
	var msg2 = struct {
		Name   string `json:"name1" yaml:"name2" xml:"name3"`
		Age    int
		Number int
	}{
		Name:   "tom",
		Age:    11,
		Number: 100,
	}
	//json
	// {"name1":"tom","Age":11,"Number":100}
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, msg2)
	})
	/*
		<map>
		    <message>
		        <name3>tom</name3>
		        <Age>11</Age>
		        <Number>100</Number>
		    </message>
		    <name>xml</name>
		</map>
	*/
	r.GET("/xml1", func(c *gin.Context) {
		c.XML(200, gin.H{
			"name":    "xml",
			"message": msg2,
		})
	})
	/*
		message:
		    name2: tom
		    age: 11
		    number: 100
		name: yaml
	*/
	r.GET("/yaml1", func(c *gin.Context) {
		c.YAML(200, gin.H{
			"name":    "yaml",
			"message": msg2,
		})
	})
	///////////////////////////////
	// 上传文件
	/// 单文件
	//略
	/// 多个文件
	//略
	//////////////////////////////////
	r.GET("/getDataFormReader", func(c *gin.Context) {
		// 获取请求信息
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		//获取reader
		reader := response.Body
		contextLength := response.ContentLength
		contextType := response.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"content-Disposition": `attachment;filename="gopher.png"`,
		}
		c.DataFromReader(200, contextLength, contextType, reader, extraHeaders)

	})

	////////////////////////////////
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
