package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err error
)

func init()  {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/dataset_manager"

	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db

	dbConn, err = sql.Open("mysql", dsn)
	if err !=nil{
		return
	}

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = dbConn.Ping()
	if err != nil{
		return
	}

}