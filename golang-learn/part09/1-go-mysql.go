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


func insert()  {
	r, err := DB.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert succ:", id)
}

func selects()  {
	var person []Person
	err := DB.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
	HandlerError(err, "查询失败")
	fmt.Println("select succ:", person)
}

func update()  {
	res, err := DB.Exec("update person set username=? where user_id=?", "aaaa", 2)
	HandlerError(err, "更新失败")

	row, err := res.RowsAffected()
	HandlerError(err, "rows failed")
	fmt.Println("update succ:", row)
}

func delete()  {
	res, err := DB.Exec("delete from person where user_id=?", 3)
	HandlerError(err, "exec failed")
	row,err := res.RowsAffected()
	HandlerError(err, "rows failed")
	fmt.Println("delete succ: ",row)
}


func main() {
	defer DB.Close()
	//insert()
	//selects()
	//update()
	delete()

}