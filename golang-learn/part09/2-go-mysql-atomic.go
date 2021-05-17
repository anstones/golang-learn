package main

import (
"fmt"
_ "github.com/go-sql-driver/mysql"
"github.com/jmoiron/sqlx"
)


type Person struct {
	UserId int 	`db:"user_id"`
	UserName string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}


type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}




var DB *sqlx.DB



func init() {
	database, err := sqlx.Open("mysql", "root:mysql123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	DB = database
	//defer DB.Close()
}

func HandlerError(err error, why string)  {
	if err != nil{
		fmt.Println(why, err)
	}
	return
}

func main() {
	conn, err := DB.Begin()
	HandlerError(err, "begin failed ")

	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil{
		fmt.Println("exec failed", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil{
		fmt.Println("exec failed", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()

}
