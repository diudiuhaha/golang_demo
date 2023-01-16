package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局的DB，是个连接池对象
var db *sql.DB

func initMysql() (err error) { //连接数据库
	//数据库信息
	dsn := "root:123456@tcp(10.0.0.10:3306)/demo"
	//连接,(数据库驱动名，数据库远程连接信息)
	db, err = sql.Open("mysql", dsn)
	//这里是 = 不能是:= ，否则会报错panic: runtime error: invalid memory address or nil pointer dereference
	if err != nil {
		fmt.Println("link to database failed,err:", err)
		return
	}
	//尝试连接
	err = db.Ping()
	if err != nil {
		fmt.Println("open database failed ,err", err)
		return
	}
	fmt.Println("连接数据库成功")
	return //函数返回
}

// 定义一个结构体，用于与数据库的对应
// 因为这个结构体后面都要使用，所以移除了作为全局变量吧
type StuInfo struct {
	Id   string `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 单行查询
func queryDemo() {
	sqlStr := "select id,name,age from user where id = ?"
	prepare, err := db.Prepare(sqlStr) //发送命令进行预处理
	if err != nil {
		fmt.Println("prepare failed,err:", err)
		return
	}
	defer prepare.Close() //执行完查询后要关闭预处理
	rows, err := prepare.Query(1)
	if err != nil {
		fmt.Println("query failed ,err:", err)
		return
	}
	defer func(rows *sql.Rows) { ////同样的 ，rows也要及时的关闭
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var u StuInfo
	var stuId = [3]int{6, 7}
	for i := 0; i < len(stuId); i++ {

		err := db.QueryRow(sqlStr, stuId[i]).Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Println("scan failed ,err", err)
			return
		}
		fmt.Println(u)
	}
}

// 预处理插入处理
func prepareInsert() {
	sqlStr := "insert into user(name, age) values (?,?)"
	prepare, err := db.Prepare(sqlStr)
	if err != nil {
		return
	}
	//应该立刻去写一个defer 关闭 预处理
	defer prepare.Close()

	//传输数据
	_, err = prepare.Exec("tom", 22)
	if err != nil {
		fmt.Printf("%v insert failed ,err:", err)
		return
	}
	_, err = prepare.Exec("jack", 23)
	if err != nil {
		fmt.Println("insert failed ,err :", err)
		return
	}
	fmt.Println("insert data sucess!!")

}

func main() {
	err := initMysql()
	if err != nil {
		fmt.Println("Mysql初始化失败")
	}
	defer queryDemo()
	prepareInsert()
}
