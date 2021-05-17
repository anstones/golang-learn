package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code string   `gorm:"comment:'编码'"`
	Price uint 	  `gorm:"comment:'价格'"`
}

func main() {
	db , err := gorm.Open("mysql", "root:mysql123@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// 热更新
	db.AutoMigrate(&Product{})

	// 增
	db.Create(&Product{Code:"L1212", Price: 1000})

	// 查
	var product Product
	db.First(&product, 1) // where id=1
	//db.First(&product, "code=?", "L1212") // where code=L1212
	fmt.Println( product.Code, product.Price)

	//// 改
	//db.Model(&product).Update("Price", 2000)
	//
	////删
	//db.Delete(&product)
}