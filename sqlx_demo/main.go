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
	dsn := "root:123456@tcp(10.0.0.10:3306)/sqlx_test"
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
	sqlStr := "select name,age from user_info where id = ?"
	var user StuInfo
	// dest是用户声明变量接收查询结果，query为查询 SQL 语句，args为绑定参数的赋值。
	err = db.Get(&user, sqlStr, "1")
	if err != nil {
		fmt.Println("query failed , err:", err)
		return
	}
	fmt.Println(user, user.Name, user.Age)

	//查询多行数据
	//sqlStr2 := "select name,id,age from user_info where stu_num > 10 and where name = ? "
	sqlStr2 := "select id, name, age, class, stu_num from user_info where stu_num > ?  "
	var user2 []StuInfo
	if err := db.Select(&user2, sqlStr2, "20"); err != nil {
		fmt.Println("get the data failed ,err :", err)
		return
	} else {
		for i := 0; i < len(user2); i++ {
			//fmt.Println(user2)
			fmt.Printf("id:%d, name:%s, age:%d\n", user2[i].ID, user2[i].Name, user2[i].Age)
		}
	}

	// 插入、更新、删除操作
	//采用Exec函数来实现,和原生sql库实现是一致的
	//  插入
	//sqlStr3 := "insert into user_info(id, name, age, class, stu_num) values (?,?,?,?,?)"
	//result, err := db.Exec(sqlStr3, 6, "张大", 12, 2, 22)
	//if err != nil {
	//	fmt.Println("insert data failed ,err :", err)
	//	return
	//}
	//insertId, err := result.LastInsertId()
	//if err != nil {
	//	fmt.Println("get insert id failed ,err :", err)
	//	return
	//}
	//fmt.Printf("insert data success ,insertid :", insertId)
	//// 更新
	//sqlStr4 := "update user_info set age =? where name = ?"
	//updateResult, err := db.Exec(sqlStr4, 11, "张大")
	//if err != nil {
	//	fmt.Println("update data failed ,err:", err)
	//	return
	//}
	//affectedRow, err := updateResult.RowsAffected()
	//if err != nil {
	//	fmt.Println("get affected row failed")
	//	return
	//}
	//fmt.Println("update data sucess,affected row :", affectedRow)

	//删除
	//sqlStr5 := "delete from user_info where name = ?"
	//deleteResult, err := db.Exec(sqlStr5, "张三")
	//if err != nil {
	//	fmt.Printf("del data failed ,err: %v", err)
	//	return
	//}
	//affectedRow, err := deleteResult.RowsAffected()
	//if err != nil {
	//	fmt.Println("get the affected row failed ,err :", err)
	//	return
	//}
	//fmt.Println("delete data failed ,affected row :", affectedRow)

	sqlStr6 := "select * from user_info where class = :class"
	rows, err := db.NamedQuery(sqlStr6, map[string]interface{}{"class": 2})
	if err != nil {
		fmt.Println("query failed ,err:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user2 StuInfo
		err := rows.StructScan(&user2)
		if err != nil {
			fmt.Println("scan failed ,err :", err)
		}
		continue
	}
	fmt.Println("user:", user2)
}
