package part01_example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-learn/part01-example/dbops"
	"net/http"
)

func GetBook(c *gin.Context){
	books, err := dbops.GetAllbook()
	if err != nil{
		fmt.Println("获取参数失败", err)
	}
	for i, ele := range books{
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": books})

}