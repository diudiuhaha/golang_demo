package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析URL传递的参数，对于post则解析响应包的主题（request body）
	// 注意 ：如果没有调用parseForm方法，则下面无法获取表单的数据
	fmt.Println("r.Form:", r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("url_log:", r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:value", k, ":", v)
	}
	fmt.Fprintf(w, "hello golang!")
}

// 新增表单
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		// 必须显示的调用一下 r.ParseForm()
		r.ParseForm()

		if len(r.Form["username"][0]) == 0 || len(r.Form["password"][0]) == 0 {
			println("用户名或密码未输入")
			return
		}

		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}

	//// 下拉菜单
	//s := []string{"apple", "pear", "banana"}
	//v := r.Form.Get("fruit")
	//for _, item := range s {
	//	if item == v {
	//		fmt.Println("下拉菜单的元素存在", v)
	//		return
	//	}
	//}
	//println("下拉菜单的元素不存在：", v)

	// 单选按钮
	s1 := []string{"1", "2"}
	dxV := r.Form.Get("gender")
	for _, v := range s1 {
		if v == dxV {
			fmt.Println("单选的按钮值为：", v)
		}
	}
	//println("单选的按钮不存在：", dxV)

}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求的方法method:", r.Method)
	if r.Method == "GET" {

	}
}

func main() {
	http.HandleFunc("/", sayHelloName) // 路由
	http.HandleFunc("/login", login)
	http.HandleFunc("upload", upload)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("listen failed ,err:", err)
		return
	}
}
