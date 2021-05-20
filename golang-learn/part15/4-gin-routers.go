package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	return r
}



func main() {
	//r := gin.Default()
	//r.GET("/topgoer", helloHandler)
	r := setupRouter()
	if err := r.Run(); err != nil{
		fmt.Println("startup service failed, err:%v\n", err)
	}
}