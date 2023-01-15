package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// func initMysql() {
// }
// func QueryRow() {
// }

type db *sqlx.DB      // 定义一个全局变量
type StuInfo struct { // 定义数据结构体
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Age    int    `db:"age"`
	Class  int    `db:"class"`
	StuNum int    `db:"stu_num"`
}

func main() {
	dsn := "root:123456@tcp(42.192.92.160:13306)/sqlx_test"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println("connect mysql failed,err:", err)
		db.SetMaxOpenConns(200) // 最大连接数
		db.SetMaxIdleConns(10)  // 最大空闲数
		return
	} else {
		log.Println("连接sql成功")
	}
	// 查询一行数据
	sqlStr := "select name,age from user_info where name = ?"
	var user StuInfo
	// dest是用户声明变量接收查询结果，query为查询 SQL 语句，args为绑定参数的赋值。
	err = db.Get(&user, sqlStr, "tom")
	if err != nil {
		fmt.Println("query failed , err:", err)
		return
	}
	fmt.Println(user, user.Name, user.Age)
}
