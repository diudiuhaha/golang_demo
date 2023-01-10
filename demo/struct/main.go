package main

import (
	"log"
	"reflect"
)

func main() {
	//定义一个结构体
	var StuInfo1 struct {
		Name  string //结构体的属性
		Age   int
		NoNum string
	}
	//结构体需要赋值，也就是实例化/实体化
	//方法1
	StuInfo1.Name = "TOM"
	StuInfo1.Age = 22
	StuInfo1.NoNum = "001"

	log.Printf("结构体SruInfo1的数据类型为：%v，值为：%v", reflect.TypeOf(StuInfo1), reflect.ValueOf(StuInfo1))
	// 结构体SruInfo1的数据类型为：struct { Name string; Age int; NoNum string }，值为：{TOM 22 001}

	// 方法2
	StuInfo2 := struct {
		Name string
		Age  int
	}{
		Name: "Jay",
		Age:  22,
	}
	log.Printf("结构体StuInfo2的数据类型为：%v，值为：%v", reflect.TypeOf(StuInfo2), reflect.ValueOf(StuInfo2))
	// 结构体StuInfo2的数据类型为：struct { Name string; Age int }，值为：{Jay 22}

}
